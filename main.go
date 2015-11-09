package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

type ExampleClient struct {
	ApiKey string
	Url    string
}

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{ // Source https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/provider.go#L20-L43
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your provider. Source https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/schema.go#L29-L142
			"api_key": &schema.Schema{ // The provider understands that there is a field called "api_key"
				Type:        schema.TypeString, // Other supported types https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/schema.go#L36-L42
				Required:    true,
				Description: "API Key used to authenticate with the service provider",
			},
			"url": &schema.Schema{ // The provider understands that there is a field called "url"
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL to the API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{ // Source https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/resource.go#L17-L81
			"example_server": &schema.Resource{
				SchemaVersion: 1,
				Create:        createFunc,
				Read:          readFunc,
				Update:        updateFunc,
				Delete:        deleteFunc,
				Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
					"hostname": &schema.Schema{ // The provider understands that the resource "example_server" has a field called "hostname"
						Type:     schema.TypeString,
						Required: true,
					},
					"cpus": &schema.Schema{ // The provider understands that the resource "example_server" has a field called "cpus"
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			client := ExampleClient{
				ApiKey: d.Get("api_key").(string),
				Url:    d.Get("url").(string),
			}

			// You could have some field validations here, like checking that
			// the API Key is has not expired or that the username/password
			// combination is valid, etc.

			return client, nil
		},
	}
}

/*\
|*| The methods defined below will get called for each resource that needs to
|*| get created (createFunc), read (readFunc), updated (updateFunc) and deleted (deleteFunc).
|*| For example, if 10 resources need to be created then `createFunc`
|*| will get called 10 times every time with the information for the proper
|*| resource that is being mapped.
|*|
|*| If at some point any of these functions returns an error, Terraform will
|*| imply that something went wrong with the modification of the resource and it
|*| will prevent the execution of further calls that depend on that resource
|*| that failed to be created/updated/deleted.
\*/

func createFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
