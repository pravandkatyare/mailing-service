package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pravandkatyare/mailing-service/app"
	mjclient "github.com/pravandkatyare/mailing-service/mailingClients/mailjet"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	setupDefaults()
	app.StartApplication()

}

// setting default variables
func setupDefaults() {

	// --------logrus default setup-------
	host, err := os.Hostname()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Service Startup")

	// Show the version and build info
	logrus.Infof("Golang OS             : %s", runtime.GOOS)
	logrus.Infof("Golang Arch           : %s", runtime.GOARCH)
	logrus.Infof("Service Host          : %s", host)

	// -------logrus default setup complete------

	// -------viper default setup-------

	viper.AutomaticEnv()

	defaults := map[string]interface{}{
		"PORT":           8080,
		"header.api-key": "x-default",
	}

	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	// Read from application.yml file
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Config file not found: %s", err))
		} else {
			panic(fmt.Errorf("Error reading config file: %s", err))
		}
	}
	// --------viper default setup complete-------

	// --------Initialize MailJet Api Keys-------
	mjclient.Init()
	// --------Initializing MailJet Api Keys complete-------

}
