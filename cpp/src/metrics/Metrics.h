/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#pragma once

#include "utils/Error.h"
#include <memory>
#include <vector>


#include <prometheus/registry.h>
#include <prometheus/exposer.h>
#include "server/ServerConfig.h"

#define METRICS_NOW_TIME std::chrono::system_clock::now()
#define METRICS_INSTANCE server::GetInstance()
#define METRICS_MICROSECONDS(a,b) (std::chrono::duration_cast<std::chrono::microseconds> (b-a)).count();


namespace zilliz {
namespace vecwise {
namespace server {

class MetricsBase{
 public:
    static MetricsBase&
    GetInstance(){
        static MetricsBase instance;
        return instance;
    }

    virtual ServerError Init() {};
    virtual void AddGroupSuccessTotalIncrement(double value = 1) {};
    virtual void AddGroupFailTotalIncrement(double value = 1) {};
    virtual void HasGroupSuccessTotalIncrement(double value = 1) {};
    virtual void HasGroupFailTotalIncrement(double value = 1) {};
    virtual void GetGroupSuccessTotalIncrement(double value = 1) {};
    virtual void GetGroupFailTotalIncrement(double value = 1) {};
    virtual void GetGroupFilesSuccessTotalIncrement(double value = 1) {};
    virtual void GetGroupFilesFailTotalIncrement(double value = 1) {};
    virtual void AddVectorsSuccessTotalIncrement(double value = 1) {};
    virtual void AddVectorsFailTotalIncrement(double value = 1) {};
    virtual void AddVectorsDurationHistogramOberve(double value) {};
    virtual void SearchSuccessTotalIncrement(double value = 1) {};
    virtual void SearchFailTotalIncrement(double value = 1) {};
    virtual void SearchDurationHistogramObserve(double value) {};
    virtual void RawFileSizeHistogramObserve(double value) {};
    virtual void IndexFileSizeHistogramObserve(double value) {};
    virtual void BuildIndexDurationSecondsHistogramObserve(double value) {};
    virtual void AllBuildIndexDurationSecondsHistogramObserve(double value) {};
    virtual void CacheUsageGaugeIncrement(double value = 1) {};
    virtual void CacheUsageGaugeDecrement(double value = 1) {};
    virtual void CacheUsageGaugeSet(double value) {};
    virtual void MetaVisitTotalIncrement(double value = 1) {};
    virtual void MetaVisitDurationSecondsHistogramObserve(double value) {};
    virtual void MemUsagePercentGaugeSet(double value) {};
    virtual void MemUsagePercentGaugeIncrement(double value = 1) {};
    virtual void MemUsagePercentGaugeDecrement(double value = 1) {};
    virtual void MemUsageTotalGaugeSet(double value) {};
    virtual void MemUsageTotalGaugeIncrement(double value = 1) {};
    virtual void MemUsageTotalGaugeDecrement(double value = 1) {};
    virtual void MetaAccessTotalIncrement(double value = 1) {};
    virtual void MetaAccessDurationSecondsHistogramObserve(double value) {};
    virtual void FaissDiskLoadDurationSecondsHistogramObserve(double value) {};
    virtual void FaissDiskLoadSizeBytesHistogramObserve(double value) {};
    virtual void FaissDiskLoadIOSpeedHistogramObserve(double value) {};
    virtual void CacheAccessTotalIncrement(double value = 1) {};
    virtual void MemTableMergeDurationSecondsHistogramObserve(double value) {};
    virtual void SearchIndexDataDurationSecondsHistogramObserve(double value) {};
    virtual void SearchRawDataDurationSecondsHistogramObserve(double value) {};
    virtual void IndexFileSizeTotalIncrement(double value = 1) {};
    virtual void RawFileSizeTotalIncrement(double value = 1) {};
    virtual void IndexFileSizeGaugeSet(double value) {};
    virtual void RawFileSizeGaugeSet(double value) {};



};

enum class MetricCollectorType{
    INVALID,
    PROMETHEUS,
    ZABBIX
};



class PrometheusMetrics: public MetricsBase {

 public:
    static PrometheusMetrics &
    GetInstance() {
//        switch(MetricCollectorType) {
//            case: prometheus::
//                static
//        }
        static PrometheusMetrics instance;
        return instance;
    }

    ServerError
    Init();

 private:
    std::shared_ptr<prometheus::Exposer> exposer_ptr_;
    std::shared_ptr<prometheus::Registry> registry_ = std::make_shared<prometheus::Registry>();
    bool startup_ = false;
 public:

    void AddGroupSuccessTotalIncrement(double value = 1.0) override { if(startup_) add_group_success_total_.Increment(value);};
    void AddGroupFailTotalIncrement(double value = 1.0) override { if(startup_) add_group_fail_total_.Increment(value);};
    void HasGroupSuccessTotalIncrement(double value = 1.0) override { if(startup_) has_group_success_total_.Increment(value);};
    void HasGroupFailTotalIncrement(double value = 1.0) override { if(startup_) has_group_fail_total_.Increment(value);};
    void GetGroupSuccessTotalIncrement(double value = 1.0) override { if(startup_) get_group_success_total_.Increment(value);};
    void GetGroupFailTotalIncrement(double value = 1.0) override { if(startup_) get_group_fail_total_.Increment(value);};
    void GetGroupFilesSuccessTotalIncrement(double value = 1.0) override { if(startup_) get_group_files_success_total_.Increment(value);};
    void GetGroupFilesFailTotalIncrement(double value = 1.0) override { if(startup_) get_group_files_fail_total_.Increment(value);};
    void AddVectorsSuccessTotalIncrement(double value = 1.0) override { if(startup_) add_vectors_success_total_.Increment(value);};
    void AddVectorsFailTotalIncrement(double value = 1.0) override { if(startup_) add_vectors_fail_total_.Increment(value);};
    void AddVectorsDurationHistogramOberve(double value) override { if(startup_) add_vectors_duration_histogram_.Observe(value);};
    void SearchSuccessTotalIncrement(double value = 1.0) override { if(startup_) search_success_total_.Increment(value);};
    void SearchFailTotalIncrement(double value = 1.0) override { if(startup_) search_fail_total_.Increment(value); };
    void SearchDurationHistogramObserve(double value) override { if(startup_) search_duration_histogram_.Observe(value);};
    void RawFileSizeHistogramObserve(double value) override { if(startup_) raw_files_size_histogram_.Observe(value);};
    void IndexFileSizeHistogramObserve(double value) override { if(startup_) index_files_size_histogram_.Observe(value);};
    void BuildIndexDurationSecondsHistogramObserve(double value) override { if(startup_) build_index_duration_seconds_histogram_.Observe(value);};
    void AllBuildIndexDurationSecondsHistogramObserve(double value) override { if(startup_) all_build_index_duration_seconds_histogram_.Observe(value);};
    void CacheUsageGaugeIncrement(double value = 1.0) override { if(startup_) cache_usage_gauge_.Increment(value);};
    void CacheUsageGaugeDecrement(double value = 1.0) override { if(startup_) cache_usage_gauge_.Decrement(value);};
    void CacheUsageGaugeSet(double value) override { if(startup_) cache_usage_gauge_.Set(value);};
//    void MetaVisitTotalIncrement(double value = 1.0) override { meta_visit_total_.Increment(value);};
//    void MetaVisitDurationSecondsHistogramObserve(double value) override { meta_visit_duration_seconds_histogram_.Observe(value);};
    void MemUsagePercentGaugeSet(double value) override { if(startup_) mem_usage_percent_gauge_.Set(value);};
    void MemUsagePercentGaugeIncrement(double value = 1.0) override { if(startup_) mem_usage_percent_gauge_.Increment(value);};
    void MemUsagePercentGaugeDecrement(double value = 1.0) override { if(startup_) mem_usage_percent_gauge_.Decrement(value);};
    void MemUsageTotalGaugeSet(double value) override { if(startup_) mem_usage_total_gauge_.Set(value);};
    void MemUsageTotalGaugeIncrement(double value = 1.0) override { if(startup_) mem_usage_total_gauge_.Increment(value);};
    void MemUsageTotalGaugeDecrement(double value = 1.0) override { if(startup_) mem_usage_total_gauge_.Decrement(value);};

    void MetaAccessTotalIncrement(double value = 1) { if(startup_) meta_access_total_.Increment(value);};
    void MetaAccessDurationSecondsHistogramObserve(double value) { if(startup_) meta_access_duration_seconds_histogram_.Observe(value);};

    void FaissDiskLoadDurationSecondsHistogramObserve(double value) { if(startup_) faiss_disk_load_duration_seconds_histogram_.Observe(value);};
    void FaissDiskLoadSizeBytesHistogramObserve(double value) { if(startup_) faiss_disk_load_size_bytes_histogram_.Observe(value);};
    void FaissDiskLoadIOSpeedHistogramObserve(double value) { if(startup_) faiss_disk_load_IO_speed_histogram_.Observe(value);};

    void CacheAccessTotalIncrement(double value = 1) { if(startup_) cache_access_total_.Increment(value);};
    void MemTableMergeDurationSecondsHistogramObserve(double value) { if(startup_) mem_table_merge_duration_seconds_histogram_.Observe(value);};
    void SearchIndexDataDurationSecondsHistogramObserve(double value) { if(startup_) search_index_data_duration_seconds_histogram_.Observe(value);};
    void SearchRawDataDurationSecondsHistogramObserve(double value) { if(startup_) search_raw_data_duration_seconds_histogram_.Observe(value);};
    void IndexFileSizeTotalIncrement(double value = 1) { if(startup_) index_file_size_total_.Increment(value);};
    void RawFileSizeTotalIncrement(double value = 1) { if(startup_) raw_file_size_total_.Increment(value);};
    void IndexFileSizeGaugeSet(double value) { if(startup_) index_file_size_gauge_.Set(value);};
    void RawFileSizeGaugeSet(double value) { if(startup_) raw_file_size_gauge_.Set(value);};





//    prometheus::Counter &connection_total() {return connection_total_; }
//
//    prometheus::Counter &add_group_success_total()  { return add_group_success_total_; }
//    prometheus::Counter &add_group_fail_total()  { return add_group_fail_total_; }
//
//    prometheus::Counter &get_group_success_total()  { return get_group_success_total_;}
//    prometheus::Counter &get_group_fail_total()  { return get_group_fail_total_;}
//
//    prometheus::Counter &has_group_success_total()  { return has_group_success_total_;}
//    prometheus::Counter &has_group_fail_total()  { return has_group_fail_total_;}
//
//    prometheus::Counter &get_group_files_success_total()  { return get_group_files_success_total_;};
//    prometheus::Counter &get_group_files_fail_total()  { return get_group_files_fail_total_;}
//
//    prometheus::Counter &add_vectors_success_total() { return add_vectors_success_total_; }
//    prometheus::Counter &add_vectors_fail_total() { return add_vectors_fail_total_; }
//
//    prometheus::Histogram &add_vectors_duration_histogram() { return add_vectors_duration_histogram_;}
//
//    prometheus::Counter &search_success_total() { return search_success_total_; }
//    prometheus::Counter &search_fail_total() { return search_fail_total_; }
//
//    prometheus::Histogram &search_duration_histogram() { return search_duration_histogram_; }
//    prometheus::Histogram &raw_files_size_histogram() { return raw_files_size_histogram_; }
//    prometheus::Histogram &index_files_size_histogram() { return index_files_size_histogram_; }
//
//    prometheus::Histogram &build_index_duration_seconds_histogram() { return build_index_duration_seconds_histogram_; }
//
//    prometheus::Histogram &all_build_index_duration_seconds_histogram() { return all_build_index_duration_seconds_histogram_; }
//
//    prometheus::Gauge &cache_usage_gauge() { return cache_usage_gauge_; }
//
//    prometheus::Counter &meta_visit_total() { return meta_visit_total_; }
//
//    prometheus::Histogram &meta_visit_duration_seconds_histogram() { return meta_visit_duration_seconds_histogram_; }
//
//    prometheus::Gauge &mem_usage_percent_gauge() { return mem_usage_percent_gauge_; }
//
//    prometheus::Gauge &mem_usage_total_gauge() { return mem_usage_total_gauge_; }




    std::shared_ptr<prometheus::Exposer> &exposer_ptr() {return exposer_ptr_; }
//    prometheus::Exposer& exposer() { return exposer_;}
    std::shared_ptr<prometheus::Registry> &registry_ptr() {return registry_; }

    // .....
 private:
    ////all from db_connection.cpp
//    prometheus::Family<prometheus::Counter> &connect_request_ = prometheus::BuildCounter()
//        .Name("connection_total")
//        .Help("total number of connection has been made")
//        .Register(*registry_);
//    prometheus::Counter &connection_total_ = connect_request_.Add({});



    ////all from DBImpl.cpp
    using BucketBoundaries = std::vector<double>;
    //record add_group request
    prometheus::Family<prometheus::Counter> &add_group_request_ = prometheus::BuildCounter()
        .Name("add_group_request_total")
        .Help("the number of add_group request")
        .Register(*registry_);

    prometheus::Counter &add_group_success_total_ = add_group_request_.Add({{"outcome", "success"}});
    prometheus::Counter &add_group_fail_total_ = add_group_request_.Add({{"outcome", "fail"}});


    //record get_group request
    prometheus::Family<prometheus::Counter> &get_group_request_ = prometheus::BuildCounter()
        .Name("get_group_request_total")
        .Help("the number of get_group request")
        .Register(*registry_);

    prometheus::Counter &get_group_success_total_ = get_group_request_.Add({{"outcome", "success"}});
    prometheus::Counter &get_group_fail_total_ = get_group_request_.Add({{"outcome", "fail"}});


    //record has_group request
    prometheus::Family<prometheus::Counter> &has_group_request_ = prometheus::BuildCounter()
        .Name("has_group_request_total")
        .Help("the number of has_group request")
        .Register(*registry_);

    prometheus::Counter &has_group_success_total_ = has_group_request_.Add({{"outcome", "success"}});
    prometheus::Counter &has_group_fail_total_ = has_group_request_.Add({{"outcome", "fail"}});


    //record get_group_files
    prometheus::Family<prometheus::Counter> &get_group_files_request_ = prometheus::BuildCounter()
        .Name("get_group_files_request_total")
        .Help("the number of get_group_files request")
        .Register(*registry_);

    prometheus::Counter &get_group_files_success_total_ = get_group_files_request_.Add({{"outcome", "success"}});
    prometheus::Counter &get_group_files_fail_total_ = get_group_files_request_.Add({{"outcome", "fail"}});


    //record add_vectors count and average time
    //need to be considered
    prometheus::Family<prometheus::Counter> &add_vectors_request_ = prometheus::BuildCounter()
        .Name("add_vectors_request_total")
        .Help("the number of vectors added")
        .Register(*registry_);
    prometheus::Counter &add_vectors_success_total_ = add_vectors_request_.Add({{"outcome", "success"}});
    prometheus::Counter &add_vectors_fail_total_ = add_vectors_request_.Add({{"outcome", "fail"}});

    prometheus::Family<prometheus::Histogram> &add_vectors_duration_seconds_ = prometheus::BuildHistogram()
        .Name("add_vector_duration_seconds")
        .Help("average time of adding every vector")
        .Register(*registry_);
    prometheus::Histogram &add_vectors_duration_histogram_ = add_vectors_duration_seconds_.Add({}, BucketBoundaries{0, 0.01, 0.02, 0.03, 0.04, 0.05, 0.08, 0.1, 0.5, 1});


    //record search count and average time
    prometheus::Family<prometheus::Counter> &search_request_ = prometheus::BuildCounter()
        .Name("search_request_total")
        .Help("the number of search request")
        .Register(*registry_);
    prometheus::Counter &search_success_total_ = search_request_.Add({{"outcome","success"}});
    prometheus::Counter &search_fail_total_ = search_request_.Add({{"outcome","fail"}});

    prometheus::Family<prometheus::Histogram> &search_request_duration_seconds_ = prometheus::BuildHistogram()
        .Name("search_request_duration_second")
        .Help("histogram of processing time for each search")
        .Register(*registry_);
    prometheus::Histogram &search_duration_histogram_ = search_request_duration_seconds_.Add({}, BucketBoundaries{0.1, 1.0, 10.0});

    //record raw_files size histogram
    prometheus::Family<prometheus::Histogram> &raw_files_size_ = prometheus::BuildHistogram()
        .Name("search_raw_files_bytes")
        .Help("histogram of raw files size by bytes")
        .Register(*registry_);
    prometheus::Histogram &raw_files_size_histogram_ = raw_files_size_.Add({}, BucketBoundaries{0.1, 1.0, 10.0});

    //record index_files size histogram
    prometheus::Family<prometheus::Histogram> &index_files_size_ = prometheus::BuildHistogram()
        .Name("search_index_files_bytes")
        .Help("histogram of index files size by bytes")
        .Register(*registry_);
    prometheus::Histogram &index_files_size_histogram_ = index_files_size_.Add({}, BucketBoundaries{0.1, 1.0, 10.0});

    //record index and raw files size counter
    prometheus::Family<prometheus::Counter> &file_size_total_ = prometheus::BuildCounter()
        .Name("search_file_size_total")
        .Help("searched index and raw file size")
        .Register(*registry_);
    prometheus::Counter &index_file_size_total_ = file_size_total_.Add({{"type", "index"}});
    prometheus::Counter &raw_file_size_total_ = file_size_total_.Add({{"type", "raw"}});

    //record index and raw files size counter
    prometheus::Family<prometheus::Gauge> &file_size_gauge_ = prometheus::BuildGauge()
        .Name("search_file_size_gauge")
        .Help("searched current index and raw file size")
        .Register(*registry_);
    prometheus::Gauge &index_file_size_gauge_ = file_size_gauge_.Add({{"type", "index"}});
    prometheus::Gauge &raw_file_size_gauge_ = file_size_gauge_.Add({{"type", "raw"}});

    //record processing time for building index
    prometheus::Family<prometheus::Histogram> &build_index_duration_seconds_ = prometheus::BuildHistogram()
        .Name("build_index_duration_seconds")
        .Help("histogram of processing time for building index")
        .Register(*registry_);
    prometheus::Histogram &build_index_duration_seconds_histogram_ = build_index_duration_seconds_.Add({}, BucketBoundaries{0.1, 1.0, 10.0});


    //record processing time for all building index
    prometheus::Family<prometheus::Histogram> &all_build_index_duration_seconds_ = prometheus::BuildHistogram()
        .Name("all_build_index_duration_seconds")
        .Help("histogram of processing time for building index")
        .Register(*registry_);
    prometheus::Histogram &all_build_index_duration_seconds_histogram_ = all_build_index_duration_seconds_.Add({}, BucketBoundaries{0.1, 1.0, 10.0});

    //record duration of merging mem table
    prometheus::Family<prometheus::Histogram> &mem_table_merge_duration_seconds_ = prometheus::BuildHistogram()
        .Name("mem_table_merge_duration_seconds")
        .Help("histogram of processing time for merging mem tables")
        .Register(*registry_);
    prometheus::Histogram &mem_table_merge_duration_seconds_histogram_ = mem_table_merge_duration_seconds_.Add({}, BucketBoundaries{0.1, 1.0, 10.0});

    //record search index and raw data duration
    prometheus::Family<prometheus::Histogram> &search_data_duration_seconds_ = prometheus::BuildHistogram()
        .Name("search_data_duration_seconds")
        .Help("histograms of processing time for search index and raw data")
        .Register(*registry_);
    prometheus::Histogram &search_index_data_duration_seconds_histogram_ = search_data_duration_seconds_.Add({{"type", "index"}}, BucketBoundaries{0.1, 1.0, 10.0});
    prometheus::Histogram &search_raw_data_duration_seconds_histogram_ = search_data_duration_seconds_.Add({{"type", "raw"}}, BucketBoundaries{0.1, 1.0, 10.0});


    ////all form Cache.cpp
    //record cache usage, when insert/erase/clear/free
    prometheus::Family<prometheus::Gauge> &cache_usage_ = prometheus::BuildGauge()
        .Name("cache_usage")
        .Help("total bytes that cache used")
        .Register(*registry_);
    prometheus::Gauge &cache_usage_gauge_ = cache_usage_.Add({});


    ////all from Meta.cpp
    //record meta visit count and time
//    prometheus::Family<prometheus::Counter> &meta_visit_ = prometheus::BuildCounter()
//        .Name("meta_visit_total")
//        .Help("the number of accessing Meta")
//        .Register(*registry_);
//    prometheus::Counter &meta_visit_total_ = meta_visit_.Add({{}});
//
//    prometheus::Family<prometheus::Histogram> &meta_visit_duration_seconds_ = prometheus::BuildHistogram()
//        .Name("meta_visit_duration_seconds")
//        .Help("histogram of processing time to get data from mata")
//        .Register(*registry_);
//    prometheus::Histogram &meta_visit_duration_seconds_histogram_ = meta_visit_duration_seconds_.Add({{}}, BucketBoundaries{0.1, 1.0, 10.0});


    ////all from MemManager.cpp
    //record memory usage percent
    prometheus::Family<prometheus::Gauge> &mem_usage_percent_ = prometheus::BuildGauge()
        .Name("memory_usage_percent")
        .Help("memory usage percent")
        .Register(*registry_);
    prometheus::Gauge &mem_usage_percent_gauge_ = mem_usage_percent_.Add({});

    //record memory usage toal
    prometheus::Family<prometheus::Gauge> &mem_usage_total_ = prometheus::BuildGauge()
        .Name("memory_usage_total")
        .Help("memory usage total")
        .Register(*registry_);
    prometheus::Gauge &mem_usage_total_gauge_ = mem_usage_total_.Add({});



    ////all from DBMetaImpl.cpp
    //record meta access count
    prometheus::Family<prometheus::Counter> &meta_access_ = prometheus::BuildCounter()
        .Name("meta_access_total")
        .Help("the number of meta accessing")
        .Register(*registry_);
    prometheus::Counter &meta_access_total_ = meta_access_.Add({});

    //record meta access duration
    prometheus::Family<prometheus::Histogram> &meta_access_duration_seconds_ = prometheus::BuildHistogram()
        .Name("meta_access_duration_seconds")
        .Help("histogram of processing time for accessing mata")
        .Register(*registry_);
    prometheus::Histogram &meta_access_duration_seconds_histogram_ = meta_access_duration_seconds_.Add({}, BucketBoundaries{0.1, 1.0, 10.0});



    ////all from FaissExecutionEngine.cpp
    //record data loading from disk count, size, duration, IO speed
    prometheus::Family<prometheus::Histogram> &disk_load_duration_second_ = prometheus::BuildHistogram()
        .Name("disk_load_duration_seconds")
        .Help("Histogram of processing time for loading data from disk")
        .Register(*registry_);
    prometheus::Histogram &faiss_disk_load_duration_seconds_histogram_ = disk_load_duration_second_.Add({{"DB","Faiss"}},BucketBoundaries{0.1, 1.0, 10.0});

    prometheus::Family<prometheus::Histogram> &disk_load_size_bytes_ = prometheus::BuildHistogram()
        .Name("disk_load_size_bytes")
        .Help("Histogram of data size by bytes for loading data from disk")
        .Register(*registry_);
    prometheus::Histogram &faiss_disk_load_size_bytes_histogram_ = disk_load_size_bytes_.Add({{"DB","Faiss"}},BucketBoundaries{0.1, 1.0, 10.0});

    prometheus::Family<prometheus::Histogram> &disk_load_IO_speed_ = prometheus::BuildHistogram()
        .Name("disk_load_IO_speed_byte_per_sec")
        .Help("Histogram of IO speed for loading data from disk")
        .Register(*registry_);
    prometheus::Histogram &faiss_disk_load_IO_speed_histogram_ = disk_load_IO_speed_.Add({{"DB","Faiss"}},BucketBoundaries{0.1, 1.0, 10.0});

    ////all from CacheMgr.cpp
    //record cache access count
    prometheus::Family<prometheus::Counter> &cache_access_ = prometheus::BuildCounter()
        .Name("cache_access_total")
        .Help("the count of accessing cache ")
        .Register(*registry_);
    prometheus::Counter &cache_access_total_ = cache_access_.Add({});

};

static MetricsBase& CreateMetricsCollector(MetricCollectorType collector_type) {
    switch(collector_type) {
        case MetricCollectorType::PROMETHEUS:
            static PrometheusMetrics instance = PrometheusMetrics::GetInstance();
            return instance;
        default:
            return MetricsBase::GetInstance();
    }
}

static MetricsBase& GetInstance(){
    ConfigNode& config = ServerConfig::GetInstance().GetConfig(CONFIG_METRIC);
    std::string collector_typr_str = config.GetValue(CONFIG_METRIC_COLLECTOR);
    if(collector_typr_str == "prometheus") {
        return CreateMetricsCollector(MetricCollectorType::PROMETHEUS);
    } else if(collector_typr_str == "zabbix"){
        return CreateMetricsCollector(MetricCollectorType::ZABBIX);
    } else {
        return CreateMetricsCollector(MetricCollectorType::INVALID);
    }
}


}
}
}



