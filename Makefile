# Copyright (C) 2019-2020 Zilliz. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under the License
# is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
# or implied. See the License for the specific language governing permissions and limitations under the License.

GO		?= go
PWD 	:= $(shell pwd)
GOPATH 	:= $(shell $(GO) env GOPATH)

INSTALL_PATH := $(PWD)/bin
LIBRARY_PATH := $(PWD)/lib

all: build-cpp build-go

get-build-deps:
	@(env bash $(PWD)/scripts/install_deps.sh)

getdeps:
	@mkdir -p ${GOPATH}/bin
	@which golangci-lint 1>/dev/null || (echo "Installing golangci-lint" && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.27.0)
	@which ruleguard 1>/dev/null || (echo "Installing ruleguard" && GO111MODULE=off go get github.com/quasilyte/go-ruleguard/...)

tools/bin/revive: tools/check/go.mod
	cd tools/check; \
	$(GO) build -o ../bin/revive github.com/mgechev/revive

cppcheck:
	@(env bash ${PWD}/scripts/core_build.sh -l)

generated-proto-go:export protoc:=${PWD}/cmake_build/thirdparty/protobuf/protobuf-build/protoc
generated-proto-go: build-cpp
	@mkdir -p ${GOPATH}/bin
	@which protoc-gen-go 1>/dev/null || (echo "Installing protoc-gen-go" && go get github.com/golang/protobuf/protoc-gen-go@v1.3.2)
	@(env bash $(PWD)/scripts/proto_gen_go.sh)

check-proto-product: generated-proto-go
	@(env bash $(PWD)/scripts/check_proto_product.sh)

fmt:
ifdef GO_DIFF_FILES
	@echo "Running $@ check"
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh $(GO_DIFF_FILES)
else
	@echo "Running $@ check"
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh cmd/
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh internal/
	@GO111MODULE=on env bash $(PWD)/scripts/gofmt.sh tests/go/
endif

lint:tools/bin/revive
	@echo "Running $@ check"
	@tools/bin/revive -formatter friendly -config tools/check/revive.toml ./...

get-rocksdb:
	@go env -w CGO_CFLAGS="-I$(PWD)/internal/kv/rocksdb/cwrapper/output/include"
	@go env -w CGO_LDFLAGS="-L$(PWD)/internal/kv/rocksdb/cwrapper/output/lib -l:librocksdb.a -lstdc++ -lm -lz"
	@(env bash $(PWD)/internal/kv/rocksdb/cwrapper/build.sh)
	@go get github.com/tecbot/gorocksdb

#TODO: Check code specifications by golangci-lint
static-check:get-rocksdb
	@echo "Running $@ check"
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint cache clean
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint run --timeout=30m --config ./.golangci.yml ./internal/...
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint run --timeout=30m --config ./.golangci.yml ./cmd/...
	@GO111MODULE=on ${GOPATH}/bin/golangci-lint run --timeout=30m --config ./.golangci.yml ./tests/go/...

ruleguard:
ifdef GO_DIFF_FILES
	@echo "Running $@ check"
	@${GOPATH}/bin/ruleguard -rules ruleguard.rules.go $(GO_DIFF_FILES)
else
	@echo "Running $@ check"
	@${GOPATH}/bin/ruleguard -rules ruleguard.rules.go ./internal/...
	@${GOPATH}/bin/ruleguard -rules ruleguard.rules.go ./cmd/...
	@${GOPATH}/bin/ruleguard -rules ruleguard.rules.go ./tests/go/...
endif

verifiers: getdeps cppcheck fmt static-check

# Builds various components locally.
build-go: build-cpp get-rocksdb
	@echo "Building each component's binary to './bin'"
	@echo "Building master ..."
	@mkdir -p $(INSTALL_PATH) && go env -w CGO_ENABLED="0" && GO111MODULE=on $(GO) build -o $(INSTALL_PATH)/master $(PWD)/cmd/master/main.go 1>/dev/null
	@echo "Building proxy ..."
	@mkdir -p $(INSTALL_PATH) && go env -w CGO_ENABLED="0" && GO111MODULE=on $(GO) build -o $(INSTALL_PATH)/proxy $(PWD)/cmd/proxy/proxy.go 1>/dev/null
	@echo "Building query node ..."
	@mkdir -p $(INSTALL_PATH) && go env -w CGO_ENABLED="1" && GO111MODULE=on $(GO) build -o $(INSTALL_PATH)/querynode $(PWD)/cmd/querynode/query_node.go 1>/dev/null
	@echo "Building indexbuilder ..."
	@mkdir -p $(INSTALL_PATH) && go env -w CGO_ENABLED="1" && GO111MODULE=on $(GO) build -o $(INSTALL_PATH)/indexbuilder $(PWD)/cmd/indexbuilder/indexbuilder.go 1>/dev/null
	@echo "Building write node ..."
	@mkdir -p $(INSTALL_PATH) && go env -w CGO_ENABLED="1" && GO111MODULE=on $(GO) build -o $(INSTALL_PATH)/writenode $(PWD)/cmd/writenode/writenode.go 1>/dev/null
	@echo "Building binlog ..."
	@mkdir -p $(INSTALL_PATH) && go env -w CGO_ENABLED="1" && GO111MODULE=on $(GO) build -o $(INSTALL_PATH)/binlog $(PWD)/cmd/binlog/main.go 1>/dev/null

build-cpp:
	@(env bash $(PWD)/scripts/core_build.sh -f "$(CUSTOM_THIRDPARTY_PATH)")
	@(env bash $(PWD)/scripts/cwrapper_build.sh -t Release -f "$(CUSTOM_THIRDPARTY_PATH)")

build-cpp-with-unittest:
	@(env bash $(PWD)/scripts/core_build.sh -u -f "$(CUSTOM_THIRDPARTY_PATH)")
	@(env bash $(PWD)/scripts/cwrapper_build.sh -t Release -f "$(CUSTOM_THIRDPARTY_PATH)")

# Runs the tests.
unittest: test-cpp test-go

#TODO: proxy master query node writer's unittest
test-go:get-rocksdb
	@echo "Running go unittests..."
	@(env bash $(PWD)/scripts/run_go_unittest.sh)

test-cpp: build-cpp-with-unittest
	@echo "Running cpp unittests..."
	@(env bash $(PWD)/scripts/run_cpp_unittest.sh)

#TODO: build each component to docker
docker: verifiers
	@echo "Building query node docker image '$(TAG)'"
	@echo "Building proxy docker image '$(TAG)'"
	@echo "Building master docker image '$(TAG)'"

# Builds each component and installs it to $GOPATH/bin.
install: all
	@echo "Installing binary to './bin'"
	@mkdir -p $(GOPATH)/bin && cp -f $(PWD)/bin/querynode $(GOPATH)/bin/querynode
	@mkdir -p $(GOPATH)/bin && cp -f $(PWD)/bin/master $(GOPATH)/bin/master
	@mkdir -p $(GOPATH)/bin && cp -f $(PWD)/bin/proxy $(GOPATH)/bin/proxy
	@mkdir -p $(GOPATH)/bin && cp -f $(PWD)/bin/writenode $(GOPATH)/bin/writenode
	@mkdir -p $(GOPATH)/bin && cp -f $(PWD)/bin/indexbuilder $(GOPATH)/bin/indexbuilder
	@mkdir -p $(LIBRARY_PATH) && cp -f $(PWD)/internal/core/output/lib/* $(LIBRARY_PATH)
	@echo "Installation successful."

clean:
	@echo "Cleaning up all the generated files"
	@find . -name '*.test' | xargs rm -fv
	@find . -name '*~' | xargs rm -fv
	@rm -rf bin/
	@rm -rf lib/
	@rm -rf $(GOPATH)/bin/master
	@rm -rf $(GOPATH)/bin/proxy
	@rm -rf $(GOPATH)/bin/querynode
	@rm -rf $(GOPATH)/bin/writenode
	@rm -rf $(GOPATH)/bin/indexbuilder
