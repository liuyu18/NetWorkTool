package plugins

import (
	"fmt"

	"myTools/srack/models"
	"myTools/srack/vars"

	"github.com/go-redis/redis"
)

func ScanRedis(s models.Service) (result models.ScanResult, err error) {
	result.Service = s
	opt := redis.Options{Addr: fmt.Sprintf("%v:%v", s.Ip, s.Port),
		Password: s.Password, DB: 0, DialTimeout: vars.TimeOut}
	client := redis.NewClient(&opt)
	_, err = client.Ping().Result()
	if err != nil {
		return result, err
	}

	result.Result = true

	defer func() {
		if client != nil {
			_ = client.Close()
		}
	}()

	return result, err
}
