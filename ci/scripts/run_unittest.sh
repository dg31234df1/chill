#!/usr/bin/env bash

set -e

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
SCRIPTS_DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"

MILVUS_CORE_DIR="${SCRIPTS_DIR}/../../core"
MILVUS_PROXY_DIR="${SCRIPTS_DIR}/../../proxy"
CORE_INSTALL_PREFIX="${MILVUS_CORE_DIR}/milvus"
PROXY_INSTALL_PREFIX="${MILVUS_PROXY_DIR}/milvus"
UNITTEST_DIRS=("${CORE_INSTALL_PREFIX}/unittest" "${PROXY_INSTALL_PREFIX}/unittest")

# Currently core will install target lib to "core/lib"
if [ -d "${MILVUS_CORE_DIR}/lib" ]; then
	export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:${MILVUS_CORE_DIR}/lib
fi

# run unittest
for UNITTEST_DIR in "${UNITTEST_DIRS[@]}"; do
  if [ ! -d "${UNITTEST_DIR}" ]; then
	echo "The unittest folder does not exist!"
    exit 1
  fi
  for test in `ls ${UNITTEST_DIR}`; do
    echo $test " running..."
    # run unittest
#    ${UNITTEST_DIR}/${test}
    if [ $? -ne 0 ]; then
        echo ${UNITTEST_DIR}/${test} "run failed"
        exit 1
    fi
  done
done

# ignore Minio,S3 unittes
MILVUS_DIR="${SCRIPTS_DIR}/../../"
echo $MILVUS_DIR
go test "${MILVUS_DIR}/storage/internal/tikv/..." "${MILVUS_DIR}/reader/..." "${MILVUS_DIR}/writer/..." "${MILVUS_DIR}/pkg/master/..." -failfast
