# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pb.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x08pb.proto\x12\x05manga\">\n\x0f\x43\x61llbackRequest\x12\x0c\n\x04user\x18\x01 \x01(\t\x12\x0e\n\x06\x61\x63tion\x18\x02 \x01(\t\x12\r\n\x05param\x18\x03 \x01(\t\"=\n\rCallbackReply\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x1e\n\x07\x62uttons\x18\x02 \x03(\x0b\x32\r.manga.Button\"$\n\x06\x42utton\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\t2C\n\x05Manga\x12:\n\x08\x43hannel1\x12\x16.manga.CallbackRequest\x1a\x14.manga.CallbackReply\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'pb_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_CALLBACKREQUEST']._serialized_start=19
  _globals['_CALLBACKREQUEST']._serialized_end=81
  _globals['_CALLBACKREPLY']._serialized_start=83
  _globals['_CALLBACKREPLY']._serialized_end=144
  _globals['_BUTTON']._serialized_start=146
  _globals['_BUTTON']._serialized_end=182
  _globals['_MANGA']._serialized_start=184
  _globals['_MANGA']._serialized_end=251
# @@protoc_insertion_point(module_scope)
