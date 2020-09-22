package calc

import "gorone_api/redis"

func Publish(values []string) error {
	q := redis.GetQueue("calc")
	for _, v := range values {
		q.Publish(v)
	}
	return nil
}
