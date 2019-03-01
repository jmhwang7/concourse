package resource

import (
	"context"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/resource/v2"
	"github.com/concourse/concourse/atc/worker"
)

//go:generate counterfeiter . Resource

type Resource interface {
	Get(context.Context, worker.Volume, atc.IOConfig, atc.Source, atc.Params, atc.Space, atc.Version) error
	Put(context.Context, v2.PutEventHandler, atc.IOConfig, atc.Source, atc.Params) ([]atc.SpaceVersion, error)
	Check(context.Context, v2.CheckEventHandler, atc.Source, map[atc.Space]atc.Version) error
	Container() worker.Container
}

type ResourceType string

type Session struct {
	Metadata db.ContainerMetadata
}

type Metadata interface {
	Env() []string
}