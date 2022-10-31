package calculator

import (
	"fmt"
	"net"

	stores_postgres "go-calculatorchallenge/stores/postgres"
	stores_redis "go-calculatorchallenge/stores/redis"

	"github.com/apex/log"
	"github.com/go-redis/redis"
	"github.com/knetic/govaluate"
)

var (
	clients                              []string
	each_client_information              = make(map[string]string)
	client_how_many_calculations_perform = make(map[string]int64)
	TotalExpression                      = 0
)

// type Equation struct {
// 	Expression string
// 	Result     interface{}
// 	TimeStamp  time.Time
// }

func AdditionService(c net.Conn, value1, value2, userName string, redisClient *redis.Client) {
	log.Info("Addition Service Called...\n")

	expression := value1 + " + " + value2
	Process(expression, userName, redisClient)
}

func SubtractionService(c net.Conn, value1, value2, userName string, redisClient *redis.Client) {
	log.Info("Subtraction Service Called...\n")

	expression := value1 + " - " + value2
	Process(expression, userName, redisClient)
}

func MultiplicationService(c net.Conn, value1, value2, userName string, redisClient *redis.Client) {
	log.Info("Multiplication Service Called...\n")

	expression := value1 + " * " + value2
	Process(expression, userName, redisClient)
}

func DivisionService(c net.Conn, value1, value2, userName string, redisClient *redis.Client) {
	log.Info("Division Service Called...\n")

	expression := value1 + " / " + value2
	Process(expression, userName, redisClient)
}

func ExpressionService(c net.Conn, exp, userName string, redisClient *redis.Client) {
	log.Info("Expression Service Method called...\n")

	Process(exp, userName, redisClient)
}

func Process(exp, userName string, redisClient *redis.Client) (string, error) {
	log.Info("Process Function called...\n")

	expressionss, err := govaluate.NewEvaluableExpression(exp)
	if err != nil {
		log.WithError(err).Error("Failed to Handle Expression")
	}

	result, err := expressionss.Evaluate(nil)
	if err != nil {
		log.WithError(err).Error("Failed to Evaluate Expression")
	}

	exp_result := exp + " = " + fmt.Sprint(result)
	fmt.Println("Expression_ result:- ", exp_result)

	stores_redis.AddUsersRedis(userName, redisClient)
	stores_redis.AddExpressionsRedis(userName, exp_result, redisClient)
	// stores_redis.DisplayDataRedis(userName, redisClient)
	stores_postgres.CreateClient(userName, exp_result)
	return exp_result, err
}
