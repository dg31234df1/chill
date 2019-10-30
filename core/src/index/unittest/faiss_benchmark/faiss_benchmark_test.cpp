// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

#include <gtest/gtest.h>

#include <cassert>
#include <cmath>
#include <cstdio>
#include <cstdlib>
#include <cstring>

#include <faiss/AutoTune.h>
#include <faiss/Index.h>
#include <faiss/IndexIVF.h>
#include <faiss/gpu/GpuAutoTune.h>
#include <faiss/gpu/GpuIndexFlat.h>
#ifdef CUSTOMIZATION
#include <faiss/gpu/GpuIndexIVFSQHybrid.h>
#endif
#include <faiss/gpu/StandardGpuResources.h>
#include <faiss/index_io.h>
#include <faiss/utils.h>

#include <hdf5.h>

#include <sys/stat.h>
#include <sys/time.h>
#include <sys/types.h>
#include <unistd.h>
#include <vector>

/*****************************************************
 * To run this test, please download the HDF5 from
 *  https://support.hdfgroup.org/ftp/HDF5/releases/
 * and install it to /usr/local/hdf5 .
 *****************************************************/
#define DEBUG_VERBOSE 0

const std::string HDF5_POSTFIX = ".hdf5";
const std::string HDF5_DATASET_TRAIN = "train";
const std::string HDF5_DATASET_TEST = "test";
const std::string HDF5_DATASET_NEIGHBORS = "neighbors";
const std::string HDF5_DATASET_DISTANCES = "distances";

enum QueryMode {
    MODE_CPU = 0,
    MODE_MIX,
    MODE_GPU
};

double
elapsed() {
    struct timeval tv;
    gettimeofday(&tv, nullptr);
    return tv.tv_sec + tv.tv_usec * 1e-6;
}

void
normalize(float* arr, size_t nq, size_t dim) {
    for (size_t i = 0; i < nq; i++) {
        double vecLen = 0.0, inv_vecLen = 0.0;
        for (size_t j = 0; j < dim; j++) {
            double val = arr[i * dim + j];
            vecLen += val * val;
        }
        inv_vecLen = 1.0 / std::sqrt(vecLen);
        for (size_t j = 0; j < dim; j++) {
            arr[i * dim + j] = (float)(arr[i * dim + j] * inv_vecLen);
        }
    }
}

void*
hdf5_read(const std::string& file_name, const std::string& dataset_name, H5T_class_t dataset_class,
          size_t& d_out, size_t& n_out) {
    hid_t file, dataset, datatype, dataspace, memspace;
    H5T_class_t t_class;   /* data type class */
    hsize_t dimsm[3];      /* memory space dimensions */
    hsize_t dims_out[2];   /* dataset dimensions */
    hsize_t count[2];      /* size of the hyperslab in the file */
    hsize_t offset[2];     /* hyperslab offset in the file */
    hsize_t count_out[3];  /* size of the hyperslab in memory */
    hsize_t offset_out[3]; /* hyperslab offset in memory */
    void* data_out; /* output buffer */

    /* Open the file and the dataset. */
    file = H5Fopen(file_name.c_str(), H5F_ACC_RDONLY, H5P_DEFAULT);
    dataset = H5Dopen2(file, dataset_name.c_str(), H5P_DEFAULT);

    /* Get datatype and dataspace handles and then query
     * dataset class, order, size, rank and dimensions. */
    datatype = H5Dget_type(dataset); /* datatype handle */
    t_class = H5Tget_class(datatype);
    assert(t_class == dataset_class || !"Illegal dataset class type");

    dataspace = H5Dget_space(dataset); /* dataspace handle */
    H5Sget_simple_extent_dims(dataspace, dims_out, NULL);
    n_out = dims_out[0];
    d_out = dims_out[1];

    /* Define hyperslab in the dataset. */
    offset[0] = offset[1] = 0;
    count[0] = dims_out[0];
    count[1] = dims_out[1];
    H5Sselect_hyperslab(dataspace, H5S_SELECT_SET, offset, NULL, count, NULL);

    /* Define the memory dataspace. */
    dimsm[0] = dims_out[0];
    dimsm[1] = dims_out[1];
    dimsm[2] = 1;
    memspace = H5Screate_simple(3, dimsm, NULL);

    /* Define memory hyperslab. */
    offset_out[0] = offset_out[1] = offset_out[2] = 0;
    count_out[0] = dims_out[0];
    count_out[1] = dims_out[1];
    count_out[2] = 1;
    H5Sselect_hyperslab(memspace, H5S_SELECT_SET, offset_out, NULL, count_out, NULL);

    /* Read data from hyperslab in the file into the hyperslab in memory and display. */
    switch (t_class) {
        case H5T_INTEGER:
            data_out = new int[dims_out[0] * dims_out[1]];
            H5Dread(dataset, H5T_NATIVE_INT, memspace, dataspace, H5P_DEFAULT, data_out);
            break;
        case H5T_FLOAT:
            data_out = new float[dims_out[0] * dims_out[1]];
            H5Dread(dataset, H5T_NATIVE_FLOAT, memspace, dataspace, H5P_DEFAULT, data_out);
            break;
        default:
            printf("Illegal dataset class type\n");
            break;
    }

    /* Close/release resources. */
    H5Tclose(datatype);
    H5Dclose(dataset);
    H5Sclose(dataspace);
    H5Sclose(memspace);
    H5Fclose(file);

    return data_out;
}

std::string
get_index_file_name(const std::string& ann_test_name, const std::string& index_key, int32_t data_loops) {
    size_t pos = index_key.find_first_of(',', 0);
    std::string file_name = ann_test_name;
    file_name = file_name + "_" + index_key.substr(0, pos) + "_" + index_key.substr(pos + 1);
    file_name = file_name + "_" + std::to_string(data_loops) + ".index";
    return file_name;
}

bool
parse_ann_test_name(const std::string& ann_test_name, size_t& dim, faiss::MetricType& metric_type) {
    size_t pos1, pos2;

    if (ann_test_name.empty())
        return false;

    pos1 = ann_test_name.find_first_of('-', 0);
    if (pos1 == std::string::npos)
        return false;
    pos2 = ann_test_name.find_first_of('-', pos1 + 1);
    if (pos2 == std::string::npos)
        return false;

    dim = std::stoi(ann_test_name.substr(pos1 + 1, pos2 - pos1 - 1));
    std::string metric_str = ann_test_name.substr(pos2 + 1);
    if (metric_str == "angular") {
        metric_type = faiss::METRIC_INNER_PRODUCT;
    } else if (metric_str == "euclidean") {
        metric_type = faiss::METRIC_L2;
    } else {
        return false;
    }

    return true;
}

int32_t
GetResultHitCount(const faiss::Index::idx_t* ground_index, const faiss::Index::idx_t* index, size_t ground_k, size_t k,
                  size_t nq, int32_t index_add_loops) {
    assert(ground_k <= k);
    int hit = 0;
    for (int i = 0; i < nq; i++) {
        // count the num of results exist in ground truth result set
        // each result replicates INDEX_ADD_LOOPS times
        for (int j_c = 0; j_c < ground_k; j_c++) {
            int r_c = index[i * k + j_c];
            for (int j_g = 0; j_g < ground_k / index_add_loops; j_g++) {
                if (ground_index[i * ground_k + j_g] == r_c) {
                    hit++;
                    continue;
                }
            }
        }
    }
    return hit;
}

#if DEBUG_VERBOSE
void
print_array(const char* header, bool is_integer, const void* arr, size_t nq, size_t k) {
    const int ROW = 10;
    const int COL = 10;
    assert(ROW <= nq);
    assert(COL <= k);
    printf("%s\n", header);
    printf("==============================================\n");
    for (int i = 0; i < 10; i++) {
        for (int j = 0; j < 10; j++) {
            if (is_integer) {
                printf("%7ld ", ((int64_t*)arr)[i * k + j]);
            } else {
                printf("%.6f ", ((float*)arr)[i * k + j]);
            }
        }
        printf("\n");
    }
    printf("\n");
}
#endif

void
load_base_data(faiss::Index* &index, const std::string& ann_test_name, const std::string& index_key,
               faiss::gpu::StandardGpuResources& res, const faiss::MetricType metric_type, const size_t dim,
               int32_t index_add_loops, QueryMode mode = MODE_CPU) {
    double t0 = elapsed();

    const std::string ann_file_name = ann_test_name + HDF5_POSTFIX;
    const int GPU_DEVICE_IDX = 0;

    faiss::Index *cpu_index = nullptr, *gpu_index = nullptr;
    faiss::distance_compute_blas_threshold = 800;

    std::string index_file_name = get_index_file_name(ann_test_name, index_key, index_add_loops);

    try {
        printf("[%.3f s] Reading index file: %s\n", elapsed() - t0, index_file_name.c_str());
        cpu_index = faiss::read_index(index_file_name.c_str());

        if (mode != MODE_CPU) {
            faiss::gpu::GpuClonerOptions option;
            option.allInGpu = true;

            faiss::IndexComposition index_composition;
            index_composition.index = cpu_index;
            index_composition.quantizer = nullptr;

            switch (mode) {
                case MODE_CPU:
                    assert(false);
                    break;
                case MODE_MIX:
                    index_composition.mode = 1;  // 0: all data, 1: copy quantizer, 2: copy data
                    break;
                case MODE_GPU:
                    index_composition.mode = 0;  // 0: all data, 1: copy quantizer, 2: copy data
                    break;
            }

            printf("[%.3f s] Cloning CPU index to GPU\n", elapsed() - t0);
            gpu_index = faiss::gpu::index_cpu_to_gpu(&res, GPU_DEVICE_IDX, &index_composition, &option);
        }
    } catch (...) {
        size_t nb, d;
        printf("[%.3f s] Loading HDF5 file: %s\n", elapsed() - t0, ann_file_name.c_str());
        float* xb = (float*)hdf5_read(ann_file_name, HDF5_DATASET_TRAIN, H5T_FLOAT, d, nb);
        assert(d == dim || !"dataset does not have correct dimension");

        if (metric_type == faiss::METRIC_INNER_PRODUCT) {
            printf("[%.3f s] Normalizing base data set \n", elapsed() - t0);
            normalize(xb, nb, d);
        }

        printf("[%.3f s] Creating CPU index \"%s\" d=%ld\n", elapsed() - t0, index_key.c_str(), d);
        cpu_index = faiss::index_factory(d, index_key.c_str(), metric_type);

        printf("[%.3f s] Cloning CPU index to GPU\n", elapsed() - t0);
        gpu_index = faiss::gpu::index_cpu_to_gpu(&res, GPU_DEVICE_IDX, cpu_index);

        printf("[%.3f s] Training on %ld vectors\n", elapsed() - t0, nb);
        gpu_index->train(nb, xb);

        // add index multiple times to get ~1G data set
        for (int i = 0; i < index_add_loops; i++) {
            printf("[%.3f s] No.%d Indexing database, size %ld*%ld\n", elapsed() - t0, i, nb, d);
            gpu_index->add(nb, xb);
        }

        printf("[%.3f s] Coping GPU index to CPU\n", elapsed() - t0);
        delete cpu_index;
        cpu_index = faiss::gpu::index_gpu_to_cpu(gpu_index);

        faiss::IndexIVF *cpu_ivf_index = dynamic_cast<faiss::IndexIVF *>(cpu_index);
        if (cpu_ivf_index != nullptr) {
            cpu_ivf_index->to_readonly();
        }

        printf("[%.3f s] Writing index file: %s\n", elapsed() - t0, index_file_name.c_str());
        faiss::write_index(cpu_index, index_file_name.c_str());

        delete[] xb;
    }

    switch (mode) {
        case MODE_CPU:
        case MODE_MIX:
            index = cpu_index;
            if (gpu_index) {
                delete gpu_index;
            }
            break;
        case MODE_GPU:
            index = gpu_index;
            if (cpu_index) {
                delete cpu_index;
            }
            break;
    }
}

void
load_query_data(faiss::Index::distance_t* &xq, size_t& nq, const std::string& ann_test_name,
                const faiss::MetricType metric_type, const size_t dim) {
    double t0 = elapsed();
    size_t d;

    const std::string ann_file_name = ann_test_name + HDF5_POSTFIX;

    xq = (float *) hdf5_read(ann_file_name, HDF5_DATASET_TEST, H5T_FLOAT, d, nq);
    assert(d == dim || !"query does not have same dimension as train set");

    if (metric_type == faiss::METRIC_INNER_PRODUCT) {
        printf("[%.3f s] Normalizing query data \n", elapsed() - t0);
        normalize(xq, nq, d);
    }
}

void
load_ground_truth(faiss::Index::idx_t* &gt, size_t& k, const std::string& ann_test_name, const size_t nq) {
    const std::string ann_file_name = ann_test_name + HDF5_POSTFIX;

    // load ground-truth and convert int to long
    size_t nq2;
    int *gt_int = (int *) hdf5_read(ann_file_name, HDF5_DATASET_NEIGHBORS, H5T_INTEGER, k, nq2);
    assert(nq2 == nq || !"incorrect nb of ground truth index");

    gt = new faiss::Index::idx_t[k * nq];
    for (int i = 0; i < k * nq; i++) {
        gt[i] = gt_int[i];
    }
    delete[] gt_int;

#if DEBUG_VERBOSE
    faiss::Index::distance_t* gt_dist;  // nq * k matrix of ground-truth nearest-neighbors distances
    gt_dist = (float*)hdf5_read(ann_file_name, HDF5_DATASET_DISTANCES, H5T_FLOAT, k, nq2);
    assert(nq2 == nq || !"incorrect nb of ground truth distance");

    std::string str;
    str = ann_test_name + " ground truth index";
    print_array(str.c_str(), true, gt, nq, k);
    str = ann_test_name + " ground truth distance";
    print_array(str.c_str(), false, gt_dist, nq, k);

    delete gt_dist;
#endif
}

void
test_with_nprobes(const std::string& ann_test_name, const std::string& index_key, faiss::Index* index,
                  faiss::gpu::StandardGpuResources& res, const QueryMode query_mode,
                  const faiss::Index::distance_t *xq, const faiss::Index::idx_t *gt, const std::vector<size_t> nprobes,
                  const int32_t index_add_loops, const int32_t search_loops) {
    const size_t NQ = 1000, NQ_START = 10, NQ_STEP = 10;
    const size_t K = 1000, K_START = 100, K_STEP = 10;
    const size_t GK = 100;  // topk of ground truth

    std::unordered_map<size_t, std::string> mode_str_map =
            {{MODE_CPU, "MODE_CPU"}, {MODE_MIX, "MODE_MIX"}, {MODE_GPU, "MODE_GPU"}};

    for (auto nprobe : nprobes) {
        switch (query_mode) {
            case MODE_CPU:
            case MODE_MIX: {
                faiss::ParameterSpace params;
                std::string nprobe_str = "nprobe=" + std::to_string(nprobe);
                params.set_index_parameters(index, nprobe_str.c_str());
                break;
            }
            case MODE_GPU: {
                faiss::gpu::GpuIndexIVF *gpu_index_ivf = dynamic_cast<faiss::gpu::GpuIndexIVF*>(index);
                gpu_index_ivf->setNumProbes(nprobe);
            }
        }

        // output buffers
        faiss::Index::idx_t *I = new faiss::Index::idx_t[NQ * K];
        faiss::Index::distance_t *D = new faiss::Index::distance_t[NQ * K];

        printf("\n%s | %s - %s | nprobe=%lu\n", ann_test_name.c_str(), index_key.c_str(),
                mode_str_map[query_mode].c_str(), nprobe);
        printf("======================================================================================\n");
        for (size_t t_nq = NQ_START; t_nq <= NQ; t_nq *= NQ_STEP) {   // nq = {10, 100, 1000}
            for (size_t t_k = K_START; t_k <= K; t_k *= K_STEP) {  //  k = {100, 1000}
                faiss::indexIVF_stats.quantization_time = 0.0;
                faiss::indexIVF_stats.search_time = 0.0;

                double t_start = elapsed(), t_end;
                for (int i = 0; i < search_loops; i++) {
                    index->search(t_nq, xq, t_k, D, I);
                }
                t_end = elapsed();

#if DEBUG_VERBOSE
                std::string str;
                str = "I (" + index_key + ", nq=" + std::to_string(t_nq) + ", k=" + std::to_string(t_k) + ")";
                print_array(str.c_str(), true, I, t_nq, t_k);
                str = "D (" + index_key + ", nq=" + std::to_string(t_nq) + ", k=" + std::to_string(t_k) + ")";
                print_array(str.c_str(), false, D, t_nq, t_k);
#endif

                // k = 100 for ground truth
                int32_t hit = GetResultHitCount(gt, I, GK, t_k, t_nq, index_add_loops);

                printf("nq = %4ld, k = %4ld, elapse = %.4fs (quant = %.4fs, search = %.4fs), R@ = %.4f\n", t_nq, t_k,
                       (t_end - t_start) / search_loops, faiss::indexIVF_stats.quantization_time / 1000 / search_loops,
                       faiss::indexIVF_stats.search_time / 1000 / search_loops,
                       (hit / float(t_nq * GK / index_add_loops)));
            }
        }
        printf("======================================================================================\n");

        delete[] I;
        delete[] D;
    }
}

void
test_ann_hdf5(const std::string& ann_test_name, const std::string& index_key, const QueryMode query_mode,
              int32_t index_add_loops, const std::vector<size_t>& nprobes, int32_t search_loops) {
    double t0 = elapsed();

    faiss::gpu::StandardGpuResources res;

    faiss::MetricType metric_type;
    size_t dim;

    if (query_mode == MODE_MIX && index_key.find("SQ8Hybrid") == std::string::npos) {
        printf("Only SQ8Hybrid support MODE_MIX\n");
        return;
    }

    if (!parse_ann_test_name(ann_test_name, dim, metric_type)) {
        printf("Invalid ann test name: %s\n", ann_test_name.c_str());
        return;
    }

    size_t nq, k;
    faiss::Index* index;
    faiss::Index::distance_t* xq;
    faiss::Index::idx_t* gt;  // ground-truth index

    printf("[%.3f s] Loading base data\n", elapsed() - t0);
    load_base_data(index, ann_test_name, index_key, res, metric_type, dim, index_add_loops, query_mode);

    printf("[%.3f s] Loading queries\n", elapsed() - t0);
    load_query_data(xq, nq, ann_test_name, metric_type, dim);

    printf("[%.3f s] Loading ground truth for %ld queries\n", elapsed() - t0, nq);
    load_ground_truth(gt, k, ann_test_name, nq);

    test_with_nprobes(ann_test_name, index_key, index, res, query_mode, xq, gt, nprobes, index_add_loops, search_loops);
    printf("[%.3f s] Search test done\n\n", elapsed() - t0);

    delete[] xq;
    delete[] gt;
    delete index;
}

/************************************************************************************
 * https://github.com/erikbern/ann-benchmarks
 *
 * Dataset 	Dimensions 	Train_size 	Test_size 	Neighbors 	Distance 	Download
 * Fashion-
 *  MNIST   784         60,000      10,000 	    100         Euclidean   HDF5 (217MB)
 * GIST     960         1,000,000   1,000       100         Euclidean   HDF5 (3.6GB)
 * GloVe    100         1,183,514   10,000      100         Angular     HDF5 (463MB)
 * GloVe    200         1,183,514   10,000      100         Angular     HDF5 (918MB)
 * MNIST    784         60,000 	    10,000      100         Euclidean   HDF5 (217MB)
 * NYTimes  256         290,000     10,000      100         Angular     HDF5 (301MB)
 * SIFT     128         1,000,000   10,000      100         Euclidean   HDF5 (501MB)
 *************************************************************************************/

TEST(FAISSTEST, BENCHMARK) {
    std::vector<size_t> param_nprobes = {8, 128};
    const int32_t SEARCH_LOOPS = 5;
    const int32_t SIFT_INSERT_LOOPS = 2;  // insert twice to get ~1G data set
    const int32_t GLOVE_INSERT_LOOPS = 1;

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    test_ann_hdf5("sift-128-euclidean", "IVF16384,Flat", MODE_CPU, SIFT_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
    test_ann_hdf5("sift-128-euclidean", "IVF16384,Flat", MODE_GPU, SIFT_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);

    test_ann_hdf5("sift-128-euclidean", "IVF16384,SQ8", MODE_CPU, SIFT_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
    test_ann_hdf5("sift-128-euclidean", "IVF16384,SQ8", MODE_GPU, SIFT_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);

#ifdef CUSTOMIZATION
    test_ann_hdf5("sift-128-euclidean", "IVF16384,SQ8Hybrid", MODE_CPU, SIFT_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
    test_ann_hdf5("sift-128-euclidean", "IVF16384,SQ8Hybrid", MODE_GPU, SIFT_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
//    test_ann_hdf5("sift-128-euclidean", "IVF16384,SQ8Hybrid", MODE_MIX, SIFT_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
#endif

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    test_ann_hdf5("glove-200-angular", "IVF16384,Flat", MODE_CPU, GLOVE_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
    test_ann_hdf5("glove-200-angular", "IVF16384,Flat", MODE_GPU, GLOVE_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);

    test_ann_hdf5("glove-200-angular", "IVF16384,SQ8", MODE_CPU, GLOVE_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
    test_ann_hdf5("glove-200-angular", "IVF16384,SQ8", MODE_GPU, GLOVE_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);

#ifdef CUSTOMIZATION
    test_ann_hdf5("glove-200-angular", "IVF16384,SQ8Hybrid", MODE_CPU, GLOVE_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
    test_ann_hdf5("glove-200-angular", "IVF16384,SQ8Hybrid", MODE_GPU, GLOVE_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
//    test_ann_hdf5("glove-200-angular", "IVF16384,SQ8Hybrid", MODE_MIX, GLOVE_INSERT_LOOPS, param_nprobes, SEARCH_LOOPS);
#endif
}
