.PHONY: clean

SOURCES=$(wildcard ../src/*.cpp)
TEST_SOURCES=$(wildcard *.cpp)
OBJECTS=$(SOURCES:../src/%.cpp=build/objects/%.o)
TESTS=$(TEST_SOURCES:%.cpp=build/%)
CXXFLAGS += -std=c++11

all: $(OBJECTS) $(TESTS)

build:
	mkdir build

build/objects: build
	mkdir build/objects

build/objects/%.o: ../src/%.cpp build/objects
	$(CXX) $(CXXFLAGS) -c $< -o $@

build/%: %.cpp $(OBJECTS)
	$(CXX) $(CXXFLAGS) $(OBJECTS) $< -o $@

clean:
	rm -rf build