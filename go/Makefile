go/simple/simple.pb.go: src/simple/simple.proto go
	protoc -I src/ --go_out=go/ src/simple/simple.proto
go/enum_example/enum_example.pb.go: src/enum_example/enum_example.proto
	protoc -I src/ --go_out=go/ src/enum_example/enum_example.proto
go/simple/simple.pb.go: src/simple/simple.proto go
	protoc -I src/ --go_out=go/ src/complex/complex.proto
rust:
	protoc -I src/ --rust_out=go/ src/simple/simple.proto
	protoc -I src/ --rust_out=go/ src/enum_example/enum_example.proto
	protoc -I src/ --rust_out=go/ src/complex/complex.proto

go:
	mkdir -p go
