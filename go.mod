module github.com/romberli/go-template

go 1.16

replace github.com/spf13/pflag v1.0.5 => github.com/romberli/pflag v1.0.6-alpha

require (
	github.com/asaskevich/govalidator v0.0.0-20200819183940-29e1ff8eb0bb
	github.com/hashicorp/go-multierror v1.1.0
	github.com/romberli/go-util v0.3.11-0.20210902115345-b4398d4d2d7d
	github.com/romberli/log v1.0.20
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
)
