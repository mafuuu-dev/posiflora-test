package scope

import (
	"backend/core/config"
	"backend/core/pkg/repository"
	"backend/core/pkg/storage"
	"context"
)

type Factory struct {
	Repository *repository.Factory
}

type Support struct {
	Factory *Factory
}

type Scope struct {
	Context context.Context
	Config  *config.Config
	DB      *storage.Storage
	Support *Support
}

func New(
	ctx context.Context,
	cfg *config.Config,
	database *storage.Storage,
	support Support,
) *Scope {
	return &Scope{
		Context: ctx,
		Config:  cfg,
		DB:      database,
		Support: &support,
	}
}
