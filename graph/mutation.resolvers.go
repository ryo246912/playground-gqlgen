package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.72

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ryo246912/playground-gqlgen/graph/db"
	"github.com/ryo246912/playground-gqlgen/graph/model"
	"github.com/ryo246912/playground-gqlgen/internal"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))

	// todo := &model.Todo{
	// 	Text: input.Text,
	// 	ID:   fmt.Sprintf("T%d", randNumber),
	// 	User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	// }
	// NOTE:userIDをresolveする
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", randNumber),
		UserID: input.UserID,
	}

	r.todos = append(r.todos, todo)
	return todo, nil
}

// CreateCustomer is the resolver for the createCustomer field.
func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomer) (*bool, error) {
	_, err := r.DB.NewInsert().Model(&db.Customer{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Email:      sql.NullString{String: input.Email, Valid: input.Email != ""},
		Active:     true,
		CreateDate: time.Now(),
		LastUpdate: sql.NullTime{Time: time.Now(), Valid: true},
		StoreID:    uint8(input.StoreID),
		// 仮でベタ打ち
		AddressID: 605,
	}).Exec(ctx)

	if err != nil {
		log.Println("error!!", err)
		return nil, err
	}

	result := true
	return &result, nil
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
