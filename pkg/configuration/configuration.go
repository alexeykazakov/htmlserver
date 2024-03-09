// Package configuration is in charge of the validation and extraction of all
// the configuration details from a configuration file or environment variables.
package configuration

import (
	"time"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	// Commit current build commit set by build script.
	Commit = "0"
	// BuildTime set by build script in ISO 8601 (UTC) format:
	// YYYY-MM-DDThh:mm:ssTZD (see https://www.w3.org/TR/NOTE-datetime for
	// details).
	BuildTime = "0"
	// StartTime in ISO 8601 (UTC) format.
	StartTime = time.Now().UTC().Format("2006-01-02T15:04:05Z")
)

var logger = logf.Log.WithName("configuration")

const (
	GracefulTimeout       = time.Second * 15
	HTTPAddress           = "0.0.0.0:8090"
	HTTPCompressResponses = true
	HTTPIdleTimeout       = time.Second * 60
	HTTPReadTimeout       = time.Second * 60
	HTTPWriteTimeout      = time.Second * 60
)
