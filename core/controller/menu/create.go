// Copyright 2019 Axetroy. All rights reserved. MIT license.
package menu

import (
	"errors"
	"github.com/axetroy/go-server/core/controller"
	"github.com/axetroy/go-server/core/exception"
	"github.com/axetroy/go-server/core/helper"
	"github.com/axetroy/go-server/core/model"
	"github.com/axetroy/go-server/core/schema"
	"github.com/axetroy/go-server/core/service/database"
	"github.com/axetroy/go-server/core/validator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"time"
)

type CreateMenuParams struct {
	Name      string    `json:"name" valid:"required~请填写菜单名"` // 菜单名
	Url       *string   `json:"url"`                          // 菜单链接的 URL 地址
	Icon      *string   `json:"icon"`                         // 菜单的图标
	Accession *[]string `json:"accession"`                    // 该菜单所需要的权限
	Sort      int       `json:"sort" `                        // 菜单排序, 越大的越靠前
	ParentId  *string   `json:"parent_id"`                    // 该菜单的父级 ID
}

func Create(c controller.Context, input CreateMenuParams) (res schema.Response) {
	var (
		err  error
		data schema.Menu
		tx   *gorm.DB
	)

	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = exception.Unknown
			}
		}

		if tx != nil {
			if err != nil {
				_ = tx.Rollback().Error
			} else {
				err = tx.Commit().Error
			}
		}

		helper.Response(&res, data, err)
	}()

	// 参数校验
	if err = validator.ValidateStruct(input); err != nil {
		return
	}

	tx = database.Db.Begin()

	adminInfo := model.Admin{
		Id: c.Uid,
	}

	if err = tx.First(&adminInfo).Error; err != nil {
		// 没有找到管理员
		if err == gorm.ErrRecordNotFound {
			err = exception.AdminNotExist
		}
		return
	}

	menuInfo := model.Menu{
		Name: input.Name,
	}

	if input.Url != nil {
		menuInfo.Url = *input.Url
	}

	if input.Icon != nil {
		menuInfo.Icon = *input.Icon
	}

	if input.Accession != nil {
		menuInfo.Accession = *input.Accession
	} else {
		menuInfo.Accession = []string{}
	}

	if input.ParentId != nil {
		menuInfo.ParentId = *input.ParentId

		// 查询是否有这个 parentId
		if err = tx.Where(&model.Menu{Id: *input.ParentId}).Find(&model.Menu{}).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				err = exception.NoData
				return
			}
			return
		}
	}

	if err = tx.Create(&menuInfo).Error; err != nil {
		return
	}

	if er := mapstructure.Decode(menuInfo, &data.MenuPure); er != nil {
		err = er
		return
	}

	data.CreatedAt = menuInfo.CreatedAt.Format(time.RFC3339Nano)
	data.UpdatedAt = menuInfo.UpdatedAt.Format(time.RFC3339Nano)
	return
}

func CreateRouter(c *gin.Context) {
	var (
		input CreateMenuParams
		err   error
		res   = schema.Response{}
	)

	defer func() {
		if err != nil {
			res.Data = nil
			res.Message = err.Error()
		}
		c.JSON(http.StatusOK, res)
	}()

	if err = c.ShouldBindJSON(&input); err != nil {
		err = exception.InvalidParams
		return
	}

	res = Create(controller.NewContext(c), input)
}
