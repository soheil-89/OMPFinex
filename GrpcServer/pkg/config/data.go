package config

import (
	"github.com/soheil-89/OMPFinex/GrpcServer/pkg/config/environment"
	"time"
)

var Env = newEnvData()

type envData struct {
	App       app
	Db        db
	Jwt       jwt
	Sms       sms
	Grpc      grpc
	Sentry    sentry
	Recaptcha recaptcha
	RabbitMq  rabbitMq
}
type app struct {
	TimeZone string
	Port     string
	Url      string
	GinMode  string
}
type jwt struct {
	SecretKey string
	ExpireDay time.Duration
	Key       string
}

type db struct {
	Host      string
	Port      string
	User      string
	Password  string
	DbSslMode string
	DbName    string
}

type sms struct {
	UserName string
	Password string
	From     string
	Url      string
	Content  string
}
type grpc struct {
	Port string
}
type sentry struct {
	Dsn string
}
type recaptcha struct {
	SecretKey string
	Server    string
}
type rabbitMq struct {
	Url          string
	ChannelsName struct {
		Sms string
	}
}

func newEnvData() *envData {
	env := environment.Load()
	return &envData{
		App: app{
			TimeZone: env.Get("app.TimeZone").(string),
			Port:     env.Get("app.Port").(string),
			Url:      env.Get("app.Url").(string),
			GinMode:  env.Get("app.GIN_MODE").(string),
		},
		Db: db{
			Host:      env.Get("db.host").(string),
			Port:      env.Get("db.port").(string),
			User:      env.Get("db.user").(string),
			Password:  env.Get("db.password").(string),
			DbSslMode: env.Get("db.dbSslMode").(string),
			DbName:    env.Get("db.dbName").(string),
		},
		Jwt: jwt{
			SecretKey: env.Get("jwt.SecretKey").(string),
			ExpireDay: time.Duration(env.Get("jwt.ExpireDay").(int)),
			Key:       env.Get("jwt.Key").(string),
		},
		Sms: sms{
			UserName: env.Get("sms.UserName").(string),
			Password: env.Get("sms.Password").(string),
			From:     env.Get("sms.From").(string),
			Url:      env.Get("sms.Url").(string),
			Content:  env.Get("sms.Content").(string),
		},
		Grpc: grpc{
			Port: env.Get("grpc.port").(string),
		},
		Sentry: sentry{
			Dsn: env.Get("sentry.Dsn").(string),
		},
		Recaptcha: recaptcha{
			SecretKey: env.Get("recaptcha.SecretKey").(string),
			Server:    env.Get("recaptcha.Server").(string),
		},
		RabbitMq: rabbitMq{
			Url: env.Get("rabbitMq.Url").(string),
			ChannelsName: struct{ Sms string }{
				Sms: env.Get("rabbitMq.ChannelsName.sms").(string),
			},
		},
	}
}
