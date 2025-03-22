package decorator

import (
	"context"
	"user/app/command"
)

type CommandHandler[C, R any] interface {
	Handle(ctx context.Context, cmd C) (R, error)
}

func (c CommandHandler[C, R]) Handle(ctx context.Context, cmd command.SignInByEmailPassword) (*command.SignInResult, error) {
	//TODO implement me
	panic("implement me")
}

func ApplyCommandDecorator[C, R any](handler CommandHandler[C, R]) CommandHandler[C, R] {
	return handler
}
