package api

import (
	"fmt"
	"log"

	"github.com/Amirrezaghanavati/golang-web-api/api/middlewares"
	"github.com/Amirrezaghanavati/golang-web-api/api/routers"
	"github.com/Amirrezaghanavati/golang-web-api/api/validations"
	"github.com/Amirrezaghanavati/golang-web-api/common"
	"github.com/Amirrezaghanavati/golang-web-api/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
	passwordPolicy, err := common.NewPasswordPolicy(common.PasswordRules(cfg.Password))
	if err != nil {
		log.Fatalf("invalid password configuration: %v", err)
	}

	r := gin.New()
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		validate.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err := validate.RegisterValidation("passwordValidator", validations.PasswordValidator(passwordPolicy)); err != nil {
			log.Fatalf("register password validator: %v", err)
		}
	}
	r.Use(gin.Logger(), gin.Recovery(), middlewares.CORS(cfg.Cors.AllowOrigins), middlewares.LimitByRequest())

	api := r.Group("api")

	v1 := api.Group("/v1")
	routers.HealthRouter(v1.Group("/health"))

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
