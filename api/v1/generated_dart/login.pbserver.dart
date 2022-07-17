///
//  Generated code. Do not modify.
//  source: login.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'login.pb.dart' as $0;
import 'login.pbjson.dart';

export 'login.pb.dart';

abstract class LoginSvcServiceBase extends $pb.GeneratedService {
  $async.Future<$0.LoginResp> login($pb.ServerContext ctx, $0.LoginReq request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Login': return $0.LoginReq();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Login': return this.login(ctx, request as $0.LoginReq);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => LoginSvcServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => LoginSvcServiceBase$messageJson;
}

