package ioc

import (
	igrpc "github.com/lazywoo/mercury/captcha/grpc"
	"github.com/lazywoo/mercury/pkg/grpcx"
	"github.com/lazywoo/mercury/pkg/logger"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitGRPCxServer(captchaSrv *igrpc.CaptchaServiceServer, l logger.Logger) *grpcx.Server {
	type Config struct {
		Port int    `yaml:"port"`
		Etcd string `yaml:"etcd"`
		TTL  int64  `yaml:"ttl"`
	}
	var cfg Config
	err := viper.UnmarshalKey("grpc.server", &cfg)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	captchaSrv.Register(srv)
	return grpcx.NewServer(srv, "captcha", cfg.Port, []string{cfg.Etcd}, cfg.TTL, l)
}
