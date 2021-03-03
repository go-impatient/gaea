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

func (m *CreateArticleRequest) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type BlogEchoReqValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e BlogEchoReqValidationError) Error() string {
	return fmt.Sprintf(
		"invalid BlogEchoReq.%s: %s",
		e.field,
		e.reason)
}

func (m *CreateArticleReply) validate() error {
	if m == nil {
		return nil
	}

	return nil
}

type BlogEchoRespValidationError struct {
	field  string
	reason string
}

// Error satisfies the builtin error interface
func (e BlogEchoRespValidationError) Error() string {
	return fmt.Sprintf(
		"invalid BlogEchoResp.%s: %s",
		e.field,
		e.reason)
}
