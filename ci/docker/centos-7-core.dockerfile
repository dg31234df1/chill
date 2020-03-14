ARG arch=amd64
FROM ${arch}/centos:7

# pipefail is enabled for proper error detection in the `wget`
# step
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN yum install -y epel-release centos-release-scl-rh && yum install -y wget curl which && \
    wget -qO- "https://cmake.org/files/v3.14/cmake-3.14.3-Linux-x86_64.tar.gz" | tar --strip-components=1 -xz -C /usr/local && \
    yum install -y ccache make automake git python3-pip libcurl-devel python3-devel boost-static mysql-devel \
    devtoolset-7-gcc devtoolset-7-gcc-c++ devtoolset-7-gcc-gfortran llvm-toolset-7.0-clang llvm-toolset-7.0-clang-tools-extra \
    mysql lcov openblas-devel lapack-devel \
    && \
    rm -rf /var/cache/yum/*

RUN echo "source scl_source enable devtoolset-7" >> /etc/profile.d/devtoolset-7.sh
RUN echo "source scl_source enable llvm-toolset-7.0" >> /etc/profile.d/llvm-toolset-7.sh

ENV CLANG_TOOLS_PATH="/opt/rh/llvm-toolset-7.0/root/usr/bin"

# use login shell to activate environment un the RUN commands
SHELL [ "/bin/bash", "-c", "-l" ]

# use login shell when running the container
ENTRYPOINT [ "/bin/bash", "-c", "-l" ]
