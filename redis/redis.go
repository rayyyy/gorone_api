package redis

import "github.com/adjust/rmq"

func GetQueue(tag string) rmq.Queue {
	connection := rmq.OpenConnection("gorone redis", "tcp", "redis:6379", 0)
	queue := connection.OpenQueue(tag)
	return queue
}
