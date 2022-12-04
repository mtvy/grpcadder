# GRPC Template

## Proto Build
    ```bash
    protoc -I api/proto --go-grpc_out=pkg api/proto/adder.proto
    protoc -I api/proto --go_out=pkg api/proto/adder.proto

    python3 -m grpc_tools.protoc -I api/proto/ --python_out=pkg/api api/proto/adder.proto
    python3 -m grpc_tools.protoc -I api/proto/ --grpc_python_out=pkg/api api/proto/adder.proto
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
    
    nano ~/.bash_profile 
    export GO_PATH=~/go
    export PATH=$PATH:/$GO_PATH/bin
    source ~/.bash_profile