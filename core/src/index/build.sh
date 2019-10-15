#!/bin/bash

BUILD_TYPE="Debug"
BUILD_UNITTEST="OFF"
INSTALL_PREFIX=$(pwd)/cmake_build
MAKE_CLEAN="OFF"
PROFILING="OFF"
BUILD_FAISS_WITH_MKL="OFF"
USE_JFROG_CACHE="OFF"

while getopts "p:d:t:uhrcgmj" arg
do
        case $arg in
             t)
                BUILD_TYPE=$OPTARG # BUILD_TYPE
                ;;
             u)
                echo "Build and run unittest cases" ;
                BUILD_UNITTEST="ON";
                ;;
             p)
                INSTALL_PREFIX=$OPTARG
                ;;
             r)
                if [[ -d cmake_build ]]; then
                    rm ./cmake_build -r
                    MAKE_CLEAN="ON"
                fi
                ;;
             g)
                PROFILING="ON"
                ;;
             m)
                BUILD_FAISS_WITH_MKL="ON"
                ;;
             j)
                USE_JFROG_CACHE="ON"
                ;;
             h) # help
                echo "

parameter:
-t: build type(default: Debug)
-u: building unit test options(default: OFF)
-p: install prefix(default: $(pwd)/knowhere)
-r: remove previous build directory(default: OFF)
-g: profiling(default: OFF)
-m: build faiss with MKL(default: OFF)

usage:
./build.sh -t \${BUILD_TYPE} [-u] [-h] [-g] [-r] [-c] [-m]
                "
                exit 0
                ;;
             ?)
                echo "unknown argument"
        exit 1
        ;;
        esac
done

if [[ ! -d cmake_build ]]; then
	mkdir cmake_build
	MAKE_CLEAN="ON"
fi

cd cmake_build

CUDA_COMPILER=/usr/local/cuda/bin/nvcc

if [[ ${MAKE_CLEAN} == "ON" ]]; then
    CMAKE_CMD="cmake -DBUILD_UNIT_TEST=${BUILD_UNITTEST} \
    -DCMAKE_INSTALL_PREFIX=${INSTALL_PREFIX}
    -DCMAKE_BUILD_TYPE=${BUILD_TYPE} \
    -DCMAKE_CUDA_COMPILER=${CUDA_COMPILER} \
    -DMILVUS_ENABLE_PROFILING=${PROFILING} \
    -DBUILD_FAISS_WITH_MKL=${BUILD_FAISS_WITH_MKL} \
    -DUSE_JFROG_CACHE=${USE_JFROG_CACHE} \
    ../"
    echo ${CMAKE_CMD}

    ${CMAKE_CMD}
    make clean
fi

make -j 8 || exit 1

make install || exit 1
