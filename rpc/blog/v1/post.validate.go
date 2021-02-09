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

func (m *GetPostRequest) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type GetPostRequestValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e GetPostRequestValidationError) Error() string {
	return fmt.Sprintf(
		"invalid GetPostRequest.%s: %s",
		e.field,
		e.reason)
}

func (m *GetPostReply) validate() error {
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
