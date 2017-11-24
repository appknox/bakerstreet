# go get -u github.com/golang/protobuf/protoc-gen-go
# go get github.com/nstogner/protoc-gen-grpc-go-service
BAKERSTREET="bakerstreet"
PROTOFILE=$BAKERSTREET/$BAKERSTREET.proto
rm $BAKERSTREET/$BAKERSTREET*{.go,.py}
protoc -I . \
	--go_out=plugins=grpc:. \
	$PROTOFILE

python -m grpc_tools.protoc \
	-I . \
	--python_out=. \
	--grpc_python_out=. \
	$PROTOFILE
