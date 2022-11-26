package adapter_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-wonk/si/sihttp"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/product/adapter"
)

func Test_groupHttp_ReadGroup(t *testing.T) {
	if !onlinetest {
		t.Skip("skipping online tests")
	}

	ctx := context.Background()

	groupHttp := adapter.NewGroupHttp(sihttp.DefaultInsecureClient(),
		"http://localhost:49002", "ab2316584873095f017f6dfa7a9415794f563fcc473eb3fe65b9167e37fd5a4b")

	res, err := groupHttp.ReadGroup(ctx, group1.ID)
	assert.Nil(t, err)

	fmt.Println(res.String())
}
