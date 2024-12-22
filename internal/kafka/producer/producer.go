package producer

import (
	"fmt"
	"my_lib/lib/env"

	"github.com/IBM/sarama"
)

type brokers []string

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func getBrokers() brokers {
	return brokers{
		fmt.Sprintf("localhost:%s", env.GetKafkaPort()),
	}
}
