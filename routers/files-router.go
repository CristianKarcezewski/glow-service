package routers

import "github.com/labstack/echo"

const (
	uploadProfileImagePath = "/files/profile-image"
	uploadFilePath         = "/files/:companyId"
	fetchCompanyFilesPath  = "/files/:companyId"
	removeFilePath         = "/files/:companyId/:fileId"
)

type (
	FilesRouter struct {
		Echo               *echo.Echo
		UploadProfileImage echo.HandlerFunc
		UploadFile         echo.HandlerFunc
		FetchCompanyFiles  echo.HandlerFunc
		RemoveFile         echo.HandlerFunc
	}
)

func (fr *FilesRouter) Wire() {
	fr.Echo.POST(uploadProfileImagePath, fr.UploadProfileImage)
	fr.Echo.POST(uploadFilePath, fr.UploadFile)
	fr.Echo.GET(fetchCompanyFilesPath, fr.FetchCompanyFiles)
	fr.Echo.DELETE(removeFilePath, fr.RemoveFile)
}
