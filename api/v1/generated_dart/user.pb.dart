///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'google/protobuf/timestamp.pb.dart' as $1;
import 'google/protobuf/empty.pb.dart' as $2;

class CreateUserReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CreateUserReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'api.user.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'username')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'password')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'email')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'firstName')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lastName')
    ..hasRequiredFields = false
  ;

  CreateUserReq._() : super();
  factory CreateUserReq({
    $core.String? username,
    $core.String? password,
    $core.String? email,
    $core.String? firstName,
    $core.String? lastName,
  }) {
    final _result = create();
    if (username != null) {
      _result.username = username;
    }
    if (password != null) {
      _result.password = password;
    }
    if (email != null) {
      _result.email = email;
    }
    if (firstName != null) {
      _result.firstName = firstName;
    }
    if (lastName != null) {
      _result.lastName = lastName;
    }
    return _result;
  }
  factory CreateUserReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateUserReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateUserReq clone() => CreateUserReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateUserReq copyWith(void Function(CreateUserReq) updates) => super.copyWith((message) => updates(message as CreateUserReq)) as CreateUserReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CreateUserReq create() => CreateUserReq._();
  CreateUserReq createEmptyInstance() => create();
  static $pb.PbList<CreateUserReq> createRepeated() => $pb.PbList<CreateUserReq>();
  @$core.pragma('dart2js:noInline')
  static CreateUserReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateUserReq>(create);
  static CreateUserReq? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get username => $_getSZ(0);
  @$pb.TagNumber(1)
  set username($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUsername() => $_has(0);
  @$pb.TagNumber(1)
  void clearUsername() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get password => $_getSZ(1);
  @$pb.TagNumber(2)
  set password($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPassword() => $_has(1);
  @$pb.TagNumber(2)
  void clearPassword() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get email => $_getSZ(2);
  @$pb.TagNumber(3)
  set email($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasEmail() => $_has(2);
  @$pb.TagNumber(3)
  void clearEmail() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get firstName => $_getSZ(3);
  @$pb.TagNumber(4)
  set firstName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFirstName() => $_has(3);
  @$pb.TagNumber(4)
  void clearFirstName() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get lastName => $_getSZ(4);
  @$pb.TagNumber(5)
  set lastName($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLastName() => $_has(4);
  @$pb.TagNumber(5)
  void clearLastName() => clearField(5);
}

class DeleteUserReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DeleteUserReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'api.user.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  DeleteUserReq._() : super();
  factory DeleteUserReq() => create();
  factory DeleteUserReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteUserReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteUserReq clone() => DeleteUserReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteUserReq copyWith(void Function(DeleteUserReq) updates) => super.copyWith((message) => updates(message as DeleteUserReq)) as DeleteUserReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteUserReq create() => DeleteUserReq._();
  DeleteUserReq createEmptyInstance() => create();
  static $pb.PbList<DeleteUserReq> createRepeated() => $pb.PbList<DeleteUserReq>();
  @$core.pragma('dart2js:noInline')
  static DeleteUserReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteUserReq>(create);
  static DeleteUserReq? _defaultInstance;
}

class GetUserReq extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetUserReq', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'api.user.v1'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  GetUserReq._() : super();
  factory GetUserReq() => create();
  factory GetUserReq.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetUserReq.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetUserReq clone() => GetUserReq()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetUserReq copyWith(void Function(GetUserReq) updates) => super.copyWith((message) => updates(message as GetUserReq)) as GetUserReq; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetUserReq create() => GetUserReq._();
  GetUserReq createEmptyInstance() => create();
  static $pb.PbList<GetUserReq> createRepeated() => $pb.PbList<GetUserReq>();
  @$core.pragma('dart2js:noInline')
  static GetUserReq getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetUserReq>(create);
  static GetUserReq? _defaultInstance;
}

class UserResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserResp', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'api.user.v1'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'username')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'email')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'firstName')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'lastName')
    ..aOM<$1.Timestamp>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'createdAt', subBuilder: $1.Timestamp.create)
    ..aOM<$1.Timestamp>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'updatedAt', subBuilder: $1.Timestamp.create)
    ..hasRequiredFields = false
  ;

  UserResp._() : super();
  factory UserResp({
    $core.String? username,
    $core.String? email,
    $core.String? firstName,
    $core.String? lastName,
    $1.Timestamp? createdAt,
    $1.Timestamp? updatedAt,
  }) {
    final _result = create();
    if (username != null) {
      _result.username = username;
    }
    if (email != null) {
      _result.email = email;
    }
    if (firstName != null) {
      _result.firstName = firstName;
    }
    if (lastName != null) {
      _result.lastName = lastName;
    }
    if (createdAt != null) {
      _result.createdAt = createdAt;
    }
    if (updatedAt != null) {
      _result.updatedAt = updatedAt;
    }
    return _result;
  }
  factory UserResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserResp clone() => UserResp()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserResp copyWith(void Function(UserResp) updates) => super.copyWith((message) => updates(message as UserResp)) as UserResp; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserResp create() => UserResp._();
  UserResp createEmptyInstance() => create();
  static $pb.PbList<UserResp> createRepeated() => $pb.PbList<UserResp>();
  @$core.pragma('dart2js:noInline')
  static UserResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserResp>(create);
  static UserResp? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get username => $_getSZ(0);
  @$pb.TagNumber(1)
  set username($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUsername() => $_has(0);
  @$pb.TagNumber(1)
  void clearUsername() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get email => $_getSZ(1);
  @$pb.TagNumber(2)
  set email($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasEmail() => $_has(1);
  @$pb.TagNumber(2)
  void clearEmail() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get firstName => $_getSZ(2);
  @$pb.TagNumber(3)
  set firstName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFirstName() => $_has(2);
  @$pb.TagNumber(3)
  void clearFirstName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get lastName => $_getSZ(3);
  @$pb.TagNumber(4)
  set lastName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasLastName() => $_has(3);
  @$pb.TagNumber(4)
  void clearLastName() => clearField(4);

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

class UserSvcApi {
  $pb.RpcClient _client;
  UserSvcApi(this._client);

  $async.Future<UserResp> create_($pb.ClientContext? ctx, CreateUserReq request) {
    var emptyResponse = UserResp();
    return _client.invoke<UserResp>(ctx, 'UserSvc', 'Create', request, emptyResponse);
  }
  $async.Future<UserResp> update($pb.ClientContext? ctx, CreateUserReq request) {
    var emptyResponse = UserResp();
    return _client.invoke<UserResp>(ctx, 'UserSvc', 'Update', request, emptyResponse);
  }
  $async.Future<UserResp> get($pb.ClientContext? ctx, GetUserReq request) {
    var emptyResponse = UserResp();
    return _client.invoke<UserResp>(ctx, 'UserSvc', 'Get', request, emptyResponse);
  }
  $async.Future<$2.Empty> delete($pb.ClientContext? ctx, DeleteUserReq request) {
    var emptyResponse = $2.Empty();
    return _client.invoke<$2.Empty>(ctx, 'UserSvc', 'Delete', request, emptyResponse);
  }
}

