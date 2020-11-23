// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chaos/serverexperimentation/v1/serverexperimentation.proto

package serverexperimentationv1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _serverexperimentation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TestConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TestConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetClusterPair() == nil {
		return TestConfigValidationError{
			field:  "ClusterPair",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetClusterPair()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TestConfigValidationError{
				field:  "ClusterPair",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Fault.(type) {

	case *TestConfig_Abort:

		if v, ok := interface{}(m.GetAbort()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TestConfigValidationError{
					field:  "Abort",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TestConfig_Latency:

		if v, ok := interface{}(m.GetLatency()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TestConfigValidationError{
					field:  "Latency",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return TestConfigValidationError{
			field:  "Fault",
			reason: "value is required",
		}

	}

	return nil
}

// TestConfigValidationError is the validation error returned by
// TestConfig.Validate if the designated constraints aren't met.
type TestConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestConfigValidationError) ErrorName() string { return "TestConfigValidationError" }

// Error satisfies the builtin error interface
func (e TestConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestConfigValidationError{}

// Validate checks the field values on ClusterPairTarget with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ClusterPairTarget) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetDownstreamCluster()) < 1 {
		return ClusterPairTargetValidationError{
			field:  "DownstreamCluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetUpstreamCluster()) < 1 {
		return ClusterPairTargetValidationError{
			field:  "UpstreamCluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	if _, ok := FaultInjectionCluster_name[int32(m.GetFaultInjectionCluster())]; !ok {
		return ClusterPairTargetValidationError{
			field:  "FaultInjectionCluster",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// ClusterPairTargetValidationError is the validation error returned by
// ClusterPairTarget.Validate if the designated constraints aren't met.
type ClusterPairTargetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterPairTargetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterPairTargetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterPairTargetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterPairTargetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterPairTargetValidationError) ErrorName() string {
	return "ClusterPairTargetValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterPairTargetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterPairTarget.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterPairTargetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterPairTargetValidationError{}

// Validate checks the field values on AbortFaultConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AbortFaultConfig) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetPercent(); val <= 0 || val > 100 {
		return AbortFaultConfigValidationError{
			field:  "Percent",
			reason: "value must be inside range (0, 100]",
		}
	}

	if val := m.GetHttpStatus(); val <= 99 || val >= 600 {
		return AbortFaultConfigValidationError{
			field:  "HttpStatus",
			reason: "value must be inside range (99, 600)",
		}
	}

	return nil
}

// AbortFaultConfigValidationError is the validation error returned by
// AbortFaultConfig.Validate if the designated constraints aren't met.
type AbortFaultConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AbortFaultConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AbortFaultConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AbortFaultConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AbortFaultConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AbortFaultConfigValidationError) ErrorName() string { return "AbortFaultConfigValidationError" }

// Error satisfies the builtin error interface
func (e AbortFaultConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAbortFaultConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AbortFaultConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AbortFaultConfigValidationError{}

// Validate checks the field values on LatencyFaultConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LatencyFaultConfig) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetPercent(); val <= 0 || val > 100 {
		return LatencyFaultConfigValidationError{
			field:  "Percent",
			reason: "value must be inside range (0, 100]",
		}
	}

	if m.GetDurationMs() <= 0 {
		return LatencyFaultConfigValidationError{
			field:  "DurationMs",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// LatencyFaultConfigValidationError is the validation error returned by
// LatencyFaultConfig.Validate if the designated constraints aren't met.
type LatencyFaultConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LatencyFaultConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LatencyFaultConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LatencyFaultConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LatencyFaultConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LatencyFaultConfigValidationError) ErrorName() string {
	return "LatencyFaultConfigValidationError"
}

// Error satisfies the builtin error interface
func (e LatencyFaultConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLatencyFaultConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LatencyFaultConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LatencyFaultConfigValidationError{}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Fault.(type) {

	case *Config_Abort:

		if v, ok := interface{}(m.GetAbort()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  "Abort",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Config_Latency:

		if v, ok := interface{}(m.GetLatency()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  "Latency",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ConfigValidationError{
			field:  "Fault",
			reason: "value is required",
		}

	}

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on Abort with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Abort) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetFaultTargeting() == nil {
		return AbortValidationError{
			field:  "FaultTargeting",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFaultTargeting()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AbortValidationError{
				field:  "FaultTargeting",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetFaultParcentage() == nil {
		return AbortValidationError{
			field:  "FaultParcentage",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFaultParcentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AbortValidationError{
				field:  "FaultParcentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetAbortStatus() == nil {
		return AbortValidationError{
			field:  "AbortStatus",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAbortStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AbortValidationError{
				field:  "AbortStatus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AbortValidationError is the validation error returned by Abort.Validate if
// the designated constraints aren't met.
type AbortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AbortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AbortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AbortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AbortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AbortValidationError) ErrorName() string { return "AbortValidationError" }

// Error satisfies the builtin error interface
func (e AbortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAbort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AbortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AbortValidationError{}

// Validate checks the field values on Latency with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Latency) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetFaultTargeting() == nil {
		return LatencyValidationError{
			field:  "FaultTargeting",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFaultTargeting()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LatencyValidationError{
				field:  "FaultTargeting",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetFaultParcentage() == nil {
		return LatencyValidationError{
			field:  "FaultParcentage",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFaultParcentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LatencyValidationError{
				field:  "FaultParcentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetLatencyDuration() == nil {
		return LatencyValidationError{
			field:  "LatencyDuration",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetLatencyDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LatencyValidationError{
				field:  "LatencyDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LatencyValidationError is the validation error returned by Latency.Validate
// if the designated constraints aren't met.
type LatencyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LatencyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LatencyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LatencyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LatencyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LatencyValidationError) ErrorName() string { return "LatencyValidationError" }

// Error satisfies the builtin error interface
func (e LatencyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLatency.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LatencyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LatencyValidationError{}

// Validate checks the field values on FaultPercentage with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FaultPercentage) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetPercentage(); val <= 0 || val > 100 {
		return FaultPercentageValidationError{
			field:  "Percentage",
			reason: "value must be inside range (0, 100]",
		}
	}

	return nil
}

// FaultPercentageValidationError is the validation error returned by
// FaultPercentage.Validate if the designated constraints aren't met.
type FaultPercentageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultPercentageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultPercentageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultPercentageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultPercentageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultPercentageValidationError) ErrorName() string { return "FaultPercentageValidationError" }

// Error satisfies the builtin error interface
func (e FaultPercentageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultPercentage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultPercentageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultPercentageValidationError{}

// Validate checks the field values on AbortStatus with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AbortStatus) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetHttpStatusCode(); val <= 99 || val >= 600 {
		return AbortStatusValidationError{
			field:  "HttpStatusCode",
			reason: "value must be inside range (99, 600)",
		}
	}

	return nil
}

// AbortStatusValidationError is the validation error returned by
// AbortStatus.Validate if the designated constraints aren't met.
type AbortStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AbortStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AbortStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AbortStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AbortStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AbortStatusValidationError) ErrorName() string { return "AbortStatusValidationError" }

// Error satisfies the builtin error interface
func (e AbortStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAbortStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AbortStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AbortStatusValidationError{}

// Validate checks the field values on LatencyDuration with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LatencyDuration) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDurationMs() <= 0 {
		return LatencyDurationValidationError{
			field:  "DurationMs",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// LatencyDurationValidationError is the validation error returned by
// LatencyDuration.Validate if the designated constraints aren't met.
type LatencyDurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LatencyDurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LatencyDurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LatencyDurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LatencyDurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LatencyDurationValidationError) ErrorName() string { return "LatencyDurationValidationError" }

// Error satisfies the builtin error interface
func (e LatencyDurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLatencyDuration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LatencyDurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LatencyDurationValidationError{}

// Validate checks the field values on FaultTargeting with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FaultTargeting) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Enforcer.(type) {

	case *FaultTargeting_Upstream:

		if v, ok := interface{}(m.GetUpstream()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultTargetingValidationError{
					field:  "Upstream",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *FaultTargeting_Downstream:

		if v, ok := interface{}(m.GetDownstream()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultTargetingValidationError{
					field:  "Downstream",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return FaultTargetingValidationError{
			field:  "Enforcer",
			reason: "value is required",
		}

	}

	return nil
}

// FaultTargetingValidationError is the validation error returned by
// FaultTargeting.Validate if the designated constraints aren't met.
type FaultTargetingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultTargetingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultTargetingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultTargetingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultTargetingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultTargetingValidationError) ErrorName() string { return "FaultTargetingValidationError" }

// Error satisfies the builtin error interface
func (e FaultTargetingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultTargeting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultTargetingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultTargetingValidationError{}

// Validate checks the field values on UpstreamEnforcing with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpstreamEnforcing) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Upstream.(type) {

	case *UpstreamEnforcing_UpstreamCluster:

		if v, ok := interface{}(m.GetUpstreamCluster()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpstreamEnforcingValidationError{
					field:  "UpstreamCluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *UpstreamEnforcing_UpstreamPartialSingleCluster:

		if v, ok := interface{}(m.GetUpstreamPartialSingleCluster()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpstreamEnforcingValidationError{
					field:  "UpstreamPartialSingleCluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return UpstreamEnforcingValidationError{
			field:  "Upstream",
			reason: "value is required",
		}

	}

	switch m.Downstream.(type) {

	case *UpstreamEnforcing_DownstreamCluster:

		if v, ok := interface{}(m.GetDownstreamCluster()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpstreamEnforcingValidationError{
					field:  "DownstreamCluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return UpstreamEnforcingValidationError{
			field:  "Downstream",
			reason: "value is required",
		}

	}

	return nil
}

// UpstreamEnforcingValidationError is the validation error returned by
// UpstreamEnforcing.Validate if the designated constraints aren't met.
type UpstreamEnforcingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamEnforcingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamEnforcingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamEnforcingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamEnforcingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamEnforcingValidationError) ErrorName() string {
	return "UpstreamEnforcingValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamEnforcingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamEnforcing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamEnforcingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamEnforcingValidationError{}

// Validate checks the field values on DownstreamEnforcing with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DownstreamEnforcing) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Upstream.(type) {

	case *DownstreamEnforcing_UpstreamCluster:

		if v, ok := interface{}(m.GetUpstreamCluster()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DownstreamEnforcingValidationError{
					field:  "UpstreamCluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return DownstreamEnforcingValidationError{
			field:  "Upstream",
			reason: "value is required",
		}

	}

	switch m.Downstream.(type) {

	case *DownstreamEnforcing_DownstreamCluster:

		if v, ok := interface{}(m.GetDownstreamCluster()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DownstreamEnforcingValidationError{
					field:  "DownstreamCluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return DownstreamEnforcingValidationError{
			field:  "Downstream",
			reason: "value is required",
		}

	}

	return nil
}

// DownstreamEnforcingValidationError is the validation error returned by
// DownstreamEnforcing.Validate if the designated constraints aren't met.
type DownstreamEnforcingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownstreamEnforcingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownstreamEnforcingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownstreamEnforcingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownstreamEnforcingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownstreamEnforcingValidationError) ErrorName() string {
	return "DownstreamEnforcingValidationError"
}

// Error satisfies the builtin error interface
func (e DownstreamEnforcingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownstreamEnforcing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownstreamEnforcingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownstreamEnforcingValidationError{}

// Validate checks the field values on SingleCluster with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SingleCluster) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return SingleClusterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// SingleClusterValidationError is the validation error returned by
// SingleCluster.Validate if the designated constraints aren't met.
type SingleClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SingleClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SingleClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SingleClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SingleClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SingleClusterValidationError) ErrorName() string { return "SingleClusterValidationError" }

// Error satisfies the builtin error interface
func (e SingleClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSingleCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SingleClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SingleClusterValidationError{}

// Validate checks the field values on PartialSingleCluster with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PartialSingleCluster) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return PartialSingleClusterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	if val := m.GetPercentage(); val <= 0 || val > 100 {
		return PartialSingleClusterValidationError{
			field:  "Percentage",
			reason: "value must be inside range (0, 100]",
		}
	}

	return nil
}

// PartialSingleClusterValidationError is the validation error returned by
// PartialSingleCluster.Validate if the designated constraints aren't met.
type PartialSingleClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PartialSingleClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PartialSingleClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PartialSingleClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PartialSingleClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PartialSingleClusterValidationError) ErrorName() string {
	return "PartialSingleClusterValidationError"
}

// Error satisfies the builtin error interface
func (e PartialSingleClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPartialSingleCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PartialSingleClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PartialSingleClusterValidationError{}
