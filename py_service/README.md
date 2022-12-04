# Python GRPC Template

### Proto Build
    ```bash
    python3 -m grpc_tools.protoc -I ../api/proto/ --python_out=. --pyi_out=. --grpc_python_out=. ../api/proto/adder.proto
    ```

### Server run
    ```bash
    python3 server.py
    ```

### Client run
    ```bash
    python3 client.py
    ```

### Repo init
    pip install grpcio-tools
