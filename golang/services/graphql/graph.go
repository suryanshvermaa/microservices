package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/suryanshvermaa/microservices/golang/services/account"
	"github.com/suryanshvermaa/microservices/golang/services/catalog"
	"github.com/suryanshvermaa/microservices/golang/services/order"
)

type Server struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func NewGraphQlServer(accountUrl string, catalogUrl string, orderUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	catalogClient, err := catalog.NewClient()
	if err != nil {
		accountClient.Close()
		return nil, err
	}

	orderClient, err := order.NewClient()
	if err != nil {
		accountClient.Close()
		catalogClient.Close()
		return nil, err
	}
	return &Server{
		accountClient: accountClient,
		catalogClient: catalogClient,
		orderClient:   orderClient,
	}, nil
}

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}
// }

func (s *Server) ToExecuteSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(
		Config{
			Resolvers: s,
		})
}
