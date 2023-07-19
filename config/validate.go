package config

import (
	"path/filepath"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/pingcap/errors"
	"github.com/romberli/go-multierror"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/romberli/go-template-web/pkg/message"
)

var (
	ValidLogLevels                  = []string{"debug", "info", "warn", "warning", "error", "fatal"}
	ValidLogFormats                 = []string{"text", "json"}
	ValidServerRouterHTTPErrorCodes = []int{200, 500}
)

// ValidateConfig validates if the configuration is valid
func ValidateConfig() (err error) {
	merr := &multierror.Error{}

	// validate daemon section
	err = ValidateDaemon()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	// validate log section
	err = ValidateLog()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	// validate server section
	err = ValidateServer()
	if err != nil {
		merr = multierror.Append(merr, err)
	}

	return errors.Trace(merr.ErrorOrNil())
}

// ValidateDaemon validates if daemon section is valid
func ValidateDaemon() error {
	_, err := cast.ToBoolE(viper.Get(DaemonKey))

	return errors.Trace(err)
}

// ValidateLog validates if log section is valid.
func ValidateLog() error {
	var valid bool

	merr := &multierror.Error{}

	// validate log.FileName
	logFileName, err := cast.ToStringE(viper.Get(LogFileNameKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	logFileName = strings.TrimSpace(logFileName)
	if logFileName == constant.EmptyString {
		merr = multierror.Append(merr, message.NewMessage(message.ErrEmptyLogFileName))
	}
	isAbs := filepath.IsAbs(logFileName)
	if !isAbs {
		logFileName, err = filepath.Abs(logFileName)
		if err != nil {
			merr = multierror.Append(merr, errors.Trace(err))
		}
	}
	valid, _ = govalidator.IsFilePath(logFileName)
	if !valid {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogFileName, logFileName))
	}

	// validate log.level
	logLevel, err := cast.ToStringE(viper.Get(LogLevelKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	valid, err = common.ElementInSlice(ValidLogLevels, logLevel)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if !valid {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogLevel, logLevel))
	}

	// validate log.format
	logFormat, err := cast.ToStringE(viper.Get(LogFormatKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	valid, err = common.ElementInSlice(ValidLogFormats, logFormat)
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	if !valid {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogFormat, logFormat))
	}

	// validate log.maxSize
	logMaxSize, err := cast.ToIntE(viper.Get(LogMaxSizeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if logMaxSize < MinLogMaxSize || logMaxSize > MaxLogMaxSize {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogMaxSize, MinLogMaxSize, MaxLogMaxSize, logMaxSize))
	}

	// validate log.maxDays
	logMaxDays, err := cast.ToIntE(viper.Get(LogMaxDaysKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if logMaxDays < MinLogMaxDays || logMaxDays > MaxLogMaxDays {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogMaxDays, MinLogMaxDays, MaxLogMaxDays, logMaxDays))
	}

	// validate log.maxBackups
	logMaxBackups, err := cast.ToIntE(viper.Get(LogMaxBackupsKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if logMaxBackups < MinLogMaxDays || logMaxBackups > MaxLogMaxDays {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogMaxBackups, MinLogMaxBackups, MaxLogMaxBackups, logMaxBackups))
	}

	// validate log.rotateOnStartup
	_, err = cast.ToBoolE(viper.Get(LogRotateOnStartupKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	// validate log.stdout
	_, err = cast.ToBoolE(viper.Get(LogStdoutKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	return merr.ErrorOrNil()
}

// ValidateServer validates if server section is valid
func ValidateServer() error {
	merr := &multierror.Error{}

	// validate server.addr
	serverAddr, err := cast.ToStringE(viper.Get(ServerAddrKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	serverAddrList := strings.Split(serverAddr, constant.ColonString)

	switch len(serverAddrList) {
	case 2:
		port := serverAddrList[1]
		if !govalidator.IsPort(port) {
			merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidServerPort, constant.MinPort, constant.MaxPort, port))
		}
	default:
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidServerAddr, serverAddr))
	}

	// validate server.pidFile
	serverPidFile, err := cast.ToStringE(viper.Get(ServerPidFileKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	isAbs := filepath.IsAbs(serverPidFile)
	if !isAbs {
		serverPidFile, err = filepath.Abs(serverPidFile)
		if err != nil {
			merr = multierror.Append(merr, errors.Trace(err))
		}
	}
	ok, _ := govalidator.IsFilePath(serverPidFile)
	if !ok {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidPidFile, serverPidFile))
	}

	// validate server.readTimeout
	serverReadTimeout, err := cast.ToIntE(viper.Get(ServerReadTimeoutKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if serverReadTimeout < MinServerReadTimeout || serverReadTimeout > MaxServerReadTimeout {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidServerReadTimeout, MinServerReadTimeout, MaxServerWriteTimeout, serverReadTimeout))
	}

	// validate server.writeTimeout
	serverWriteTimeout, err := cast.ToIntE(viper.Get(ServerWriteTimeoutKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if serverWriteTimeout < MinServerWriteTimeout || serverWriteTimeout > MaxServerWriteTimeout {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidServerWriteTimeout, MinServerWriteTimeout, MaxServerWriteTimeout, serverWriteTimeout))
	}

	// validate server.pprof.enabled
	_, err = cast.ToBoolE(viper.Get(ServerPProfEnabledKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	// validate server.router.alternativeBaseURL
	serverRouterAlternativeBasePath, err := cast.ToStringE(viper.Get(ServerRouterAlternativeBasePathKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if serverRouterAlternativeBasePath != constant.EmptyString && !strings.HasPrefix(serverRouterAlternativeBasePath, constant.SlashString) {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidServerRouterAlternativeBasePath, serverRouterAlternativeBasePath))
	}

	// validate server.router.bodyPath
	_, err = cast.ToStringE(viper.Get(ServerRouterAlternativeBodyPathKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	// validate server.router.httpErrorCode
	serverRouterHttpErrorCode, err := cast.ToIntE(viper.Get(ServerRouterHTTPErrorCodeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		valid, err := common.ElementInSlice(ValidServerRouterHTTPErrorCodes, serverRouterHttpErrorCode)
		if err != nil {
			merr = multierror.Append(merr, err)
		}
		if !valid {
			merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidServerRouterHTTPErrorCode, serverRouterHttpErrorCode))
		}
	}

	return merr.ErrorOrNil()
}
