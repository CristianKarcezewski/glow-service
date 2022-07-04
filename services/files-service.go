package services

import (
	"errors"
	"glow-service/models"
	"glow-service/repository"
)

type (
	IFilesService interface {
		SaveProfileImage(log *models.StackLog, user *models.User, image string) (string, error)
		SaveCompanyFile(log *models.StackLog, user *models.User, companyId int64, fileUrl string) (*models.File, error)
		RemoveCompanyFile(log *models.StackLog, user *models.User, companyId int64, fileId int64) error
		FetchCompanyFiles(log *models.StackLog, companyId int64) (*[]models.File, error)
	}
	filesService struct {
		usersService     IUsersService
		companiesService ICompaniesService
		filesRepository  repository.IFilesRepository
	}
)

func NewFilesService(userService IUsersService, companiesService ICompaniesService, filesRepository repository.IFilesRepository) IFilesService {
	return &filesService{userService, companiesService, filesRepository}
}

func (ss *filesService) SaveProfileImage(log *models.StackLog, user *models.User, image string) (string, error) {
	log.AddStep("StorageService-SaveProfileImage")

	if user.Uid == "" {
		return "", errors.New("user not found")
	}

	user.ImageUrl = image
	firebaseError := ss.usersService.SetProfileImage(log, user, image)
	if firebaseError != nil {
		return "", firebaseError
	}
	_, databaseError := ss.usersService.Update(log, user)
	if databaseError != nil {
		return "", databaseError
	}

	return image, nil
}

func (fr *filesService) SaveCompanyFile(log *models.StackLog, user *models.User, companyId int64, fileUrl string) (*models.File, error) {
	log.AddStep("FilesService-SaveCompanyFile")

	company, companyError := fr.companiesService.GetById(log, companyId)
	if companyError != nil {
		return nil, companyError
	}

	if company.UserId != user.UserId {
		return nil, errors.New("method not authorized for this user")
	}

	file, fileError := fr.filesRepository.SaveCompanyFile(log, companyId, fileUrl)
	if fileError != nil {
		return nil, fileError
	}

	return file.ToModel(), nil
}

func (fr *filesService) RemoveCompanyFile(log *models.StackLog, user *models.User, companyId int64, fileId int64) error {
	log.AddStep("FilesService-RemoveCompanyFile")

	company, companyError := fr.companiesService.GetById(log, companyId)
	if companyError != nil {
		return companyError
	}

	if company.UserId != user.UserId {
		return errors.New("method not authorized for this user")
	}

	return fr.filesRepository.RemoveCompanyFile(log, fileId)
}

func (fr *filesService) FetchCompanyFiles(log *models.StackLog, companyId int64) (*[]models.File, error) {
	log.AddStep("FilesService-FetchCompanyFiles")

	_, companyError := fr.companiesService.GetById(log, companyId)
	if companyError != nil {
		return nil, companyError
	}

	filesDao, filesError := fr.filesRepository.FetchCompanyFiles(log, companyId)
	if filesError != nil {
		return nil, filesError
	}

	var files []models.File
	for i := range *filesDao {
		files = append(files, *(*filesDao)[i].ToModel())
	}

	return &files, nil
}
