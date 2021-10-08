package message

import (
	"github.com/romberli/go-util/config"
)

const (
	ErrPrintHelpInfo            = 400001
	ErrEmptyLogFileName         = 400002
	ErrNotValidLogFileName      = 400003
	ErrNotValidLogLevel         = 400004
	ErrNotValidLogFormat        = 400005
	ErrNotValidLogMaxSize       = 400006
	ErrNotValidLogMaxDays       = 400007
	ErrNotValidLogMaxBackups    = 400008
	ErrNotValidServerPort       = 400009
	ErrNotValidPidFile          = 400010
	ErrValidateConfig           = 400011
	ErrInitDefaultConfig        = 400012
	ErrReadConfigFile           = 400013
	ErrOverrideCommandLineArgs  = 400014
	ErrAbsoluteLogFilePath      = 400015
	ErrInitLogger               = 400016
	ErrBaseDir                  = 400017
	ErrInitConfig               = 400018
	ErrCheckServerPid           = 400019
	ErrCheckServerRunningStatus = 400020
	ErrServerIsRunning          = 400021
	ErrStartAsForeground        = 400022
	ErrSavePidToFile            = 400023
	ErrKillServerWithPid        = 400024
	ErrKillServerWithPidFile    = 400025
	ErrGetPidFromPidFile        = 400026
	ErrSetSid                   = 400027
	ErrRemovePidFile            = 400028
)

func initErrorMessage() {
	Messages[ErrPrintHelpInfo] = config.NewErrMessage(DefaultMessageHeader, ErrPrintHelpInfo, "got message when printing help information")
	Messages[ErrEmptyLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyLogFileName, "Log file name could not be an empty string")
	Messages[ErrNotValidLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFileName, "Log file name must be either unix or windows path format, %s is not valid")
	Messages[ErrNotValidLogLevel] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogLevel, "Log level must be one of [debug, info, warn, message, fatal], %s is not valid")
	Messages[ErrNotValidLogFormat] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFormat, "Log level must be either text or json, %s is not valid")
	Messages[ErrNotValidLogMaxSize] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxSize, "Log max size must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxDays] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxDays, "Log max days must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxBackups] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxBackups, "Log max backups must be between %d and %d, %d is not valid")
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
	Messages[ErrRemovePidFile] = config.NewErrMessage(DefaultMessageHeader, ErrRemovePidFile, "remove pid file failed")
}
