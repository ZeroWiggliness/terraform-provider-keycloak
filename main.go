package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/mrparkers/terraform-provider-keycloak/provider"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug: debugMode,
		// TODO: update this string with the full name of your provider as used in your configs
		ProviderAddr: "registry.terraform.io/zerowiggliness/keycloak",
		ProviderFunc: func() *schema.Provider {
			return provider.KeycloakProvider(nil)
		},
	}

	plugin.Serve(opts)
}
