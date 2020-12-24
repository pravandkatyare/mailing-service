package mjclient

import (
	"fmt"

	"github.com/pravandkatyare/mailing-service/logging"
	"github.com/spf13/viper"
)

var (
	publicKey  string
	privateKey string
)

// Init Mailjet Service API keys
func Init() {
	setKeys()
	setRegisteredAddress()
}

func setKeys() {
	if viper.GetString("MAILJET_PUBLIC_API") == "" || viper.GetString("MAILJET_PRIVATE_API") == "" {
		logging.Errorf("MailJet API keys not configured, Public key: %s & Private key: %s")
		logging.Errorf("Please provide public and private keys in configuration or export it as an environment variables")
		panic(fmt.Errorf("Mailjet API Keys not configured "))
	}
	publicKey = viper.GetString("MAILJET_PUBLIC_API")
	privateKey = viper.GetString("MAILJET_PRIVATE_API")
}

func getPublicKey() string {
	return publicKey
}

func getPrivateKey() string {
	return privateKey
}

func setRegisteredAddress() {
	if viper.GetString("mailjet.email") == "" || viper.GetString("mailjet.name") == "" {
		logging.Errorf("MailJet Registered email not configured, Email: %s & Name: %s", viper.GetString("mailjet.email"), viper.GetString("mailjet.name"))
		panic(fmt.Errorf("Mailjet Registered email not configured"))
	}
}
