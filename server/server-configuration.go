package server

import (
	"encoding/json"
	"fmt"
	"glow-service/models"
	"io/ioutil"
	"sync"

	firebaseConfig "glow-service/server/firebase-config"

	"firebase.google.com/go/auth"
	"firebase.google.com/go/storage"
)

const (
	configFilePath         = "./server/server-config.json"
	errorMessagesFilePath  = "./server/server-error-messages.json"
	databaseConfigFilePath = "./server/database-config.json"
)

type (
	Configuration struct {
		Port                  int    `json:"port" default:"8080"`
		Environment           string `json:"environment" default:"dev"`
		DatabaseHandler       IDatabaseHandler
		database              *databaseConfiguration
		ServerErrorMessages   *models.ServerErrorMessages
		FirebaseAuthClient    *auth.Client
		FirebaseStorageClient *storage.Client
		initialized           bool
	}

	databaseConfiguration struct {
		DatabaseProvider string `json:"databaseProvider,omitempty"`
		DatabaseUser     string `json:"databaseUser,omitempty"`
		DatabasePassword string `json:"databasePassword,omitempty"`
		DatabaseAddress  string `json:"databaseAddress,omitempty"`
		DatabasePort     string `json:"databasePort,omitempty"`
		Database         string `json:"database,omitempty"`
	}
)

var config *Configuration

func ConfigurationInstance() (*Configuration, error) {

	if config == nil || !config.initialized {
		var serverErr, messagesError, databaseError error
		c := &Configuration{}
		wg := &sync.WaitGroup{}

		// Read all config files on async mode
		wg.Add(3)
		go c.readServerConfig(wg, &serverErr)
		go c.readErrorMessagesConfig(wg, &messagesError)
		go c.readDatabaseConfig(wg, &databaseError)
		wg.Wait()

		if serverErr != nil {
			return nil, serverErr
		}
		if messagesError != nil {
			return nil, messagesError
		}
		if databaseError != nil {
			return nil, databaseError
		}

		//read firebase config file and start firebase application
		firebase, firebaseError := firebaseConfig.NewFirebase()
		if firebaseError != nil {
			return nil, firebaseError
		}

		//get the firebase authentication client
		firebaseAuth, firebaseAuthError := firebase.GetAuth()
		if firebaseAuthError != nil {
			return nil, firebaseAuthError
		}

		//get the firebase storage client
		firebaseStorage, firebaseStorageError := firebase.GetStorage()
		if firebaseStorageError != nil {
			return nil, firebaseStorageError
		}

		c.FirebaseAuthClient = firebaseAuth
		c.FirebaseStorageClient = firebaseStorage

		c.DatabaseHandler = SetubDatabase(
			&c.database.DatabaseProvider,
			&c.database.DatabaseUser,
			&c.database.DatabasePassword,
			&c.database.DatabaseAddress,
			&c.database.DatabasePort,
			&c.database.Database,
		)

		c.initialized = true
		config = c
		fmt.Println("\nServer configuration success loaded.")
	}

	return config, nil
}

func (c *Configuration) readServerConfig(wg *sync.WaitGroup, err *error) {

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

func (c *Configuration) readErrorMessagesConfig(wg *sync.WaitGroup, err *error) {

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

func (c *Configuration) readDatabaseConfig(wg *sync.WaitGroup, err *error) {

	file, fileErr := ioutil.ReadFile(databaseConfigFilePath)
	if fileErr != nil {
		err = &fileErr
	}
	readError := json.Unmarshal([]byte(file), &c.database)
	if readError != nil {
		err = &readError
	}
	wg.Done()
}
