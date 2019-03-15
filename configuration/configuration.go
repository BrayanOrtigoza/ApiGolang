//
package configuration

import (
    "bytes"
    "os"
    "strconv"

    "API/constants"
)

var (
    //
    Debug bool

    //
    UrlRedisServer      string
    PasswordRedisServer string
    DatabaseRedisServer int

    //
    UrlRethinkDBServer      string
    UserRethinkDBServer     string
    PasswordRethinkDBServer string
    DatabaseRethinkDBServer string

    //
    KeyTokenRedis string

    //
    EndpointWebSocket string
    NameToken         string
    ReadBufferSize    int
    WriteBufferSize   int
    UrlWebSocket      string
)

//
func concatHostPort(host string, port string) string {
    var buffer bytes.Buffer
    buffer.WriteString(host)
    buffer.WriteString(constants.Colon)
    buffer.WriteString(port)
    urlRedisServer := buffer.String()
    return urlRedisServer
}

//
func verifyUrlServer(hostServer string, portServer string, defaultPort string ) (string, string) {
    if hostServer == constants.Empty {
        hostServer = constants.DefaultHost
    }
    if portServer == constants.Empty {
        portServer = defaultPort
    }
    return hostServer, portServer
}

//
func verifyModeDebug(modeDebug bool, errorDebug error) bool {
    if errorDebug == nil {
        return modeDebug
    } else {
        return constants.NoDebug
    }
}

//
func readModeDebug() {
    modeDebug := os.Getenv(nameEnviromentDebug)
    Debug = verifyModeDebug(strconv.ParseBool(modeDebug))
}

//
func readConfigurationRedis() {
    hostRedisServer := os.Getenv(nameHostRedis)
    portRedisServer := os.Getenv(namePortRedis)
    PasswordRedisServer = os.Getenv(namePasswordRedis)
    DatabaseRedisServer = databaseRedis
    hostRedisServer, portRedisServer = verifyUrlServer(hostRedisServer, portRedisServer, constants.DefaultPortRedis)
    UrlRedisServer = concatHostPort(hostRedisServer, portRedisServer)
}

//
func readConfigurationRethinkDB() {
    hostRethinkDBServer := os.Getenv(nameHostRethinkDB)
    portRethinkDBServer := os.Getenv(namePortRethinkDB)
    PasswordRethinkDBServer = os.Getenv(namePasswordRethinkDB)
    DatabaseRethinkDBServer = databaseRethinkDB
    hostRethinkDBServer, portRethinkDBServer = verifyUrlServer(hostRethinkDBServer, portRethinkDBServer, constants.DefaultPortRethinkDB)
    UrlRethinkDBServer = concatHostPort(hostRethinkDBServer, portRethinkDBServer)
}

//
func Init() {
    readModeDebug()
    readConfigurationRedis()
    readConfigurationRethinkDB()
}
