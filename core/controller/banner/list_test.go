// Copyright 2019 Axetroy. All rights reserved. MIT license.
package banner_test

import (
	"encoding/json"
	"github.com/axetroy/go-server/core/controller"
	"github.com/axetroy/go-server/core/controller/banner"
	"github.com/axetroy/go-server/core/model"
	"github.com/axetroy/go-server/core/schema"
	"github.com/axetroy/go-server/core/service/token"
	"github.com/axetroy/go-server/tester"
	"github.com/axetroy/mocker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetList(t *testing.T) {
	adminInfo, _ := tester.LoginAdmin()

	context := controller.Context{
		Uid: adminInfo.Id,
	}

	{
		var (
			image    = "test"
			href     = "test"
			platform = model.BannerPlatformApp
		)

		r := banner.Create(controller.Context{
			Uid: adminInfo.Id,
		}, banner.CreateParams{
			Image:    image,
			Href:     href,
			Platform: platform,
		})

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)

		n := schema.Banner{}

		assert.Nil(t, tester.Decode(r.Data, &n))

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)

		defer banner.DeleteBannerById(n.Id)
	}

	// 获取列表
	{
		r := banner.GetBannerList(context, banner.Query{})

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)

		banners := make([]schema.Banner, 0)

		assert.Nil(t, tester.Decode(r.Data, &banners))

		assert.Equal(t, schema.DefaultLimit, r.Meta.Limit)
		assert.Equal(t, schema.DefaultPage, r.Meta.Page)
		assert.IsType(t, 1, r.Meta.Num)
		assert.IsType(t, int64(1), r.Meta.Total)

		assert.True(t, len(banners) >= 1)

		for _, b := range banners {
			assert.IsType(t, "string", b.Image)
			assert.IsType(t, "string", b.Href)
			assert.IsType(t, model.BannerPlatformApp, b.Platform)
			assert.IsType(t, "string", b.CreatedAt)
			assert.IsType(t, "string", b.UpdatedAt)
		}
	}
}

func TestGetListRouter(t *testing.T) {
	adminInfo, _ := tester.LoginAdmin()

	header := mocker.Header{
		"Authorization": token.Prefix + " " + adminInfo.Token,
	}

	{
		var (
			image    = "test"
			href     = "test"
			platform = model.BannerPlatformApp
		)

		r := banner.Create(controller.Context{
			Uid: adminInfo.Id,
		}, banner.CreateParams{
			Image:    image,
			Href:     href,
			Platform: platform,
		})

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)

		n := schema.Banner{}

		assert.Nil(t, tester.Decode(r.Data, &n))

		assert.Equal(t, schema.StatusSuccess, r.Status)
		assert.Equal(t, "", r.Message)

		defer banner.DeleteBannerById(n.Id)
	}

	{
		r := tester.HttpAdmin.Get("/v1/banner", nil, &header)

		res := schema.Response{}

		if !assert.Nil(t, json.Unmarshal(r.Body.Bytes(), &res)) {
			return
		}

		if !assert.Equal(t, schema.StatusSuccess, res.Status) {
			return
		}

		if !assert.Equal(t, "", res.Message) {
			return
		}

		banners := make([]schema.Banner, 0)

		assert.Nil(t, tester.Decode(res.Data, &banners))

		for _, b := range banners {
			assert.IsType(t, "string", b.Image)
			assert.IsType(t, "string", b.Href)
			assert.IsType(t, model.BannerPlatformApp, b.Platform)
			assert.IsType(t, "string", b.CreatedAt)
			assert.IsType(t, "string", b.UpdatedAt)
		}
	}
}
