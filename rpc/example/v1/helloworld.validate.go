package example_v1

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

func (m *HelloworldEchoReq) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type HelloworldEchoReqValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e HelloworldEchoReqValidationError) Error() string {
	return fmt.Sprintf(
		"invalid HelloworldEchoReq.%s: %s",
		e.field,
		e.reason)
}

func (m *HelloworldEchoResp) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type HelloworldEchoRespValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e HelloworldEchoRespValidationError) Error() string {
	return fmt.Sprintf(
		"invalid HelloworldEchoResp.%s: %s",
		e.field,
		e.reason)
}
