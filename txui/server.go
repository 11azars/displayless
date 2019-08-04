package txui

import (
  "net/http"
  "sync"

  "github.com/rs/zerolog"
  goji "goji.io"
  "goji.io/pat"
  "github.com/99designs/gqlgen/handler"

  "gitlab.com/newrx/ha/txui/gql"
)

// Server is a http (GraphQL) server used to host
// an API for credential/config exchange over AP.
type Server struct{
  httpServer *http.Server

  stopLock *sync.Mutex
  attemptingStop bool
}

// Start creates and starts an http server to host the
// API providing wifi management functions (usually provided
// over AP, in order to pass creds to connect to wireless networks).
func Start(addr string, autorestart, devmode bool, logger zerolog.Logger) *Server {
  // TODO add TLS
  mux := goji.NewMux()
  server := &Server{
    httpServer: &http.Server{Addr: addr, Handler: mux},
    stopLock: &sync.Mutex{},
  }

  // add graphql endpoint
  mux.Handle(
    pat.New("/graphql"),
    handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &gql.Resolver{
      // pass in network manager
    }})),
  )

  // optionally add graphql playground
  if devmode {
    mux.Handle(pat.New("/"), handler.Playground("GraphQL playground", "/query"))
  }

  go func() {
    restartLoop:
    for {
      if err := server.httpServer.ListenAndServe(); err != http.ErrServerClosed {
        logger.Error().Err(err).Msg("http listener crashed")
      }

      if server.attemptingStop || !autorestart {
        break restartLoop
      }
    }
  }()

  return server
}

// Close stops the server running in the background
func (s *Server) Close() error {
  s.stopLock.Lock()
  s.attemptingStop = true
  s.stopLock.Unlock()

  return s.httpServer.Shutdown(context.TODO())
}
