package pkg

import (
	_ "moocss.com/gaea/pkg/conf" // init conf

	"moocss.com/gaea/pkg/log"
)

// GatherMetrics 收集一些被动指标
func GatherMetrics() {
}

// Reset all utils
func Reset() {
	log.Reset()
}

// Stop all utils
func Stop() {
}
