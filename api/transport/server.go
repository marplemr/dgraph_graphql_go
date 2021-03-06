package transport

import (
	"context"
	"log"
	"time"

	"github.com/romshark/dgraph_graphql_go/api/graph"
	"github.com/romshark/dgraph_graphql_go/store"
)

// OnGraphQuery defines the graph query callback function
type OnGraphQuery func(context.Context, graph.Query) (graph.Response, error)

// OnAuth defines the client authentication callback function
type OnAuth func(ctx context.Context, sessionKey string) (store.ID, time.Time)

// OnDebugAuth defines the debug authentication callback function
type OnDebugAuth func(ctx context.Context, sessionKey string) bool

// OnDebugSess defines the debug session creation callback function
type OnDebugSess func(ctx context.Context, username, password string) []byte

// Server defines the interface of the server transport layer implementation.
// Run and Init are not intended to be thread-safe and shall only be used
// by a single goroutine!
type Server interface {
	// Init initializes the server transport implementation.
	// The provided callbacks must be registered and invoked accordingly
	Init(
		onGraphQuery OnGraphQuery,
		onAuth OnAuth,
		onDebugAuth OnDebugAuth,
		onDebugSess OnDebugSess,
		errorLog *log.Logger,
	) error

	// Run starts serving. Blocks until the underlying server is shut down
	Run() error

	// Shutdown instructs the underlying server to shut down and blocks until
	// it is
	Shutdown(ctx context.Context) error
}
