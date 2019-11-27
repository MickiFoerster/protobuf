dir=com.github.MickiFoerster.protobuf
if [ ! -d $dir ]; then
    mkdir -p $dir
fi
protobuf_ver=$(protoc --version | cut -d' ' -f2)
if [ ! -f protobuf-java-3.10.0.jar  ];then
    wget -O protobuf-java-3.10.0.jar \
    https://repo1.maven.org/maven2/com/google/protobuf/protobuf-java/${protobuf_ver}/protobuf-java-3.10.0.jar
fi
protoc -I simple/ --java_out=example/ simple/simple.proto
#protoc -I enums/ --java_out=enums/ enums/enum_example.proto
#protoc -I complex/ --java_out=complex/ complex/complex.proto
