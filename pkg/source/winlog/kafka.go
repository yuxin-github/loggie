package winlog

import (
	"fmt"
	"github.com/loggie-io/loggie/pkg/core/api"
	"github.com/loggie-io/loggie/pkg/core/event"
	"github.com/loggie-io/loggie/pkg/pipeline"
)

const (
	Type = "winLog"
)

func init() {
	pipeline.Register(api.SOURCE, Type, makeSource)
}

func makeSource(info pipeline.Info) api.Component {
	return &Source{
		done:      make(chan struct{}),
		config:    &Config{},
		eventPool: info.EventPool,
	}
}

type Source struct {
	name      string
	done      chan struct{}
	config    *Config
	eventPool *event.Pool
}

func (k *Source) Config() interface{} {
	return k.config
}

func (k *Source) Category() api.Category {
	return api.SOURCE
}

func (k *Source) Type() api.Type {
	return Type
}

func (k *Source) String() string {
	return fmt.Sprintf("%s/%s", api.SOURCE, Type)
}

func (k *Source) Init(context api.Context) error {
	k.name = context.Name()
	return nil
}

func (k *Source) Start() error {
	return nil
}

func (k *Source) Stop() {
	close(k.done)
}

func (k *Source) ProductLoop(productFunc api.ProductFunc) {
	return
}

func (k *Source) Commit(events []api.Event) {
	// commit when sink ack
	k.eventPool.PutAll(events)
}
