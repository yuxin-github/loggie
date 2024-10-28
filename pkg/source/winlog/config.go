package winlog

import (
	"time"

	kafkaSink "github.com/loggie-io/loggie/pkg/sink/kafka"
)

type Config struct {
	Brokers            []string       `yaml:"brokers,omitempty" validate:"required"`
	Topic              string         `yaml:"topic,omitempty"` // reserved for compatibility
	Topics             []string       `yaml:"topics,omitempty"`
	GroupId            string         `yaml:"groupId,omitempty" default:"loggie"`
	ClientId           string         `yaml:"clientId,omitempty"`
	Worker             int            `yaml:"worker,omitempty" default:"1"`
	QueueCapacity      int            `yaml:"queueCapacity" default:"100"`
	MinAcceptedBytes   int            `yaml:"minAcceptedBytes" default:"1"`
	MaxAcceptedBytes   int            `yaml:"maxAcceptedBytes" default:"1024000"`
	ReadMaxAttempts    int            `yaml:"readMaxAttempts" default:"3"`
	MaxReadWait        time.Duration  `yaml:"maxPollWait" default:"10s"`
	ReadBackoffMin     time.Duration  `yaml:"readBackoffMin" default:"100ms"`
	ReadBackoffMax     time.Duration  `yaml:"readBackoffMax" default:"1s"`
	EnableAutoCommit   bool           `yaml:"enableAutoCommit"`
	AutoCommitInterval time.Duration  `yaml:"autoCommitInterval" default:"1s"`
	AutoOffsetReset    string         `yaml:"autoOffsetReset" default:"latest" validate:"oneof=earliest latest"`
	SASL               kafkaSink.SASL `yaml:"sasl,omitempty"`
	AddonMeta          *bool          `yaml:"addonMeta,omitempty" default:"true"`
}
