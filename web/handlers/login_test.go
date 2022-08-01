package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sandbox/service"
	"sandbox/web"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	web.RunTest(func(c *service.Container) {
		headers := map[string]string{
			"Authorization": "abc123",
		}

		body := `{
			"username": "something", "password": "somethingelse"
		}`
		w, err := web.MakeRequest(c.Web, http.MethodPost, "/login", bytes.NewReader([]byte(body)), headers)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp struct {
			User    string `json:"user"`
			Session string `json:"session"`
		}
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "something", resp.User)
	})
}
