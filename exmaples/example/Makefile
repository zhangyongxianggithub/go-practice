HOMEDIR 	:=	$(shell pwd)
OUTPUT_DIR 	:=	$(HOMEDIR)/output
MOD_NAME	:=	rule-engine

export GOENV =  $(HOMEDIR)/go.env
GOROOT 		:= $(or $(GO_1_19_HOME),$(shell go env GOROOT))
GO        	:= $(GOROOT)/bin/go
GOPATH    	:= $(shell $(GO) env GOPATH)
GOINSTALL 	:= $(GO) install
GOGET     	:= $(GO) get
GOMOD     	:= $(GO) mod
GOBUILD   	:= $(GO) build
GOTEST    	:= $(GO) test -gcflags="-N -l"
GOPKGS    	:= $$($(GO) list ./...| grep -vE "vendor")
SWAG      	:= $(GOPATH)/bin/swag

# test cover files
COVPROF 	:= $(HOMEDIR)/covprof.out  # coverage profile
COVFUNC 	:= $(HOMEDIR)/covfunc.txt  # coverage profile information for each function
COVHTML 	:= $(HOMEDIR)/covhtml.html # HTML representation of coverage profile


# flags to inject into binaries
APP_VERSION := 0.0.0
GIT_BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT  := $(shell git rev-parse HEAD)
BUILD_DATE  := $(shell date "+%Y-%m-%d %H:%M:%S")
VERSION_VAR := icode.baidu.com/baidu/themis/server-go/version
GOMODULE    := $(shell $(GO) list -m)

# gin controller directories that swag scans for openapi comments
SWAG_SCAN_DIRS  := $(shell ls -d ../../controller/* | tr '\n' ',' | sed 's/,$$//')
SWAG_DOC_PREFIX := RuleEngine
SWAGGER_VAR     := icode.baidu.com/baidu/themis/server-go/swagger

# 这里有用吗
LD_FLAGS := " \
    -X main.CliName=${CLINAME} \
	-X '$(VERSION_VAR).AppVersion=${APP_VERSION}' \
    -X '$(VERSION_VAR).GitBranch=${GIT_BRANCH}' \
	-X '$(VERSION_VAR).GitCommit=${GIT_COMMIT}' \
	-X '$(VERSION_VAR).BuildDate=${BUILD_DATE}' \
    -X '$(VERSION_VAR).BuildPipeline=${AGILE_PIPELINE_NAME}' \
	-X '$(VERSION_VAR).BuildNumber=${AGILE_PIPELINE_BUILD_NUMBER}' \
	-X '$(SWAGGER_VAR).DocPrefix=${SWAG_DOC_PREFIX}' \
    "

default: clean prepare compile package

clean:
	$(GO) clean
	rm -rf $(OUTPUT_DIR)
	@echo "clean accomplished"

prepare: check-var git set-env install-doc doc
	@echo "prepare accomplished"

check-var:
	$(info GOROOT: $(GOROOT))
	$(info GOPATH: $(GOPATH))
	$(info GOBIN: $(GOBIN))
	$(info SWAG: $(SWAG))

git:
	git config --global http.sslVerify false

set-env:
	$(GO) env -w GO111MODULE=on
	$(GO) env -w GONOPROXY=\*\*.baidu.com\*\*
	$(GO) env -w GOPROXY=https://goproxy.baidu-int.com
	$(GO) env -w GONOSUMDB=\*
install-doc:
	$(GOINSTALL) github.com/swaggo/swag/cmd/swag@v1.16.1

compile:
	$(GOMOD) tidy
	@echo "synchronized dependencies accomplished"

package:
	mkdir -p $(OUTPUT_DIR)
# cp -rf ./* $(OUTPUT_DIR) cannot copy a directory, './output', into itself

doc:
	$(SWAG) init -g server.go -d $(HOMEDIR)/cmd/server,$(SWAG_SCAN_DIRS) --instanceName=${SWAG_DOC_PREFIX} --parseDependency --parseInternal --generatedTime -o $(HOMEDIR)/doc
	@echo "generate swagger doc accomplished"

doc-fmt:
	@$(SWAG) fmt
	@echo "swagger format accomplished"

.PHONY: default clean prepare check-var git set-env install-doc compile package doc doc-fmt