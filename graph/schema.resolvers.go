package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ChizShortener/database"
	"ChizShortener/graph/generated"
	"ChizShortener/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	return db.SaveLink(ctx, &input), nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	return db.GetAllLinks(ctx), nil
}

func (r *queryResolver) GetLongLink(ctx context.Context, shortLink string) (*model.Link, error) {
	return db.GetLink(ctx,"shortLink",shortLink)
}

func (r *queryResolver) GetShortLink(ctx context.Context, longLink string) (*model.Link, error) {
	return db.GetLink(ctx,"longLink",longLink)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Link(ctx context.Context, shortLink *string, longLink *string) (*model.Link, error) {
	panic(fmt.Errorf("not implemented"))
}

var db = database.ConnectDb()
