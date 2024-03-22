package services

import (
	"context"
	"errors"
	"strings"

	"github.com/go-kit/log"
)

type stringService struct {
	logger log.Logger
}

type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) uint64
}

// NewService func initializes a service
func NewService(logger log.Logger) StringService {
	return &stringService{
		logger: logger,
	}
}

func (s stringService) Uppercase(ctx context.Context, str string) (string, error) {
	if str == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(str), nil
}

func (s stringService) Count(ctx context.Context, str string) uint64 {
	return uint64(len(str))
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
