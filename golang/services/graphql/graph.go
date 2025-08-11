package main

type Server struct {
	accountClient
	catalogClient
	orderClient
}

func NewGraphQlServer(accountUrl string) {

}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}
