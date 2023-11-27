package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/logger"
	"github.com/vinoMamba/lazy-doc-end/middlewares"
	"github.com/vinoMamba/lazy-doc-end/params/request"
	"github.com/vinoMamba/lazy-doc-end/storage"
	"github.com/vinoMamba/lazy-doc-end/utils"
)

func HandleProject(r *gin.Engine) {
	r.Use(middlewares.AuthMiddleware).POST("/project", projectCreate)
	r.Use(middlewares.AuthMiddleware).GET("/project", projectGetList)
}

func projectCreate(c *gin.Context) {
	db := storage.NewQuery()
	log := logger.New(c)
	var body request.ProjectCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Bad Request",
			"data":    nil,
		})
		return
	}
	if body.ProjectName == "" {
		log.Error("project name is empty")
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "password verify failed",
			"data":    nil,
		})
		return
	}
	userId := utils.GetCurrentUserId(c)

	_, err := db.CreateProject(c, storage.CreateProjectParams{
		ProjectName: body.ProjectName,
		ProjectDescription: sql.NullString{
			String: body.ProjectDesc,
			Valid:  true,
		},
		CreatedBy: userId,
		IsPublic: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		IsDeleted: sql.NullBool{
			Bool:  false,
			Valid: true,
		},
	})

	if err != nil {
		log.WithError(err).Errorln("create project failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    1,
			"message": "server error",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    nil,
	})
}

func projectGetList(c *gin.Context) {

}
