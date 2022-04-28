package sessions

import(
	"github.com/go-redis/redis/v8"
)

type SessionsManager struct{
	rdb redis.Client
}

func NewSessionsManager(){
	
}

