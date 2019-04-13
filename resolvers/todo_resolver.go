package resolvers

import (
	"context"

	"github.com/Keijun-KUMAGAI/graphql-server/gqlgen"
	"github.com/Keijun-KUMAGAI/graphql-server/prisma-client"
)

// -------------------- Mutation --------------------

func (r *mutationResolver) TodoCreate(ctx context.Context, params gqlgen.TodoCreateInput) (*prisma.Todo, error) {
	todo, err := r.Prisma.CreateTodo(prisma.TodoCreateInput{
		Message: params.Message,
		User: prisma.UserCreateOneWithoutTodosInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &params.UserID,
			},
		},
	}).Exec(ctx)

	return todo, err
}
func (r *mutationResolver) TodoUpdate(ctx context.Context, params gqlgen.TodoUpdateInput) (*prisma.Todo, error) {
	todo, err := r.Prisma.UpdateTodo(prisma.TodoUpdateParams{
		Where: prisma.TodoWhereUniqueInput{
			ID: &params.ID,
		},
		Data: prisma.TodoUpdateInput{
			Message: &params.Message,
			Done:    &params.Done,
		},
	}).Exec(ctx)

	return todo, err
}
func (r *mutationResolver) TodoDelete(ctx context.Context, params gqlgen.TodoDeleteInput) (*prisma.Todo, error) {
	todo, err := r.Prisma.DeleteTodo(prisma.TodoWhereUniqueInput{
		ID: &params.ID,
	}).Exec(ctx)

	return todo, err
}

// -------------------- TodoResolver --------------------

func (r *todoResolver) User(ctx context.Context, obj *prisma.Todo) (*prisma.User, error) {
	user, err := r.Prisma.Todo(prisma.TodoWhereUniqueInput{ID: &obj.ID}).User().Exec(ctx)
	if err != nil {
		return nil, nil
	}

	return user, err
}
