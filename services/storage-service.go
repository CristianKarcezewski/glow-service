package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"glow-service/models"
	"io"

	"firebase.google.com/go/storage"
	"github.com/google/uuid"
)

const (
	firebaseStorageBucketName = "glow-files"
)

type (
	IStorageService interface {
		SaveProfileImage(user *models.User, image string) (string, error)
	}
	storageService struct {
		FirebaseStorageClient *storage.Client
		UsersService          IUsersService
	}
)

func NewStorageService(firebaseStorageClient *storage.Client, userService IUsersService) IStorageService {
	return &storageService{firebaseStorageClient, userService}
}

func (ss *storageService) SaveProfileImage(user *models.User, image string) (string, error) {

	if user.Uid == "" {
		return "", errors.New("user not found")
	}

	uuid := uuid.New()

	firebaseStorageBucket, firebaseStorageBucketerror := ss.FirebaseStorageClient.Bucket(firebaseStorageBucketName)
	if firebaseStorageBucketerror != nil {
		return "", firebaseStorageBucketerror
	}

	object := firebaseStorageBucket.Object(fmt.Sprintf("%s/%s", user.Uid, uuid))
	writer := object.NewWriter(context.Background())

	//Set the attribute
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": uuid.String()}
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader([]byte(image))); err != nil {
		return "", err
	}

	// if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
	// 	return "", err
	// }

	return "", nil
}

// func (ss *storageService) RemoveFile() (*string, error) {
// 	start firebase storage bucket
// 	firebaseStorageBucket, firebaseStorageBucketerror := ss.FirebaseStorageClient.Bucket(firebaseStorageBucketName)
// 	if firebaseStorageBucketerror != nil {
// 		return nil, firebaseStorageBucketerror
// 	}

// 	firebaseStorageBucket.ACL()
// }

func (ss *storageService) toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
