GOBUILD:=go build -trimpath -ldflags="-s -w" -a -v -o
BIN:=bin
NANE:=dnsec

build:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(BIN)/$(NANE)_darwin_arm64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(BIN)/$(NANE)_darwin_amd64
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(BIN)/$(NANE)_linux_amd64
	GOOS=linux GOARCH=mips $(GOBUILD) $(BIN)/$(NANE)_linux_mips
	GOOS=linux GOARCH=mips64le $(GOBUILD) $(BIN)/$(NANE)_linux_mips64le
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(BIN)/$(NANE)_windows_amd64.exe