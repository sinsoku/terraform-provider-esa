package esa

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceEsaMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceEsaMemberCreate,
		Read:   resourceEsaMemberRead,
		Delete: resourceEsaMemberDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"code": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceEsaMemberCreate(d *schema.ResourceData, m interface{}) error {
	api := m.(*Api)
	email := d.Get("email").(string)

	emails := []string{email}
	invitations, _, err := api.SendInvitation(emails[:])
	if err != nil {
		return err
	}

	invitation := invitations.Invitations[0]
	d.SetId(api.Team + ":" + email)
	d.Set("code", invitation.Code)
	return resourceEsaMemberRead(d, m)
}

func resourceEsaMemberRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*Api)
	log.Printf("[DEBUG] id: %s", d.Id())
	parts := strings.SplitN(d.Id(), ":", 2)
	email := parts[1]

	d.Set("email", email)
	members, _, _ := api.Members()
	for _, member := range members.Members {
		if member.Email == email {
			return nil
		}
	}

	invitations, _, err := api.PendingInvitations()
	if err != nil {
		return err
	}

	for _, invitation := range invitations.Invitations {
		if invitation.Email == email {
			d.Set("code", invitation.Code)
		}
	}

	return nil
}

func resourceEsaMemberDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*Api)
	email := d.Get("email").(string)

	members, _, _ := api.Members()
	for _, member := range members.Members {
		if member.Email == email {
			_, err := api.DeleteMember(member.ScreenName)
			if err != nil {
				return err
			}

			return nil
		}
	}

	code := d.Get("code").(string)
	_, err := api.CancelInvitation(code)
	if err != nil {
		return err
	}

	return nil
}
