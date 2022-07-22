///
//  Generated code. Do not modify.
//  source: todo.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $1;
import 'google/protobuf/empty.pb.dart' as $2;

class CreateTodoReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CreateTodoReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'api.user.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'todo')
    ..aOM<$1.Timestamp>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'dueAt', subBuilder: $1.Timestamp.create)
    ..aOM<$1.Timestamp>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notifyAt', subBuilder: $1.Timestamp.create)
    ..hasRequiredFields = false
  ;

  CreateTodoReq._() : super();
  factory CreateTodoReq({
    $core.String? todo,
    $1.Timestamp? dueAt,
    $1.Timestamp? notifyAt,
  }) {
    final _result = create();
    if (todo != null) {
      _result.todo = todo;
    }
    if (dueAt != null) {
      _result.dueAt = dueAt;
    }
    if (notifyAt != null) {
      _result.notifyAt = notifyAt;
    }
    return _result;
  }
  factory CreateTodoReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateTodoReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateTodoReq clone() => CreateTodoReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateTodoReq copyWith(void Function(CreateTodoReq) updates) => super.copyWith((message) => updates(message as CreateTodoReq)) as CreateTodoReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CreateTodoReq create() => CreateTodoReq._();
  CreateTodoReq createEmptyInstance() => create();
  static $pb.PbList<CreateTodoReq> createRepeated() => $pb.PbList<CreateTodoReq>();
  @$core.pragma('dart2js:noInline')
  static CreateTodoReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateTodoReq>(create);
  static CreateTodoReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todo => $_getSZ(0);
  @$pb.TagNumber(1)
  set todo($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodo() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodo() => clearField(1);

  @$pb.TagNumber(2)
  $1.Timestamp get dueAt => $_getN(1);
  @$pb.TagNumber(2)
  set dueAt($1.Timestamp v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasDueAt() => $_has(1);
  @$pb.TagNumber(2)
  void clearDueAt() => clearField(2);
  @$pb.TagNumber(2)
  $1.Timestamp ensureDueAt() => $_ensure(1);

  @$pb.TagNumber(3)
  $1.Timestamp get notifyAt => $_getN(2);
  @$pb.TagNumber(3)
  set notifyAt($1.Timestamp v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasNotifyAt() => $_has(2);
  @$pb.TagNumber(3)
  void clearNotifyAt() => clearField(3);
  @$pb.TagNumber(3)
  $1.Timestamp ensureNotifyAt() => $_ensure(2);
}

class DeleteTodoReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DeleteTodoReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'api.user.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'todoId')
    ..hasRequiredFields = false
  ;

  DeleteTodoReq._() : super();
  factory DeleteTodoReq({
    $core.String? todoId,
  }) {
    final _result = create();
    if (todoId != null) {
      _result.todoId = todoId;
    }
    return _result;
  }
  factory DeleteTodoReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteTodoReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteTodoReq clone() => DeleteTodoReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteTodoReq copyWith(void Function(DeleteTodoReq) updates) => super.copyWith((message) => updates(message as DeleteTodoReq)) as DeleteTodoReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteTodoReq create() => DeleteTodoReq._();
  DeleteTodoReq createEmptyInstance() => create();
  static $pb.PbList<DeleteTodoReq> createRepeated() => $pb.PbList<DeleteTodoReq>();
  @$core.pragma('dart2js:noInline')
  static DeleteTodoReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteTodoReq>(create);
  static DeleteTodoReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);
}

class GetTodoReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetTodoReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'api.user.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'todoId')
    ..hasRequiredFields = false
  ;

  GetTodoReq._() : super();
  factory GetTodoReq({
    $core.String? todoId,
  }) {
    final _result = create();
    if (todoId != null) {
      _result.todoId = todoId;
    }
    return _result;
  }
  factory GetTodoReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetTodoReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetTodoReq clone() => GetTodoReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetTodoReq copyWith(void Function(GetTodoReq) updates) => super.copyWith((message) => updates(message as GetTodoReq)) as GetTodoReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetTodoReq create() => GetTodoReq._();
  GetTodoReq createEmptyInstance() => create();
  static $pb.PbList<GetTodoReq> createRepeated() => $pb.PbList<GetTodoReq>();
  @$core.pragma('dart2js:noInline')
  static GetTodoReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetTodoReq>(create);
  static GetTodoReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);
}

class TodoResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'TodoResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'api.user.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'todoId')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'todo')
    ..aOM<$1.Timestamp>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'dueAt', subBuilder: $1.Timestamp.create)
    ..aOM<$1.Timestamp>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'notifyAt', subBuilder: $1.Timestamp.create)
    ..aOM<$1.Timestamp>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'createdAt', subBuilder: $1.Timestamp.create)
    ..aOM<$1.Timestamp>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'updatedAt', subBuilder: $1.Timestamp.create)
    ..hasRequiredFields = false
  ;

  TodoResp._() : super();
  factory TodoResp({
    $core.String? todoId,
    $core.String? todo,
    $1.Timestamp? dueAt,
    $1.Timestamp? notifyAt,
    $1.Timestamp? createdAt,
    $1.Timestamp? updatedAt,
  }) {
    final _result = create();
    if (todoId != null) {
      _result.todoId = todoId;
    }
    if (todo != null) {
      _result.todo = todo;
    }
    if (dueAt != null) {
      _result.dueAt = dueAt;
    }
    if (notifyAt != null) {
      _result.notifyAt = notifyAt;
    }
    if (createdAt != null) {
      _result.createdAt = createdAt;
    }
    if (updatedAt != null) {
      _result.updatedAt = updatedAt;
    }
    return _result;
  }
  factory TodoResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory TodoResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  TodoResp clone() => TodoResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  TodoResp copyWith(void Function(TodoResp) updates) => super.copyWith((message) => updates(message as TodoResp)) as TodoResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static TodoResp create() => TodoResp._();
  TodoResp createEmptyInstance() => create();
  static $pb.PbList<TodoResp> createRepeated() => $pb.PbList<TodoResp>();
  @$core.pragma('dart2js:noInline')
  static TodoResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<TodoResp>(create);
  static TodoResp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get todoId => $_getSZ(0);
  @$pb.TagNumber(1)
  set todoId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTodoId() => $_has(0);
  @$pb.TagNumber(1)
  void clearTodoId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get todo => $_getSZ(1);
  @$pb.TagNumber(2)
  set todo($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTodo() => $_has(1);
  @$pb.TagNumber(2)
  void clearTodo() => clearField(2);

  @$pb.TagNumber(3)
  $1.Timestamp get dueAt => $_getN(2);
  @$pb.TagNumber(3)
  set dueAt($1.Timestamp v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasDueAt() => $_has(2);
  @$pb.TagNumber(3)
  void clearDueAt() => clearField(3);
  @$pb.TagNumber(3)
  $1.Timestamp ensureDueAt() => $_ensure(2);

  @$pb.TagNumber(4)
  $1.Timestamp get notifyAt => $_getN(3);
  @$pb.TagNumber(4)
  set notifyAt($1.Timestamp v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasNotifyAt() => $_has(3);
  @$pb.TagNumber(4)
  void clearNotifyAt() => clearField(4);
  @$pb.TagNumber(4)
  $1.Timestamp ensureNotifyAt() => $_ensure(3);

  @$pb.TagNumber(5)
  $1.Timestamp get createdAt => $_getN(4);
  @$pb.TagNumber(5)
  set createdAt($1.Timestamp v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasCreatedAt() => $_has(4);
  @$pb.TagNumber(5)
  void clearCreatedAt() => clearField(5);
  @$pb.TagNumber(5)
  $1.Timestamp ensureCreatedAt() => $_ensure(4);

  @$pb.TagNumber(6)
  $1.Timestamp get updatedAt => $_getN(5);
  @$pb.TagNumber(6)
  set updatedAt($1.Timestamp v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasUpdatedAt() => $_has(5);
  @$pb.TagNumber(6)
  void clearUpdatedAt() => clearField(6);
  @$pb.TagNumber(6)
  $1.Timestamp ensureUpdatedAt() => $_ensure(5);
}

class TodoSvcApi {
  $pb.RpcClient _client;
  TodoSvcApi(this._client);

  $async.Future<TodoResp> create_($pb.ClientContext? ctx, CreateTodoReq request) {
    var emptyResponse = TodoResp();
    return _client.invoke<TodoResp>(ctx, 'TodoSvc', 'Create', request, emptyResponse);
  }
  $async.Future<TodoResp> update($pb.ClientContext? ctx, CreateTodoReq request) {
    var emptyResponse = TodoResp();
    return _client.invoke<TodoResp>(ctx, 'TodoSvc', 'Update', request, emptyResponse);
  }
  $async.Future<TodoResp> get($pb.ClientContext? ctx, GetTodoReq request) {
    var emptyResponse = TodoResp();
    return _client.invoke<TodoResp>(ctx, 'TodoSvc', 'Get', request, emptyResponse);
  }
  $async.Future<$2.Empty> delete($pb.ClientContext? ctx, DeleteTodoReq request) {
    var emptyResponse = $2.Empty();
    return _client.invoke<$2.Empty>(ctx, 'TodoSvc', 'Delete', request, emptyResponse);
  }
}

