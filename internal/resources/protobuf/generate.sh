cd ../../..
protoc -I internal/ --go_out=./ internal/resources/protobuf/simple.proto
