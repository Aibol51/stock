# Custom configuration | 独立配置
# Service name | 项目名称
SERVICE=Core
# Service name in specific style | 项目经过style格式化的名称
SERVICE_STYLE=core
# Service name in lowercase | 项目名称全小写格式
SERVICE_LOWER=core
# Service name in snake format | 项目名称下划线格式
SERVICE_SNAKE=core
# Service name in snake format | 项目名称短杠格式
SERVICE_DASH=core

# The project version, if you don't use git, you should set it manually | 项目版本，如果不使用git请手动设置
VERSION=$(shell git describe --tags --always)

# The project file name style | 项目文件命名风格
PROJECT_STYLE=go_zero

# Whether to use i18n | 是否启用 i18n
PROJECT_I18N=true

# Swagger type, support yml,json | Swagger 文件类型，支持yml,json
SWAGGER_TYPE=json

# Ent enabled features | Ent 启用的官方特性
ENT_FEATURE=sql/execquery,intercept

# The arch of the build | 构建的架构
GOARCH=amd64

# ---- You may not need to modify the codes below | 下面的代码大概率不需要更改 ----

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
LDFLAGS := -s -w

.PHONY: gen-api-proto-logic
gen-api-proto-logic: # Generate logic code from Proto, need model and group params
	goctls api proto --proto=../rpc/core.proto --style=go_zero --api_service_name=core --rpc_service_name=Core --output=./ --model=$(model) --rpc_name=Core --grpc_package=github.com/suyuan32/simple-admin-core/rpc/types/core --i18n=true
	@echo "Generate logic codes from Ent successfully"