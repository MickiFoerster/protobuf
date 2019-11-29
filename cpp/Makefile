CXXFLAGS += -I. -I/home/micki/programs/protobuf/bld/include
BIN=simple_demo

%.cc : %.h
%.o : %.cc
	$(CXX) $(CXXFLAGS) -c -o $@ $<

all : $(BIN)

simple_demo: simple_demo.o
	$(CXX) -o $@ $<
simple_demo.o : simple/simple.pb.o 
simple/simple.pb.o : simple/simple.pb.h


clean:
	rm -f $(BIN) *.o

.PHONY: clean all
