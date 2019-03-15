package constants

const (
	//
	Colon = ":"
	Empty = ""
	Ping  = "PING"
	Pong  = "PONG"

	//
	NoDebug         = false
	PongReceived    = true
	PongNotReceived = false

	//
	Endpoint                  = "/socket_test"
	DefaultHost               = "127.0.0.1"
	DefaultPortRedis          = "6379"
	DefaultPortRethinkDB      = "28015"
	DefaultPortWebSocket      = "4000"
	KeyTokenRedis             = "TokenAuth"
	NameToken                 = "access_token"
	SecondsReconnectRedis     = 30
	SecondsReconnectRethinkDB = 30
	SecondsTimeoutPing        = 30
	SecondsTimeoutPong        = 30
)
