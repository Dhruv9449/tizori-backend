package cmd

import (
	"log"
	"os"

	"github.com/GDGVIT/Tizori-backend/api"
	tizoriCli "github.com/GDGVIT/Tizori-backend/cli"
	"github.com/GDGVIT/Tizori-backend/internal/auth"
	tizoriCrypto "github.com/GDGVIT/Tizori-backend/internal/crypto"
	"github.com/GDGVIT/Tizori-backend/internal/database"
	"github.com/GDGVIT/Tizori-backend/internal/models"
	"github.com/GDGVIT/Tizori-backend/internal/models/seeds"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
)

// Tizori CLI App
type TizoriApp struct {
	env Env

	cliApp *cli.App
	webApp *fiber.App
}

// Environment variables
type Env struct {
	// Fiber Variables
	fiberPort string
	debug     string

	// Database Variables
	postgresUrl string

	// Auth Variables
	jwtSecret string
	aesKey    string
}

// Method to create a new TizoriApp
func NewTizoriCliApp() *TizoriApp {
	var vittyApp TizoriApp
	vittyApp.init()
	return &vittyApp
}

// Method to set environment variables
func (t *TizoriApp) setEnv() {
	t.env.fiberPort = os.Getenv("FIBER_PORT")
	t.env.debug = os.Getenv("DEBUG")
	t.env.postgresUrl = os.Getenv("POSTGRES_URL")
	t.env.jwtSecret = os.Getenv("JWT_SECRET")
	t.env.aesKey = os.Getenv("AES_KEY")
}

// Method to initialize CLI app
func (t *TizoriApp) initCliApp() {
	t.cliApp = tizoriCli.NewCliApp()

	// Adding Run command
	t.cliApp.Commands = append(t.cliApp.Commands, &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run the server",
		Action: func(c *cli.Context) error {
			t.webApp.Listen(t.env.fiberPort)
			return nil
		},
	})

	t.cliApp.Commands = append(t.cliApp.Commands, &cli.Command{
		Name:    "seed",
		Aliases: []string{"s"},
		Usage:   "Seed the DB",
		Action: func(c *cli.Context) error {
			seeds.IntializeSeeds()
			return nil
		},
	})
}

// Method to initialize Web app
func (t *TizoriApp) initWebApp() {
	t.webApp = api.NewWebApi()
}

// Method to initialize the TizoriApp
func (t *TizoriApp) init() {
	// Set environment variables
	t.setEnv()

	// Connect to database
	database.Connect(t.env.debug, t.env.postgresUrl)

	// Initialize models
	models.InitializeModels()

	// Initialize auth
	auth.InitializeAuth(t.env.jwtSecret)

	// Initialize AES Key
	tizoriCrypto.InitializeAESKey(t.env.aesKey)

	// Initialize Web app
	t.initWebApp()

	// Initialize CLI app
	t.initCliApp()
}

func (t *TizoriApp) Run() {
	err := t.cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
