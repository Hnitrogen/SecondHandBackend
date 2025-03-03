package main

import (
	"context"
	"flag"
	"os"
	"time"

	"stuff/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"

	_ "go.uber.org/automaxprocs"

	userpb "stuff/api/other/user/v1"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "stuffService"
	// Version is the version of the compiled software.
	Version string = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	id string = "stuffService-1"
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	// 从环境变量或配置文件中读取 Consul 地址
	consulAddr := os.Getenv("CONSUL_ADDR")
	if consulAddr == "" {
		consulAddr = "127.0.0.1:8500" // 默认地址
	}

	cfg := api.DefaultConfig()
	cfg.Address = consulAddr

	// 添加重试逻辑
	var client *api.Client
	var err error
	for i := 0; i < 3; i++ {
		client, err = api.NewClient(cfg)
		if err == nil {
			break
		}
		logger.Log(log.LevelWarn, "msg", "Failed to connect to Consul, retrying...", "error", err)
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		panic(err)
	}
	logger.Log(log.LevelInfo, "msg", "Consul client initialized", "address", consulAddr)

	reg := consul.New(client)
	logger.Log(log.LevelInfo, "msg", "Consul registrar initialized")

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{
			"env":    os.Getenv("APP_ENV"),
			"region": os.Getenv("APP_REGION"),
		}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(reg),
	)
}

// UserRpcClient 连接依赖
func createUserClient(logger log.Logger, client *api.Client) userpb.UserClient {
	endpoint := "discovery:///userService"

	// 创建 Consul 服务发现
	dis := consul.New(client)

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis), // 添加 Consul 服务发现
	)
	if err != nil {
		logger.Log(log.LevelError, "msg", "Failed to create gRPC connection", "error", err)
		panic(err)
	}

	logger.Log(log.LevelInfo, "msg", "gRPC client initialized", "endpoint", endpoint)
	return userpb.NewUserClient(conn)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// 创建 Consul 客户端
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// 创建 UserClient
	userClient := createUserClient(logger, consulClient)

	// 初始化应用
	app, cleanup, err := wireApp(bc.Server, bc.Data, logger, userClient)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 启动并等待停止信号
	if err := app.Run(); err != nil {
		panic(err)
	}
}
