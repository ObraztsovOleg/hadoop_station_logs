BUILD_DIR="../bin"

build:
	go mod init reducer.go
	go build -o ${BUILD_DIR}/h_reducer reducer.go

clean:
	rm -rf go.sum go.mod ${BUILD_DIR}/h_reducer
