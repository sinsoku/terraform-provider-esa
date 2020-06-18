package esa

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ESA_ACCESS_TOKEN", nil),
				Description: descriptions["token"],
			},
			"team": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ESA_TEAM", nil),
				Description: descriptions["team"],
			},
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ESA_API_ENDPOINT", "https://api.esa.io/"),
				Description: descriptions["api_endpoint"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"esa_member": resourceEsaMember(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"esa_user": dataSourceEsaUser(),
		},
		ConfigureFunc: configureProvider,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"token":        "The OAuth token",
		"team":         "The Team name",
		"api_endpoint": "The API endpoint",
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	api := NewApi(
		d.Get("token").(string),
		d.Get("team").(string),
		d.Get("api_endpoint").(string),
	)

	return api, nil
}
