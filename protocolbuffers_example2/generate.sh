protoc -I src/ --go_out=src/simple/ src/simple/simple.proto
protoc -I src/ --go_out=src/enum_example/ src/enum_example/enum_example.proto
protoc -I src/ --go_out=src/complex/ src/complex/complex.proto