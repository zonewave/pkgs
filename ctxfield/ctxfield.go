package ctxfield

type ctxKey string

var (
	// TraceIDKey trace Key
	TraceIDKey ctxKey = "pkgs-trace-id"
	// UserKitKey userKit key
	UserKitKey ctxKey = "pkgs-user-kit"
)
