package firebaseConfig

import (
	"context"
	"errors"
	"path/filepath"

	firebase "firebase.google.com/go"
	authentication "firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

const (
	firebasecConfigFile = "server/firebase-config/firebase-config.json"
)

var auth *authentication.Client

func GetFirebase() (*authentication.Client, error) {

	if auth == nil {

		// is necessary mount the absolute path to firebase-config.json file
		path, pathErr := filepath.Abs(firebasecConfigFile)
		if pathErr != nil {
			return nil, errors.New("error mounting path to config file")
		}

		app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(path))
		if err != nil {
			return nil, errors.New("error starting firebase app")
		}

		aut, authError := app.Auth(context.Background())
		if authError != nil {
			return nil, errors.New("error starting firebase client")
		}

		auth = aut
	}

	return auth, nil
}
