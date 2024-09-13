# testid [![Go Reference](https://pkg.go.dev/badge/github.com/newmo-oss/testido.svg)](https://pkg.go.dev/github.com/newmo-oss/testid)[![Go Report Card](https://goreportcard.com/badge/github.com/newmo-oss/testid)](https://goreportcard.com/report/github.com/newmo-oss/testid)

testid provides utilities of test id for unit testing of Go.

## Usage

```go
package a_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/newmo-oss/newmotime"
	"github.com/newmo-oss/newmotime/newmotimetest"
	"github.com/newmo-oss/testid"
)

func Test(t *testing.T) {
	tid := uuid.NewString()
	ctx := testid.WithValue(context.Background(), tid)
	now := newmotime.Now(ctx)
	newmotimetest.SetFixedNow(t, ctx, now)
}
```

## License
MIT
