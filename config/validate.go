package config

import (
	"path/filepath"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/pingcap/errors"
	"github.com/romberli/go-multierror"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/crypto"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/romberli/go-template-web/pkg/message"
)

var (
	ValidLogLevels                  = []string{"debug", "info", "warn", "warning", "error", "fatal"}
	ValidLogFormats                 = []string{"text", "json"}
	ValidServerRouterHTTPErrorCodes = []int{200, 500}

	dbStringItems = []string{DBMySQLNameKey, DBMySQLUserKey, DBMySQLPassKey, DBRedisUserKey, DBRedisPassKey}
)

const defaultTestMessage = "hello, world1"

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

	// validate db section
	err = ValidateDB()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate auth section
	err = ValidateAuth()
	if err != nil {
		merr = multierror.Append(merr, err)
	}
	// validate sm section
	err = ValidateSM()
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
	if !common.ElementInSlice(ValidLogLevels, logLevel) {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidLogLevel, logLevel))
	}

	// validate log.format
	logFormat, err := cast.ToStringE(viper.Get(LogFormatKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}
	if !common.ElementInSlice(ValidLogFormats, logFormat) {
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
		if !common.ElementInSlice(ValidServerRouterHTTPErrorCodes, serverRouterHttpErrorCode) {
			merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidServerRouterHTTPErrorCode, serverRouterHttpErrorCode))
		}
	}

	return merr.ErrorOrNil()
}

// ValidateDB validates if db section is valid
func ValidateDB() error {
	merr := &multierror.Error{}

	// validate db.mysql.addr
	dbMySQLAddr, err := cast.ToStringE(viper.Get(DBMySQLAddrKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		e := validateAddr(dbMySQLAddr)
		if e != nil {
			merr = multierror.Append(merr, e)
		}
	}

	// validate db.redis.addrs, empty means redis is disabled
	redisAddrs, err := cast.ToStringSliceE(viper.Get(DBRedisAddrsKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		for _, addr := range redisAddrs {
			e := validateAddr(addr)
			if e != nil {
				merr = multierror.Append(merr, e)
			}
		}
	}

	// validate db.pool.maxConnections
	maxConnections, err := cast.ToIntE(viper.Get(DBPoolMaxConnectionsKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if maxConnections < MinDBPoolMaxConnections || maxConnections > MaxDBPoolMaxConnections {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidDBPoolMaxConnections, MinDBPoolMaxConnections, MaxDBPoolMaxConnections, maxConnections))
	}

	// validate db.pool.initConnections
	initConnections, err := cast.ToIntE(viper.Get(DBPoolInitConnectionsKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if initConnections < MinDBPoolInitConnections || initConnections > MaxDBPoolInitConnections {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidDBPoolInitConnections, MinDBPoolInitConnections, MaxDBPoolInitConnections, initConnections))
	}

	// validate db.pool.maxIdleConnections
	maxIdleConnections, err := cast.ToIntE(viper.Get(DBPoolMaxIdleConnectionsKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if maxIdleConnections < MinDBPoolMaxIdleConnections || maxIdleConnections > MaxDBPoolMaxIdleConnections {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidDBPoolMaxIdleConnections, MinDBPoolMaxIdleConnections, MaxDBPoolMaxIdleConnections, maxIdleConnections))
	}

	// validate db.pool.maxIdleTime
	maxIdleTime, err := cast.ToIntE(viper.Get(DBPoolMaxIdleTimeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if maxIdleTime < MinDBPoolMaxIdleTime || maxIdleTime > MaxDBPoolMaxIdleTime {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidDBPoolMaxIdleTime, MinDBPoolMaxIdleTime, MaxDBPoolMaxIdleTime, maxIdleTime))
	}

	// validate db.pool.maxWaitTime
	maxWaitTime, err := cast.ToIntE(viper.Get(DBPoolMaxWaitTimeKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if maxWaitTime < MinDBPoolMaxWaitTime || maxWaitTime > MaxDBPoolMaxWaitTime {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidDBPoolMaxWaitTime, MinDBPoolMaxWaitTime, MaxDBPoolMaxWaitTime, maxWaitTime))
	}

	// validate db.pool.maxRetryCount
	maxRetryCount, err := cast.ToIntE(viper.Get(DBPoolMaxRetryCountKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if maxRetryCount < MinDBPoolMaxRetryCount || maxRetryCount > MaxDBPoolMaxRetryCount {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidDBPoolMaxRetryCount, MinDBPoolMaxRetryCount, MaxDBPoolMaxRetryCount, maxRetryCount))
	}

	// validate db.pool.keepAliveInterval
	keepAliveInterval, err := cast.ToIntE(viper.Get(DBPoolKeepAliveIntervalKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if keepAliveInterval < MinDBPoolKeepAliveInterval || keepAliveInterval > MaxDBPoolKeepAliveInterval {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidDBPoolKeepAliveInterval, MinDBPoolKeepAliveInterval, MaxDBPoolKeepAliveInterval, keepAliveInterval))
	}

	// validate string items
	for _, item := range dbStringItems {
		_, e := cast.ToStringE(viper.Get(item))
		if e != nil {
			merr = multierror.Append(merr, errors.Trace(e))
		}
	}

	return errors.Trace(merr.ErrorOrNil())
}

// validateAddr validates if the address is formatted as host:port
func validateAddr(addr string) error {
	addrList := strings.Split(addr, constant.ColonString)
	if len(addrList) != 2 {
		return message.NewMessage(message.ErrNotValidDBAddr, addr)
	}
	if !govalidator.IsHost(strings.TrimSpace(addrList[0])) || !govalidator.IsPort(strings.TrimSpace(addrList[1])) {
		return message.NewMessage(message.ErrNotValidDBAddr, addr)
	}

	return nil
}

// validateSM2Key validates if the key is a usable sm2 private key hex string
// by encrypting and decrypting a test message
func validateSM2Key(key string) error {
	sm2, err := crypto.NewSM2WithPrivateKeyHexString(key)
	if err != nil {
		return message.NewMessage(message.ErrNotValidSMGTWKey, err, key)
	}
	cipher, err := sm2.Encrypt(defaultTestMessage)
	if err != nil {
		return message.NewMessage(message.ErrNotValidSMGTWKey, err, key)
	}
	text, err := sm2.Decrypt(cipher)
	if err != nil {
		return message.NewMessage(message.ErrNotValidSMGTWKey, err, key)
	}
	if text != defaultTestMessage {
		return message.NewMessage(message.ErrNotValidSMGTWKey, key)
	}

	return nil
}

// ValidateSM validates if sm section is valid
func ValidateSM() error {
	merr := &multierror.Error{}

	// validate sm.gtwKey, it must be a valid sm2 private key hex string
	smGTWKey, err := cast.ToStringE(viper.Get(SMGTWKeyKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else {
		e := validateSM2Key(smGTWKey)
		if e != nil {
			merr = multierror.Append(merr, e)
		}
	}

	return merr.ErrorOrNil()
}

// ValidateAuth validates if auth section is valid
func ValidateAuth() error {
	merr := &multierror.Error{}

	// validate auth.enabled
	authEnabled, err := cast.ToBoolE(viper.Get(AuthEnabledKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	}

	// validate auth.timestampWindow
	timestampWindow, err := cast.ToIntE(viper.Get(AuthTimestampWindowKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if timestampWindow < MinAuthTimestampWindow || timestampWindow > MaxAuthTimestampWindow {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidAuthTimestampWindow, MinAuthTimestampWindow, MaxAuthTimestampWindow, timestampWindow))
	}

	// validate auth.secretTTL
	secretTTL, err := cast.ToIntE(viper.Get(AuthSecretTTLKey))
	if err != nil {
		merr = multierror.Append(merr, errors.Trace(err))
	} else if secretTTL < MinAuthSecretTTL || secretTTL > MaxAuthSecretTTL {
		merr = multierror.Append(merr, message.NewMessage(message.ErrNotValidAuthSecretTTL, MinAuthSecretTTL, MaxAuthSecretTTL, secretTTL))
	}

	// auth relies on the client secret table in mysql and the nonce store in redis
	if authEnabled {
		redisAddrs, err := cast.ToStringSliceE(viper.Get(DBRedisAddrsKey))
		if err != nil {
			merr = multierror.Append(merr, errors.Trace(err))
		} else if len(redisAddrs) == constant.ZeroInt {
			merr = multierror.Append(merr, message.NewMessage(message.ErrAuthRedisRequired))
		}
	}

	return merr.ErrorOrNil()
}
