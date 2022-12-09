package eureka

import "testing"

func TestClient_Register(t *testing.T) {
	eurekaConfig := GetDefaultEurekaClientConfig()
	eurekaConfig.UseDnsForFetchingServiceUrls = false
	eurekaConfig.ServiceUrl = map[string]string{
		DEFAULT_ZONE: "http://hddata:hddata$2019@192.168.1.83:8002/eureka,http://hddata:hddata$2019@192.168.1.59:8002/eureka",
	}
	eurekaClient := new(Client).Config(eurekaConfig).Register("gw-test-eureka", 80)
	eurekaClient.Run()
	select {}
}
