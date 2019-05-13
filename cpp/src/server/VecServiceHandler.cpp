/*******************************************************************************
 * Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 ******************************************************************************/
#include "VecServiceHandler.h"
#include "VecServiceTask.h"
#include "ServerConfig.h"
#include "VecIdMapper.h"
#include "utils/Log.h"
#include "utils/CommonUtil.h"
#include "utils/TimeRecorder.h"

#include "db/DB.h"
#include "db/Env.h"

namespace zilliz {
namespace vecwise {
namespace server {

namespace {
    class TimeRecordWrapper {
    public:
        TimeRecordWrapper(const std::string& func_name)
        : recorder_(func_name), func_name_(func_name) {
            SERVER_LOG_TRACE << func_name << " called";
        }

        ~TimeRecordWrapper() {
            recorder_.Elapse("cost");
            SERVER_LOG_TRACE << func_name_ << " finished";
        }

    private:
        TimeRecorder recorder_;
        std::string func_name_;
    };
    void TimeRecord(const std::string& func_name) {

    }

    const std::map<ServerError, zilliz::VecErrCode::type>& ErrorMap() {
        static const std::map<ServerError, zilliz::VecErrCode::type> code_map = {
            {SERVER_UNEXPECTED_ERROR, zilliz::VecErrCode::ILLEGAL_ARGUMENT},
            {SERVER_NULL_POINTER, zilliz::VecErrCode::ILLEGAL_ARGUMENT},
            {SERVER_INVALID_ARGUMENT, zilliz::VecErrCode::ILLEGAL_ARGUMENT},
            {SERVER_FILE_NOT_FOUND, zilliz::VecErrCode::ILLEGAL_ARGUMENT},
            {SERVER_NOT_IMPLEMENT, zilliz::VecErrCode::ILLEGAL_ARGUMENT},
            {SERVER_BLOCKING_QUEUE_EMPTY, zilliz::VecErrCode::ILLEGAL_ARGUMENT},
            {SERVER_GROUP_NOT_EXIST, zilliz::VecErrCode::GROUP_NOT_EXISTS},
            {SERVER_INVALID_TIME_RANGE, zilliz::VecErrCode::ILLEGAL_TIME_RANGE},
            {SERVER_INVALID_VECTOR_DIMENSION, zilliz::VecErrCode::ILLEGAL_VECTOR_DIMENSION},
        };

        return code_map;
    }

    const std::map<ServerError, std::string>& ErrorMessage() {
        static const std::map<ServerError, std::string> msg_map = {
            {SERVER_UNEXPECTED_ERROR, "unexpected error occurs"},
            {SERVER_NULL_POINTER, "null pointer error"},
            {SERVER_INVALID_ARGUMENT, "invalid argument"},
            {SERVER_FILE_NOT_FOUND, "file not found"},
            {SERVER_NOT_IMPLEMENT, "not implemented"},
            {SERVER_BLOCKING_QUEUE_EMPTY, "queue empty"},
            {SERVER_GROUP_NOT_EXIST, "group not exist"},
            {SERVER_INVALID_TIME_RANGE, "invalid time range"},
            {SERVER_INVALID_VECTOR_DIMENSION, "invalid vector dimension"},
        };

        return msg_map;
    }

    void ExecTask(BaseTaskPtr& task_ptr) {
        if(task_ptr == nullptr) {
            return;
        }

        VecServiceScheduler& scheduler = VecServiceScheduler::GetInstance();
        scheduler.ExecuteTask(task_ptr);

        if(!task_ptr->IsAsync()) {
            task_ptr->WaitToFinish();
            ServerError err = task_ptr->ErrorCode();
            if (err != SERVER_SUCCESS) {
                zilliz::VecException ex;
                ex.__set_code(ErrorMap().at(err));
                std::string msg = task_ptr->ErrorMsg();
                if(msg.empty()){
                    msg = ErrorMessage().at(err);
                }
                ex.__set_reason(msg);
                throw ex;
            }
        }
    }
}

void
VecServiceHandler::add_group(const VecGroup &group) {
    TimeRecordWrapper rc("add_group()");
    SERVER_LOG_TRACE << "group.id = " << group.id << ", group.dimension = " << group.dimension
                        << ", group.index_type = " << group.index_type;

    BaseTaskPtr task_ptr = AddGroupTask::Create(group.dimension, group.id);
    ExecTask(task_ptr);
}

void
VecServiceHandler::get_group(VecGroup &_return, const std::string &group_id) {
    TimeRecordWrapper rc("get_group()");
    SERVER_LOG_TRACE << "group_id = " << group_id;

    _return.id = group_id;
    BaseTaskPtr task_ptr = GetGroupTask::Create(group_id, _return.dimension);
    ExecTask(task_ptr);
}

void
VecServiceHandler::del_group(const std::string &group_id) {
    TimeRecordWrapper rc("del_group()");
    SERVER_LOG_TRACE << "group_id = " << group_id;

    BaseTaskPtr task_ptr = DeleteGroupTask::Create(group_id);
    ExecTask(task_ptr);
}


void
VecServiceHandler::add_vector(std::string& _return, const std::string &group_id, const VecTensor &tensor) {
    TimeRecordWrapper rc("add_vector()");
    SERVER_LOG_TRACE << "group_id = " << group_id << ", vector size = " << tensor.tensor.size();

    BaseTaskPtr task_ptr = AddVectorTask::Create(group_id, &tensor, _return);
    ExecTask(task_ptr);
}

void
VecServiceHandler::add_vector_batch(std::vector<std::string> & _return,
                                    const std::string &group_id,
                                    const VecTensorList &tensor_list) {
    TimeRecordWrapper rc("add_vector_batch()");
    SERVER_LOG_TRACE << "group_id = " << group_id << ", vector list size = "
                     << tensor_list.tensor_list.size();

    BaseTaskPtr task_ptr = AddBatchVectorTask::Create(group_id, &tensor_list, _return);
    ExecTask(task_ptr);
}

void
VecServiceHandler::add_binary_vector(std::string& _return,
                                     const std::string& group_id,
                                     const VecBinaryTensor& tensor) {
    TimeRecordWrapper rc("add_binary_vector()");
    SERVER_LOG_TRACE << "group_id = " << group_id << ", vector size = " << tensor.tensor.size()/4;

    BaseTaskPtr task_ptr = AddVectorTask::Create(group_id, &tensor, _return);
    ExecTask(task_ptr);
}

void
VecServiceHandler::add_binary_vector_batch(std::vector<std::string> & _return,
                                           const std::string& group_id,
                                           const VecBinaryTensorList& tensor_list) {
    TimeRecordWrapper rc("add_binary_vector_batch()");
    SERVER_LOG_TRACE << "group_id = " << group_id << ", vector list size = "
                     << tensor_list.tensor_list.size();

    BaseTaskPtr task_ptr = AddBatchVectorTask::Create(group_id, &tensor_list, _return);
    ExecTask(task_ptr);
}

void
VecServiceHandler::search_vector(VecSearchResult &_return,
                                 const std::string &group_id,
                                 const int64_t top_k,
                                 const VecTensor &tensor,
                                 const VecSearchFilter& filter) {
    TimeRecordWrapper rc("search_vector()");
    SERVER_LOG_TRACE << "group_id = " << group_id << ", top_k = " << top_k
                        << ", vector dimension = " << tensor.tensor.size();

    VecTensorList tensor_list;
    tensor_list.tensor_list.push_back(tensor);
    VecSearchResultList result;
    BaseTaskPtr task_ptr = SearchVectorTask::Create(group_id, top_k, &tensor_list, filter, result);
    ExecTask(task_ptr);

    if(!result.result_list.empty()) {
        _return = result.result_list[0];
    } else {
        SERVER_LOG_ERROR << "No search result returned";
    }
}

void
VecServiceHandler::search_vector_batch(VecSearchResultList &_return,
                                       const std::string &group_id,
                                       const int64_t top_k,
                                       const VecTensorList &tensor_list,
                                       const VecSearchFilter& filter) {
    TimeRecordWrapper rc("search_vector_batch()");
    SERVER_LOG_TRACE << "group_id = " << group_id << ", top_k = " << top_k
                     << ", vector list size = " << tensor_list.tensor_list.size();

    BaseTaskPtr task_ptr = SearchVectorTask::Create(group_id, top_k, &tensor_list, filter, _return);
    ExecTask(task_ptr);
}

void
VecServiceHandler::search_binary_vector(VecSearchResult& _return,
                                        const std::string& group_id,
                                        const int64_t top_k,
                                        const VecBinaryTensor& tensor,
                                        const VecSearchFilter& filter) {
    TimeRecordWrapper rc("search_binary_vector()");
    SERVER_LOG_TRACE << "group_id = " << group_id << ", top_k = " << top_k
                     << ", vector dimension = " << tensor.tensor.size();

    VecBinaryTensorList tensor_list;
    tensor_list.tensor_list.push_back(tensor);
    VecSearchResultList result;
    BaseTaskPtr task_ptr = SearchVectorTask::Create(group_id, top_k, &tensor_list, filter, result);
    ExecTask(task_ptr);

    if(!result.result_list.empty()) {
        _return = result.result_list[0];
    } else {
        SERVER_LOG_ERROR << "No search result returned";
    }
}

void
VecServiceHandler::search_binary_vector_batch(VecSearchResultList& _return,
                                              const std::string& group_id,
                                              const int64_t top_k,
                                              const VecBinaryTensorList& tensor_list,
                                              const VecSearchFilter& filter) {
    TimeRecordWrapper rc("search_binary_vector_batch()");
    SERVER_LOG_TRACE << "group_id = " << group_id << ", top_k = " << top_k
                     << ", vector list size = " << tensor_list.tensor_list.size();

    BaseTaskPtr task_ptr = SearchVectorTask::Create(group_id, top_k, &tensor_list, filter, _return);
    ExecTask(task_ptr);
}


}
}
}
