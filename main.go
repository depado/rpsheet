package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/Depado/rpsheet/conf"
	"github.com/Depado/rpsheet/models"
)

func main() {
	var err error
	var nogaj models.Character

	// Loading character
	if err = nogaj.Load("nogaj.yml"); err != nil {
		logrus.WithError(err).WithField("file", "nogaj.yml").Fatal("Couldn't load char")
	}

	// Loading configuration
	if err = conf.Load("conf.yml"); err != nil {
		logrus.WithError(err).WithField("file", "conf.yml").Fatal("Couldn't load configuration file")
	}

	// Set gin mode
	if !conf.C.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Router setup
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nogaj)
	})

	// Running server
	if err = r.Run(fmt.Sprintf("%s:%d", conf.C.Host, conf.C.Port)); err != nil {
		logrus.WithError(err).Fatal("Couldn't start listening")
	}
}
