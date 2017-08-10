package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"github.com/Depado/rpsheet/models"
)

func main() {
	var err error
	var c []byte
	var nogaj models.Character

	if c, err = ioutil.ReadFile("nogaj.yml"); err != nil {
		logrus.WithError(err).Fatal("Couldn't load file")
	}
	if err = yaml.Unmarshal(c, &nogaj); err != nil {
		logrus.WithError(err).Fatal("Couldn't unmarshal")
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nogaj)
	})
	r.Run(":8080")
}
