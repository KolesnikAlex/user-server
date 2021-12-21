package http

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
	"net/http"
	"strconv"
	"user-server/app"
	"user-server/app/service"
	"user-server/config"
	"user-server/internal/database"
)

func Setup() {
	routes(app.App.Echo, app.App.Config, app.App.PostgresClient)
}

func routes(e *echo.Echo, config *config.Config, postgresClient *sqlx.DB) {

	NewUserHandler(config, postgresClient).Routes(e)

}


type UserHandler struct {
	config  *config.Config
	service service.UserService
}

func NewUserHandler(config *config.Config, postgresClient *sqlx.DB) *UserHandler {
	return &UserHandler{
		config: config,
		service: database.NewSQLRepo(postgresClient),
	}
}

func (h UserHandler) Routes(e *echo.Echo) {
	reference := e.Group("/user")
	reference.GET("/:id", h.GetUser)
	reference.POST("/", h.AddUser)
	reference.DELETE("/:id", h.RemoveUser)
	reference.PUT("/", h.UpdateUser)
}

func (h UserHandler) GetUser(c echo.Context) error {
	id := strToInt(c.Param("id"))

	res, err := h.service.GetUser(c.Request().Context(), int64(id))
	if err != nil {
		log.Error().Err(err).Msg("(http) err get user")
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h UserHandler) AddUser(c echo.Context) error {
	var user service.User

	err := c.Bind(&user)
	if err != nil {
		log.Error().Err(err).Msg("bad parse struct")
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.service.AddUser(c.Request().Context(), user)
	if err != nil {
		log.Error().Err(err).Msg("invalid save user")
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("User with id: %d was add", user.ID))
}

func (h UserHandler) RemoveUser(c echo.Context) error {
	id := strToInt(c.Param("id"))
	err := h.service.RemoveUser(c.Request().Context(), int64(id))
	if err != nil {
		log.Error().Err(err).Msg("invalid delete user")
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("User with id: %d was deleted", id))
}

func (h UserHandler) UpdateUser(c echo.Context) error {
	var user service.User
	err := c.Bind(&user)
	if err != nil {
		log.Error().Err(err).Msg("bad parse struct")
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.service.UpdateUser(c.Request().Context(), user)
	if err != nil {
		log.Error().Err(err).Msg("invalid save user")
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("User with id: %d was updated", user.ID))
}


func strToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Error().Err(err).Msgf("Don't convert, %s", s)
	}
	return res
}