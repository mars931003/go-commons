package redis

import (
	"github.com/go-redis/redis"
)

type GoRedisClient interface {
	GetString()
	SetString()
	GetList()
	SetList()
}

type rdsClient redis.Client

var RdsGoClient *rdsClient

func init() {
	//读配置文件获取redis 参数
	//client := redis.NewClient(&redis.Options{
	//	Addr:     host,
	//	Password: pwd,
	//	DB:       database,
	//})
	//result, err := client.Ping().Result()
	//log.Println(result)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//RdsGoClient = (*rdsClient)(client)
}

func main() {

}

//func (c *rdsClient) GetString() {
//
//}
//
//func (c *rdsClient) SetString() {
//
//}
//
//func (c *rdsClient) GetList() {}
//
//func (c *rdsClient) SetList() {
//
//}
