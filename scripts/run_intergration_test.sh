#!/usr/bin/env bash

# Licensed to the LF AI & Data foundation under one
# or more contributor license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership. The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# run integration test
echo "Running integration test under ./tests/integration"

BASEDIR=$(dirname "$0")
source $BASEDIR/setenv.sh

set -ex

# starting the timer
beginTime=`date +%s`
if [[ $(uname -s) == "Darwin" && "$(uname -m)" == "arm64" ]]; then
    APPLE_SILICON_FLAG="-tags dynamic"
fi

for d in $(go list ./tests/integration/...); do
    echo "$d"
    go test -race ${APPLE_SILICON_FLAG} -v "$d"
done

endTime=`date +%s`

echo "Total time for go integration test:" $(($endTime-$beginTime)) "s"

