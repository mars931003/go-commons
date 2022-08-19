package go_commons

import (
	"github.com/go-redis/redis"
	"log"
)

type GoRedisClient[T any] interface {
	GetString(key string) string
	SetString(key string)
	GetList(key string) []T
	SetList(key string)
	GetKey(key string)
}

type rdsClient redis.Client

func (r rdsClient) GetString(key string) string {
	return r.Get(key).String()
}

func (r rdsClient) SetString(key string) {
	//TODO implement me
	panic("implement me")
}

func (r rdsClient) GetList(key string) []string {
	//TODO implement me
	panic("implement me")
}

func (r rdsClient) SetList(key string) {
	//TODO implement me
	panic("implement me")
}

func (r rdsClient) GetKey(key string) {
	//TODO implement me
	panic("implement me")
}

//var RdsGoClient *rdsClient[]

func init() {
	//读配置文件获取redis 参数
	config := NewConfigReader("")
	v := config.YmlFileRead()
	host := v.GetString("redis-client.host")
	pwd := v.GetString("redis-client.pwd")
	database := v.GetInt("redis-client.database")
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pwd,
		DB:       database,
	})
	result, err := client.Ping().Result()
	log.Println(result)
	if err != nil {
		log.Fatal(err.Error())
	}
	//RdsGoClient = (*rdsClient)(client)
}

//func (rc *rdsClient) readConfigFile(key string) interface{} {
//	config := NewConfigReader("./redis-config.yaml")
//	fileInfo := config.YmlFileRead()
//	return fileInfo.Get(key)
//}
