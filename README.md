### Setup

1. `brew update`
2. `brew install protobuf protoc-gen-go protoc-gen-go-grpc`

### Generate Cliend and Server code

```bash
make protogen
```

### Start the server

```bash
make run
```
