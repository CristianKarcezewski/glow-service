package routers

import "github.com/labstack/echo"

const (
	uploadProfileImagePath = "/files/profile-image"
)

type (
	FilesRouter struct {
		Echo               *echo.Echo
		UploadProfileImage echo.HandlerFunc
		// UploadFile         echo.HandlerFunc
	}
)

func (fr *FilesRouter) Wire() {
	fr.Echo.POST(uploadProfileImagePath, fr.UploadProfileImage)
}
