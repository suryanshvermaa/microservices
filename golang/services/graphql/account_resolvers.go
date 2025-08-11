package main

import "context"

type accountResolvers struct {
	server *Server
}

func (r *queryResolver) Orders(ctx context.Context, pagination *PaginationInput, id *string) ([]Order, error) {

}
