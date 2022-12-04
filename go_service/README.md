# Golang GRPC Template

### Proto Build
    protoc -I ../api/proto --go_out=pkg --go-grpc_out=pkg ../api/proto/adder.proto

### Server run
    go build -v ./cmd/server/main.go

### Client run
- Go-server connection
    ```bash
    go run ./cmd/client/main.go localhost:8080
    ```
- Py-server connection
    ```bash
    go run ./cmd/client/main.go localhost:8081
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
