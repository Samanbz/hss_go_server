package test

import (
	"bytes"
	"hss/internal/models"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func TestPost(app *fiber.App, url string, body models.Serializable) (int, []byte, error) {
	req := httptest.NewRequest(http.MethodPost, "/company", bytes.NewBuffer(body.ToJSON()))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		return 0, nil, err
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, resBody, nil
}

func TestGet(app *fiber.App, url string, params *map[string]string) (int, []byte, error) {
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")

	if params == nil {
		params = &map[string]string{}
	}

	q := req.URL.Query()
	for k, v := range *params {
		q.Add(k, v)
	}

	resp, err := app.Test(req)
	if err != nil {
		return 0, nil, err
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, resBody, nil
}
