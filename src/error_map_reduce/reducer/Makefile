BUILD_DIR="../../bin"

build:
	go mod init reducer.go
	go build -o ${BUILD_DIR}/error_reducer reducer.go

clean:
	rm -rf go.sum go.mod ${BUILD_DIR}/error_reducer
