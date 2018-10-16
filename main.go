package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/promiseofcake/datadog-fetch-hcl/client"
	"github.com/promiseofcake/datadog-fetch-hcl/convert"
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
	debug := flag.Bool("debug", false, "Debug Datadog API Output")
	flag.Parse()

	bts, err := dd.FetchJSON(*id)
	if err != nil {
		log.Fatalf("Error retrieving Dashboard: %s", err)
	}

	// debug the API response
	if *debug {
		var pretty bytes.Buffer
		l := log.New(os.Stderr, "", 0)

		json.Indent(&pretty, bts, "", "\t")
		l.Printf("%s\n", pretty.Bytes())
	}

	dash, err := convert.BuildDashbard(bts)
	if err != nil {
		log.Fatalf("Error building Dashboard: %s", err)
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
