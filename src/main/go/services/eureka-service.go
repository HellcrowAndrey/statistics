package services

import (
	"../config"
	"github.com/hudl/fargo"
	"log"
	"time"
)

type EurekaService struct {
	config *config.Config
}

func NewEurekaService(config *config.Config) *EurekaService {
	return &EurekaService{config: config}
}

func (service *EurekaService) Run() {
	instance := service.createInstance()
	connection := fargo.NewConn("http://localhost:3000/eureka")
	err := connection.DeregisterInstance(&instance)
	if err != nil {
		panic("Can not connect to eureka.")
	}
	err = connection.RegisterInstance(&instance)
	if err != nil {
		panic("Can not connect to eureka.")
	} else {
		HeartBeat(connection, &instance)
	}
}

func (service *EurekaService) createInstance() fargo.Instance {
	return fargo.Instance{
		InstanceId:     "statistics-golang",
		HostName:       "127.0.0.1",
		Port:           8080,
		SecurePort:     8443,
		App:            "people",
		IPAddr:         "127.0.0.1",
		VipAddress:     "people",
		HomePageUrl:    "http://127.0.0.1:8080/",
		StatusPageUrl:  "http://127.0.0.1:8080/info",
		HealthCheckUrl: "http://127.0.0.1:8080/health",
		DataCenterInfo: fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:         fargo.UP,
		LeaseInfo: fargo.LeaseInfo{
			DurationInSecs: 90,
		},
	}
}

func HeartBeat(ec fargo.EurekaConnection, i *fargo.Instance) {
	ticker := time.Tick(time.Duration(30*1000) * time.Millisecond)
	for {
		select {
		case <-ticker:
			if err := ec.HeartBeatInstance(i); err != nil {
				log.Println("Lost connection to eureka")
			}
		}
	}
}
