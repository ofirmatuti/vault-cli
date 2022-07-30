package api

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	oktaAuth "github.com/hashicorp/vault/builtin/credential/okta"
	tokenHelper "github.com/hashicorp/vault/command/token"
)

func OktaLogin(client *vault.Client) (*vault.Client) {
	okta := oktaAuth.CLIHandler{}
	authConfig := make(map[string]string)
	authConfig["username"] = "segev.matuti@riskified.com"
	authConfig["mount"] = "okta"

	// Create token helper for storing token in local disk
	myTokenHelper, _ := tokenHelper.NewInternalTokenHelper()
	// Get cached token if exists
	token, _ := myTokenHelper.Get()
	if token == "" {
		fmt.Println("no token, fetching from Okta")
		// Token is not exists, fetch from Okta
		oktaSecret, _ := okta.Auth(client, authConfig)
		token = oktaSecret.Auth.ClientToken
		myTokenHelper.Store(token)
	}
	// Set token for client
	client.SetToken(token)
	return client
}
