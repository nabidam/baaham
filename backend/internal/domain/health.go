package domain

import "context"

type HealthRepository interface {
	Check(ctx context.Context) (bool, error)
}

type HealthService interface {
	HealthCheck(ctx context.Context) (bool, error)
}
