package app

import (
	"database/sql"
	"embed"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/phuslu/log"
	"github.com/pressly/goose/v3"
	"net"
	"net/http"
	"os"
	"user-server/config"
	myGrpc "user-server/grpc"
	"user-server/internal/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Application struct {
		Config         *config.Config
		Echo           *echo.Echo
		PostgresClient *sqlx.DB
		grpcController myGrpc.GrpcUserServiceServer
	//	HTTPClient     *http.Client
	}
)

var (
	App *Application
)

func New() *Application {
	App = &Application{
		Config: &config.Config{},
	}

	App.Config = config.LoadConfig()
	App.setupEcho()
	App.setupPostgresClient()
	App.setupGrpcController()

	return App
}

func (app *Application) setupEcho() {
	app.Echo = echo.New()
}


func (app *Application) setupPostgresClient() {

	err := setupEmbedMigrations(app.Config.PostgresDB.DbString)
	if err != nil {
		log.Fatal().Err(err).Msg("migration failed")
	}
	client, err := sqlx.Connect("postgres", app.Config.PostgresDB.DbString)
	if err != nil {
		log.Fatal().Err(err).Msg("postgres client creation:")
	}
	//defer client.Close()
	//may be delete client.Ping(), because sqlx.Connect() already contains verify with a ping?
	if err = client.Ping(); err != nil {
		log.Fatal().Err(err).Msg("postgres client ping:")
	}
	app.PostgresClient = client
}

//go:embed migrate/*.sql
var embedMigrations embed.FS


func setupEmbedMigrations(dbConnect string) error {
	db, err := sql.Open("postgres", dbConnect)
	if err != nil {
		return err
	}
	defer db.Close()
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(db, "migrate"); err != nil {
		return err
	}
	return nil
}

func (app *Application) setupGrpcController() {
	sqlRepo := database.NewSQLRepo(app.PostgresClient)
	app.grpcController = myGrpc.NewUserServiceController(sqlRepo)

}

func (app *Application) Run() {
	//starting http server
	log.Info().Msgf("listening on %s", app.Config.HttpServer.Port)
	err := http.ListenAndServe(app.Config.HttpServer.Port, app.Echo)
	if err != nil {
		panic(err)
	}

	//starting grpc server
	server := grpc.NewServer()
	myGrpc.RegisterGrpcUserServiceServer(server, app.grpcController)
	reflection.Register(server)

	con, err := net.Listen("tcp", app.Config.GrpcServer.Port)
	if err != nil {
		panic(err)
	}

	log.Printf("Starting gRPC server on %s...\n", con.Addr().String())
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
