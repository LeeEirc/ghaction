GHRNAME=ghrcreate
BUILDDIR=build

BASEPATH := $(shell pwd)
GHRSRCFILE := $(BASEPATH)/cmd/ghrcreate/

GOBUILD=CGO_ENABLED=0 go build -trimpath

PLATFORM_LIST = \
	darwin-amd64 \
	linux-amd64 \
	linux-arm64

WINDOWS_ARCH_LIST = \
	windows-amd64

all-arch: $(PLATFORM_LIST) $(WINDOWS_ARCH_LIST)

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BUILDDIR)/$(GHRNAME)-$@ $(GHRSRCFILE)

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BUILDDIR)/$(GHRNAME)-$@ $(GHRSRCFILE)

linux-arm64:
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(BUILDDIR)/$(GHRNAME)-$@ $(GHRSRCFILE)

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BUILDDIR)/$(GHRNAME)-$@.exe $(GHRSRCFILE)

.PHONY: clean
clean:
	-rm -rf $(BUILDDIR)
