package csl_configuration_client

import (
	"encoding/json"
	"errors"
	"fmt"
	discoveryClient "github.com/carrot-systems/csl-discovery-client"
	"net/http"
	"os"
)

type ConfigurationClient struct {
	discovery *discoveryClient.DiscoveryClient
}

func (client *ConfigurationClient) LoadConfiguration() error {
	service, err := client.discovery.GetService("config")

	if err != nil {
		return err
	}

	configurationClientUrl := service[0].ExternalUrl
	url := fmt.Sprintf("http://%s/config/%s", configurationClientUrl, client.discovery.RegisteredName)

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)

	var configurationResponse ConfigResponse
	err = decoder.Decode(&configurationResponse)

	if err != nil {
		return err
	}

	if !configurationResponse.Status.Success {
		return errors.New(configurationResponse.Status.Message)
	}

	for _, entry := range configurationResponse.Values {
		err := os.Setenv(entry.Key, entry.Value)
		if err != nil {
			return err
		}
	}

	return nil
	/*
		url := fmt.Sprintf("http://%s/services", serverConf.ServerUrl)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(dt))

		if err != nil {
			return err
		}

		decoder := json.NewDecoder(resp.Body)

		var registerResponse RegisterResponse
		err = decoder.Decode(&registerResponse)

		if err != nil {
			return err
		}

		client.registered = registerResponse.Status.Success
		println(registerResponse.Status.Message)
		if !registerResponse.Status.Success {
			return errors.New(registerResponse.Status.Message)
		}
		return nil*/
}

func NewClient(discovery *discoveryClient.DiscoveryClient) *ConfigurationClient {
	return &ConfigurationClient{
		discovery: discovery,
	}
}
