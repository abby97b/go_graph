package graph

import "github.com/abbyb97/simple_go_graphql/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	users []*model.User
}
