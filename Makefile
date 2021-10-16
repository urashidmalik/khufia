.ONESHELL: # Applies to every targets in the file!

BINARY_NAME=khufia
FREEBSD=freebsd
WINDOWS=windows
LINUX=linux
DARWIN=darwin


ARCH_AMD64=amd64
ARCH_ARM64=arm64


compile-windows-amd64:
	cd src && GOOS=${WINDOWS} GOARCH=${ARCH_AMD64} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${WINDOWS}-${ARCH_AMD64} main.go
compile-windows-arm64:
	cd src && GOOS=${WINDOWS} GOARCH=${ARCH_ARM64} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${WINDOWS}-${ARCH_ARM64} main.go

compile-linux-amd64:
	cd src && GOOS=${LINUX} GOARCH=${ARCH_AMD64} go build -o ../bin/${BINARY_NAME}-${LINUX}-${ARCH_AMD64} main.go
compile-linux-arm64:
	cd src && GOOS=${LINUX} GOARCH=${ARCH_ARM64} go build -o ../bin/${BINARY_NAME}-${LINUX}-${ARCH_ARM64} main.go
	
compile-freebsd-amd64:	
	cd src && GOOS=${FREEBSD} GOARCH=${ARCH_AMD64}  go build -o ../bin/${BINARY_NAME}-${FREEBSD}-${ARCH_AMD64} main.go
compile-freebsd-arm64:
	cd src && GOOS=${FREEBSD} GOARCH=${ARCH_ARM64}  go build -o ../bin/${BINARY_NAME}-${FREEBSD}-${ARCH_ARM64} main.go

compile-mac-intel:
	cd src && GOOS=${DARWIN} GOARCH=${ARCH_AMD64} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${DARWIN}-${ARCH_AMD64} main.go

compile-mac-apple-silicon:
	cd src && GOOS=${DARWIN} GOARCH=${ARCH_ARM64} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${DARWIN}-${ARCH_ARM64} main.go


build: clean compile-windows-amd64 compile-windows-arm64 compile-freebsd-amd64 compile-freebsd-arm64 compile-linux-amd64 compile-linux-arm64 compile-mac-intel compile-mac-apple-silicon
	
build-mac: clean compile-mac-intel compile-mac-apple-silicon

test-coverprofile:
	cd src && go test --coverprofile=coverage.out ./...

test: test-coverprofile
	cd src && go tool cover -func=coverage.out

test-html: test
	cd src && go tool cover -html=coverage.out	
clean: 
	rm -rf bin

run:
	go run main.go