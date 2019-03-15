package configuration

const (
	// Names of enviroment variables
	nameEnviromentDebug   = "DEBUG"
	nameHostRedis         = "HOST_REDIS"
	nameHostRethinkDB     = "HOST_RETHINKDB"
	namePasswordRedis     = "PASSWORD_REDIS"
	namePasswordRethinkDB = "PASSWORD_RETHINKDB"
	namePortApp           = "PORT_APP"
	namePortRedis         = "PORT_REDIS"
	namePortRethinkDB     = "PORT_RETHINKDB"

	// Constants for connections to databases (default)
	databaseRedis     = 0
	databaseRethinkDB = "test"

	// Constants for Socket
	hostWebSocket   = "0.0.0.0"
	readBufferSize  = 1024
	writeBufferSize = 1024
)
