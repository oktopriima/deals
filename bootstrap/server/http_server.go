package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/bootstrap/config"
)

type (
	EchoInstance struct {
		Router *echo.Echo
		Config config.AppConfig
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

var trans ut.Translator

func NewEchoInstance(r *echo.Echo, cfg config.AppConfig) *EchoInstance {
	return &EchoInstance{
		Router: r,
		Config: cfg,
	}
}

func (server *EchoInstance) runHttp() (err error) {
	port := fmt.Sprintf(":%s", server.Config.App.Port)

	// Set up validator
	v := validator.New()

	// Set up English translator
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")

	entranslations.RegisterDefaultTranslations(v, trans)

	server.Router.Validator = &CustomValidator{
		validator: v,
	}
	if err = server.Router.Start(port); err != nil {
		return err
	}

	return
}

func (server *EchoInstance) RunWithGracefullyShutdown() {
	// run group on another thread
	go func() {
		err := server.runHttp()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Router.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}

func (c *CustomValidator) Validate(i interface{}) error {
	err := c.validator.Struct(i)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"code":    http.StatusBadRequest,
		})
	}

	e := make(map[string]string)
	for _, msg := range validationErrors {
		e[strings.ToLower(msg.Field())] = msg.Translate(trans)
	}

	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": e,
	})
}
