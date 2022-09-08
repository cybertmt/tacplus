/*
 Copyright (c) Facebook, Inc. and its affiliates.

 This source code is licensed under the MIT license found in the
 LICENSE file in the root directory of this source tree.
*/

// Code generated by codegen.go. Any changes will be overwritten.
// To regenerate run: go generate.

package tacquito

import (
	"fmt"
	"strings"
)

// AcctReplyStatus is the status of the accounting action.
type AcctReplyStatus uint8

const (
	// AcctReplyStatusSuccess per rfc
	AcctReplyStatusSuccess AcctReplyStatus = 0x01
	// AcctReplyStatusError per rfc
	AcctReplyStatusError AcctReplyStatus = 0x02
)

// Validate characterics of type based on rfc and usage.
func (t AcctReplyStatus) Validate(condition interface{}) error {
	switch t {
	case AcctReplyStatusSuccess, AcctReplyStatusError:
		return nil
	}
	return fmt.Errorf("unknown AcctReplyStatus value [%v]", t)
}

// Len returns the length of AcctReplyStatus.
func (t AcctReplyStatus) Len() int {
	return 1
}

// String returns AcctReplyStatus as a string.
func (t AcctReplyStatus) String() string {
	switch t {
	case AcctReplyStatusSuccess:
		return "AcctReplyStatusSuccess"
	case AcctReplyStatusError:
		return "AcctReplyStatusError"
	}
	return fmt.Sprintf("unknown AcctReplyStatus[%d]", uint8(t))
}

// AcctServerMsg is a string that may be presented to the user.  The
// server_msg_len indicates the length of the server_msg field, in
// bytes.  For details of text encoding, see "Treatment of Text
// Strings"
type AcctServerMsg string

// Validate characterics of type based on rfc and usage.
func (t AcctServerMsg) Validate(condition interface{}) error {
	// https://datatracker.ietf.org/doc/html/rfc8907#section-3.6
	if isAllASCII(string(t)) {
		return nil
	}
	return fmt.Errorf("AcctServerMsg is not all ascii, but it must be, [%v]", t)
}

// Len returns the length of AcctServerMsg.
func (t AcctServerMsg) Len() int {
	return len(t)
}

// String returns AcctServerMsg as a string.
func (t AcctServerMsg) String() string {
	return string(t)
}

// AcctData is a string that may be presented on an administrative
// display, console, or log.  The decision to present this message is
// client specific.  The data_len indicates the length of the data
// field, in bytes.  For details of text encoding, see "Treatment of
// Text Strings"
type AcctData string

// Validate characterics of type based on rfc and usage.
func (t AcctData) Validate(condition interface{}) error {
	// https://datatracker.ietf.org/doc/html/rfc8907#section-3.6
	if isAllASCII(string(t)) {
		return nil
	}
	return fmt.Errorf("AcctData is not all ascii, but it must be, [%v]", t)
}

// Len returns the length of AcctData.
func (t AcctData) Len() int {
	return len(t)
}

// String returns AcctData as a string.
func (t AcctData) String() string {
	return string(t)
}

// AcctStartTime The time the action started (in seconds since the epoch).
type AcctStartTime int

// Validate characterics of type based on rfc and usage.
func (t AcctStartTime) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctStartTime.
func (t AcctStartTime) Len() int {
	return 1
}

// String returns AcctStartTime as a string.
func (t AcctStartTime) String() string {
	return fmt.Sprint(int(t))
}

// AcctStopTime The time the action stopped (in seconds since the epoch).
type AcctStopTime int

// Validate characterics of type based on rfc and usage.
func (t AcctStopTime) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctStopTime.
func (t AcctStopTime) Len() int {
	return 1
}

// String returns AcctStopTime as a string.
func (t AcctStopTime) String() string {
	return fmt.Sprint(int(t))
}

// AcctTaskID - Start and stop records for the same event MUST have matching
// task_id argument values. The client MUST ensure that active
// task_ids are not duplicated; a client MUST NOT reuse a task_id in
// a start record until it has sent a stop record for that task_id.
// Servers MUST NOT make assumptions about the format of a task_id.
type AcctTaskID string

// Validate characterics of type based on rfc and usage.
func (t AcctTaskID) Validate(condition interface{}) error {
	// https://datatracker.ietf.org/doc/html/rfc8907#section-3.6
	if isAllASCII(string(t)) {
		return nil
	}
	return fmt.Errorf("AcctTaskID is not all ascii, but it must be, [%v]", t)
}

// Len returns the length of AcctTaskID.
func (t AcctTaskID) Len() int {
	return len(t)
}

// String returns AcctTaskID as a string.
func (t AcctTaskID) String() string {
	return string(t)
}

// AcctTimezone The time zone abbreviation for all timestamps included in this packet
type AcctTimezone string

// Validate characterics of type based on rfc and usage.
func (t AcctTimezone) Validate(condition interface{}) error {
	// https://datatracker.ietf.org/doc/html/rfc8907#section-3.6
	if isAllASCII(string(t)) {
		return nil
	}
	return fmt.Errorf("AcctTimezone is not all ascii, but it must be, [%v]", t)
}

// Len returns the length of AcctTimezone.
func (t AcctTimezone) Len() int {
	return len(t)
}

// String returns AcctTimezone as a string.
func (t AcctTimezone) String() string {
	return string(t)
}

// AcctEvent is Used only when "service=system". Current values are "net_acct",
// "cmd_acct", "conn_acct", "shell_acct", "sys_acct", and
// "clock_change". These indicate system-level changes. The flags
// field SHOULD indicate whether the service started or stopped.
type AcctEvent string

// Validate characterics of type based on rfc and usage.
func (t AcctEvent) Validate(condition interface{}) error {
	// https://datatracker.ietf.org/doc/html/rfc8907#section-3.6
	if isAllASCII(string(t)) {
		return nil
	}
	return fmt.Errorf("AcctEvent is not all ascii, but it must be, [%v]", t)
}

// Len returns the length of AcctEvent.
func (t AcctEvent) Len() int {
	return len(t)
}

// String returns AcctEvent as a string.
func (t AcctEvent) String() string {
	return string(t)
}

// AcctReason Accompanies an event argument. It describes why the event occurred.
type AcctReason string

// Validate characterics of type based on rfc and usage.
func (t AcctReason) Validate(condition interface{}) error {
	// https://datatracker.ietf.org/doc/html/rfc8907#section-3.6
	if isAllASCII(string(t)) {
		return nil
	}
	return fmt.Errorf("AcctReason is not all ascii, but it must be, [%v]", t)
}

// Len returns the length of AcctReason.
func (t AcctReason) Len() int {
	return len(t)
}

// String returns AcctReason as a string.
func (t AcctReason) String() string {
	return string(t)
}

// AcctErrMsg A string describing the status of the action. For details of text
// encoding, see "Treatment of Text Strings" (https://datatracker.ietf.org/doc/html/rfc8907#section-3.7).
type AcctErrMsg string

// Validate characterics of type based on rfc and usage.
func (t AcctErrMsg) Validate(condition interface{}) error {
	// https://datatracker.ietf.org/doc/html/rfc8907#section-3.6
	if isAllASCII(string(t)) {
		return nil
	}
	return fmt.Errorf("AcctErrMsg is not all ascii, but it must be, [%v]", t)
}

// Len returns the length of AcctErrMsg.
func (t AcctErrMsg) Len() int {
	return len(t)
}

// String returns AcctErrMsg as a string.
func (t AcctErrMsg) String() string {
	return string(t)
}

// AcctElapsedTime The elapsed time in seconds for the action.
type AcctElapsedTime int

// Validate characterics of type based on rfc and usage.
func (t AcctElapsedTime) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctElapsedTime.
func (t AcctElapsedTime) Len() int {
	return 1
}

// String returns AcctElapsedTime as a string.
func (t AcctElapsedTime) String() string {
	return fmt.Sprint(int(t))
}

// AcctBytes The number of bytes transferred by this action.
type AcctBytes int

// Validate characterics of type based on rfc and usage.
func (t AcctBytes) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctBytes.
func (t AcctBytes) Len() int {
	return 1
}

// String returns AcctBytes as a string.
func (t AcctBytes) String() string {
	return fmt.Sprint(int(t))
}

// AcctBytesIn The number of bytes transferred by this action from the endstation
// to the client port.
type AcctBytesIn int

// Validate characterics of type based on rfc and usage.
func (t AcctBytesIn) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctBytesIn.
func (t AcctBytesIn) Len() int {
	return 1
}

// String returns AcctBytesIn as a string.
func (t AcctBytesIn) String() string {
	return fmt.Sprint(int(t))
}

// AcctBytesOut The number of bytes transferred by this action from the client to
// the endstation port.
type AcctBytesOut int

// Validate characterics of type based on rfc and usage.
func (t AcctBytesOut) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctBytesOut.
func (t AcctBytesOut) Len() int {
	return 1
}

// String returns AcctBytesOut as a string.
func (t AcctBytesOut) String() string {
	return fmt.Sprint(int(t))
}

// AcctPaks The number of packets transferred by this action.
type AcctPaks int

// Validate characterics of type based on rfc and usage.
func (t AcctPaks) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctPaks.
func (t AcctPaks) Len() int {
	return 1
}

// String returns AcctPaks as a string.
func (t AcctPaks) String() string {
	return fmt.Sprint(int(t))
}

// AcctPaksIn The number of input packets transferred by this action from the
// endstation to the client port.
type AcctPaksIn int

// Validate characterics of type based on rfc and usage.
func (t AcctPaksIn) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctPaksIn.
func (t AcctPaksIn) Len() int {
	return 1
}

// String returns AcctPaksIn as a string.
func (t AcctPaksIn) String() string {
	return fmt.Sprint(int(t))
}

// AcctPaksOut The number of output packets transferred by this action from the
// endstation to the client port.
type AcctPaksOut int

// Validate characterics of type based on rfc and usage.
func (t AcctPaksOut) Validate(condition interface{}) error {
	return nil
}

// Len returns the length of AcctPaksOut.
func (t AcctPaksOut) Len() int {
	return 1
}

// String returns AcctPaksOut as a string.
func (t AcctPaksOut) String() string {
	return fmt.Sprint(int(t))
}

// AcctRequestFlag bitmapped values
type AcctRequestFlag uint8

const (
	// AcctFlagStart per rfc
	AcctFlagStart AcctRequestFlag = 0x02
	// AcctFlagStop per rfc
	AcctFlagStop AcctRequestFlag = 0x04
	// AcctFlagWatchdog per rfc
	AcctFlagWatchdog AcctRequestFlag = 0x08
	// AcctFlagWatchdogWithUpdate with update per rfc
	AcctFlagWatchdogWithUpdate AcctRequestFlag = 0x0A
)

// Set AcctRequestFlag's f bit.
func (b *AcctRequestFlag) Set(f AcctRequestFlag) { *b = *b | f }

// Clear AcctRequestFlag's f bit.
func (b *AcctRequestFlag) Clear(f AcctRequestFlag) { *b = *b &^ f }

// Toggle AcctRequestFlag's f bit.
func (b *AcctRequestFlag) Toggle(f AcctRequestFlag) { *b = *b ^ f }

// Has returns true when b has the f bit set.
func (b *AcctRequestFlag) Has(f AcctRequestFlag) bool { return *b&f != 0 }

// String to satisfy Fields interface
func (t AcctRequestFlag) String() string {
	flags := make([]string, 0, 4) // 4 supported flags
	if t.Has(AcctFlagStart) {
		flags = append(flags, "AcctFlagStart")
	}
	if t.Has(AcctFlagStop) {
		flags = append(flags, "AcctFlagStop")
	}
	if t.Has(AcctFlagWatchdog) {
		flags = append(flags, "AcctFlagWatchdog")
	}
	if t.Has(AcctFlagWatchdogWithUpdate) {
		flags = append(flags, "AcctFlagWatchdogWithUpdate")
	}
	return strings.Join(flags, "|")
}

// Validate checks for the correct flags to be set
func (t AcctRequestFlag) Validate(condition interface{}) error {
	if t.Has(AcctFlagStop) && t.Has(AcctFlagWatchdog) {
		return fmt.Errorf("stop and watchdog flags are both set")
	}
	return nil
}

// Len returns the length of AcctRequestFlag.
func (t AcctRequestFlag) Len() int {
	return 1
}