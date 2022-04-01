package api

import (
	"fmt"
	"github.com/cloudstruct/cargo/config"
	"github.com/cloudstruct/cargo/logging"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) error {
	// Disable gin debug output
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	// Configure router
	router := gin.New()
	// Catch panics and return a 500
	router.Use(gin.Recovery())
	// Access logging
	accessLogger := logging.GetAccessLogger()
	router.Use(ginzap.Ginzap(accessLogger, "", true))
	router.Use(ginzap.RecoveryWithZap(accessLogger, true))

	// Configure healthcheck route
	router.GET("/healthcheck", handleHealthcheck)

	// Start listener
	err := router.Run(fmt.Sprintf("%s:%d", cfg.Api.Address, cfg.Api.Port))
	return err
}

func handleHealthcheck(c *gin.Context) {
	// TODO: add some actual health checking here
	c.JSON(200, gin.H{"failed": false})
}
