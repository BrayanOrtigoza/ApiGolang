package conection


import (
	"time"

	"API/configuration"
	"API/constants"

	"github.com/go-redis/redis"
	rethinkDB "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var (
	client  *redis.Client
	session *rethinkDB.Session
)

//
func newConnectionRedisServer() {
	client = redis.NewClient(&redis.Options{
		Addr:     configuration.UrlRedisServer,
		Password: configuration.PasswordRedisServer,
		DB:       configuration.DatabaseRedisServer,
	})
}

//
func newConnectionRethinkDBServer() error {
	var err error
	session, err = rethinkDB.Connect(rethinkDB.ConnectOpts{
		Address:  configuration.UrlRethinkDBServer,
		Database: configuration.DatabaseRethinkDBServer,
		Username: configuration.UserRethinkDBServer,
		Password: configuration.PasswordRethinkDBServer,
	})
	return err
}

//
func verifyConnectionRedisServer() bool {
	pong, err := client.Ping().Result()
	if err == nil && pong == constants.Pong {
		return constants.PongReceived
	} else {
		return constants.PongNotReceived
	}
}

//
func verifyConnectionRethinkDBServer() bool {
	return session.IsConnected()
}

//
func reconnectRedisServer() {
	newConnectionRedisServer()
	for {
		if !verifyConnectionRedisServer() {
			newConnectionRedisServer()
		}
		waitTimeSeconds := time.Duration(constants.SecondsReconnectRedis)
		time.Sleep(waitTimeSeconds * time.Second)
	}
}

//
func reconnectRethinkDBServer() {
	err := newConnectionRethinkDBServer()
	for {
		if !verifyConnectionRethinkDBServer() || err != nil {
			err = newConnectionRethinkDBServer()
		}
		waitTimeSeconds := time.Duration(constants.SecondsReconnectRethinkDB)
		time.Sleep(waitTimeSeconds * time.Second)
	}
}

//
func Init() {
	go reconnectRedisServer()
	go reconnectRethinkDBServer()
}
