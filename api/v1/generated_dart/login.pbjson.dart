///
//  Generated code. Do not modify.
//  source: login.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use loginReqDescriptor instead')
const LoginReq$json = const {
  '1': 'LoginReq',
  '2': const [
    const {'1': 'username', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'username'},
    const {'1': 'password', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'password'},
  ],
};

/// Descriptor for `LoginReq`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginReqDescriptor = $convert.base64Decode('CghMb2dpblJlcRIjCgh1c2VybmFtZRgBIAEoCUIH+kIEcgIQA1IIdXNlcm5hbWUSIwoIcGFzc3dvcmQYAiABKAlCB/pCBHICEANSCHBhc3N3b3Jk');
@$core.Deprecated('Use loginRespDescriptor instead')
const LoginResp$json = const {
  '1': 'LoginResp',
  '2': const [
    const {'1': 'token', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'token'},
  ],
};

/// Descriptor for `LoginResp`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginRespDescriptor = $convert.base64Decode('CglMb2dpblJlc3ASHQoFdG9rZW4YASABKAlCB/pCBHICEANSBXRva2Vu');
const $core.Map<$core.String, $core.dynamic> LoginSvcServiceBase$json = const {
  '1': 'LoginSvc',
  '2': const [
    const {'1': 'Login', '2': '.api.user.v1.LoginReq', '3': '.api.user.v1.LoginResp', '4': const {}},
  ],
};

@$core.Deprecated('Use loginSvcServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> LoginSvcServiceBase$messageJson = const {
  '.api.user.v1.LoginReq': LoginReq$json,
  '.api.user.v1.LoginResp': LoginResp$json,
};

/// Descriptor for `LoginSvc`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List loginSvcServiceDescriptor = $convert.base64Decode('CghMb2dpblN2YxJMCgVMb2dpbhIVLmFwaS51c2VyLnYxLkxvZ2luUmVxGhYuYXBpLnVzZXIudjEuTG9naW5SZXNwIhSC0+STAg4iCS9sb2dpbi92MToBKg==');
