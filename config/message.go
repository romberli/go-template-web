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

var (
	// info
	InfoServerStart      = "go-template-200001: go-template started successfully. pid: %d, pid file: %s"
	InfoServerStop       = "go-template-200002: go-template stopped successfully. pid: %d, pid file: %s"
	InfoServerIsRunning  = "go-template-200003: go-template is running. pid: %d"
	InfoServerNotRunning = "go-template-200004: go-template is not running.pid: %d"

	// error
	ErrPrintHelpInfo            = "go-template-400001: got error when printing help information.\n%s"
	ErrEmptyLogFileName         = "go-template-400002: Log file name could not be an empty string."
	ErrNotValidLogFileName      = "go-template-400003: Log file name must be either unix or windows path format, %s is not valid."
	ErrNotValidLogLevel         = "go-template-400004: Log level must be one of [debug, info, warn, error, fatal], %s is not valid."
	ErrNotValidLogFormat        = "go-template-400005: Log level must be either text or json, %s is not valid."
	ErrNotValidLogMaxSize       = "go-template-400006: Log max size must be between %d and %d, %d is not valid."
	ErrNotValidLogMaxDays       = "go-template-400007: Log max days must be between %d and %d, %d is not valid."
	ErrNotValidLogMaxBackups    = "go-template-400008: Log max backups must be between %d and %d, %d is not valid."
	ErrNotValidServerPort       = "go-template-400009: Server port must be between %d and %d, %d is not valid."
	ErrNotValidPidFile          = "go-template-400010: pid file name must be either unix or windows path format, %s is not valid."
	ErrValidateConfig           = "go-template-400011: validate config failed.\n%s"
	ErrInitDefaultConfig        = "go-template-400012: init default configuration failed.\n%s"
	ErrReadConfigFile           = "go-template-400013: read config file failed.\n%s"
	ErrOverrideCommandLineArgs  = "go-template-400014: override command line arguments failed.\n%s"
	ErrAbsoluteLogFilePath      = "go-template-400015: get absolute path of log file failed. log file: %s\n%s"
	ErrInitLogger               = "go-template-400016: initialize logger failed.\n%s"
	ErrBaseDir                  = "go-template-400017: get base dir of %s failed.\n%s"
	ErrInitConfig               = "go-template-400018: init config failed.\n%s"
	ErrCheckServerPid           = "go-template-400019: check server pid file failed.\n%s"
	ErrCheckServerRunningStatus = "go-template-400020: check server running status failed.\n%s"
	ErrServerIsRunning          = "go-template-400021: pid file exists or server is still running. pid file: %s."
	ErrStartAsForeground        = "go-template-400022: got error when starting go-template as foreground.\n%s"
	ErrSavePidToFile            = "go-template-400023: got error when writing pid to file.\n%s"
	ErrKillServerWithPid        = "go-template-400024: kill server failed. pid: %d\n%s"
	ErrKillServerWithPidFile    = "go-template-400025: kill server with pid file failed. pid: %d\n, pid file: %s.\n%s"
	ErrGetPidFromPidFile        = "go-template-400026: get pid from pid file failed. pid file: %s\n%s"
	ErrSetSid                   = "go-template-400027: set sid failed when daemonizing server.\n%s"
	ErrRemovePidFile            = "go-template-400028: remove pid file failed.\n%s"
)
