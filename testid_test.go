package testid_test

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/newmo-oss/testid"
)

func Test(t *testing.T) {
	t.Parallel()

	{
		want := false
		_, got := testid.FromContext(context.Background())
		if got != want {
			t.Errorf("FromContext(context.Background()) must return %v but got %v", want, got)
		}
	}

	{
		want := uuid.NewString()
		ctx := testid.WithValue(context.Background(), want)
		got, ok := testid.FromContext(ctx)
		if !ok {
			t.Errorf("FromContext must return true as the second return value when ctx is associated with test id")
		}

		if got != want {
			t.Errorf("got, want = %v, %v", got, want)
		}
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	{
		got := testid.New(t)
		t.Log("testid =", got)
		if !strings.HasPrefix(got, "TestNew_") {
			t.Fatal(`testid.New must return string which has t.Name() + "_" prefix:`, got)
		}

		_, err := uuid.Parse(strings.TrimPrefix(got, "TestNew_"))
		if err != nil {
			t.Error("unexpected error:", err)
		}
	}

	t.Run("Sub", func(t *testing.T) {
		got := testid.New(t)
		t.Log("testid =", got)
		if !strings.HasPrefix(got, "TestNew/Sub_") {
			t.Fatal(`testid.New must return string which has t.Name() + "_" prefix:`, got)
		}

		_, err := uuid.Parse(strings.TrimPrefix(got, "TestNew/Sub_"))
		if err != nil {
			t.Error("unexpected error:", err)
		}
	})
}
