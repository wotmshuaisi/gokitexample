package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(BugsService) BugsService

type loggingMiddleware struct {
	logger log.Logger
	next   BugsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a BugsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next BugsService) BugsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, bug string) (err error) {
	defer func() {
		l.logger.Log("method", "Create", "bug", bug, "err", err)
	}()
	return l.next.Create(ctx, bug)
}
