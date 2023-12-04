package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/lazy-doc-end/logger"
	"github.com/vinoMamba/lazy-doc-end/middlewares"
	"github.com/vinoMamba/lazy-doc-end/models"
	"github.com/vinoMamba/lazy-doc-end/params/request"
	"github.com/vinoMamba/lazy-doc-end/storage"
	"github.com/vinoMamba/lazy-doc-end/utils"
)

func HandleProject(r *gin.Engine) {
	pg := r.Group("/project")
	pg.Use(middlewares.AuthMiddleware).POST("", projectCreate)
}

func projectCreate(c *gin.Context) {
	log := logger.New(c)
	var body request.ProjectCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("Bind json failed in userRegister")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Bad Request",
		})
		return
	}

	if err := storage.CreateProject(c, &models.Project{
		ProjectName: body.ProjectName,
		ProejctDesc: body.ProjectDesc,
		IsPublic:    body.IsPublic,
		CreatedBy:   utils.GetCurrentUserId(c),
	}); err != nil {
		log.WithError(err).Errorln("Create project failed in projectCreate")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "OK",
	})
}
