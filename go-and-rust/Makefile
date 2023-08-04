OUTPUT_DIR=proto/go proto/rust

go: proto/go/simple/simple.pb.go \
	proto/go/enum_example/enum_example.pb.go \
	proto/go/complex/complex.pb.go 

proto/go/simple/simple.pb.go: src/simple/simple.proto $(OUTPUT_DIR)
	protoc -I src/ --go_out=proto/go/ src/simple/simple.proto
proto/go/enum_example/enum_example.pb.go: src/enum_example/enum_example.proto
	protoc -I src/ --go_out=proto/go/ src/enum_example/enum_example.proto
proto/go/complex/complex.pb.go: src/complex/complex.proto
	protoc -I src/ --go_out=proto/go/ src/complex/complex.proto

rust:
	mkdir -p proto/rust/simple proto/rust/enum_example proto/rust/complex
	protoc -I src/ --rust_out=proto/rust/simple       src/simple/simple.proto
	protoc -I src/ --rust_out=proto/rust/enum_example src/enum_example/enum_example.proto
	protoc -I src/ --rust_out=proto/rust/complex      src/complex/complex.proto

$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)
