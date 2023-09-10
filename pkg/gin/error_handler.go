package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type HttpStatus int

var _ error = &HttpError{}

type HttpError struct {
	Timestamp time.Time  `json:"timestamp"`
	Status    HttpStatus `json:"status"`
	Meta      any        `json:"meta"`
	Path      string     `json:"path"`
	Message   string     `json:"message"`
}

func (r HttpError) Error() string {
	return r.Message
}

func HttpErrorHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			log.Default().Printf("Error occurred: %v\n", err)
		}
		if len(c.Errors) > 0 {
			target := HttpError{}
			if errors.As(c.Errors[0], &target) {
				c.JSON(int(target.Status), target)
			} else {
				statusError := HttpError{
					Timestamp: time.Now(),
					Status:    http.StatusInternalServerError,
					Meta:      c.Errors,
					Path:      c.FullPath(),
					Message:   c.Errors[0].Error(),
				}
				wrappedErr := c.Error(statusError)
				c.JSON(http.StatusInternalServerError, &wrappedErr)
			}
		}
	}
}
