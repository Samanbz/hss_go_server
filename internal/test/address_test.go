package test

import (
	"hss/internal/models"
	"hss/internal/test/helpers"
	"hss/internal/test/mocks"
	"hss/internal/test/mocks/mock_data"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
)

func TestAddress(t *testing.T) {
	t.Log("Testing Address endpoints...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	container, pool, app, err := SetupTestContainer(ctx)
	defer TeardownTestContainer(ctx, container, pool, app)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("TestInsertAddress", func(t *testing.T) {
		t.Log("TestInsertAddress")

		err := mocks.NewMocks(
			ctx, pool,
			mocks.NewCompanyMockGroup(mock_data.MockCompany),
		)
		if err != nil {
			t.Fatal(err)
		}
		mock_data.MockAddress.CompanyID = mock_data.MockCompany.ID

		testInsertAddress(t, app)
	})
}

func testInsertAddress(t *testing.T, app *fiber.App) {
	inputAddress := mock_data.MockAddress
	statusCode, body, err := helpers.TestPost(app, "/address", inputAddress)
	if err != nil {
		t.Fatal(err)
	}

	if statusCode != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, statusCode)
		t.Errorf("response body: %s", body)
		t.FailNow()
	}

	var outputAddress models.Address
	err = outputAddress.FromJSON(body)
	if err != nil {
		t.Fatal(err)
	}

	err = helpers.CheckStruct(&outputAddress, inputAddress, false)
	if err != nil {
		t.Fatal(err)
	}
}
