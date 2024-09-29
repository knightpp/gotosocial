package util_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"codeberg.org/superseriousbusiness/oauth2/v4"
	"github.com/gin-gonic/gin"
	"github.com/superseriousbusiness/gotosocial/internal/api/util"
	"github.com/superseriousbusiness/gotosocial/internal/middleware"
)

func TestGH1361(t *testing.T) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest("GET", "/test", nil).WithContext(context.Background())
	c.Request.Header.Set("Authorization", "Bearer Test123")
	want := errors.New("test error")
	mw := middleware.TokenCheck(nil, func(r *http.Request) (oauth2.TokenInfo, error) {
		return nil, want
	})

	mw(c)

	auth, err := util.TokenAuth(c, true, false, false, false)
	if err == nil {
		t.Fatal("err should have occurred")
		return
	}
	if auth != nil {
		t.Fatal("auth should be nil", auth)
		return
	}
	if !errors.Is(err, want) {
		t.Fatal("unexpected error message:", err)
		return
	}
}

