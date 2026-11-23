package cmd

import (
	"os"
	"strings"

	"github.com/pingcap/errors"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/romberli/go-template-web/config"
	"github.com/romberli/go-template-web/pkg/message"
)

// OverrideConfigByEnv overrides the configuration by environment variables, it has top priority
func OverrideConfigByEnv() error {
	podName := os.Getenv(config.EnvPodNameKey)
	viper.Set(config.EnvPodNameKey, podName)
	podIP := os.Getenv(config.EnvPodIPKey)
	viper.Set(config.EnvPodIPKey, podIP)

	// validate configuration
	err := config.ValidateConfig()
	if err != nil {
		return message.NewMessage(message.ErrValidateConfig, err)
	}

	return nil
}

// OverrideConfigByCLI read configuration from command line interface, it will override the config file configuration
func OverrideConfigByCLI() error {
	// override config
	overrideConfigByCLI()
	// override daemon
	err := overrideDaemonByCLI()
	if err != nil {
		return err
	}
	// override log
	err = overrideLogByCLI()
	if err != nil {
		return err
	}
	// override server
	err = overrideServerByCLI()
	if err != nil {
		return err
	}
	// override db
	overrideDBByCLI()
	// override auth
	err = overrideAuthByCLI()
	if err != nil {
		return err
	}
	// override sm
	overrideSMByCLI()
	// validate configuration
	err = config.ValidateConfig()
	if err != nil {
		return message.NewMessage(message.ErrValidateConfig, err)
	}

	return nil
}

// overrideConfigByCLI overrides the config section by command line interface
func overrideConfigByCLI() {
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.Set(config.ConfKey, cfgFile)
	}
}

// overrideDaemonByCLI overrides the daemon section by command line interface
func overrideDaemonByCLI() error {
	if daemonStr != constant.DefaultRandomString {
		daemon, err := cast.ToBoolE(daemonStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.DaemonKey, daemon)
	}

	return nil
}

// overrideLogByCLI overrides the log section by command line interface
func overrideLogByCLI() error {
	if logFileName != constant.DefaultRandomString {
		viper.Set(config.LogFileNameKey, logFileName)
	}
	if logLevel != constant.DefaultRandomString {
		logLevel = strings.ToLower(logLevel)
		viper.Set(config.LogLevelKey, logLevel)
	}
	if logFormat != constant.DefaultRandomString {
		logLevel = strings.ToLower(logFormat)
		viper.Set(config.LogFormatKey, logFormat)
	}
	if logMaxSize != constant.DefaultRandomInt {
		viper.Set(config.LogMaxSizeKey, logMaxSize)
	}
	if logMaxDays != constant.DefaultRandomInt {
		viper.Set(config.LogMaxDaysKey, logMaxDays)
	}
	if logMaxBackups != constant.DefaultRandomInt {
		viper.Set(config.LogMaxBackupsKey, logMaxBackups)
	}
	if logRotateOnStartupStr != constant.DefaultRandomString {
		rotateOnStartup, err := cast.ToBoolE(logRotateOnStartupStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.LogRotateOnStartupKey, rotateOnStartup)
	}
	if logStdoutStr != constant.DefaultRandomString {
		logStdout, err := cast.ToBoolE(logStdoutStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.LogStdoutKey, logStdout)
	}

	return nil
}

// overrideServerByCLI overrides the server section by command line interface
func overrideServerByCLI() error {
	if serverAddr != constant.DefaultRandomString {
		viper.Set(config.ServerAddrKey, serverAddr)
	}
	if serverPidFile != constant.DefaultRandomString {
		viper.Set(config.ServerPidFileKey, serverPidFile)
	}
	if serverReadTimeout != constant.DefaultRandomInt {
		viper.Set(config.ServerReadTimeoutKey, serverReadTimeout)
	}
	if serverWriteTimeout != constant.DefaultRandomInt {
		viper.Set(config.ServerWriteTimeoutKey, serverWriteTimeout)
	}
	if serverPProfEnabledStr != constant.DefaultRandomString {
		pprofEnabled, err := cast.ToBoolE(serverPProfEnabledStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.ServerPProfEnabledKey, pprofEnabled)
	}
	if serverRouterAlternativeBasePath != constant.DefaultRandomString {
		viper.Set(config.ServerRouterAlternativeBasePathKey, serverRouterAlternativeBasePath)
	}
	if serverRouterAlternativeBodyPath != constant.DefaultRandomString {
		viper.Set(config.ServerRouterAlternativeBodyPathKey, serverRouterAlternativeBodyPath)
	}
	if serverRouterHTTPErrorCode != constant.DefaultRandomInt {
		viper.Set(config.ServerRouterHTTPErrorCodeKey, serverRouterHTTPErrorCode)
	}

	return nil
}

// overrideDBByCLI overrides the db section by command line interface
func overrideDBByCLI() {
	if dbMySQLAddr != constant.DefaultRandomString {
		viper.Set(config.DBMySQLAddrKey, dbMySQLAddr)
	}
	if dbMySQLName != constant.DefaultRandomString {
		viper.Set(config.DBMySQLNameKey, dbMySQLName)
	}
	if dbMySQLUser != constant.DefaultRandomString {
		viper.Set(config.DBMySQLUserKey, dbMySQLUser)
	}
	if dbMySQLPass != constant.DefaultRandomString {
		viper.Set(config.DBMySQLPassKey, dbMySQLPass)
	}
	if dbRedisAddrs != constant.DefaultRandomString {
		viper.Set(config.DBRedisAddrsKey, dbRedisAddrs)
	}
	if dbRedisUser != constant.DefaultRandomString {
		viper.Set(config.DBRedisUserKey, dbRedisUser)
	}
	if dbRedisPass != constant.DefaultRandomString {
		viper.Set(config.DBRedisPassKey, dbRedisPass)
	}
	if dbPoolMaxConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxConnectionsKey, dbPoolMaxConnections)
	}
	if dbPoolInitConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolInitConnectionsKey, dbPoolInitConnections)
	}
	if dbPoolMaxIdleConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxIdleConnectionsKey, dbPoolMaxIdleConnections)
	}
	if dbPoolMaxIdleTime != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxIdleTimeKey, dbPoolMaxIdleTime)
	}
	if dbPoolMaxWaitTime != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxWaitTimeKey, dbPoolMaxWaitTime)
	}
	if dbPoolMaxRetryCount != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxRetryCountKey, dbPoolMaxRetryCount)
	}
	if dbPoolKeepAliveInterval != constant.DefaultRandomInt {
		viper.Set(config.DBPoolKeepAliveIntervalKey, dbPoolKeepAliveInterval)
	}
}

// overrideAuthByCLI overrides the auth section by command line interface
func overrideAuthByCLI() error {
	if authEnabledStr != constant.DefaultRandomString {
		authEnabled, err := cast.ToBoolE(authEnabledStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.AuthEnabledKey, authEnabled)
	}
	if authTimestampWindow != constant.DefaultRandomInt {
		viper.Set(config.AuthTimestampWindowKey, authTimestampWindow)
	}
	if authSecretTTL != constant.DefaultRandomInt {
		viper.Set(config.AuthSecretTTLKey, authSecretTTL)
	}

	return nil
}

// overrideSMByCLI overrides the sm section by command line interface
func overrideSMByCLI() {
	if smGTWKey != constant.DefaultRandomString {
		viper.Set(config.SMGTWKeyKey, smGTWKey)
	}
}
