package testid_test

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/newmo-oss/gotestingmock"

	"github.com/newmo-oss/testid"
)

func Test(t *testing.T) {
	t.Parallel()

	t.Run("empty test id", func(t *testing.T) {
		want := false
		_, got := testid.FromContext(context.Background())
		if got != want {
			t.Errorf("FromContext(context.Background()) must return %v but got %v", want, got)
		}
	})

	t.Run("get test id from associated context", func(t *testing.T) {
		want := uuid.NewString()
		ctx := testid.WithValue(context.Background(), want)
		got, ok := testid.FromContext(ctx)
		if !ok {
			t.Errorf("FromContext must return true as the second return value when ctx is associated with test id")
		}

		if got != want {
			t.Errorf("got, want = %v, %v", got, want)
		}
	})

	t.Run("duplicated check on (default in test)", func(t *testing.T) {
		r := gotestingmock.Run(func(tb *gotestingmock.TB) {
			ctx := testid.WithValue(context.Background(), uuid.NewString())
			ctx = testid.WithValue(ctx, uuid.NewString())
		})

		if r.PanicValue == nil {
			t.Error("expected panic did not occur")
		}
	})

	t.Run("duplicated check off", func(t *testing.T) {
		testid.SetCheckDuplicated(false)
		r := gotestingmock.Run(func(tb *gotestingmock.TB) {
			ctx := testid.WithValue(context.Background(), uuid.NewString())
			ctx = testid.WithValue(ctx, uuid.NewString())
		})

		if r.PanicValue != nil {
			t.Error("unexpected panic:", r.PanicValue)
		}
	})
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
