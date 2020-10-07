package redis

import (
	"gorone_api/lib/util"

	"github.com/adjust/rmq"
)

func GetQueue(tag string) rmq.Queue {
	connection := rmq.OpenConnection("gorone redis", "tcp", util.GetEnv("REDIS_URL", "redis:6379"), 0)
	queue := connection.OpenQueue(tag)
	return queue
}
