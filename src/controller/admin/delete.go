package admin

import (
	"fmt"
	"github.com/axetroy/go-server/src/service"
	"github.com/jinzhu/gorm"
)

func DeleteByField(field, value string) {
	var (
		err error
		tx  *gorm.DB
	)

	defer func() {
		if tx != nil {
			if err != nil {
				_ = tx.Rollback()
			} else {
				_ = tx.Commit()
			}
		}
	}()

	tx = service.Db.Begin()

	raw := fmt.Sprintf("DELETE FROM \"%v\" WHERE %s = '%v'", "admin", field, value)

	if err = tx.Exec(raw).Error; err != nil {
		return
	}
}

func DeleteAdminByAccount(account string) {
	DeleteByField("username", account)
}