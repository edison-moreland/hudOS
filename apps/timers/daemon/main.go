package main

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/edison-moreland/nreal-hud/go-sdk/app/service"
	"github.com/edison-moreland/nreal-hud/proto/timers"
)

// Daemon listens for protobuf messages, and manages the timer windows
func main() {
	service.ServiceGRPC("timers", &timers.Timers_ServiceDesc, func(logger zerolog.Logger) (interface{}, error) {
		return &TimersService{
			Logger: logger,
		}, nil
	})
}

type TimersService struct {
	timers.UnimplementedTimersServer
	zerolog.Logger
}

func (t *TimersService) NewTimer(ctx context.Context, request *timers.NewTimerRequest) (*timers.NewTimerResponse, error) {
	t.Info().Msg("New timer!!!")

	return &timers.NewTimerResponse{}, nil
}

func (t *TimersService) StopTimer(ctx context.Context, request *timers.StopTimerRequest) (*timers.StopTimerResponse, error) {
	t.Info().Msg("Stop timer!!!")

	return &timers.StopTimerResponse{}, nil
}

func (t *TimersService) ListTimers(ctx context.Context, request *timers.ListTimersRequest) (*timers.ListTimersResponse, error) {
	t.Info().Msg("List timers!!!")

	return &timers.ListTimersResponse{}, nil
}
