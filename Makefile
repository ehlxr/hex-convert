.PHONY: default init build install release dep get_deps clean build_amd64 build_386 upx

# https://golang.org/doc/install/source#environment
GOOS := $(shell go2 env | awk -F= '$$1=="GOOS" {print $$2}' | awk -F '"' '{print $$2}') # 此处 awk 需使用两个 $
GOARCH := $(shell go2 env | awk -F= '$$1=="GOARCH" {print $$2}' | awk -F '"' '{print $$2}')
OSS = darwin dragonfly freebsd linux netbsd openbsd plan9 solaris windows
PKG =
# ifeq ($(strip $(GOOS)), windows)
# 	GOARCH := $(strip $(GOARCH)).exe
# endif

default:
	@echo "hc info: please choose a target for 'make'"
	@echo "available target: build install release clean build_amd64 build_386 upx"

build:
	@ go2 build -ldflags "-s -w" -o dist/hc_$(strip $(GOOS))_$(strip $(if \
    $(findstring windows,$(GOOS)),\
    $(strip $(GOARCH)).exe,\
    $(strip $(GOARCH))\
	))

install:
	go2 install -ldflags "-s -w"

release: build_amd64 build_386 upx

clean:
	go2 clean -i
	rm -rf dist/hc* hc* hex-convert

build_amd64:
	@ $(foreach OS,\
	$(OSS),\
	$(shell CGO_ENABLED=0 GOOS=$(OS) GOARCH=amd64 go2 build -ldflags "-s -w" -o dist/hc_$(OS)_amd64$(if $(findstring windows,$(OS)),.exe)))
	@ echo done

build_386:
	@ $(foreach OS,\
	$(OSS),\
	$(shell CGO_ENABLED=0 GOOS=$(OS) GOARCH=386 go2 build -ldflags "-s -w" -o dist/hc_$(OS)_386$(if $(findstring windows,$(OS)),.exe)))
	@ echo done

# 压缩。需要安装 https://github.com/upx/upx
upx:
	upx $(if $(PKG),$(PKG),dist/hc*)