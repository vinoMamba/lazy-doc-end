package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
	pg.Use(middlewares.AuthMiddleware).GET("", projectList)
	pg.Use(middlewares.AuthMiddleware).PUT("", projectUpdate)
	pg.Use(middlewares.AuthMiddleware).DELETE("/:id", projecDelete)
}

func projectCreate(c *gin.Context) {
	log := logger.New(c)
	var body request.ProjectCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("Bind json failed in userRegister")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
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
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func projectList(c *gin.Context) {
	log := logger.New(c)
	projects, err := storage.GetPorjectList(c)
	if err != nil {
		log.WithError(err).Errorln("Get project list failed in projectList")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func projectUpdate(c *gin.Context) {
	log := logger.New(c)
	var body request.ProjectUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("Bind json failed in projectUpdate")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	p, err := storage.GetProjectById(c, body.Id)

	if err != nil {
		log.WithError(err).Errorln("Get project by id failed in projectUpdate")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	p.ProjectName = body.ProjectName
	p.ProejctDesc = body.ProjectDesc
	p.IsPublic = body.IsPublic
	if err := storage.UpdateProject(c, p); err != nil {
		log.WithError(err).Errorln("Update project failed in projectUpdate")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func projecDelete(c *gin.Context) {
	log := logger.New(c)
	id := c.Param("id")
	if err := storage.DeleteProjectById(c, cast.ToInt(id)); err != nil {
		log.WithError(err).Errorln("Delete project by id failed in projecDelete")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
