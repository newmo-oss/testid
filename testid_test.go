package testid_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/newmo-oss/testid"
)

func Test(t *testing.T) {
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
