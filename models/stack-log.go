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
		UserName string `json:"username,omitempty" validate:"required"`
		Email    string `json:"email,omitempty" validate:"required"`
	}
)

func (log *StackLog) SetUser(email string) {
	log.User.Email = email
}

// Add a stackTrace step log into object.
func (log *StackLog) AddStep(stackLog string) {

	if len(log.StackTrace) == 0 {
		fmt.Printf("\n")
	}

	if len(log.User.Email) > 0 {
		log.StackTrace = append(log.StackTrace, fmt.Sprintf("{%s: %s}(%s) STEP: %s", log.Platform, log.User.Email, log.dateToString(), stackLog))
	} else {
		log.StackTrace = append(log.StackTrace, fmt.Sprintf("{%s: %s}(%s) STEP: %s", log.Platform, "Anonymous User", log.dateToString(), stackLog))
	}
	fmt.Println(log.StackTrace[(len(log.StackTrace) - 1)])
}

// Add a stackTrace info log into object.
func (log *StackLog) AddInfo(stackLog string) {
	if len(log.StackTrace) == 0 {
		fmt.Printf("\n")
	}

	if len(log.User.Email) > 0 {
		log.StackTrace = append(log.StackTrace, fmt.Sprintf("{%s: %s}(%s) INFO: %s", log.Platform, log.User.Email, log.dateToString(), stackLog))
	} else {
		log.StackTrace = append(log.StackTrace, fmt.Sprintf("{%s: %s}(%s) INFO: %s", log.Platform, "Anonymous User", log.dateToString(), stackLog))
	}
	fmt.Println(log.StackTrace[(len(log.StackTrace) - 1)])
}

// Add a stackTrace info log into object.
func (log *StackLog) AddError(stackLog string) *ErrorResponse {
	log.StackTrace = append(log.StackTrace, fmt.Sprintf("(%s) ERROR: %s", log.dateToString(), stackLog))
	fmt.Println(log.StackTrace[(len(log.StackTrace) - 1)])
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
}
