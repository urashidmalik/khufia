.ONESHELL: # Applies to every targets in the file!

BINARY_NAME=khufia
FREEBSD=freebsd
WINDOWS=windows
LINUX=linux
DARWIN=darwin

ARCH_386=386
ARCH_AMD64=amd64
ARCH_ARM64=arm64


compile-windows:
	cd src && GOOS=${WINDOWS} GOARCH=${ARCH_386} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${WINDOWS}-${ARCH_386} main.go
	cd src && GOOS=${WINDOWS} GOARCH=${ARCH_AMD64} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${WINDOWS}-${ARCH_AMD64} main.go
	cd src && GOOS=${WINDOWS} GOARCH=${ARCH_ARM64} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${WINDOWS}-${ARCH_ARM64} main.go

compile-linux:
	cd src && GOOS=${LINUX} GOARCH=${ARCH_386} go build -o ../bin/${BINARY_NAME}-${LINUX}-${ARCH_386} main.go
	cd src && GOOS=${LINUX} GOARCH=${ARCH_AMD64} go build -o ../bin/${BINARY_NAME}-${LINUX}-${ARCH_AMD64} main.go
	cd src && GOOS=${LINUX} GOARCH=${ARCH_ARM64} go build -o ../bin/${BINARY_NAME}-${LINUX}-${ARCH_ARM64} main.go
	
compile-freebsd:	
	cd src && GOOS=${FREEBSD} GOARCH=${ARCH_386} go build -o ../bin/${BINARY_NAME}-${FREEBSD}-${ARCH_386} main.go
	cd src && GOOS=${FREEBSD} GOARCH=${ARCH_AMD64}  go build -o ../bin/${BINARY_NAME}-${FREEBSD}-${ARCH_AMD64} main.go
	cd src && GOOS=${FREEBSD} GOARCH=${ARCH_ARM64}  go build -o ../bin/${BINARY_NAME}-${FREEBSD}-${ARCH_ARM64} main.go

compile-mac-intel:
	cd src && GOOS=${DARWIN} GOARCH=${ARCH_AMD64} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${DARWIN}-${ARCH_AMD64} main.go

compile-mac-apple-silicon:
	cd src && GOOS=${DARWIN} GOARCH=${ARCH_ARM64} CGO_ENABLED=1 go build -o ../bin/${BINARY_NAME}-${DARWIN}-${ARCH_ARM64} main.go


build: clean compile-windows compile-windows compile-windows compile-freebsd compile-linux compile-mac-intel compile-mac-apple-silicon
	
build-mac: clean compile-mac-intel compile-mac-apple-silicon

test:
	cd src && go test --coverprofile=coverage.out ./...
	cd src && go tool cover -func=coverage.out

test-html: test
	cd src && go tool cover -html=coverage.out	
clean: 
	rm -rf bin

run:
	go run main.go