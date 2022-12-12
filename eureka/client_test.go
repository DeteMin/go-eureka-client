package eureka

import "testing"

func TestClient_Register(t *testing.T) {
	eurekaConfig := GetDefaultEurekaClientConfig()
	eurekaConfig.UseDnsForFetchingServiceUrls = false
	eurekaConfig.ServiceUrl = map[string]string{
		DEFAULT_ZONE: "http://hddata:hddata$2019@192.168.1.116:8002/eureka,http://hddata:hddata$2019@192.168.1.59:8002/eureka",
	}
	eurekaClient := new(Client).Config(eurekaConfig).Register("gw-test-eureka", 80)
	eurekaClient.Run()

	//eurekaClient2 := new(Client).Config(eurekaConfig).Register("gw-test-eureka", 8013)
	//eurekaClient2.Run()
	select {}
}

func Test_Del(t *testing.T) {
	eurekaConfig := GetDefaultEurekaClientConfig()
	eurekaConfig.UseDnsForFetchingServiceUrls = false
	eurekaConfig.ServiceUrl = map[string]string{
		DEFAULT_ZONE: "http://hddata:hddata$2019@192.168.1.116:8002/eureka,http://hddata:hddata$2019@192.168.1.59:8002/eureka",
	}
	eurekaClient := new(Client).Config(eurekaConfig).Register("GW-VIDEOSY-SERVICE", 8013)

	vo := eurekaClient.GetInstance()
	vo.Hostname = "192.168.1.144"
	vo.IppAddr = "192.168.1.144"
	eurekaClient = eurekaClient.RegisterVo(vo)

	eurekaClient.Run()

	apis, err := eurekaClient.Api()
	if err != nil {
		log.Errorf("Failed to get EurekaServerApi instance, de-register %s failed, err=%s", eurekaClient.instance.InstanceId, err.Error())
		return
	}

	for _, api := range apis {
		err = api.DeRegisterInstance(eurekaClient.instance.App, eurekaClient.instance.InstanceId)
		if err != nil {
			log.Errorf("Failed to de-register %s, err=%s", eurekaClient.instance.InstanceId, err.Error())
			return
		}
	}

	log.Infof("de-register %s success.", eurekaClient.instance.InstanceId)
	select {}
}
