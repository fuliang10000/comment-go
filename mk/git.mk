
GIT_COMMIT ?= $(shell git rev-parse --short HEAD || echo "0.0.0")

GIT_TAG := $(shell git describe --exact-match --tags --abbrev=0  2> /dev/null || echo untagged)

#当前build状态
GIT_STATE:=-dirty
ifeq (, $(shell git status --porcelain 2>/dev/null))
GIT_STATE=
endif