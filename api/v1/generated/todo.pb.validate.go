// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: todo.proto

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

// Validate checks the field values on CreateTodoReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateTodoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTodoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateTodoReqMultiError, or
// nil if none found.
func (m *CreateTodoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTodoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTodo()) < 1 {
		err := CreateTodoReqValidationError{
			field:  "Todo",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDueAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTodoReqValidationError{
					field:  "DueAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTodoReqValidationError{
					field:  "DueAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDueAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTodoReqValidationError{
				field:  "DueAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNotifyAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTodoReqValidationError{
					field:  "NotifyAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTodoReqValidationError{
					field:  "NotifyAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNotifyAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTodoReqValidationError{
				field:  "NotifyAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTodoReqMultiError(errors)
	}

	return nil
}

// CreateTodoReqMultiError is an error wrapping multiple validation errors
// returned by CreateTodoReq.ValidateAll() if the designated constraints
// aren't met.
type CreateTodoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTodoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTodoReqMultiError) AllErrors() []error { return m }

// CreateTodoReqValidationError is the validation error returned by
// CreateTodoReq.Validate if the designated constraints aren't met.
type CreateTodoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTodoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTodoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTodoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTodoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTodoReqValidationError) ErrorName() string { return "CreateTodoReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateTodoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTodoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTodoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTodoReqValidationError{}

// Validate checks the field values on DeleteTodoReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteTodoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTodoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteTodoReqMultiError, or
// nil if none found.
func (m *DeleteTodoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTodoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTodoId()) < 1 {
		err := DeleteTodoReqValidationError{
			field:  "TodoId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteTodoReqMultiError(errors)
	}

	return nil
}

// DeleteTodoReqMultiError is an error wrapping multiple validation errors
// returned by DeleteTodoReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteTodoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTodoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTodoReqMultiError) AllErrors() []error { return m }

// DeleteTodoReqValidationError is the validation error returned by
// DeleteTodoReq.Validate if the designated constraints aren't met.
type DeleteTodoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTodoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTodoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTodoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTodoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTodoReqValidationError) ErrorName() string { return "DeleteTodoReqValidationError" }

// Error satisfies the builtin error interface
func (e DeleteTodoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTodoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTodoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTodoReqValidationError{}

// Validate checks the field values on GetTodoReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTodoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTodoReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTodoReqMultiError, or
// nil if none found.
func (m *GetTodoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTodoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetTodoReqMultiError(errors)
	}

	return nil
}

// GetTodoReqMultiError is an error wrapping multiple validation errors
// returned by GetTodoReq.ValidateAll() if the designated constraints aren't met.
type GetTodoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTodoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTodoReqMultiError) AllErrors() []error { return m }

// GetTodoReqValidationError is the validation error returned by
// GetTodoReq.Validate if the designated constraints aren't met.
type GetTodoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTodoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTodoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTodoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTodoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTodoReqValidationError) ErrorName() string { return "GetTodoReqValidationError" }

// Error satisfies the builtin error interface
func (e GetTodoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTodoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTodoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTodoReqValidationError{}

// Validate checks the field values on TodoResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TodoResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TodoResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TodoRespMultiError, or nil
// if none found.
func (m *TodoResp) ValidateAll() error {
	return m.validate(true)
}

func (m *TodoResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTodoId()) < 1 {
		err := TodoRespValidationError{
			field:  "TodoId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTodo()) < 1 {
		err := TodoRespValidationError{
			field:  "Todo",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDueAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TodoRespValidationError{
					field:  "DueAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TodoRespValidationError{
					field:  "DueAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDueAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TodoRespValidationError{
				field:  "DueAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNotifyAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TodoRespValidationError{
					field:  "NotifyAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TodoRespValidationError{
					field:  "NotifyAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNotifyAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TodoRespValidationError{
				field:  "NotifyAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetCreatedAt() == nil {
		err := TodoRespValidationError{
			field:  "CreatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUpdatedAt() == nil {
		err := TodoRespValidationError{
			field:  "UpdatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TodoRespMultiError(errors)
	}

	return nil
}

// TodoRespMultiError is an error wrapping multiple validation errors returned
// by TodoResp.ValidateAll() if the designated constraints aren't met.
type TodoRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TodoRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TodoRespMultiError) AllErrors() []error { return m }

// TodoRespValidationError is the validation error returned by
// TodoResp.Validate if the designated constraints aren't met.
type TodoRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TodoRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TodoRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TodoRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TodoRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TodoRespValidationError) ErrorName() string { return "TodoRespValidationError" }

// Error satisfies the builtin error interface
func (e TodoRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTodoResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TodoRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TodoRespValidationError{}
