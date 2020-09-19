package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetOne(c)

	fmt.Println(w)

	assert.Equal(t, 200, w.Code)
}
