# Golang GRPC Template

### Proto Build
    ```bash
    protoc -I ../api/proto --go_out=pkg --go-grpc_out=pkg ../api/proto/adder.proto
    ```

### Server run
    ```bash
    go build -v ./cmd/server/main.go
    ```

### Client run
    ```bash
    go run ./cmd/client/main.go
    ```

### Repo init
    go mod init <repo>
    go mod tidy

### Issues

    - proto build with error PATH:
        ```bash
        nano ~/.bash_profile 
        export GO_PATH=~/go
        export PATH=$PATH:/$GO_PATH/bin
        source ~/.bash_profile
        ```