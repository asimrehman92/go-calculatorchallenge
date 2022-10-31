package calculator

import (
	stores_postgres "go-calculatorchallenge/stores/postgres"
	stores_redis "go-calculatorchallenge/stores/redis"

	"net"
	"testing"

	"github.com/go-redis/redis"
)

var (
	connection net.Conn
	rClient    *redis.Client
)

func init() {
	rClient = stores_redis.GetRedisClient()
	stores_postgres.DataMigration()
}

func TestAdditionService(t *testing.T) {
	type args struct {
		c                        net.Conn
		value1, value2, userName string
		redisClient              *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "addFunc",
			args: args{connection, "1", "1", "Satya", rClient},
		},
		{
			name: "addFunc",
			args: args{connection, "159", "141", "Nadella", rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AdditionService(tt.args.c, tt.args.value1, tt.args.value2, tt.args.userName, tt.args.redisClient)
		})
	}
}

func TestSubtractionService(t *testing.T) {
	type args struct {
		c                        net.Conn
		value1, value2, userName string
		redisClient              *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "subFunc",
			args: args{connection, "1", "1", "Satya", rClient},
		},
		{
			name: "subFunc",
			args: args{connection, "159", "141", "Nadella", rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SubtractionService(tt.args.c, tt.args.value1, tt.args.value2, tt.args.userName, tt.args.redisClient)
		})
	}
}

func TestMultiplicationService(t *testing.T) {
	type args struct {
		c                        net.Conn
		value1, value2, userName string
		redisClient              *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "multiFunc",
			args: args{connection, "1", "1", "Satya", rClient},
		},
		{
			name: "multiFunc",
			args: args{connection, "159", "141", "Nadella", rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MultiplicationService(tt.args.c, tt.args.value1, tt.args.value2, tt.args.userName, tt.args.redisClient)
		})
	}
}

func TestDivisionService(t *testing.T) {
	type args struct {
		c                        net.Conn
		value1, value2, userName string
		redisClient              *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "divFunc",
			args: args{connection, "1", "1", "Satya", rClient},
		},
		{
			name: "divFunc",
			args: args{connection, "159", "141", "Nadella", rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DivisionService(tt.args.c, tt.args.value1, tt.args.value2, tt.args.userName, tt.args.redisClient)
		})
	}
}

func TestExpressionService(t *testing.T) {
	type args struct {
		c             net.Conn
		exp, userName string
		redisClient   *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "expFunc",
			args: args{connection, "1 * 1", "Satya", rClient},
		},
		{
			name: "expFunc",
			args: args{connection, "2 * 2", "Nadella", rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExpressionService(tt.args.c, tt.args.exp, tt.args.userName, tt.args.redisClient)
		})
	}
}

func TestProcess(t *testing.T) {
	type args struct {
		exp, userName string
		redisClient   *redis.Client
		want          string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "processFunc",
			args: args{"1 + 1", "Satya", rClient, "1 + 1 = 2"},
		},
		{
			name: "processFunc",
			args: args{"2 + 2", "Nadella", rClient, "2 + 2 = 4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Process(tt.args.exp, tt.args.userName, tt.args.redisClient)
			if err != nil {
				t.Errorf("expecting nil err, got %v", err)
			}
			if got != tt.args.want {
				t.Errorf("expecting %s, got %s", tt.args.want, got)
			}
		})
	}
}
