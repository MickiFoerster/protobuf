protobufver=$(shell protoc --version | cut -d' ' -f2)
CLASS_PATH=-cp .:gson-2.8.5.jar:protobuf-java-$(protobufver).jar:protobuf-java-util-$(protobufver).jar:build/generated

run :  SimpleMain EnumExample ComplexExample

examples: SimpleMain EnumExample ComplexExample ProtoToJSONMain

SimpleMain : com/github/MickiFoerster/protobuf/SimpleMain.class
	java $(CLASS_PATH) com.github.MickiFoerster.protobuf.SimpleMain

EnumExample : com/github/MickiFoerster/protobuf/EnumMain.class
	java $(CLASS_PATH) com.github.MickiFoerster.protobuf.EnumMain

ComplexExample : com/github/MickiFoerster/protobuf/ComplexMain.class
	java $(CLASS_PATH) com.github.MickiFoerster.protobuf.ComplexMain

ProtoToJSONMain : com/github/MickiFoerster/protobuf/ProtoToJSONMain.class
	java $(CLASS_PATH) com.github.MickiFoerster.protobuf.ProtoToJSONMain
	java $(CLASS_PATH) com/github/MickiFoerster/protobuf/ProtoToJSONMain

com/github/MickiFoerster/protobuf/SimpleMain.class : src/main/java/com.github.MickiFoerster.protobuf/SimpleMain.java protobuf-java-$(protobufver).jar build/generated/com/example/simple/Simple.java
	javac $(CLASS_PATH) $< 
	javac -d . $(CLASS_PATH) $<

com/github/MickiFoerster/protobuf/EnumMain.class : src/main/java/com.github.MickiFoerster.protobuf/EnumMain.java protobuf-java-$(protobufver).jar build/generated/com/example/enumerations/EnumExample.java
	javac $(CLASS_PATH) $< 
	javac -d . $(CLASS_PATH) $<

com/github/MickiFoerster/protobuf/ComplexMain.class : src/main/java/com.github.MickiFoerster.protobuf/ComplexMain.java protobuf-java-$(protobufver).jar build/generated/com/example/complex/ComplexExample.java
	javac $(CLASS_PATH) $< 
	javac -d . $(CLASS_PATH) $<

com/github/MickiFoerster/protobuf/ProtoToJSONMain.class : src/main/java/com.github.MickiFoerster.protobuf/ProtoToJSONMain.java \
                                                          src/main/java/com.github.MickiFoerster.protobuf/SimpleMain.java \
                                                          protobuf-java-$(protobufver).jar \
														  protobuf-java-util-$(protobufver).jar \
														  gson-2.8.5.jar
	javac $(CLASS_PATH) $< 
	javac -d . $(CLASS_PATH) $<

build/generated/com/example/simple/Simple.java : src/main/proto/simple.proto build
	protoc -Isrc/main/proto --java_out=build/generated/ $<
build/generated/com/example/enumerations/EnumExample.java : src/main/proto/enum_example.proto build
	protoc -Isrc/main/proto --java_out=build/generated/ $<
build/generated/com/example/complex/ComplexExample.java : src/main/proto/complex.proto build
	protoc -Isrc/main/proto --java_out=build/generated/ $<

build : 
	mkdir -p build/generated

protobuf-java-$(protobufver).jar :
	wget -O $@ \
    https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java/${protobufver}/$@

protobuf-java-util-$(protobufver).jar :
	wget -O $@ \
    https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java-util/${protobufver}/$@

gson-2.8.5.jar : 
	wget -O $@ \
	https://repo1.maven.org/maven2/com/google/code/gson/gson/2.8.5/gson-2.8.5.jar

.PHONY: run
