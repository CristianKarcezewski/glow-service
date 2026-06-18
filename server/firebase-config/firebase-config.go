package firebaseConfig

import (
	"context"
	"errors"
	"path/filepath"

	firebase "firebase.google.com/go"
	authentication "firebase.google.com/go/auth"
	storage "firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

const (
	firebasecConfigFile = "server/firebase-config/firebase-config.json"
)

type (
	firebaseInstance struct {
		App *firebase.App
	}
)

func NewFirebase() (*firebaseInstance, error) {
	// is necessary mount the absolute path to firebase-config.json file
	path, pathErr := filepath.Abs(firebasecConfigFile)
	if pathErr != nil {
		return nil, errors.New("error mounting path to config file")
	}

	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(path))
	if err != nil {
		return nil, errors.New("error starting firebase app")
	}

	return &firebaseInstance{app}, nil
}

func (fi *firebaseInstance) GetAuth() (*authentication.Client, error) {

	aut, authError := fi.App.Auth(context.Background())
	if authError != nil {
		return nil, errors.New("error starting firebase client")
	}

	return aut, nil
}

func (fi *firebaseInstance) GetStorage() (*storage.Client, error) {

	st, stError := fi.App.Storage(context.Background())
	if stError != nil {
		return nil, errors.New("error starting firebase storage")
	}

	return st, nil
}
