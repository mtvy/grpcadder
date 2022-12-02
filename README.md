# GRPC Template

## Proto Build
    ```bash
    protoc -I api/proto --go-grpc_out=pkg api/proto/adder.proto
    ```

# Server run
    ```bash
    go build -v ./cmd/server
    ./server
    ```

# Client run
    ```bash
    go run ./cmd/client/main.go
    ```

# Repo init
    go mod init <repo>
    go mod tidy
