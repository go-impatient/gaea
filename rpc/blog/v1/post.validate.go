package blog_v1

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

func (m *PostEchoReq) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type PostEchoReqValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e PostEchoReqValidationError) Error() string {
	return fmt.Sprintf(
		"invalid PostEchoReq.%s: %s",
		e.field,
		e.reason)
}

func (m *PostEchoResp) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type PostEchoRespValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e PostEchoRespValidationError) Error() string {
	return fmt.Sprintf(
		"invalid PostEchoResp.%s: %s",
		e.field,
		e.reason)
}
