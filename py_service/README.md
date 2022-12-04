# Python GRPC Template

### Proto Build
    python3 -m grpc_tools.protoc -I ../api/proto/ --python_out=. --pyi_out=. --grpc_python_out=. ../api/proto/adder.proto

### Server run
    python3 server.py

### Client run
- Go-server connection
```bash
python3 client.py --host=localhost:8080
```
- Py-server connection
```bash
python3 client.py --host=localhost:8081
```

### Repo init
    pip install grpcio-tools
