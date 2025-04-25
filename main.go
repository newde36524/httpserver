package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("p", 1189, "listen port")
	download := flag.String("f", "", "download file path")
	text := flag.String("s", "", "return text")
	directory := flag.String("d", "", "directory path")
	flag.Parse()

	e := gin.New()
	if *download != "" {
		e.GET("/", func(c *gin.Context) {
			c.Header("Access-Control-Expose-Headers", "Content-Disposition")
			c.Header("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filepath.Base(*download)))
			c.Header("Content-Type", "application/octet-stream")
			c.File(*download)
		})
	}
	if *text != "" {
		e.GET("/", func(c *gin.Context) {
			c.String(200, *text)
		})
	}
	if *directory != "" {
		e.Static("/", *directory)
	}
	_ = e.Run(fmt.Sprintf(":%d", *port))
}
