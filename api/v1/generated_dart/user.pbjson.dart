///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import 'google/protobuf/timestamp.pbjson.dart' as $1;

@$core.Deprecated('Use createUserReqDescriptor instead')
const CreateUserReq$json = const {
  '1': 'CreateUserReq',
  '2': const [
    const {'1': 'username', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'username'},
    const {'1': 'password', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'password'},
    const {'1': 'email', '3': 3, '4': 1, '5': 9, '8': const {}, '10': 'email'},
    const {'1': 'first_name', '3': 4, '4': 1, '5': 9, '10': 'firstName'},
    const {'1': 'last_name', '3': 5, '4': 1, '5': 9, '10': 'lastName'},
  ],
  '7': const {},
};

/// Descriptor for `CreateUserReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createUserReqDescriptor = $convert.base64Decode('Cg1DcmVhdGVVc2VyUmVxEiMKCHVzZXJuYW1lGAEgASgJQgf6QgRyAhADUgh1c2VybmFtZRIjCghwYXNzd29yZBgCIAEoCUIH+kIEcgIQA1IIcGFzc3dvcmQSHQoFZW1haWwYAyABKAlCB/pCBHICYAFSBWVtYWlsEh0KCmZpcnN0X25hbWUYBCABKAlSCWZpcnN0TmFtZRIbCglsYXN0X25hbWUYBSABKAlSCGxhc3ROYW1lOgW6RgIgAQ==');
@$core.Deprecated('Use deleteUserReqDescriptor instead')
const DeleteUserReq$json = const {
  '1': 'DeleteUserReq',
};

/// Descriptor for `DeleteUserReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteUserReqDescriptor = $convert.base64Decode('Cg1EZWxldGVVc2VyUmVx');
@$core.Deprecated('Use deleteUserRespDescriptor instead')
const DeleteUserResp$json = const {
  '1': 'DeleteUserResp',
};

/// Descriptor for `DeleteUserResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteUserRespDescriptor = $convert.base64Decode('Cg5EZWxldGVVc2VyUmVzcA==');
@$core.Deprecated('Use getUserReqDescriptor instead')
const GetUserReq$json = const {
  '1': 'GetUserReq',
};

/// Descriptor for `GetUserReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUserReqDescriptor = $convert.base64Decode('CgpHZXRVc2VyUmVx');
@$core.Deprecated('Use userRespDescriptor instead')
const UserResp$json = const {
  '1': 'UserResp',
  '2': const [
    const {'1': 'username', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'username'},
    const {'1': 'email', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'email'},
    const {'1': 'first_name', '3': 3, '4': 1, '5': 9, '10': 'firstName'},
    const {'1': 'last_name', '3': 4, '4': 1, '5': 9, '10': 'lastName'},
    const {'1': 'created_at', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '8': const {}, '10': 'createdAt'},
    const {'1': 'updated_at', '3': 6, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '8': const {}, '10': 'updatedAt'},
  ],
};

/// Descriptor for `UserResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userRespDescriptor = $convert.base64Decode('CghVc2VyUmVzcBIjCgh1c2VybmFtZRgBIAEoCUIH+kIEcgIQA1IIdXNlcm5hbWUSHQoFZW1haWwYAiABKAlCB/pCBHICYAFSBWVtYWlsEh0KCmZpcnN0X25hbWUYAyABKAlSCWZpcnN0TmFtZRIbCglsYXN0X25hbWUYBCABKAlSCGxhc3ROYW1lEkMKCmNyZWF0ZWRfYXQYBSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wQgj6QgWyAQIIAVIJY3JlYXRlZEF0EkMKCnVwZGF0ZWRfYXQYBiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wQgj6QgWyAQIIAVIJdXBkYXRlZEF0');
const $core.Map<$core.String, $core.dynamic> UserSvcServiceBase$json = const {
  '1': 'UserSvc',
  '2': const [
    const {'1': 'Create', '2': '.api.user.v1.CreateUserReq', '3': '.api.user.v1.UserResp', '4': const {}},
    const {'1': 'Update', '2': '.api.user.v1.CreateUserReq', '3': '.api.user.v1.UserResp', '4': const {}},
    const {'1': 'Get', '2': '.api.user.v1.GetUserReq', '3': '.api.user.v1.UserResp', '4': const {}},
    const {'1': 'Delete', '2': '.api.user.v1.DeleteUserReq', '3': '.api.user.v1.DeleteUserResp', '4': const {}},
  ],
};

@$core.Deprecated('Use userSvcServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> UserSvcServiceBase$messageJson = const {
  '.api.user.v1.CreateUserReq': CreateUserReq$json,
  '.api.user.v1.UserResp': UserResp$json,
  '.google.protobuf.Timestamp': $1.Timestamp$json,
  '.api.user.v1.GetUserReq': GetUserReq$json,
  '.api.user.v1.DeleteUserReq': DeleteUserReq$json,
  '.api.user.v1.DeleteUserResp': DeleteUserResp$json,
};

/// Descriptor for `UserSvc`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List userSvcServiceDescriptor = $convert.base64Decode('CgdVc2VyU3ZjElAKBkNyZWF0ZRIaLmFwaS51c2VyLnYxLkNyZWF0ZVVzZXJSZXEaFS5hcGkudXNlci52MS5Vc2VyUmVzcCITgtPkkwINIggvdXNlci92MToBKhI9CgZVcGRhdGUSGi5hcGkudXNlci52MS5DcmVhdGVVc2VyUmVxGhUuYXBpLnVzZXIudjEuVXNlclJlc3AiABI3CgNHZXQSFy5hcGkudXNlci52MS5HZXRVc2VyUmVxGhUuYXBpLnVzZXIudjEuVXNlclJlc3AiABJDCgZEZWxldGUSGi5hcGkudXNlci52MS5EZWxldGVVc2VyUmVxGhsuYXBpLnVzZXIudjEuRGVsZXRlVXNlclJlc3AiAA==');
