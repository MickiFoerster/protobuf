protobufver=$(shell protoc --version | cut -d' ' -f2)

run : com/github/MickiFoerster/protobuf/SimpleMain.class
	java -cp .:protobuf-java-3.10.0.jar com.github.MickiFoerster.protobuf.SimpleMain

com/github/MickiFoerster/protobuf/SimpleMain.class : com.github.MickiFoerster.protobuf/SimpleMain.java protobuf-java-$(protobufver).jar example/simple/Simple.java
	javac -cp .:protobuf-java-3.10.0.jar com.github.MickiFoerster.protobuf/SimpleMain.java 
	javac -d . -cp .:protobuf-java-$(protobufver).jar com.github.MickiFoerster.protobuf/SimpleMain.java 

protobuf-java-$(protobufver).jar :
	wget -O protobuf-java-$(protobufver).jar \
    https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java/${protobufver}/protobuf-java-$(protobufver).jar

example/simple/Simple.java : simple/simple.proto
	protoc -I simple/ --java_out=. simple/simple.proto

.PHONY: run
