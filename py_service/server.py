from concurrent import futures

import grpc
import adder_pb2_grpc as adder_pb2_grpc
import adder_pb2 as adder_pb2


class AdderServicer(adder_pb2_grpc.AdderServicer):
    def Add(self, request, context):
        print(request)
        return adder_pb2.AddResponse(result=request.x+request.y)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    adder_pb2_grpc.add_AdderServicer_to_server(AdderServicer(), server)
    server.add_insecure_port("localhost:8081")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()