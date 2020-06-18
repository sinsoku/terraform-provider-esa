package esa

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceEsaUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEsaUserRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEsaUserRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*Api)
	user, _, err := api.User()
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(user.Id))
	d.Set("name", user.Name)

	return nil
}
