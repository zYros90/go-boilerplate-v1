///
//  Generated code. Do not modify.
//  source: todo.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'todo.pb.dart' as $3;
import 'google/protobuf/empty.pb.dart' as $2;
import 'todo.pbjson.dart';

export 'todo.pb.dart';

abstract class TodoSvcServiceBase extends $pb.GeneratedService {
  $async.Future<$3.TodoResp> create($pb.ServerContext ctx, $3.CreateTodoReq request);
  $async.Future<$3.TodoResp> update($pb.ServerContext ctx, $3.CreateTodoReq request);
  $async.Future<$3.TodoResp> get($pb.ServerContext ctx, $3.GetTodoReq request);
  $async.Future<$2.Empty> delete($pb.ServerContext ctx, $3.DeleteTodoReq request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Create': return $3.CreateTodoReq();
      case 'Update': return $3.CreateTodoReq();
      case 'Get': return $3.GetTodoReq();
      case 'Delete': return $3.DeleteTodoReq();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Create': return this.create(ctx, request as $3.CreateTodoReq);
      case 'Update': return this.update(ctx, request as $3.CreateTodoReq);
      case 'Get': return this.get(ctx, request as $3.GetTodoReq);
      case 'Delete': return this.delete(ctx, request as $3.DeleteTodoReq);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => TodoSvcServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => TodoSvcServiceBase$messageJson;
}

