module github.com/romberli/go-template-web

go 1.16

replace github.com/spf13/pflag v1.0.5 => github.com/romberli/pflag v1.0.6-alpha

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/buger/jsonparser v1.1.1
	github.com/gin-contrib/pprof v1.4.0
	github.com/gin-gonic/gin v1.9.1
	github.com/pingcap/errors v0.11.5-0.20211224045212-9687c2b0f87c
	github.com/romberli/go-multierror v1.1.2
	github.com/romberli/go-util v0.3.17-0.20221219110924-82346681020c
	github.com/romberli/log v1.0.25
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.16.0
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.1
	go.uber.org/zap v1.24.0
)
