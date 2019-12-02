protobufver=$(shell protoc --version | cut -d' ' -f2)

run : ComplexExample

SimpleMain : com/github/MickiFoerster/protobuf/SimpleMain.class
	java -cp .:protobuf-java-3.10.0.jar com.github.MickiFoerster.protobuf.SimpleMain

EnumExample : com/github/MickiFoerster/protobuf/EnumMain.class
	java -cp .:protobuf-java-3.10.0.jar com.github.MickiFoerster.protobuf.EnumMain

ComplexExample : com/github/MickiFoerster/protobuf/ComplexMain.class
	java -cp .:protobuf-java-3.10.0.jar com.github.MickiFoerster.protobuf.ComplexMain

com/github/MickiFoerster/protobuf/SimpleMain.class : com.github.MickiFoerster.protobuf/SimpleMain.java protobuf-java-$(protobufver).jar example/simple/Simple.java
	javac -cp .:protobuf-java-3.10.0.jar $< 
	javac -d . -cp .:protobuf-java-$(protobufver).jar $<

com/github/MickiFoerster/protobuf/EnumMain.class : com.github.MickiFoerster.protobuf/EnumMain.java protobuf-java-$(protobufver).jar example/enumerations/EnumExample.java
	javac -cp .:protobuf-java-3.10.0.jar $< 
	javac -d . -cp .:protobuf-java-$(protobufver).jar $<

com/github/MickiFoerster/protobuf/ComplexMain.class : com.github.MickiFoerster.protobuf/ComplexMain.java protobuf-java-$(protobufver).jar example/complex/ComplexExample.java
	javac -cp .:protobuf-java-3.10.0.jar $< 
	javac -d . -cp .:protobuf-java-$(protobufver).jar $<

example/simple/Simple.java : simple/simple.proto
	protoc -I simple/ --java_out=. $<
example/enumerations/EnumExample.java : enums/enum_example.proto
	protoc -I enums/ --java_out=. $<
example/complex/ComplexExample.java : complex/complex.proto
	protoc -I complex/ --java_out=. $<

protobuf-java-$(protobufver).jar :
	wget -O protobuf-java-$(protobufver).jar \
    https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java/${protobufver}/protobuf-java-$(protobufver).jar


.PHONY: run
