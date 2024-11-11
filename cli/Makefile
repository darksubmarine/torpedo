.PHONY: all build tidy release installer clean

all: build


PACKAGE=github.com/darksubmarine/torpedo/cli/src
VERSION := $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP := $(shell date '+%Y%m%dT%H%M%S')

LDFLAGS := "-X '${PACKAGE}.Version=${VERSION}' -X '${PACKAGE}.CommitHash=${COMMIT_HASH}' -X '${PACKAGE}.BuildTimestamp=${BUILD_TIMESTAMP}'"

BINARIES = torpedo_linux_amd64 torpedo_linux_arm64 torpedo_osx_amd64 torpedo_osx_m1

OS = linux darwin
ARCH = amd64 arm64

build:
	go build -ldflags ${LDFLAGS} -o bin/torpedo
	for os in $(OS) ; do \
		for arch in $(ARCH) ; do \
			GOOS=$$os GOARCH=$$arch go build -ldflags ${LDFLAGS} -o bin/torpedo\_$$os\_$$arch\_${VERSION} ; \
			tar cJf torpedo\_$$os\_$$arch\_${VERSION}.tar.xz bin/torpedo\_$$os\_$$arch\_${VERSION} ; \
		done \
	done

installer:
	for os in $(OS) ; do \
        for arch in $(ARCH) ; do \
                cat installer.template.sh | \
                					sed -e "s/AUTO_REPLACE_BUILD_VERSION/${VERSION}/" | \
                					sed -e "s/AUTO_REPLACE_COMMIT_VERSION/${COMMIT_HASH}/" | \
                					sed -e "s/AUTO_REPLACE_BINARY_FILENAME/torpedo\_$$os\_$$arch\_${VERSION}/" > torpedo-$$os-$$arch-${VERSION}-installer.sh ; \
                cat torpedo\_$$os\_$$arch\_${VERSION}.tar.xz >> torpedo-$$os-$$arch-${VERSION}-installer.sh ; \
                chmod +x torpedo-$$os-$$arch-${VERSION}-installer.sh ; \
            done \
    done

clean:
	mkdir -p release
	for os in $(OS) ; do \
		for arch in $(ARCH) ; do \
			rm torpedo\_$$os\_$$arch\_${VERSION}.tar.xz ; \
			cp torpedo-$$os-$$arch-${VERSION}-installer.sh release/torpedo-$$os-$$arch-latest-installer.bin ; \
			mv torpedo-$$os-$$arch-${VERSION}-installer.sh release/torpedo-$$os-$$arch-${VERSION}-installer.bin ; \
		done \
	done

release: build installer clean

tidy:
	go mod tidy
