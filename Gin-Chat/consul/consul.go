package consul

import (
	config "SecondHandBackend-master/Gin-Chat/config"
	"fmt"

	stuffpb "SecondHandBackend-master/Gin-Chat/api/other/stuff/v1/SecondHandBackend-master/api/other/stuff/v1"
	userpb "SecondHandBackend-master/Gin-Chat/api/other/user/v1/SecondHandBackend-master/api/other/user/v1"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = config.ConsulAddr

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           config.ConsulHealth,
		Timeout:                        "5s",
		Interval:                       "60s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

// 根据ID获取服务
func GetServiceByID(id string) {
	cfg := api.DefaultConfig()
	cfg.Address = config.ConsulAddr

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == %s", id))
	if err != nil {
		panic(err)
	}

	svrHost := ""
	svrPort := 0
	for key, value := range data {
		svrHost = value.Address
		svrPort = value.Port
		fmt.Println("获取到服务：" + key)
		fmt.Println("IP:PORT", svrHost, svrPort)

	}
}

func CreateUserClient() (userpb.UserClient, error) {
	// 从 Consul 获取服务信息
	cfg := api.DefaultConfig()
	cfg.Address = config.ConsulAddr

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create consul client: %v", err)
	}

	// 获取用户服务的信息
	services, err := client.Agent().ServicesWithFilter(`Service == "userService"`)
	if err != nil {
		return nil, fmt.Errorf("failed to get service from consul: %v", err)
	}

	// 获取服务地址
	var serviceAddress string
	for _, service := range services {
		serviceAddress = fmt.Sprintf("%s:%d", service.Address, service.Port)
		break // 使用第一个可用的服务实例
	}

	if serviceAddress == "" {
		return nil, fmt.Errorf("no available user service found")
	}

	// 创建 gRPC 连接
	cc, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection: %v", err)
	}

	return userpb.NewUserClient(cc), nil
}

func CreateStuffClient() (stuffpb.StuffClient, error) {
	// 从 Consul 获取服务信息
	cfg := api.DefaultConfig()
	cfg.Address = config.ConsulAddr

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create consul client: %v", err)
	}

	// 获取用户服务的信息
	services, err := client.Agent().ServicesWithFilter(`Service == "stuffService"`)
	if err != nil {
		return nil, fmt.Errorf("failed to get service from consul: %v", err)
	}

	// 获取服务地址
	var serviceAddress string
	for _, service := range services {
		serviceAddress = fmt.Sprintf("%s:%d", service.Address, service.Port)
		break // 使用第一个可用的服务实例
	}

	if serviceAddress == "" {
		return nil, fmt.Errorf("no available stuff service found")
	}

	// 创建 gRPC 连接
	cc, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection: %v", err)
	}

	return stuffpb.NewStuffClient(cc), nil
}
