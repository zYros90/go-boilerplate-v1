///
//  Generated code. Do not modify.
//  source: todo.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import 'google/protobuf/timestamp.pbjson.dart' as $1;
import 'google/protobuf/empty.pbjson.dart' as $2;

@$core.Deprecated('Use createTodoReqDescriptor instead')
const CreateTodoReq$json = const {
  '1': 'CreateTodoReq',
  '2': const [
    const {'1': 'todo', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'todo'},
    const {'1': 'due_at', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'dueAt'},
    const {'1': 'notify_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'notifyAt'},
  ],
  '7': const {},
};

/// Descriptor for `CreateTodoReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTodoReqDescriptor = $convert.base64Decode('Cg1DcmVhdGVUb2RvUmVxEhsKBHRvZG8YASABKAlCB/pCBHICEAFSBHRvZG8SMQoGZHVlX2F0GAIgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIFZHVlQXQSNwoJbm90aWZ5X2F0GAMgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIIbm90aWZ5QXQ6BbpGAiAB');
@$core.Deprecated('Use updateTodoReqDescriptor instead')
const UpdateTodoReq$json = const {
  '1': 'UpdateTodoReq',
  '2': const [
    const {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'todoId'},
    const {'1': 'todo', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'todo'},
    const {'1': 'due_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'dueAt'},
    const {'1': 'notify_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'notifyAt'},
  ],
  '7': const {},
};

/// Descriptor for `UpdateTodoReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateTodoReqDescriptor = $convert.base64Decode('Cg1VcGRhdGVUb2RvUmVxEiAKB3RvZG9faWQYASABKAlCB/pCBHICEAFSBnRvZG9JZBIbCgR0b2RvGAIgASgJQgf6QgRyAhABUgR0b2RvEjEKBmR1ZV9hdBgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSBWR1ZUF0EjcKCW5vdGlmeV9hdBgEIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCG5vdGlmeUF0OgW6RgIgAQ==');
@$core.Deprecated('Use deleteTodoReqDescriptor instead')
const DeleteTodoReq$json = const {
  '1': 'DeleteTodoReq',
  '2': const [
    const {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'todoId'},
  ],
};

/// Descriptor for `DeleteTodoReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteTodoReqDescriptor = $convert.base64Decode('Cg1EZWxldGVUb2RvUmVxEiAKB3RvZG9faWQYASABKAlCB/pCBHICEAFSBnRvZG9JZA==');
@$core.Deprecated('Use getTodoReqDescriptor instead')
const GetTodoReq$json = const {
  '1': 'GetTodoReq',
  '2': const [
    const {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'todoId'},
  ],
};

/// Descriptor for `GetTodoReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getTodoReqDescriptor = $convert.base64Decode('CgpHZXRUb2RvUmVxEiAKB3RvZG9faWQYASABKAlCB/pCBHICEAFSBnRvZG9JZA==');
@$core.Deprecated('Use todoRespDescriptor instead')
const TodoResp$json = const {
  '1': 'TodoResp',
  '2': const [
    const {'1': 'todo_id', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'todoId'},
    const {'1': 'todo', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'todo'},
    const {'1': 'due_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'dueAt'},
    const {'1': 'notify_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'notifyAt'},
    const {'1': 'created_at', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '8': const {}, '10': 'createdAt'},
    const {'1': 'updated_at', '3': 6, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '8': const {}, '10': 'updatedAt'},
  ],
};

/// Descriptor for `TodoResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List todoRespDescriptor = $convert.base64Decode('CghUb2RvUmVzcBIgCgd0b2RvX2lkGAEgASgJQgf6QgRyAhABUgZ0b2RvSWQSGwoEdG9kbxgCIAEoCUIH+kIEcgIQAVIEdG9kbxIxCgZkdWVfYXQYAyABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgVkdWVBdBI3Cglub3RpZnlfYXQYBCABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUghub3RpZnlBdBJDCgpjcmVhdGVkX2F0GAUgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEII+kIFsgECCAFSCWNyZWF0ZWRBdBJDCgp1cGRhdGVkX2F0GAYgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEII+kIFsgECCAFSCXVwZGF0ZWRBdA==');
const $core.Map<$core.String, $core.dynamic> TodoSvcServiceBase$json = const {
  '1': 'TodoSvc',
  '2': const [
    const {'1': 'Create', '2': '.api.user.v1.CreateTodoReq', '3': '.api.user.v1.TodoResp', '4': const {}},
    const {'1': 'Update', '2': '.api.user.v1.CreateTodoReq', '3': '.api.user.v1.TodoResp', '4': const {}},
    const {'1': 'Get', '2': '.api.user.v1.GetTodoReq', '3': '.api.user.v1.TodoResp', '4': const {}},
    const {'1': 'Delete', '2': '.api.user.v1.DeleteTodoReq', '3': '.google.protobuf.Empty', '4': const {}},
  ],
};

@$core.Deprecated('Use todoSvcServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> TodoSvcServiceBase$messageJson = const {
  '.api.user.v1.CreateTodoReq': CreateTodoReq$json,
  '.google.protobuf.Timestamp': $1.Timestamp$json,
  '.api.user.v1.TodoResp': TodoResp$json,
  '.api.user.v1.GetTodoReq': GetTodoReq$json,
  '.api.user.v1.DeleteTodoReq': DeleteTodoReq$json,
  '.google.protobuf.Empty': $2.Empty$json,
};

/// Descriptor for `TodoSvc`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List todoSvcServiceDescriptor = $convert.base64Decode('CgdUb2RvU3ZjElAKBkNyZWF0ZRIaLmFwaS51c2VyLnYxLkNyZWF0ZVRvZG9SZXEaFS5hcGkudXNlci52MS5Ub2RvUmVzcCITgtPkkwINIggvdG9kby92MToBKhJQCgZVcGRhdGUSGi5hcGkudXNlci52MS5DcmVhdGVUb2RvUmVxGhUuYXBpLnVzZXIudjEuVG9kb1Jlc3AiE4LT5JMCDRoIL3RvZG8vdjE6ASoSUAoDR2V0EhcuYXBpLnVzZXIudjEuR2V0VG9kb1JlcRoVLmFwaS51c2VyLnYxLlRvZG9SZXNwIhmC0+STAhMSES90b2RvL3YxLzp0b2RvX2lkElcKBkRlbGV0ZRIaLmFwaS51c2VyLnYxLkRlbGV0ZVRvZG9SZXEaFi5nb29nbGUucHJvdG9idWYuRW1wdHkiGYLT5JMCEyoRL3RvZG8vdjEvOnRvZG9faWQ=');
