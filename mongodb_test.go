package mongodb_test

import (
	"context"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-mongodb"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUse(t *testing.T) {
	w := waitfor.New(mongodb.Use())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := w.Test(ctx, []string{"mongodb://usr:pass@localhost/my-db"})

	assert.Error(t, err)
}
