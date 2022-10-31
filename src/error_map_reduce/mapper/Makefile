BUILD_DIR="../bin"

build:
	go mod init mapper.go
	go build -o ${BUILD_DIR}/error_mapper mapper.go

clean:
	rm -rf go.sum go.mod ${BUILD_DIR}/error_mapper
