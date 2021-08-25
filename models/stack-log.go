package models

import (
	"fmt"
	"time"
)

type (
	StackLog struct {
		Platform   string   `json:"platform,omitempty" validate:"required"`
		User       user     `json:"user,omitempty"`
		StackTrace []string `json:"stackTrace,omitempty"`
	}

	user struct {
		Token    string `json:"-"`
		UserId   int64  `json:"serId,omitempty" validate:"required"`
		UserName string `json:"username,omitempty" validate:"required"`
		Email    string `json:"email,omitempty" validate:"required"`
	}
)

func (log *StackLog) SetUser(email string) {
	log.User.Email = email
}

// Add a stackTrace step log into object.
func (log *StackLog) AddStep(stackLog string) {
	log.StackTrace = append(log.StackTrace, fmt.Sprintf("(%s) STEP: %s", log.dateToString(), stackLog))
}

// Add a stackTrace info log into object.
func (log *StackLog) AddInfo(stackLog string) {
	log.StackTrace = append(log.StackTrace, fmt.Sprintf("(%s) INFO: %s", log.dateToString(), stackLog))
}

// Add a stackTrace info log into object.
func (log *StackLog) AddError(stackLog string) *ErrorResponse {
	log.StackTrace = append(log.StackTrace, fmt.Sprintf("(%s) ERROR: %s", log.dateToString(), stackLog))
	return &ErrorResponse{Message: stackLog}
}

// Print all stack trace into console.
func (log *StackLog) PrintStackOnConsole() {
	if log.User.Email != "" {
		fmt.Printf("\n{%s: %s}\n", log.Platform, log.User.Email)
	} else {
		fmt.Printf("\n{%s: %s}\n", log.Platform, "Anonymous user")
	}
	for stepIndex := range log.StackTrace {
		fmt.Printf("%s\n", log.StackTrace[stepIndex])
	}
}

func (log *StackLog) dateToString() string {
	return fmt.Sprint(time.Now().Format("02/01/2006 15:04:05.000"))
	// t := time.Now()
	// return fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
}
