package test

import (
	"fmt"
	"hss/internal/models"
	"hss/internal/test/helpers"
	"hss/internal/test/mocks"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
)

func TestCompany(t *testing.T) {
	t.Log("Testing Company endpoints...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	container, pool, app, err := SetupTestContainer(ctx)
	defer TeardownTestContainer(ctx, container, pool, app)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("TestInsertCompany", func(t *testing.T) {
		t.Log("TestInsertCompany")
		testInsertCompany(t, app)
	})

	t.Run("TestGetCompany", func(t *testing.T) {
		t.Log("TestGetCompany")

		err := mocks.NewMocks(
			ctx, pool,
			mocks.NewCompanyMockGroup(&mocks.MockCompany2),
		)
		if err != nil {
			t.Fatal(err)
		}

		testGetCompany(t, app, mocks.MockCompany2)
	})
}

func testInsertCompany(t *testing.T, app *fiber.App) {
	inputCompany := mocks.MockCompany
	statusCode, body, err := helpers.TestPost(app, "/company", inputCompany)
	if err != nil {
		t.Fatal(err)
	}

	if statusCode != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, statusCode)
		t.Errorf("response body: %s", body)
		t.FailNow()
	}

	var outputCompany models.Company
	err = outputCompany.FromJSON(body)
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.CheckStruct(&outputCompany, &inputCompany, false)
	if err != nil {
		t.Fatal(err)
	}
}

func testGetCompany(t *testing.T, app *fiber.App, company models.Company) {
	statusCode, body, err := helpers.TestGet(app, fmt.Sprintf("/company/%d", company.ID), nil)
	if err != nil {
		t.Fatal(err)
	}

	if statusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, statusCode)
		t.Errorf("response body: %s", body)
		t.FailNow()
	}

	var resCompany models.Company
	err = resCompany.FromJSON(body)
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.CheckStruct(&resCompany, &company, true)
	if err != nil {
		t.Fatal(err)
	}
}
