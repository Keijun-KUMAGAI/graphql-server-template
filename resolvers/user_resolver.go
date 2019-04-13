package resolvers

import (
	"context"

	"github.com/Keijun-KUMAGAI/graphql-server/gqlgen"
	"github.com/Keijun-KUMAGAI/graphql-server/prisma-client"
)

func (r *mutationResolver) UserCreate(ctx context.Context, params gqlgen.UserCreateInput) (*prisma.User, error) {
	user, err := r.Prisma.CreateUser(prisma.UserCreateInput{
		Name: params.Name,
	}).Exec(ctx)
	return user, err
}
