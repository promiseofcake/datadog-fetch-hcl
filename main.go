package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/promiseofcake/datadog-fetch-hcl/client"
	"github.com/rodaine/hclencoder"
)

const (
	envAPIKey = "DATADOG_API_KEY"
	envAppKey = "DATADOG_APP_KEY"
)

func main() {
	apiKey := os.Getenv(envAPIKey)
	appKey := os.Getenv(envAppKey)
	if apiKey == "" || appKey == "" {
		log.Fatalf("Please export %s, %s and try again", envAPIKey, envAppKey)
	}
	dd := client.NewDataDog(apiKey, appKey)

	id := flag.Int("id", 0, "Dashboard ID to retrieve")
	title := flag.String("title", "", "Dashboard Title for TF definition")
	flag.Parse()

	dash, err := dd.GetDashboard(*id)
	if err != nil {
		log.Fatalf("Error retrieving Dashboard: %s", err)
	}
	out, err := hclencoder.Encode(dash)
	if err != nil {
		log.Fatalf("Error encoding Dashboard: %s", err)
	}

	// hackery to build the dual string resource HCL
	hcl := processHCL(out, *title)
	fmt.Printf("\n%s", hcl)
}

func processHCL(bts []byte, title string) string {
	replace := fmt.Sprintf("resource \"datadog_timeboard\" \"%s\"", title)
	return strings.Replace(string(bts), "resource", replace, 1)
}
