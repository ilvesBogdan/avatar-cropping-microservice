import grpc
from concurrent import futures
import grpc_pb2_grpc as pb

from grpc_pb2 import ImageResponse
from img import image_processing, encode_image


class ImageService(pb.ImageServiceServicer):
    def UploadImage(self, request, context):
        base64_image = request.base64_image

        image_data = image_processing(base64_image)

        image_data = encode_image(image_data, 650, 'webp')

        return ImageResponse(image_data=image_data)


def main(port: int):
    server = grpc.server(futures.ThreadPoolExecutor())
    pb.add_ImageServiceServicer_to_server(ImageService(), server)
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    main(50051)
