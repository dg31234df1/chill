/**
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

#include <cstdio>
#include <cstdlib>

#include <faiss/IndexFlat.h>
#include <faiss/IndexIVFPQ.h>
#include <faiss/utils/BitsetView.h>


int main() {
    int d = 64;                            // dimension
    int nb = 100000;                       // database size
    int nq = 10;//10000;                        // nb of queries
    faiss::ConcurrentBitsetPtr bitset = std::make_shared<faiss::ConcurrentBitset>(nb);

    float *xb = new float[d * nb];
    float *xq = new float[d * nq];

    for(int i = 0; i < nb; i++) {
        for(int j = 0; j < d; j++)
            xb[d * i + j] = drand48();
        xb[d * i] += i / 1000.;
    }

    srand((unsigned)time(NULL));
    printf("delete ids: \n");
    for(int i = 0; i < nq; i++) {
        auto tmp = rand()%nb;
        bitset->set(tmp);
        printf("%d ", tmp);
        for(int j = 0; j < d; j++)
            xq[d * i + j] = xb[d * tmp + j];
//        xq[d * i] += i / 1000.;
    }
    printf("\n");


    int nlist = 100;
    int k = 4;
    int m = 8;                             // bytes per vector
    faiss::IndexFlatL2 quantizer(d);       // the other index
    faiss::IndexIVFPQ index(&quantizer, d, nlist, m, 8);
    // here we specify METRIC_L2, by default it performs inner-product search
    index.train(nb, xb);
    index.add(nb, xb);

    printf("------------sanity check----------------\n");
    {       // sanity check
        long *I = new long[k * 5];
        float *D = new float[k * 5];

        index.search(5, xb, k, D, I);

        printf("I=\n");
        for(int i = 0; i < 5; i++) {
            for(int j = 0; j < k; j++)
                printf("%5ld ", I[i * k + j]);
            printf("\n");
        }

        printf("D=\n");
        for(int i = 0; i < 5; i++) {
            for(int j = 0; j < k; j++)
                printf("%7g ", D[i * k + j]);
            printf("\n");
        }

        delete [] I;
        delete [] D;
    }

    printf("---------------search xq-------------\n");
    {       // search xq
        long *I = new long[k * nq];
        float *D = new float[k * nq];

        index.nprobe = 10;
        index.search(nq, xq, k, D, I);

        printf("I=\n");
        for(int i = 0; i < nq; i++) {
            for(int j = 0; j < k; j++)
                printf("%5ld ", I[i * k + j]);
            printf("\n");
        }

        delete [] I;
        delete [] D;
    }

    printf("----------------search xq with delete------------\n");
    {       // search xq with delete
        long *I = new long[k * nq];
        float *D = new float[k * nq];

        index.nprobe = 10;
        index.search(nq, xq, k, D, I, bitset);

        printf("I=\n");
        for(int i = 0; i < nq; i++) {
            for(int j = 0; j < k; j++)
                printf("%5ld ", I[i * k + j]);
            printf("\n");
        }

        delete [] I;
        delete [] D;
    }



    delete [] xb;
    delete [] xq;

    return 0;
}
