// Package testid defines utilities of test id.
package testid

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

const (
	// MetadataKey is the key of gRPC meta data for test id.
	// It can be used for setting test id to gRPC metadata via an interceptor.
	MetadataKey = "test-id"
	// HTTPHeaderKey is the key of HTTP header for test id.
	// It can be used to set test id to HTTP header via a middleware.
	HTTPHeaderKey = "X-Test-ID"
)

type contextKey struct{}

// FromContext returns a test id from ctx or false if no test id is found.
func FromContext(ctx context.Context) (string, bool) {
	tid, ok := ctx.Value(contextKey{}).(string)
	return tid, ok
}

var checkDuplicated = testing.Testing()

// WithValue returns a context associated with tid.
func WithValue(ctx context.Context, tid string) context.Context {
	beforeTestID, ok := ctx.Value(contextKey{}).(string)
	if ok && checkDuplicated {
		msg := fmt.Sprintf("test id has already been associated: %s", beforeTestID)
		panic(msg)
	}
	return context.WithValue(ctx, contextKey{}, tid)
}

// New creates a testid which is based on (testing.TB).Name and uuid.
func New(tb testing.TB) string {
	return tb.Name() + "_" + uuid.NewString()
}
