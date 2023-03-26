package service

import (
	"net"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type NewService func(logger zerolog.Logger) (interface{}, error)

func ServiceGRPC(appName string, desc *grpc.ServiceDesc, newService NewService) {
	logger := zerolog.New(zerolog.NewConsoleWriter())

	service, err := newService(logger)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Failed to create service")
	}

	grpcServer := grpc.NewServer()
	grpcServer.RegisterService(desc, service)

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Failed to listen")
	}

	if err := grpcServer.Serve(listen); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Service returned an error")
	}
}
