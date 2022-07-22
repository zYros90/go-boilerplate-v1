// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CreateUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateUserReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateUserReqMultiError, or
// nil if none found.
func (m *CreateUserReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 3 {
		err := CreateUserReqValidationError{
			field:  "Username",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 3 {
		err := CreateUserReqValidationError{
			field:  "Password",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = CreateUserReqValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for FirstName

	// no validation rules for LastName

	if len(errors) > 0 {
		return CreateUserReqMultiError(errors)
	}

	return nil
}

func (m *CreateUserReq) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateUserReq) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateUserReqMultiError is an error wrapping multiple validation errors
// returned by CreateUserReq.ValidateAll() if the designated constraints
// aren't met.
type CreateUserReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserReqMultiError) AllErrors() []error { return m }

// CreateUserReqValidationError is the validation error returned by
// CreateUserReq.Validate if the designated constraints aren't met.
type CreateUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReqValidationError) ErrorName() string { return "CreateUserReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReqValidationError{}

// Validate checks the field values on DeleteUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteUserReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteUserReqMultiError, or
// nil if none found.
func (m *DeleteUserReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteUserReqMultiError(errors)
	}

	return nil
}

// DeleteUserReqMultiError is an error wrapping multiple validation errors
// returned by DeleteUserReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserReqMultiError) AllErrors() []error { return m }

// DeleteUserReqValidationError is the validation error returned by
// DeleteUserReq.Validate if the designated constraints aren't met.
type DeleteUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserReqValidationError) ErrorName() string { return "DeleteUserReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserReqValidationError{}

// Validate checks the field values on GetUserReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserReqMultiError, or
// nil if none found.
func (m *GetUserReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetUserReqMultiError(errors)
	}

	return nil
}

// GetUserReqMultiError is an error wrapping multiple validation errors
// returned by GetUserReq.ValidateAll() if the designated constraints aren't met.
type GetUserReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserReqMultiError) AllErrors() []error { return m }

// GetUserReqValidationError is the validation error returned by
// GetUserReq.Validate if the designated constraints aren't met.
type GetUserReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReqValidationError) ErrorName() string { return "GetUserReqValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReqValidationError{}

// Validate checks the field values on UserResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserRespMultiError, or nil
// if none found.
func (m *UserResp) ValidateAll() error {
	return m.validate(true)
}

func (m *UserResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 3 {
		err := UserRespValidationError{
			field:  "Username",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = UserRespValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for FirstName

	// no validation rules for LastName

	if m.GetCreatedAt() == nil {
		err := UserRespValidationError{
			field:  "CreatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUpdatedAt() == nil {
		err := UserRespValidationError{
			field:  "UpdatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserRespMultiError(errors)
	}

	return nil
}

func (m *UserResp) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UserResp) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UserRespMultiError is an error wrapping multiple validation errors returned
// by UserResp.ValidateAll() if the designated constraints aren't met.
type UserRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRespMultiError) AllErrors() []error { return m }

// UserRespValidationError is the validation error returned by
// UserResp.Validate if the designated constraints aren't met.
type UserRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRespValidationError) ErrorName() string { return "UserRespValidationError" }

// Error satisfies the builtin error interface
func (e UserRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRespValidationError{}
