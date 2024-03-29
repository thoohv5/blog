// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/wechat/v1/base/base.proto

package base

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on QRCodeReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *QRCodeReq) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetExpireSeconds(); val <= 0 || val > 2592000 {
		return QRCodeReqValidationError{
			field:  "ExpireSeconds",
			reason: "value must be inside range (0, 2592000]",
		}
	}

	// no validation rules for ActionName

	if val := m.GetSceneId(); val < 0 || val > 100000 {
		return QRCodeReqValidationError{
			field:  "SceneId",
			reason: "value must be inside range [0, 100000]",
		}
	}

	if l := utf8.RuneCountInString(m.GetSceneStr()); l < 0 || l > 64 {
		return QRCodeReqValidationError{
			field:  "SceneStr",
			reason: "value length must be between 0 and 64 runes, inclusive",
		}
	}

	return nil
}

// QRCodeReqValidationError is the validation error returned by
// QRCodeReq.Validate if the designated constraints aren't met.
type QRCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QRCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QRCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QRCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QRCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QRCodeReqValidationError) ErrorName() string { return "QRCodeReqValidationError" }

// Error satisfies the builtin error interface
func (e QRCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQRCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QRCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QRCodeReqValidationError{}

// Validate checks the field values on QRCodeResp with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *QRCodeResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for QrCode

	return nil
}

// QRCodeRespValidationError is the validation error returned by
// QRCodeResp.Validate if the designated constraints aren't met.
type QRCodeRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QRCodeRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QRCodeRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QRCodeRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QRCodeRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QRCodeRespValidationError) ErrorName() string { return "QRCodeRespValidationError" }

// Error satisfies the builtin error interface
func (e QRCodeRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQRCodeResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QRCodeRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QRCodeRespValidationError{}

// Validate checks the field values on CheckQRCodeReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckQRCodeReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Key

	return nil
}

// CheckQRCodeReqValidationError is the validation error returned by
// CheckQRCodeReq.Validate if the designated constraints aren't met.
type CheckQRCodeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckQRCodeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckQRCodeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckQRCodeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckQRCodeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckQRCodeReqValidationError) ErrorName() string { return "CheckQRCodeReqValidationError" }

// Error satisfies the builtin error interface
func (e CheckQRCodeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckQRCodeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckQRCodeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckQRCodeReqValidationError{}

// Validate checks the field values on CheckQRCodeResp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CheckQRCodeResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	// no validation rules for Extra

	return nil
}

// CheckQRCodeRespValidationError is the validation error returned by
// CheckQRCodeResp.Validate if the designated constraints aren't met.
type CheckQRCodeRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckQRCodeRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckQRCodeRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckQRCodeRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckQRCodeRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckQRCodeRespValidationError) ErrorName() string { return "CheckQRCodeRespValidationError" }

// Error satisfies the builtin error interface
func (e CheckQRCodeRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckQRCodeResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckQRCodeRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckQRCodeRespValidationError{}
