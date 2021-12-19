package http

import (
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
	reference := e.Group("/coins")
	reference.GET("/:id", h.GetUser)
}

func (h UserHandler) GetUser(c echo.Context) error {
	//log.Info().Msg("get coin from id")
	id := strToInt(c.Param("id"))

	res, err := h.service.GetUser(c.Request().Context(), int64(id))
	if err != nil {
		log.Error().Err(err).Msg("(http) err get coins from id")
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func strToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		log.Error().Err(err).Msgf("Dont convert, %s", s)
	}
	return res
}