package services

import (
	"errors"
	"glow-service/models"
)

const (
	firebaseStorageBucketName = "glow-files"
)

type (
	IStorageService interface {
		SaveProfileImage(log *models.StackLog, user *models.User, image string) (string, error)
	}
	storageService struct {
		UsersService IUsersService
	}
)

func NewStorageService(userService IUsersService) IStorageService {
	return &storageService{userService}
}

func (ss *storageService) SaveProfileImage(log *models.StackLog, user *models.User, image string) (string, error) {
	log.AddStep("StorageService-SaveProfileImage")

	if user.Uid == "" {
		return "", errors.New("user not found")
	}

	user.ImageUrl = image
	firebaseError := ss.UsersService.SetProfileImage(log, user, image)
	if firebaseError != nil {
		return "", firebaseError
	}
	_, databaseError := ss.UsersService.Update(log, user)
	if databaseError != nil {
		return "", databaseError
	}

	return image, nil
}

// func (ss *storageService) RemoveFile() (*string, error) {
// 	start firebase storage bucket
// 	firebaseStorageBucket, firebaseStorageBucketerror := ss.FirebaseStorageClient.Bucket(firebaseStorageBucketName)
// 	if firebaseStorageBucketerror != nil {
// 		return nil, firebaseStorageBucketerror
// 	}

// 	firebaseStorageBucket.ACL()
// }

// func (ss *storageService) toBase64(b []byte) string {
// 	return base64.StdEncoding.EncodeToString(b)
// }
