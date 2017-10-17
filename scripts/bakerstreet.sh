# go get -u github.com/golang/protobuf/protoc-gen-go
# go get github.com/nstogner/protoc-gen-grpc-go-service
BAKERSTREET="bakerstreet"
PROTOFILE=$BAKERSTREET/$BAKERSTREET.proto
rm $BAKERSTREET/$BAKERSTREET*{.go,.py}
protoc -I $BAKERSTREET \
	--go_out=plugins=grpc:$BAKERSTREET \
	$PROTOFILE

python -m grpc_tools.protoc \
	-I $BAKERSTREET \
	--python_out=$BAKERSTREET \
	--grpc_python_out=$BAKERSTREET \
	$PROTOFILE
