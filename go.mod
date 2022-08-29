module github.com/romberli/go-template-web

go 1.16

replace github.com/spf13/pflag v1.0.5 => github.com/romberli/pflag v1.0.6-alpha

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/asaskevich/govalidator v0.0.0-20200819183940-29e1ff8eb0bb
	github.com/gin-gonic/gin v1.7.7
	github.com/pingcap/errors v0.11.5-0.20211224045212-9687c2b0f87c
	github.com/romberli/go-multierror v1.1.2
	github.com/romberli/go-util v0.3.16
	github.com/romberli/log v1.0.24
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	go.uber.org/zap v1.19.1
)
