// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: distpebble/v1/txn.proto

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

// Validate checks the field values on Txn with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Txn) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Txn with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TxnMultiError, or nil if none found.
func (m *Txn) ValidateAll() error {
	return m.validate(true)
}

func (m *Txn) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TxnId

	// no validation rules for TxnType

	for idx, item := range m.GetReadSet() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TxnValidationError{
						field:  fmt.Sprintf("ReadSet[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TxnValidationError{
						field:  fmt.Sprintf("ReadSet[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TxnValidationError{
					field:  fmt.Sprintf("ReadSet[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetWriteSet() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TxnValidationError{
						field:  fmt.Sprintf("WriteSet[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TxnValidationError{
						field:  fmt.Sprintf("WriteSet[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TxnValidationError{
					field:  fmt.Sprintf("WriteSet[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetReadWriteSet() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TxnValidationError{
						field:  fmt.Sprintf("ReadWriteSet[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TxnValidationError{
						field:  fmt.Sprintf("ReadWriteSet[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TxnValidationError{
					field:  fmt.Sprintf("ReadWriteSet[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TxnMultiError(errors)
	}

	return nil
}

// TxnMultiError is an error wrapping multiple validation errors returned by
// Txn.ValidateAll() if the designated constraints aren't met.
type TxnMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TxnMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TxnMultiError) AllErrors() []error { return m }

// TxnValidationError is the validation error returned by Txn.Validate if the
// designated constraints aren't met.
type TxnValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TxnValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TxnValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TxnValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TxnValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TxnValidationError) ErrorName() string { return "TxnValidationError" }

// Error satisfies the builtin error interface
func (e TxnValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTxn.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TxnValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TxnValidationError{}

// Validate checks the field values on KeyEntry with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KeyEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KeyEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KeyEntryMultiError, or nil
// if none found.
func (m *KeyEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *KeyEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Master

	// no validation rules for Counter

	if len(errors) > 0 {
		return KeyEntryMultiError(errors)
	}

	return nil
}

// KeyEntryMultiError is an error wrapping multiple validation errors returned
// by KeyEntry.ValidateAll() if the designated constraints aren't met.
type KeyEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KeyEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KeyEntryMultiError) AllErrors() []error { return m }

// KeyEntryValidationError is the validation error returned by
// KeyEntry.Validate if the designated constraints aren't met.
type KeyEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyEntryValidationError) ErrorName() string { return "KeyEntryValidationError" }

// Error satisfies the builtin error interface
func (e KeyEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyEntryValidationError{}
