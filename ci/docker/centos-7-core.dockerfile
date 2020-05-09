ARG arch=amd64
FROM ${arch}/centos:7

# pipefail is enabled for proper error detection in the `wget`
# step
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN yum install -y epel-release centos-release-scl-rh && yum install -y wget curl which && \
    wget -qO- "https://cmake.org/files/v3.14/cmake-3.14.3-Linux-x86_64.tar.gz" | tar --strip-components=1 -xz -C /usr/local && \
    yum install -y make automake git python3-pip libcurl-devel python3-devel boost-static mysql-devel \
    devtoolset-7-gcc devtoolset-7-gcc-c++ devtoolset-7-gcc-gfortran llvm-toolset-7.0-clang llvm-toolset-7.0-clang-tools-extra \
    mysql lcov && \
    rm -rf /var/cache/yum/* && \
    echo "source scl_source enable devtoolset-7" >> /etc/profile.d/devtoolset-7.sh && \
    echo "source scl_source enable llvm-toolset-7.0" >> /etc/profile.d/llvm-toolset-7.sh

ENV CLANG_TOOLS_PATH="/opt/rh/llvm-toolset-7.0/root/usr/bin"

RUN source /etc/profile.d/devtoolset-7.sh && \
    wget https://github.com/xianyi/OpenBLAS/archive/v0.3.9.tar.gz && \
    tar zxvf v0.3.9.tar.gz && cd OpenBLAS-0.3.9 && \
    make TARGET=CORE2 DYNAMIC_ARCH=1 DYNAMIC_OLDER=1 USE_THREAD=0 USE_OPENMP=0 FC=gfortran CC=gcc COMMON_OPT="-O3 -g -fPIC" FCOMMON_OPT="-O3 -g -fPIC -frecursive" NMAX="NUM_THREADS=128" LIBPREFIX="libopenblas" LAPACKE="NO_LAPACKE=1" INTERFACE64=0 NO_STATIC=1 && \
    make PREFIX=/usr install && \
    cd .. && rm -rf OpenBLAS-0.3.9 && rm v0.3.9.tar.gz

RUN yum install -y ccache && \
    rm -rf /var/cache/yum/*

# use login shell to activate environment un the RUN commands
SHELL [ "/bin/bash", "-c", "-l" ]

# use login shell when running the container
ENTRYPOINT [ "/bin/bash", "-c", "-l" ]
