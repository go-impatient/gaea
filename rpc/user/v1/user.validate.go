package user_v1

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// ensure the imports are used
var (
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = (*url.URL)(nil)
)

func (m *UserEchoReq) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type UserEchoReqValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e UserEchoReqValidationError) Error() string {
	return fmt.Sprintf(
		"invalid UserEchoReq.%s: %s",
		e.field,
		e.reason)
}

func (m *UserEchoResp) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type UserEchoRespValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e UserEchoRespValidationError) Error() string {
	return fmt.Sprintf(
		"invalid UserEchoResp.%s: %s",
		e.field,
		e.reason)
}
