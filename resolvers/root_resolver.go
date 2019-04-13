package resolvers

import (
	"context"

	"github.com/Keijun-KUMAGAI/graphql-server/gqlgen"
	"github.com/Keijun-KUMAGAI/graphql-server/prisma-client"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gqlgen.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() gqlgen.TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *queryResolver) Info(ctx context.Context) (string, error) {
	return "this api is working! ver 1.0.0", nil
}

type todoResolver struct{ *Resolver }
