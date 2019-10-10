pbdir=proto

all: clean g
	go build -o app
clean:
	go clean
	rm -rf app
	rm -rf ${pbdir}/*.go
g:
	cd ${pbdir} &&  protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  --go_out=plugins=grpc:. --grpc-gateway_out=. *.proto