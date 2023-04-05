package service

import (
	"net"
	"strconv"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type HudServiceDesc interface {
	ServiceDesc() *grpc.ServiceDesc
	Port() uint64
}

type NewService func(logger zerolog.Logger) (interface{}, error)

func ServiceGRPC(appName string, desc HudServiceDesc, newService NewService) {
	logger := zerolog.New(zerolog.NewConsoleWriter())

	service, err := newService(logger)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Failed to create service")
	}

	grpcServer := grpc.NewServer()
	grpcServer.RegisterService(desc.ServiceDesc(), service)

	listenAddr := net.JoinHostPort("localhost", strconv.FormatUint(desc.Port(), 10))

	logger.Info().
		Str("listen_addr", listenAddr).
		Str("app_name", appName).
		Msg("Starting")

	listen, err := net.Listen("tcp", listenAddr)
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
