
GO := CGO_ENABLED=0 go

# 定义所需的最低Golang版本号
MIN_GO_VERSION := 1.21.4

# 获取当前安装的Golang版本信息
CURRENT_GO_VERSION := $(shell go version)

# 提取当前Golang主版本号（去除"go"开头）
CURRENT_MAJOR_VERSION := $(subst go, ,$(word 3, $(CURRENT_GO_VERSION)))

# 将当前Golang主版本号转换为数字形式
CURRENT_MAJOR_NUMBER := $(subst .,,${CURRENT_MAJOR_VERSION})

# 将所需的最低Golang版本号转换为数字形式
MIN_MAJOR_NUMBER := $(subst .,,${MIN_GO_VERSION})

BUILD_DATE ?=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')	# Blank error: date '+%FT %T %z':"buildDate":"2023-03-31T 20:05:43 +0800"


#检测go 是否安装
ifneq ($(shell go version 1>&2 >/dev/null && echo 0|| echo 1),0)
    $(error go version 执行失败)
endif

# 判断当前Golang版本是否符合要求
ifneq ($(shell test $(CURRENT_MAJOR_NUMBER) -lt $(MIN_MAJOR_NUMBER) || echo ok),ok)
    $(error golang min version $(MIN_MAJOR_NUMBER) > $(CURRENT_MAJOR_NUMBER))
endif

GOOS ?= $(shell go version | awk '{match($$0, /([a-zA-Z0-9]+)\/([a-zA-Z0-9]+)/,arr); print arr[1]}')
GOARCH ?= $(shell go version | awk '{match($$0, /([a-zA-Z0-9]+)\/([a-zA-Z0-9]+)/,arr); print arr[2]}')
#后缀
ifeq ($(GOOS),windows)
	GO_OUT_EXT := .exe
endif

GO_LDFLAGS += -X $(VERSION_PACKAGE).Version=$(GIT_TAG)-$(GIT_COMMIT)$(GIT_STATE)\
	-X $(VERSION_PACKAGE).BuildAt=$(BUILD_DATE) \
	-s -w		# -s -w deletes debugging information and symbol tables
ifeq ($(DEBUG), 1)
	GO_BUILD_FLAGS += -gcflags "all=-N -l"
	GO_LDFLAGS=
endif


GO_BUILD_FLAGS += -ldflags "$(GO_LDFLAGS)"