package message

import (
	"github.com/romberli/go-util/config"
)

const (
	ErrPrintHelpInfo                           = 400001
	ErrEmptyLogFileName                        = 400002
	ErrNotValidLogFileName                     = 400003
	ErrNotValidLogLevel                        = 400004
	ErrNotValidLogFormat                       = 400005
	ErrNotValidLogMaxSize                      = 400006
	ErrNotValidLogMaxDays                      = 400007
	ErrNotValidLogMaxBackups                   = 400008
	ErrNotValidServerPort                      = 400009
	ErrNotValidPidFile                         = 400010
	ErrValidateConfig                          = 400011
	ErrInitDefaultConfig                       = 400012
	ErrReadConfigFile                          = 400013
	ErrOverrideCommandLineArgs                 = 400014
	ErrAbsoluteLogFilePath                     = 400015
	ErrInitLogger                              = 400016
	ErrBaseDir                                 = 400017
	ErrInitConfig                              = 400018
	ErrCheckServerPid                          = 400019
	ErrCheckServerRunningStatus                = 400020
	ErrServerIsRunning                         = 400021
	ErrStartAsForeground                       = 400022
	ErrSavePidToFile                           = 400023
	ErrKillServerWithPid                       = 400024
	ErrKillServerWithPidFile                   = 400025
	ErrGetPidFromPidFile                       = 400026
	ErrSetSid                                  = 400027
	ErrRemovePidFile                           = 400028
	ErrNotValidDBAddr                          = 400029
	ErrNotValidDBName                          = 400030
	ErrNotValidDBUser                          = 400031
	ErrNotValidDBPass                          = 400032
	ErrNotValidDBPoolMaxConnections            = 400033
	ErrNotValidDBPoolInitConnections           = 400034
	ErrNotValidDBPoolMaxIdleConnections        = 400035
	ErrNotValidDBPoolMaxIdleTime               = 400036
	ErrNotValidDBPoolKeepAliveInterval         = 400037
	ErrInitConnectionPool                      = 400038
	ErrNotValidServerReadTimeout               = 400039
	ErrNotValidServerWriteTimeout              = 400040
	ErrNotValidServerAddr                      = 400041
	ErrNotValidServerRouterAlternativeBasePath = 400042
	ErrNotValidServerRouterHTTPErrorCode       = 400043
	ErrFieldNotExists                          = 400044
	ErrGetRawData                              = 400045
	ErrUnmarshalRawData                        = 400046
	ErrGenerateNewMapWithTag                   = 400047
	ErrMarshalData                             = 400048
	ErrTypeConversion                          = 400049
	ErrFieldNotExistsOrWrongType               = 400050
	ErrNotValidTimeLayout                      = 400051
	ErrNotValidTimeDuration                    = 400052
	ErrGetResponseCode                         = 400053
	ErrSetResponseCode                         = 400054
	ErrGinRecovery                             = 400055
)

func initErrorMessage() {
	Messages[ErrPrintHelpInfo] = config.NewErrMessage(DefaultMessageHeader, ErrPrintHelpInfo, "got message when printing help information")
	Messages[ErrEmptyLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyLogFileName, "log file name could not be an empty string")
	Messages[ErrNotValidLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFileName, "log file name must be either unix or windows path format, %s is not valid")
	Messages[ErrNotValidLogLevel] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogLevel, "log level must be one of [debug, info, warn, message, fatal], %s is not valid")
	Messages[ErrNotValidLogFormat] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFormat, "log format must be either text or json, %s is not valid")
	Messages[ErrNotValidLogMaxSize] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxSize, "log max size must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxDays] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxDays, "log max days must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxBackups] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxBackups, "log max backups must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerPort] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerPort, "Server port must be between %d and %d, %d is not valid")
	Messages[ErrNotValidPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidPidFile, "pid file name must be either unix or windows path format, %s is not valid")
	Messages[ErrValidateConfig] = config.NewErrMessage(DefaultMessageHeader, ErrValidateConfig, "validate config failed")
	Messages[ErrInitDefaultConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitDefaultConfig, "init default configuration failed")
	Messages[ErrReadConfigFile] = config.NewErrMessage(DefaultMessageHeader, ErrReadConfigFile, "read config file failed")
	Messages[ErrOverrideCommandLineArgs] = config.NewErrMessage(DefaultMessageHeader, ErrOverrideCommandLineArgs, "override command line arguments failed")
	Messages[ErrAbsoluteLogFilePath] = config.NewErrMessage(DefaultMessageHeader, ErrAbsoluteLogFilePath, "get absolute path of log file failed. log file: %s")
	Messages[ErrInitLogger] = config.NewErrMessage(DefaultMessageHeader, ErrInitLogger, "initialize logger failed")
	Messages[ErrBaseDir] = config.NewErrMessage(DefaultMessageHeader, ErrBaseDir, "get base dir of %s failed")
	Messages[ErrInitConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitConfig, "init config failed")
	Messages[ErrCheckServerPid] = config.NewErrMessage(DefaultMessageHeader, ErrCheckServerPid, "check server pid file failed")
	Messages[ErrCheckServerRunningStatus] = config.NewErrMessage(DefaultMessageHeader, ErrCheckServerRunningStatus, "check server running status failed")
	Messages[ErrServerIsRunning] = config.NewErrMessage(DefaultMessageHeader, ErrServerIsRunning, "pid file exists or server is still running. pid file: %s")
	Messages[ErrStartAsForeground] = config.NewErrMessage(DefaultMessageHeader, ErrStartAsForeground, "got message when starting das as foreground")
	Messages[ErrSavePidToFile] = config.NewErrMessage(DefaultMessageHeader, ErrSavePidToFile, "got message when writing pid to file")
	Messages[ErrKillServerWithPid] = config.NewErrMessage(DefaultMessageHeader, ErrKillServerWithPid, "kill server failed. pid: %d")
	Messages[ErrKillServerWithPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrKillServerWithPidFile, "kill server with pid file failed. pid: %d, pid file: %s")
	Messages[ErrGetPidFromPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrGetPidFromPidFile, "get pid from pid file failed. pid file: %s")
	Messages[ErrSetSid] = config.NewErrMessage(DefaultMessageHeader, ErrSetSid, "set sid failed when daemonizing server")
	Messages[ErrRemovePidFile] = config.NewErrMessage(DefaultMessageHeader, ErrRemovePidFile, "remove pid file failed. pid file: %s")
	Messages[ErrNotValidDBAddr] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBAddr, "database address must be formatted as host:port, %s is not valid")
	Messages[ErrNotValidDBName] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBName, "database name must be a string, %s is not valid")
	Messages[ErrNotValidDBUser] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBUser, "database user name must be a string, %s is not valid")
	Messages[ErrNotValidDBPass] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPass, "database password must be a string, %s is not valid")
	Messages[ErrNotValidDBPoolMaxConnections] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolMaxConnections, "database max connections must be between %d and %d, %d is not valid")
	Messages[ErrNotValidDBPoolInitConnections] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolInitConnections, "database initial connections must be between %d and %d, %d is not valid")
	Messages[ErrNotValidDBPoolMaxIdleConnections] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolMaxIdleConnections, "database max idle connections must be between %d and %d, %d is not valid")
	Messages[ErrNotValidDBPoolMaxIdleTime] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolMaxIdleTime, "database max idle time must be between %d and %d, %d is not valid")
	Messages[ErrNotValidDBPoolKeepAliveInterval] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidDBPoolKeepAliveInterval, "database keep alive interval must be between %d and %d, %d is not valid")
	Messages[ErrInitConnectionPool] = config.NewErrMessage(DefaultMessageHeader, ErrInitConnectionPool, "init connection pool failed.")
	Messages[ErrNotValidServerReadTimeout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerReadTimeout, "server read timeout must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerWriteTimeout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerWriteTimeout, "server write timeout must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerAddr] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerAddr, "server addr must be formatted as host:port, %s is not valid")
	Messages[ErrFieldNotExists] = config.NewErrMessage(DefaultMessageHeader, ErrFieldNotExists, "field does not exist. field name: %s")
	Messages[ErrGetRawData] = config.NewErrMessage(DefaultMessageHeader, ErrGetRawData, "get raw data from http body failed")
	Messages[ErrUnmarshalRawData] = config.NewErrMessage(DefaultMessageHeader, ErrUnmarshalRawData, "unmarshal raw data failed")
	Messages[ErrGenerateNewMapWithTag] = config.NewErrMessage(DefaultMessageHeader, ErrGenerateNewMapWithTag, "generate new map with tag %s failed")
	Messages[ErrMarshalData] = config.NewErrMessage(DefaultMessageHeader, ErrMarshalData, "marshal service failed")
	Messages[ErrTypeConversion] = config.NewErrMessage(DefaultMessageHeader, ErrTypeConversion, "type conversion failed")
	Messages[ErrFieldNotExistsOrWrongType] = config.NewErrMessage(DefaultMessageHeader, ErrFieldNotExistsOrWrongType, "field does not exist or is wrong type. field name: %s")
	Messages[ErrNotValidTimeLayout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidTimeLayout, "time layout must be formatted as yyyy-MM-dd HH:mm:ss, %s is not valid")
	Messages[ErrNotValidTimeDuration] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidTimeDuration, "time duration must be formatted, e.g. such as 300ms, -1.5h or 2h45m, %s is not valid")
	Messages[ErrGinRecovery] = config.NewErrMessage(DefaultMessageHeader, ErrGinRecovery, "gin Recovery: panic recovered")
}
