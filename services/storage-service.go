package services

import (
	"firebase.google.com/go/storage"
)

const (
	firebaseStorageBucketName = "glow-files"
)

type (
	IStorageService interface{}
	storageService  struct {
		FirebaseStorageClient *storage.Client
	}
)

func NewStorageService(firebaseStorageClient *storage.Client) IStorageService {
	return &storageService{firebaseStorageClient}
}

// func (ss *storageService) SaveFile() (*string, error) {
// 	//start firebase storage bucket
// 	firebaseStorageBucket, firebaseStorageBucketerror := ss.FirebaseStorageClient.Bucket(firebaseStorageBucketName)
// 	if firebaseStorageBucketerror != nil {
// 		return nil, firebaseStorageBucketerror
// 	}

// 	firebaseStorageBucket.Object()
// }

// func (ss *storageService) RemoveFile() (*string, error) {
// 	//start firebase storage bucket
// 	firebaseStorageBucket, firebaseStorageBucketerror := ss.FirebaseStorageClient.Bucket(firebaseStorageBucketName)
// 	if firebaseStorageBucketerror != nil {
// 		return nil, firebaseStorageBucketerror
// 	}

// 	firebaseStorageBucket.ACL()
// }
