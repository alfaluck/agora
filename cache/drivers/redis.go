package drivers

import (
	"github.com/alfaluck/agora/cache"
	"github.com/mediocregopher/radix.v2/pool"
	"log"
	"github.com/mediocregopher/radix.v2/redis"
	"errors"
)

type Redis struct {
	prefix string
	cache.Provider
	pool *pool.Pool
}

func (r *Redis) Configure(config map[string]string) (err error) {
	net, ok := config["net"]
	if !ok {
		log.Println("Missed config param cacheHandler:net, using default value `tcp`")
		net = "tcp"
	}
	host, ok := config["host"]
	if !ok {
		log.Println("Missed config param cacheHandler:host, using default value `localhost`")
		host = "localhost"
	}
	port, ok := config["post"]
	if !ok {
		log.Println("Missed config param cacheHandler:port, using default value `6379`")
		host = "6379"
	}
	addr := host + ":" + port
	df := redis.Dial
	password, ok := config["password"]
	if !ok {
		password = ""
	}
	database, ok := config["database"]
	if !ok {
		database = "0"
	}
	if password != "" || database != "0" {
		df = func(network, addr string) (*redis.Client, error) {
			client, err := redis.Dial(network, addr)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if err = client.Cmd("AUTH", password).Err; err != nil {
					client.Close()
					return nil, err
				}
			}
			if database != "0" {
				if err = client.Cmd("SELECT", database).Err; err != nil {
					client.Close()
					return nil, err
				}
			}
			return client, nil
		}
	}
	r.pool, err = pool.NewCustom(net, addr, 10, df)

	prefix, ok := config["prefix"]
	if ok {
		r.prefix = prefix
	}

	return
}

func (r *Redis) GetItem(key string) (*cache.Item, error) {
	result := r.pool.Cmd("GET")
	return nil, errors.New("Not implemented GetItem() method")
}
func init() {
	cache.Register("redis", new(Redis))
}
