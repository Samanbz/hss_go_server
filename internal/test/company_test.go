package test

import (
	"fmt"
	"hss/internal/models"
	"hss/internal/test/mocks"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
)

func TestCompany(t *testing.T) {
	t.Log("TestCompany")

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

		err := NewMocks(
			ctx, pool,
			NewCompanyMockGroup(&mocks.CompanyInput2),
		)
		if err != nil {
			t.Fatal(err)
		}

		testGetCompany(t, app, mocks.CompanyInput2)
	})
}

func testInsertCompany(t *testing.T, app *fiber.App) {
	inputCompany := mocks.CompanyInput
	statusCode, body, err := TestPost(app, "/company", inputCompany)
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

	if outputCompany.Username != inputCompany.Username {
		t.Errorf("expected username %s, got %s", inputCompany.Username, outputCompany.Username)
	}

	if outputCompany.CompanyName != inputCompany.CompanyName {
		t.Errorf("expected company name %s, got %s", inputCompany.CompanyName, outputCompany.CompanyName)
	}

	if outputCompany.RepFirstname != inputCompany.RepFirstname {
		t.Errorf("expected rep firstname %s, got %s", inputCompany.RepFirstname, outputCompany.RepFirstname)
	}
}

func testGetCompany(t *testing.T, app *fiber.App, company models.Company) {
	statusCode, body, err := TestGet(app, fmt.Sprintf("/company/%d", company.ID), nil)
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

	if resCompany.ID != company.ID {
		t.Errorf("expected company ID %d, got %d", company.ID, resCompany.ID)
	}

	if resCompany.Username != company.Username {
		t.Errorf("expected username %s, got %s", company.Username, resCompany.Username)
	}

	if resCompany.CompanyName != company.CompanyName {
		t.Errorf("expected company name %s, got %s", company.CompanyName, resCompany.CompanyName)
	}

	if resCompany.RepFirstname != company.RepFirstname {
		t.Errorf("expected rep firstname %s, got %s", company.RepFirstname, resCompany.RepFirstname)
	}

}
