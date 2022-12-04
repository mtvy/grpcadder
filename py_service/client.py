
import grpc
import adder_pb2_grpc as adder_pb2_grpc
import adder_pb2 as adder_pb2

import sys

def run():

    host = sys.argv[1].split('=')[1] # localhost:808...

    with grpc.insecure_channel(host) as conn:
        stub = adder_pb2_grpc.AdderStub(conn)
        
        _x = int(input("x = ")); _y = int(input("y = "))
        
        req = adder_pb2.AddRequest(x=_x, y=_y)
        res = stub.Add(req)

        print(res)

if __name__ == "__main__":
    run()