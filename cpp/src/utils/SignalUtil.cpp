////////////////////////////////////////////////////////////////////////////////
// Copyright 上海赜睿信息科技有限公司(Zilliz) - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited.
// Proprietary and confidential.
////////////////////////////////////////////////////////////////////////////////
#include "SignalUtil.h"
#include "src/server/Server.h"
#include "utils/Log.h"

#include <signal.h>
#include <execinfo.h>

namespace zilliz {
namespace milvus {
namespace server {

void SignalUtil::HandleSignal(int signum){

    switch(signum){
        case SIGINT:
        case SIGUSR2:{
            server::Server* server_ptr = server::Server::Instance();
            server_ptr->Stop();

            exit(0);
        }
        default:{
            SERVER_LOG_INFO << "Server received signal:" << std::to_string(signum);
            SignalUtil::PrintStacktrace();

            std::string info = "Server encounter critical signal:";
            info += std::to_string(signum);
//            SendSignalMessage(signum, info);

            SERVER_LOG_INFO << info;

            server::Server* server_ptr = server::Server::Instance();
            server_ptr->Stop();

            exit(1);
        }
    }
}

void SignalUtil::PrintStacktrace() {
    SERVER_LOG_INFO << "Call stack:";

    const int size = 32;
    void* array[size];
    int stack_num = backtrace(array, size);
    char ** stacktrace = backtrace_symbols(array, stack_num);
    for (int i = 0; i < stack_num; ++i) {
        std::string info = stacktrace[i];
        SERVER_LOG_INFO << info;
    }
    free(stacktrace);
}


}
}
}
