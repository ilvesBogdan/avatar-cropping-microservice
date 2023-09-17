from concurrent import futures

import grpc
import grpc_pb2_grpc as pb
from grpc_pb2 import ImageResponse
from img import encode_image, image_processing
from PIL import UnidentifiedImageError


class ImageService(pb.ImageServiceServicer):
    def UploadImage(self, request, context):

        try:
            image_data = image_processing(request.raw_image)
        except UnidentifiedImageError:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("Invalid file transfer format")
            return
        except Exception as err:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f"A file reading error occurred: {err}")
            return

        for format in request.formats:
            if not context.is_active():
                break
            image = encode_image(image_data.copy(), format.size, format.format)
            yield ImageResponse(image_data=image)


def main(port: int):
    server = grpc.server(futures.ThreadPoolExecutor())
    pb.add_ImageServiceServicer_to_server(ImageService(), server)
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    main(50051)
