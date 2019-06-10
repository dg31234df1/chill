#include "LicenseCheck.h"
#include <iostream>
#include <thread>

#include <boost/archive/binary_oarchive.hpp>
#include <boost/archive/binary_iarchive.hpp>
//#include <boost/foreach.hpp>
//#include <boost/serialization/vector.hpp>
#include <boost/filesystem/path.hpp>
#include <boost/serialization/map.hpp>
#include <boost/filesystem/operations.hpp>


namespace zilliz {
namespace vecwise {
namespace server {

using IO_SERVICE = boost::asio::io_service;

namespace {
IO_SERVICE& GetIOService() {
    static IO_SERVICE io;
    return io;
}


}

// Part 1:  Legality check

ServerError
LicenseCheck::LegalityCheck(const std::string &license_file_path) {

    int device_count;
    LicenseLibrary::GetDeviceCount(device_count);
    std::vector<std::string> uuid_array;
    LicenseLibrary::GetUUID(device_count, uuid_array);

    std::vector<std::string> sha_array;
    LicenseLibrary::GetUUIDSHA256(device_count, uuid_array, sha_array);

    int output_device_count;
    std::map<int, std::string> uuid_encryption_map;
    time_t starting_time;
    time_t end_time;
    ServerError err = LicenseLibrary::LicenseFileDeserialization(license_file_path,
                                               output_device_count,
                                               uuid_encryption_map,
                                               starting_time,
                                               end_time);
    if(err !=SERVER_SUCCESS)
    {
        printf("License check error: 01\n");
        return SERVER_UNEXPECTED_ERROR;
    }
    time_t system_time;
    LicenseLibrary::GetSystemTime(system_time);

    if (device_count != output_device_count) {
        printf("License check error: 02\n");
        return SERVER_UNEXPECTED_ERROR;
    }
    for (int i = 0; i < device_count; ++i) {
        if (sha_array[i] != uuid_encryption_map[i]) {
            printf("License check error: 03\n");
            return SERVER_UNEXPECTED_ERROR;
        }
    }
    if (system_time < starting_time || system_time > end_time) {
        printf("License check error: 04\n");
        return SERVER_UNEXPECTED_ERROR;
    }
    printf("Legality Check Success\n");
    return SERVER_SUCCESS;
}

// Part 2: Timing check license

ServerError
LicenseCheck::AlterFile(const std::string &license_file_path,
                        const boost::system::error_code &ec,
                        boost::asio::deadline_timer *pt) {

    ServerError err = LegalityCheck(license_file_path);
    if(err!=SERVER_SUCCESS)
    {
        exit(1);
    }
    printf("---runing---\n");
    pt->expires_at(pt->expires_at() + boost::posix_time::hours(1));
    pt->async_wait(boost::bind(AlterFile, license_file_path, boost::asio::placeholders::error, pt));
    return SERVER_SUCCESS;

}

ServerError
LicenseCheck::StartCountingDown(const std::string &license_file_path) {

    if (!LicenseLibrary::IsFileExistent(license_file_path)) {
        printf("license file not exist\n");
        exit(1);
    }

    //create a thread to run AlterFile
    std::thread io_thread([&]() {
        boost::asio::io_service& io = GetIOService();
        boost::asio::deadline_timer t(io, boost::posix_time::hours(1));
        t.async_wait(boost::bind(AlterFile, license_file_path, boost::asio::placeholders::error, &t));
        io.run();//this thread will block here
    });
    io_thread.detach();

    return SERVER_SUCCESS;
}

ServerError
LicenseCheck::StopCountingDown() {
    boost::asio::io_service& io = GetIOService();
    if(!io.stopped()) {
        io.stop();
    }

    return SERVER_SUCCESS;
}

}
}
}