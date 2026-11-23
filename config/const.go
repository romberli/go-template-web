/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"net/http"

	"github.com/romberli/go-util/constant"
)

// global constant
const (
	DefaultCommandName = "go-template-web"
	DefaultBaseDir     = constant.CurrentDir
	// daemon
	DefaultDaemon  = false
	DaemonArg      = "--daemon"
	DaemonArgTrue  = "--daemon=true"
	DaemonArgFalse = "--daemon=false"
	// log
	DefaultLogDir          = "./log"
	MinLogMaxSize          = 1
	MaxLogMaxSize          = constant.MaxInt
	MinLogMaxDays          = 1
	MaxLogMaxDays          = constant.MaxInt
	MinLogMaxBackups       = 1
	MaxLogMaxBackups       = constant.MaxInt
	DefaultRotateOnStartup = false
	DefaultLogStdout       = false
	// server
	DefaultServerAddr                      = "0.0.0.0:80"
	DefaultServerReadTimeout               = 5
	DefaultServerWriteTimeout              = 10
	MinServerReadTimeout                   = 0
	MaxServerReadTimeout                   = 3600
	MinServerWriteTimeout                  = 1
	MaxServerWriteTimeout                  = 3600
	DefaultServerPProfEnabled              = false
	DefaultServerRouterAlternativeBasePath = constant.EmptyString
	DefaultServerRouterAlternativeBodyPath = constant.EmptyString
	DefaultServerRouterHTTPErrorCode       = http.StatusInternalServerError
	// db
	DefaultDBMySQLName          = "go_template_web"
	DefaultDBMySQLUser          = "root"
	DefaultDBMySQLPass          = constant.EmptyString
	DefaultDBRedisUser          = "default"
	DefaultDBRedisPass          = constant.EmptyString
	MinDBPoolMaxConnections     = 1
	MaxDBPoolMaxConnections     = constant.MaxInt
	MinDBPoolInitConnections    = 1
	MaxDBPoolInitConnections    = constant.MaxInt
	MinDBPoolMaxIdleConnections = 1
	MaxDBPoolMaxIdleConnections = constant.MaxInt
	MinDBPoolMaxIdleTime        = 1
	MaxDBPoolMaxIdleTime        = constant.MaxInt
	MinDBPoolMaxWaitTime        = -1
	MaxDBPoolMaxWaitTime        = constant.MaxInt
	MinDBPoolMaxRetryCount      = -1
	MaxDBPoolMaxRetryCount      = constant.MaxInt
	MinDBPoolKeepAliveInterval  = 1
	MaxDBPoolKeepAliveInterval  = constant.MaxInt
	// auth
	DefaultAuthEnabled         = false
	DefaultAuthTimestampWindow = 300
	MinAuthTimestampWindow     = 1
	MaxAuthTimestampWindow     = 3600
	DefaultAuthSecretTTL       = 600
	MinAuthSecretTTL           = 1
	MaxAuthSecretTTL           = 86400
	// sm
	DefaultSMGTWKey = "C0B4C03A2CC83DE33427A4B9504E60B52F31746FA641D89AF2B4A96A53B64446"
)

// configuration constant
const (
	// config
	ConfKey = "config"
	// daemon
	DaemonKey = "daemon"
	// log
	LogFileNameKey        = "log.fileName"
	LogLevelKey           = "log.level"
	LogFormatKey          = "log.format"
	LogMaxSizeKey         = "log.maxSize"
	LogMaxDaysKey         = "log.maxDays"
	LogMaxBackupsKey      = "log.maxBackups"
	LogRotateOnStartupKey = "log.rotateOnStartup"
	LogStdoutKey          = "log.stdout"
	// server
	ServerAddrKey                      = "server.addr"
	ServerIDKey                        = "server.id"
	ServerPidFileKey                   = "server.pidFile"
	ServerReadTimeoutKey               = "server.readTimeout"
	ServerWriteTimeoutKey              = "server.writeTimeout"
	ServerPProfEnabledKey              = "server.pprof.enabled"
	ServerRouterAlternativeBasePathKey = "server.router.alternativeBasePath"
	ServerRouterAlternativeBodyPathKey = "server.router.alternativeBodyPath"
	ServerRouterHTTPErrorCodeKey       = "server.router.httpErrorCode"
	// db
	DBMySQLAddrKey              = "db.mysql.addr"
	DBMySQLNameKey              = "db.mysql.name"
	DBMySQLUserKey              = "db.mysql.user"
	DBMySQLPassKey              = "db.mysql.pass"
	DBRedisAddrsKey             = "db.redis.addrs"
	DBRedisUserKey              = "db.redis.user"
	DBRedisPassKey              = "db.redis.pass"
	DBPoolMaxConnectionsKey     = "db.pool.maxConnections"
	DBPoolInitConnectionsKey    = "db.pool.initConnections"
	DBPoolMaxIdleConnectionsKey = "db.pool.maxIdleConnections"
	DBPoolMaxIdleTimeKey        = "db.pool.maxIdleTime"
	DBPoolMaxWaitTimeKey        = "db.pool.maxWaitTime"
	DBPoolMaxRetryCountKey      = "db.pool.maxRetryCount"
	DBPoolKeepAliveIntervalKey  = "db.pool.keepAliveInterval"
	// auth
	AuthEnabledKey         = "auth.enabled"
	AuthTimestampWindowKey = "auth.timestampWindow"
	AuthSecretTTLKey       = "auth.secretTTL"
	// sm
	SMGTWKeyKey = "sm.gtwKey"
	// env
	EnvPodNameKey = "MY_POD_NAME"
	EnvPodIPKey   = "MY_POD_IP"
)
