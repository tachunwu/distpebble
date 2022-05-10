// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: distpebble/v1/lookup.proto

package distpebblev1

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

// Validate checks the field values on LookupMasterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LookupMasterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupMasterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LookupMasterRequestMultiError, or nil if none found.
func (m *LookupMasterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupMasterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TxnId

	if len(errors) > 0 {
		return LookupMasterRequestMultiError(errors)
	}

	return nil
}

// LookupMasterRequestMultiError is an error wrapping multiple validation
// errors returned by LookupMasterRequest.ValidateAll() if the designated
// constraints aren't met.
type LookupMasterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupMasterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupMasterRequestMultiError) AllErrors() []error { return m }

// LookupMasterRequestValidationError is the validation error returned by
// LookupMasterRequest.Validate if the designated constraints aren't met.
type LookupMasterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupMasterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupMasterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupMasterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupMasterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupMasterRequestValidationError) ErrorName() string {
	return "LookupMasterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LookupMasterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupMasterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupMasterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupMasterRequestValidationError{}

// Validate checks the field values on LookupMasterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LookupMasterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LookupMasterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LookupMasterResponseMultiError, or nil if none found.
func (m *LookupMasterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LookupMasterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TxnId

	for idx, item := range m.GetKeyEntries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LookupMasterResponseValidationError{
						field:  fmt.Sprintf("KeyEntries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LookupMasterResponseValidationError{
						field:  fmt.Sprintf("KeyEntries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LookupMasterResponseValidationError{
					field:  fmt.Sprintf("KeyEntries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LookupMasterResponseMultiError(errors)
	}

	return nil
}

// LookupMasterResponseMultiError is an error wrapping multiple validation
// errors returned by LookupMasterResponse.ValidateAll() if the designated
// constraints aren't met.
type LookupMasterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LookupMasterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LookupMasterResponseMultiError) AllErrors() []error { return m }

// LookupMasterResponseValidationError is the validation error returned by
// LookupMasterResponse.Validate if the designated constraints aren't met.
type LookupMasterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupMasterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupMasterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupMasterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupMasterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupMasterResponseValidationError) ErrorName() string {
	return "LookupMasterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LookupMasterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupMasterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupMasterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupMasterResponseValidationError{}