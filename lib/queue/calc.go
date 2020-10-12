package calc

import (
	"gorone_api/models"
	"gorone_api/redis"
)

func Publish(req models.Request) error {
	q := redis.GetQueue(req.RedisTagName)
	for _, v := range req.DecodedBody() {
		q.Publish(v)
	}
	return nil
}
