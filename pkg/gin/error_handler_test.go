package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestErrorHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "givenError_whenResponding_returnResponseStatusError",
			args: args{
				c: testContext,
			},
			err: errors.New("basic error"),
		},
		{
			name: "givenResponseStatusError_whenResponding_returnResponseStatusError",
			args: args{
				c: testContext,
			},
			err: HttpError{
				Timestamp: time.Now(),
				Status:    http.StatusBadRequest,
				Meta:      nil,
				Path:      "/",
				Message:   "response status error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			// when
			_ = tt.args.c.Error(tt.err)
			HttpErrorHandler()(tt.args.c)
			got := tt.args.c.Errors.Last()
			// then
			var want HttpError
			if !errors.As(got, &want) {
				t.Errorf("Error() = %v, err %v", got, want)
			}
		})
	}
}

func TestResponseStatusError_Error(t *testing.T) {
	type fields struct {
		Timestamp time.Time
		Status    HttpStatus
		Meta      any
		Path      string
		Message   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "givenResponseStatusError_whenError()_returnErrorMessage",
			fields: fields{
				Timestamp: time.Now(),
				Status:    http.StatusBadRequest,
				Meta:      nil,
				Path:      "/",
				Message:   "I'm test",
			},
			want: "I'm test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := HttpError{
				Timestamp: tt.fields.Timestamp,
				Status:    tt.fields.Status,
				Meta:      tt.fields.Meta,
				Path:      tt.fields.Path,
				Message:   tt.fields.Message,
			}
			if got := r.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
