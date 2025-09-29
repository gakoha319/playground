package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

// func TestGORMGen(t *testing.T) {
//      user := models.User{Name: "jinzhu2"}
//      ctx := context.Background()

//      gorm.G[models.User](DB).Create(ctx, &user)

//      if u, err := gorm.G[models.User](DB).Where(g.User.ID.Eq(user.ID)).First(ctx); err != nil {
//              t.Errorf("Failed, got error: %v", err)
//      } else if u.Name != user.Name {
//              t.Errorf("Failed, got user name: %v", u.Name)
//      }
// }

func TestGORMCallbackLogging(t *testing.T) {
	theCallback := func(tx *gorm.DB) {
		fmt.Println("Hello world from theCallBack")
	}

	cb := DB.Callback().Create().Before("gorm:create")
	assert.NotNil(t, cb)
	err := cb.Replace("test:create", theCallback)
	assert.NoError(t, err)

	user := models.User{Name: "gakoha"} // garyko
	tx := DB.Create(&user)
	assert.NotNil(t, tx)
	assert.Nil(t, tx.Error)
}
