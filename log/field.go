package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/ybzhanghx/pkgs/ctxfield"
	"github.com/ybzhanghx/pkgs/userkit"
)

var ctxMarker struct{}

// Ctx creates an entry from the standard logger and adds a context to it.
// Add a single field(trace_id) to the Entry
func Ctx(ctx context.Context) *Entry {
	if entry, ok := ctx.Value(ctxMarker).(*Entry); ok && entry != nil {
		return entry
	}
	entry := logrus.WithContext(ctx)
	if traceID, ok := ctx.Value(ctxfield.TraceIDKey).(string); ok {
		entry = entry.WithField("trace_id", traceID)
	}
	if userKit, ok := ctx.Value(ctxfield.UserKitKey).(*userkit.UserKit); ok {
		entry = entry.WithField("user_kit", userKit)
	}
	return entry
}
