GENERATED_CODE_TARGET_DIR=protoc-helloworld
GENERATED_CODE=$(GENERATED_CODE_TARGET_DIR)/hello-world.pb.go
PROTO_FILE=proto/hello-world.proto
BINARY=helloworld

all: \
	$(GENERATED_CODE) \
	$(BINARY)

$(BINARY) : main.go $(GENERATED_CODE)
	go build -o $@ $<

$(GENERATED_CODE): $(PROTO_FILE)
	mkdir -p $(GENERATED_CODE_TARGET_DIR)
	protoc -I proto/ --go_out=$(GENERATED_CODE_TARGET_DIR)/ $(PROTO_FILE)


clean:
	rm -rf $(GENERATED_CODE_TARGET_DIR) $(BINARY)


.PHONY: clean
