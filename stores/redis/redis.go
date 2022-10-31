package stores

import (
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/go-redis/redis"
)

type Equation struct {
	Expression string
	Result     string
	TimeStamp  time.Time
}

func GetRedisClient() *redis.Client {

	// var (
	// 	host     = settings.GetString("REDIS_CONFIG.REDIS_HOST")
	// 	port     = settings.GetString("REDIS_CONFIG.REDIS_PORT")
	// 	password = settings.GetString("REDIS_CONFIG.REDIS_PASSWORD")
	// )

	var (
		host     = "localhost"
		port     = "6379"
		password = ""
	)

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})

	// Check we have correctly connected to our Redis server
	pong, err := client.Ping().Result()
	if err != nil {
		log.WithError(err).Error("first Run the redis-server.exe file for redis connection")
	}
	fmt.Println(pong, err)
	return client
}

func DisplayDataRedis(user string, redisClient *redis.Client) {

	eachClientExpressions := redisClient.SMembers(user)
	allUsersName := redisClient.SMembers("id")
	totalUsers := redisClient.SCard("id")

	log.WithFields(log.Fields{
		"Total Users":             totalUsers,
		"All Users Name":          allUsersName,
		"Each Client Expressions": eachClientExpressions,
	}).Info("Redis Data")
}

func AddUsersRedis(user string, redisClient *redis.Client) {
	redisClient.SAdd("id", user)
}

func AddExpressionsRedis(user, exp string, redisClient *redis.Client) {
	redisClient.SAdd(user, exp)
}
