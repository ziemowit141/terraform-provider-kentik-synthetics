//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/kentik/terraform-provider-kentik-synthetics/synthetics"
)

func main() {
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like Delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: synthetics.NewProvider}

	if debugMode {
		err := plugin.Debug(context.Background(), "kentik/automation/kentik-synthetics", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		plugin.Serve(opts)
	}
}
