package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spa-stc/newsletter/internal/run"
)

func main() {
	if err := run.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running cron app: %s \n", err.Error())
	}
}
