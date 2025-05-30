package graph

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ryo246912/playground-gqlgen/internal"
	"github.com/ryo246912/playground-gqlgen/middlewares/auth"
)

var Directive internal.DirectiveRoot = internal.DirectiveRoot{
	IsAuthenticated: IsAuthenticated,
}

func IsAuthenticated(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
	if _, ok := auth.GetUserName(ctx); ok != nil {
		return nil, errors.New("not authenticated")
	}
	return next(ctx)
}
