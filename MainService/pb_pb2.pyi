from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CallbackRequest(_message.Message):
    __slots__ = ("user", "action", "param")
    USER_FIELD_NUMBER: _ClassVar[int]
    ACTION_FIELD_NUMBER: _ClassVar[int]
    PARAM_FIELD_NUMBER: _ClassVar[int]
    user: str
    action: str
    param: str
    def __init__(self, user: _Optional[str] = ..., action: _Optional[str] = ..., param: _Optional[str] = ...) -> None: ...

class CallbackReply(_message.Message):
    __slots__ = ("text", "buttons")
    TEXT_FIELD_NUMBER: _ClassVar[int]
    BUTTONS_FIELD_NUMBER: _ClassVar[int]
    text: str
    buttons: _containers.RepeatedCompositeFieldContainer[Button]
    def __init__(self, text: _Optional[str] = ..., buttons: _Optional[_Iterable[_Union[Button, _Mapping]]] = ...) -> None: ...

class Button(_message.Message):
    __slots__ = ("text", "data")
    TEXT_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    text: str
    data: str
    def __init__(self, text: _Optional[str] = ..., data: _Optional[str] = ...) -> None: ...
