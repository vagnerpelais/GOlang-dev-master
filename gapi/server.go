package gapi

import (
	"fmt"

	db "github.com/vagnerpelais/simplebank/db/sqlc"
	"github.com/vagnerpelais/simplebank/pb"
	"github.com/vagnerpelais/simplebank/token"
	"github.com/vagnerpelais/simplebank/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

