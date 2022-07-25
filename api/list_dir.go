package api

//TODO: recieves string (path), and returns KVSecret

func ListDir() {
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
