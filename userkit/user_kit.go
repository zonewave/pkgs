package userkit

import (
	"context"
	"encoding/json"
	"github.com/ybzhanghx/pkgs/ctxfield"
)

//go:generate msgp
type UserKit struct {
	UserID     uint64 `json:"userID,omitempty" msg:"userID"`
	Locale     string `json:"locale,omitempty" msg:"locale"`
	AppVersion string `json:"appVersion,omitempty" msg:"appVersion"`
	RegionCode string `json:"regionCode,omitempty" msg:"regionCode"`
	TimeZone   string `json:"timezone,omitempty" msg:"timezone"`
	TimeOffset int    `json:"timeOffset,omitempty" msg:"timeOffset"`
	UserAgent  string `json:"userAgent,omitempty" msg:"userAgent"`
}

func (z *UserKit) String() string {
	bs, _ := json.Marshal(z)
	return string(bs)
}

func ToCtx(ctx context.Context, kit *UserKit) context.Context {
	return context.WithValue(ctx, ctxfield.UserKitKey, kit)
}

func FromCtx(ctx context.Context) *UserKit {
	v, ok := ctx.Value(ctxfield.UserKitKey).(*UserKit)
	if !ok {
		return nil
	}
	return v
}