package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	pipedrivesdk "github.com/alexellis/pipedrive-sdk"
)

// This example imports a contact into Pipedrive from a 
// number of command line arguments
//
// To add a "Person", you must first create or find the Person's 
// Organization.
// Then you add the Organization ID to the Person's creation
// requestion
// You should also search for the person to find out whether 
// they already exist in Pipedrive.
func main() {
	if len(os.Args) != 4 {
		panic("Give org name email as arguments")
	}

	if _, ok := os.LookupEnv("organization"); !ok {
		panic("environment variable organization not set")
	}

	term := os.Args[1]

	personTerm := os.Args[2]
	personEmail := os.Args[3]

	apiKey := os.Getenv("PIPE")
	client := pipedrivesdk.NewPipeDriveClient(http.DefaultClient, apiKey, os.Getenv("organization"))

	res, err := client.SearchOrg(term)
	if err != nil {
		panic(err)
	}

	orgID := -1
	if res.Data.Items != nil && len(res.Data.Items) > 0 {
		if res.Data.Items[0].ResultScore > 0.2 {
			fmt.Printf("Found: %v\n", res)
			orgID = res.Data.Items[0].Item.ID
		}
	}
	if orgID == -1 {
		createRes, err := client.CreateOrg(term)
		if err != nil {
			panic(err)
		}
		orgID = createRes.Data.ID
		fmt.Printf("Created: %v\n", createRes)
	}

	personID := -1
	pRes, err := client.SearchPerson(personTerm)
	if err != nil {
		panic(err)
	}
	if pRes.Success {
		if pRes.Data.Items != nil && len(pRes.Data.Items) > 0 {
			if pRes.Data.Items[0].ResultScore > 0.2 {
				fmt.Printf("Found: %v\n", pRes)
				personID = pRes.Data.Items[0].Item.ID
			}
		}
	} else {
		panic("searchPerson failed")
	}

	if personID == -1 {
		log.Printf("Create user with ORG: %d", orgID)

		createPersonRes, err := client.CreatePerson(personTerm, personEmail, orgID)
		if err != nil {
			panic(err)
		}

		if createPersonRes.Success {
			log.Printf("Created: %v\n", createPersonRes.Data.ID)
		}
	}
}
