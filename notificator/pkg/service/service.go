package service

import (
	"context"

	"github.com/pborman/uuid"
)

// NotificatorService describes the service.
type NotificatorService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendEmail(ctx context.Context, email, content string) (id string, err error)
}

type basicNotificatorService struct{}

func (b *basicNotificatorService) SendEmail(ctx context.Context, email string, content string) (id string, err error) {
	// TODO: send real email
	uid := uuid.NewUUID()
	err = nil
	id = uid.String()
	return id, err
}

// NewBasicNotificatorService returns a naive, stateless implementation of NotificatorService.
func NewBasicNotificatorService() NotificatorService {
	return &basicNotificatorService{}
}

// New returns a NotificatorService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificatorService {
	var svc = NewBasicNotificatorService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
