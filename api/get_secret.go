package api

//TODO:
// Get secret recieves string (path), and returns KVSecret https://github.com/hashicorp/vault/blob/69e9eb40c5a7ccc5f9b39e297901fafbe01eb14e/api/kv.go#L14

func GetSecret() {
	// config := vault.DefaultConfig()
	// config.Address = "https://vault.main2.staging.riskxint.com/"
	// //TODO: Make client available globally in the program. (Dependancy Injection)
	// client, err := vault.NewClient(config)
	// if err != nil {
	// 	log.Fatalf("unable to initialize Vault client: %v", err)
	// }

	// okta := oktaAuth.CLIHandler{}
	// authConfig := make(map[string]string)
	// authConfig["username"] = "segev.matuti@riskified.com"
	// authConfig["mount"] = "okta"

	// // Create token helper for storing token in local disk
	// myTokenHelper, _ := tokenHelper.NewInternalTokenHelper()
	// // Get cached token if exists
	// token, err := myTokenHelper.Get()
	// if token == "" {
	// 	fmt.Println("no token, fetching from Okta")
	// 	// Token is not exists, fetch from Okta
	// 	oktaSecret, _ := okta.Auth(client, authConfig)
	// 	token = oktaSecret.Auth.ClientToken
	// 	myTokenHelper.Store(token)
	// }
	// // Set token for client
	// client.SetToken(token)
}
