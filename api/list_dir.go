package api

import(
	vault "github.com/hashicorp/vault/api"
	"fmt"
	"context"
	"log"
)

//TODO: recieves string (path), and returns KVSecret

func ListDir(client *vault.Client, path string) {
	fmt.Println(path)
	secret, err := client.KVv2("secret").Get(context.Background(), path)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	fmt.Println(secret)
	//TODO: print as pretty json using MarhsalIdent - https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console
	//for key, value := range secret.Data {
	//	fmt.Printf("%v \t= %v\n", key, value)
	//}
}
