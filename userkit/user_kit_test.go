package userkit

import (
	"context"
	"github.com/r3labs/diff"
	"github.com/stretchr/testify/require"
	"github.com/zonewave/pkgs/util"
	"math/rand"
	"testing"
)

func fakeUserKit(opts ...func(item *UserKit)) *UserKit {
	item := &UserKit{
		UserID:     rand.Uint64(),
		Locale:     util.RandString(10),
		AppVersion: util.RandString(10),
		RegionCode: util.RandString(10),
		TimeZone:   util.RandString(10),
		TimeOffset: rand.Int(),
		UserAgent:  util.RandString(10),
	}
	for _, opt := range opts {
		opt(item)
	}
	return item
}

func TestUserkitMethod(t *testing.T) {
	ctx := context.Background()
	kit := fakeUserKit()

	ctx = ToCtx(ctx, kit)
	kit2 := FromCtx(ctx)
	change, err := diff.Diff(kit, kit2)
	require.NoError(t, err)
	require.Empty(t, change)

	jsonResp := kit.String()
	require.NotEmpty(t, jsonResp)

}
