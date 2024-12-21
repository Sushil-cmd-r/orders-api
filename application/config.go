package application

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Port  uint16
	DBUrl string
}

func (a *App) loadConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	cfg := config{Port: 8080}
	if p, ok := os.LookupEnv("PORT"); ok {
		if port, err := strconv.ParseInt(p, 10, 16); err != nil {
			cfg.Port = uint16(port)
		}
	}

	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		return errors.New("no DB_URL variable set")
	}
	cfg.DBUrl = dbUrl

	a.cfg = cfg
	return nil
}
