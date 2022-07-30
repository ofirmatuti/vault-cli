/*
	Copyright © 2022 Ofir Matuti and Segev matuti <ofirmatuti@gmail.com, segevmatuti1@gmail.com>
*/
package main

import "log"
import vault "github.com/hashicorp/vault/api"
import "github.com/ofirmatuti/vt/cmd"
import "github.com/ofirmatuti/vt/api"

func main() {
	config := vault.DefaultConfig()
	config.Address = "https://vault.main2.staging.riskxint.com/"
	//TODO: Make client available globally in the program. (Dependancy Injection)
	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}
	api.OktaLogin(client)
	api.ListDir(client, "riski-services/devops/argocd/secrets")
	cmd.Execute()
}
