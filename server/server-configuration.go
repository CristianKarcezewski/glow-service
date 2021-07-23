package server

import (
	"encoding/json"
	"fmt"
	"glow-service/models"
	"io/ioutil"
	"sync"
)

const (
	configFilePath        = "./server/config.json"
	errorMessagesFilePath = "./server/server-error-messages.json"
)

type (
	configuration struct {
		Port                int    `json:"port" short:"p" default:"8080" desc:"Server port"`
		Environment         string `json:"environment" short:"env" default:"dev"`
		ServerErrorMessages models.ServerErrorMessages
		initialized         bool
	}
)

var config *configuration

func ConfigurationInstance() (*configuration, error) {

	if config == nil || !config.initialized {
		var err error
		c := &configuration{}
		wg := &sync.WaitGroup{}
		wg.Add(2)
		go c.readConfigEnv(wg, &err)
		go c.readErrorMessagesConfig(wg, &err)
		wg.Wait()
		if err != nil {
			return nil, err
		}
		c.initialized = true
		config = c
		fmt.Println("\nServer configuration success loaded.")
	}

	return config, nil
}

func (c *configuration) readConfigEnv(wg *sync.WaitGroup, err *error) {

	file, fileErr := ioutil.ReadFile(configFilePath)
	if fileErr != nil {
		err = &fileErr
	}
	readError := json.Unmarshal([]byte(file), c)
	if readError != nil {
		err = &readError
	}
	wg.Done()
}

func (c *configuration) readErrorMessagesConfig(wg *sync.WaitGroup, err *error) {

	file, fileErr := ioutil.ReadFile(errorMessagesFilePath)
	if fileErr != nil {
		err = &fileErr
	}
	readError := json.Unmarshal([]byte(file), &c.ServerErrorMessages)
	if readError != nil {
		err = &readError
	}
	wg.Done()
}
