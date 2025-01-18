package test

import (
	"context"
	"testing"
	"time"
)

func TestEmployee(t *testing.T) {
	t.Log("Testing Employee endpoints...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	container, pool, app, err := SetupTestContainer(ctx)
	defer TeardownTestContainer(ctx, container, pool, app)
	if err != nil {
		t.Fatal(err)
	}

	//TODO: test cases
}
