package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := models.User{Name: "jinzhu"}

	InitializeDB()
	DB.Create(&user)

	var result models.User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

// func TestGORMGen(t *testing.T) {
//      InitializeDB()
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
	InitializeDB()

	var logBuffer bytes.Buffer
	DB.Config.Logger = logger.New(
		log.New(&logBuffer, "", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)

	callbackCalled := false
	theCallback := func(tx *gorm.DB) {
		callbackCalled = true
	}

	err := DB.Callback().Create().Replace("gorm:create", theCallback)
	assert.NoError(t, err)

	user := models.User{
		Name: "gakoha",
		Company: models.Company{
			Name: "garyko",
		},
	}
	assert.False(t, callbackCalled)
	tx := DB.Create(&user)
	assert.NotNil(t, tx)
	assert.Nil(t, tx.Error)
	assert.True(t, callbackCalled)

	fmt.Fprintf(os.Stderr, "logBuffer-->%s<--logBuffer\n", logBuffer.String())
}
