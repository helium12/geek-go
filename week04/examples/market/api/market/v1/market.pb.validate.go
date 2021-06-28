// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: market.proto

package v1

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

// Validate checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Product) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Desc

	// no validation rules for Sold

	return nil
}

// ProductValidationError is the validation error returned by Product.Validate
// if the designated constraints aren't met.
type ProductValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductValidationError) ErrorName() string { return "ProductValidationError" }

// Error satisfies the builtin error interface
func (e ProductValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProduct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductValidationError{}

// Validate checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProductRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 5 || l > 50 {
		return CreateProductRequestValidationError{
			field:  "Name",
			reason: "value length must be between 5 and 50 runes, inclusive",
		}
	}

	// no validation rules for Desc

	return nil
}

// CreateProductRequestValidationError is the validation error returned by
// CreateProductRequest.Validate if the designated constraints aren't met.
type CreateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductRequestValidationError) ErrorName() string {
	return "CreateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductRequestValidationError{}

// Validate checks the field values on CreateProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProductReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProductReplyValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateProductReplyValidationError is the validation error returned by
// CreateProductReply.Validate if the designated constraints aren't met.
type CreateProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductReplyValidationError) ErrorName() string {
	return "CreateProductReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductReplyValidationError{}

// Validate checks the field values on UpdateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateProductRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return UpdateProductRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 5 || l > 50 {
		return UpdateProductRequestValidationError{
			field:  "Name",
			reason: "value length must be between 5 and 50 runes, inclusive",
		}
	}

	// no validation rules for Desc

	return nil
}

// UpdateProductRequestValidationError is the validation error returned by
// UpdateProductRequest.Validate if the designated constraints aren't met.
type UpdateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProductRequestValidationError) ErrorName() string {
	return "UpdateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProductRequestValidationError{}

// Validate checks the field values on UpdateProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateProductReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProductReplyValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateProductReplyValidationError is the validation error returned by
// UpdateProductReply.Validate if the designated constraints aren't met.
type UpdateProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProductReplyValidationError) ErrorName() string {
	return "UpdateProductReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProductReplyValidationError{}

// Validate checks the field values on DeleteProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteProductRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteProductRequestValidationError is the validation error returned by
// DeleteProductRequest.Validate if the designated constraints aren't met.
type DeleteProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductRequestValidationError) ErrorName() string {
	return "DeleteProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductRequestValidationError{}

// Validate checks the field values on DeleteProductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteProductReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteProductReplyValidationError is the validation error returned by
// DeleteProductReply.Validate if the designated constraints aren't met.
type DeleteProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProductReplyValidationError) ErrorName() string {
	return "DeleteProductReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProductReplyValidationError{}

// Validate checks the field values on GetProductRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetProductRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetProductRequestValidationError is the validation error returned by
// GetProductRequest.Validate if the designated constraints aren't met.
type GetProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductRequestValidationError) ErrorName() string {
	return "GetProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductRequestValidationError{}

// Validate checks the field values on GetProductReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetProductReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProductReplyValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetProductReplyValidationError is the validation error returned by
// GetProductReply.Validate if the designated constraints aren't met.
type GetProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductReplyValidationError) ErrorName() string { return "GetProductReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductReplyValidationError{}

// Validate checks the field values on ListProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListProductRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListProductRequestValidationError is the validation error returned by
// ListProductRequest.Validate if the designated constraints aren't met.
type ListProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductRequestValidationError) ErrorName() string {
	return "ListProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductRequestValidationError{}

// Validate checks the field values on ListProductReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListProductReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProductReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListProductReplyValidationError is the validation error returned by
// ListProductReply.Validate if the designated constraints aren't met.
type ListProductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProductReplyValidationError) ErrorName() string { return "ListProductReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListProductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProductReplyValidationError{}