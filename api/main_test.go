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
	db "techschool/samplebank/db/sqlc"
	"techschool/samplebank/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetrickey:  util.RandomString(32),
		AccessTokenDuraion: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
