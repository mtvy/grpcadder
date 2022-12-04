
import grpc
import adder_pb2_grpc as adder_pb2_grpc
import adder_pb2 as adder_pb2

def run():
    with grpc.insecure_channel("localhost:8080") as channel:
        stub = adder_pb2_grpc.AdderStub(channel)
        
        _x = int(input("x = ")); _y = int(input("y = "))
        
        req = adder_pb2.AddRequest(x=_x, y=_y)
        print(req)
        repl = stub.Add(req)
        print(repl)

if __name__ == "__main__":
    run()