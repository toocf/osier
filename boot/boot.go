package boot

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"osier/config"
)

var (
	Vpr *viper.Viper
	Cfg *config.Config
	Log *zap.Logger
	Gdb *gorm.DB
	Rds *redis.Client
	Ege *gin.Engine
)
