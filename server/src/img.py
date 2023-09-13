import os
import re
import base64
from PIL import Image
from io import BytesIO


def image_processing(base64str: str) -> Image.Image:
    """
    Обрабатывает изображение, закодированное в формате base64.
    Аргументы:
        base64str (str): Строка, содержащая изображение в формате base64.
    Возвращает:
        Image.Image: Объект изображения типа PIL.Image.Image.
    """
    text = re.sub('^data:image/.+;base64,', '', base64str)
    img = Image.open(BytesIO(base64.b64decode(text)))
    img = crop_center(img)
    return img

    # for size, format in ((650, 'webp'), (650, 'jpeg'), (200, 'webp'), (200, 'jpeg')):
    #     save_image(path, img.copy(), size, format)
    # print(avatar_id)


def encode_image(img: Image.Image, size: int, format: str) -> bytes:
    """
    Кодирует изображение в поток байт.

    Аргументы:
      - img (PIL.Image.Image): Изображение, которое нужно закодировать.
      - size (int): Размер изображения после изменения (ширина и высота).
      - format (str): Формат кодирования изображения.

    Возвращает:
        bytes: Закодированное изображение в виде потока байт.
    """
    img.thumbnail(size=(size, size))
    buffer = BytesIO()
    img.save(buffer, format=format, quality=95)
    encoded_image = buffer.getvalue()
    return encoded_image


def crop_center(pil_img):
    """
    Функция для обрезки изображения по центру.

    Аргументы:
        pil_img (PIL.Image): Объект изображения.

    Возвращает:
        PIL.Image: Обрезанное изображение.
    """
    crop = min(pil_img.size)
    img_width, img_height = pil_img.size
    return pil_img.crop(((img_width - crop) // 2,
                        (img_height - crop) // 2,
                        (img_width + crop) // 2,
                        (img_height + crop) // 2))
