# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: grpc.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\ngrpc.proto\x12\x0cImageService\"H\n\x0cImageRequest\x12%\n\x07\x66ormats\x18\x01 \x03(\x0b\x32\x14.ImageService.Format\x12\x11\n\traw_image\x18\x02 \x01(\x0c\"#\n\rImageResponse\x12\x12\n\nimage_data\x18\x01 \x01(\x0c\"&\n\x06\x46ormat\x12\x0e\n\x06\x66ormat\x18\x01 \x01(\t\x12\x0c\n\x04size\x18\x02 \x01(\x05\x32Z\n\x0cImageService\x12J\n\x0bUploadImage\x12\x1a.ImageService.ImageRequest\x1a\x1b.ImageService.ImageResponse\"\x00\x30\x01\x42\x1fZ\x1dimgresizer/golang client/mainb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'grpc_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\035imgresizer/golang client/main'
  _globals['_IMAGEREQUEST']._serialized_start=28
  _globals['_IMAGEREQUEST']._serialized_end=100
  _globals['_IMAGERESPONSE']._serialized_start=102
  _globals['_IMAGERESPONSE']._serialized_end=137
  _globals['_FORMAT']._serialized_start=139
  _globals['_FORMAT']._serialized_end=177
  _globals['_IMAGESERVICE']._serialized_start=179
  _globals['_IMAGESERVICE']._serialized_end=269
# @@protoc_insertion_point(module_scope)
