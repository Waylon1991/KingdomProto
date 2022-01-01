// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gate/gate.proto

package gate

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

	client "github.com/Waylon1991/KingdomProto/client"

	common "github.com/Waylon1991/KingdomProto/common"
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

	_ = common.ProtocolType(0)

	_ = client.CP(0)
)

// Validate checks the field values on GateInfoRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GateInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GateInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GateInfoRequestMultiError, or nil if none found.
func (m *GateInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GateInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GateInfoRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GateInfoRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GateInfoRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Protocol

	if len(errors) > 0 {
		return GateInfoRequestMultiError(errors)
	}
	return nil
}

// GateInfoRequestMultiError is an error wrapping multiple validation errors
// returned by GateInfoRequest.ValidateAll() if the designated constraints
// aren't met.
type GateInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GateInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GateInfoRequestMultiError) AllErrors() []error { return m }

// GateInfoRequestValidationError is the validation error returned by
// GateInfoRequest.Validate if the designated constraints aren't met.
type GateInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GateInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GateInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GateInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GateInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GateInfoRequestValidationError) ErrorName() string { return "GateInfoRequestValidationError" }

// Error satisfies the builtin error interface
func (e GateInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGateInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GateInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GateInfoRequestValidationError{}

// Validate checks the field values on GateInfoReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GateInfoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GateInfoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GateInfoReplyMultiError, or
// nil if none found.
func (m *GateInfoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GateInfoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Location

	if len(errors) > 0 {
		return GateInfoReplyMultiError(errors)
	}
	return nil
}

// GateInfoReplyMultiError is an error wrapping multiple validation errors
// returned by GateInfoReply.ValidateAll() if the designated constraints
// aren't met.
type GateInfoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GateInfoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GateInfoReplyMultiError) AllErrors() []error { return m }

// GateInfoReplyValidationError is the validation error returned by
// GateInfoReply.Validate if the designated constraints aren't met.
type GateInfoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GateInfoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GateInfoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GateInfoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GateInfoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GateInfoReplyValidationError) ErrorName() string { return "GateInfoReplyValidationError" }

// Error satisfies the builtin error interface
func (e GateInfoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGateInfoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GateInfoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GateInfoReplyValidationError{}

// Validate checks the field values on BroadcastRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BroadcastRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BroadcastRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BroadcastRequestMultiError, or nil if none found.
func (m *BroadcastRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BroadcastRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pid

	// no validation rules for Msg

	if len(errors) > 0 {
		return BroadcastRequestMultiError(errors)
	}
	return nil
}

// BroadcastRequestMultiError is an error wrapping multiple validation errors
// returned by BroadcastRequest.ValidateAll() if the designated constraints
// aren't met.
type BroadcastRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BroadcastRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BroadcastRequestMultiError) AllErrors() []error { return m }

// BroadcastRequestValidationError is the validation error returned by
// BroadcastRequest.Validate if the designated constraints aren't met.
type BroadcastRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BroadcastRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BroadcastRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BroadcastRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BroadcastRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BroadcastRequestValidationError) ErrorName() string { return "BroadcastRequestValidationError" }

// Error satisfies the builtin error interface
func (e BroadcastRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBroadcastRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BroadcastRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BroadcastRequestValidationError{}

// Validate checks the field values on BroadcastReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BroadcastReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BroadcastReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BroadcastReplyMultiError,
// or nil if none found.
func (m *BroadcastReply) ValidateAll() error {
	return m.validate(true)
}

func (m *BroadcastReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BroadcastReplyMultiError(errors)
	}
	return nil
}

// BroadcastReplyMultiError is an error wrapping multiple validation errors
// returned by BroadcastReply.ValidateAll() if the designated constraints
// aren't met.
type BroadcastReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BroadcastReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BroadcastReplyMultiError) AllErrors() []error { return m }

// BroadcastReplyValidationError is the validation error returned by
// BroadcastReply.Validate if the designated constraints aren't met.
type BroadcastReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BroadcastReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BroadcastReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BroadcastReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BroadcastReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BroadcastReplyValidationError) ErrorName() string { return "BroadcastReplyValidationError" }

// Error satisfies the builtin error interface
func (e BroadcastReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBroadcastReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BroadcastReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BroadcastReplyValidationError{}