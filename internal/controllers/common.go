package controllers

import (
	"embed"
	"fmt"
	"github.com/alhaos/uselessAuthorization/internal/autorizaton"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed template/*.html
var templatesFS embed.FS

type EndpointControllers struct {
	auth *autorizaton.Auth
}

func SetTemplates(router *gin.Engine) error {
	const op = "controllers.SetTemplates"
	t, err := template.ParseFS(templatesFS, "template/*.html")
	if err != nil {
		return fmt.Errorf("%s unable parse tamplates: %w", op, err)
	}
	router.SetHTMLTemplate(t)
	return nil
}

func New(auth *autorizaton.Auth) *EndpointControllers {
	return &EndpointControllers{auth: auth}
}

func (ec *EndpointControllers) RegisterRoutes(r *gin.Engine) {
	r.GET("/login", ec.loginController)
	r.POST("/auth", ec.authController)
	r.GET("protected", ec.protectedController)
}

func (ec *EndpointControllers) loginController(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func (ec *EndpointControllers) authController(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	result, err := ec.auth.Check(username, password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if !result {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "authorization failed",
		})
	}

	context.Redirect(http.StatusFound, "/protected")
}

func (ec *EndpointControllers) protectedController(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "authorization success",
	})
}
