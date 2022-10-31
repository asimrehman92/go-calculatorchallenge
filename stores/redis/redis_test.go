package stores

import (
	"testing"

	"github.com/go-redis/redis"
)

var rClient *redis.Client

func init() {
	rClient = GetRedisClient()
}

func TestGetRedisClient(t *testing.T) {

	type args struct {
		want *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "getRedisClientFunc",
			args: args{rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := GetRedisClient(); !reflect.DeepEqual(got, tt.args.want) {
			// 	t.Errorf("GetRedisClient() = %v, want %v", got, tt.args.want)
			// }
			got := GetRedisClient()
			if got == tt.args.want {
				t.Errorf("expecting %v, got %v", tt.args.want, got)
			}
		})
	}
}

func TestDisplayDataRedis(t *testing.T) {
	type args struct {
		user        string
		redisClient *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "displayDataRedisFunc",
			args: args{"Satya", rClient},
		},
		{
			name: "displayDataRedisFunc",
			args: args{"Nadella", rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DisplayDataRedis(tt.args.user, tt.args.redisClient)
		})
	}
}

func TestAddUsersRedis(t *testing.T) {
	type args struct {
		user        string
		redisClient *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "addUsersRedisFunc",
			args: args{"Satya", rClient},
		},
		{
			name: "addUsersRedisFunc",
			args: args{"Nadella", rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddUsersRedis(tt.args.user, tt.args.redisClient)
		})
	}
}

func TestAddExpressionsRedis(t *testing.T) {
	type args struct {
		user, exp   string
		redisClient *redis.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "addExpressionsRedisFunc",
			args: args{"Satya", "8 * 8", rClient},
		},
		{
			name: "addExpressionsRedisFunc",
			args: args{"Nadella", "9 * 9", rClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddExpressionsRedis(tt.args.user, tt.args.exp, tt.args.redisClient)
		})
	}
}
