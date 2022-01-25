# Proto files

When modifying files you need to compile them
```
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative simulations/simulations.proto
```

