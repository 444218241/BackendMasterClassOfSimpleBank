/*
@time: 2021/9/6 10:56
@author: chenZouLu
@file: mian_test
@software: GoLand
@note:
*/

package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
