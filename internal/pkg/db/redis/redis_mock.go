package redis

import (
	redismock "github.com/go-redis/redismock/v9"
)

type MockClient struct {
	RedisClient Redis
	Mock        redismock.ClientMock
}

func GetMockRedisClient() MockClient {
	rc, rm := redismock.NewClientMock()
	return MockClient{
		RedisClient: Redis{
			redis: rc,
		},
		Mock: rm,
	}
}
