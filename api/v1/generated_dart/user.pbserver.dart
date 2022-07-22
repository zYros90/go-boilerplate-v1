///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'user.pb.dart' as $4;
import 'google/protobuf/empty.pb.dart' as $2;
import 'user.pbjson.dart';

export 'user.pb.dart';

abstract class UserSvcServiceBase extends $pb.GeneratedService {
  $async.Future<$4.UserResp> create($pb.ServerContext ctx, $4.CreateUserReq request);
  $async.Future<$4.UserResp> update($pb.ServerContext ctx, $4.CreateUserReq request);
  $async.Future<$4.UserResp> get($pb.ServerContext ctx, $4.GetUserReq request);
  $async.Future<$2.Empty> delete($pb.ServerContext ctx, $4.DeleteUserReq request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Create': return $4.CreateUserReq();
      case 'Update': return $4.CreateUserReq();
      case 'Get': return $4.GetUserReq();
      case 'Delete': return $4.DeleteUserReq();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Create': return this.create(ctx, request as $4.CreateUserReq);
      case 'Update': return this.update(ctx, request as $4.CreateUserReq);
      case 'Get': return this.get(ctx, request as $4.GetUserReq);
      case 'Delete': return this.delete(ctx, request as $4.DeleteUserReq);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => UserSvcServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => UserSvcServiceBase$messageJson;
}

