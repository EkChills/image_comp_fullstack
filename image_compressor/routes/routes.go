package routes

import (
	"image"
	"image/jpeg"
	"net/http"
	"os"
	"path/filepath"
    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"POST", "GET", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true,
    }))
	server.GET("hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	server.POST("/upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("picture")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		tmpDir, err := os.MkdirTemp("", "uploads-*")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	

		defer os.RemoveAll(tmpDir)

		fullPath := filepath.Join(tmpDir, fileHeader.Filename)

		if err := c.SaveUploadedFile(fileHeader, fullPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		img, err := os.Open(fullPath)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		tempOutputPath, err := os.MkdirTemp("", "outputs-*")
		
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				defer os.RemoveAll(tempOutputPath)

				outputFilePath := filepath.Join(tempOutputPath, "img-new.jpg")

		out, err := os.Create(outputFilePath)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		decodedImage, _, err := image.Decode(img)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := jpeg.Encode(out, decodedImage, &jpeg.Options{Quality: 30}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
		}

		compressedFile, err := os.Open(outputFilePath)

		if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

		defer compressedFile.Close()

		c.File(outputFilePath)


	})
}
