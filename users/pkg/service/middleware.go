package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UsersService) UsersService

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UsersService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsersService) UsersService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, email string) (err error) {
	defer func() {
		l.logger.Log("method", "Create", "email", email, "err", err)
	}()
	return l.next.Create(ctx, email)
}
