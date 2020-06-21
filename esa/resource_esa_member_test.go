package esa

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccEsaMember_basic(t *testing.T) {
	userEmail := os.Getenv("ESA_TEST_USER_EMAIL")
	if userEmail == "" {
		t.Skip("This test requires you to set the test user's email address (set it by exporting ESA_TEST_USER_EMAIL)")
	}

	rn := "esa_member.test"
	codeRe := regexp.MustCompile("^[[:alnum:]]+")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEsaMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccEsaMemberConfig(userEmail),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEsaMemberExists(rn, userEmail),
					resource.TestCheckResourceAttr(rn, "email", userEmail),
					resource.TestMatchResourceAttr(rn, "code", codeRe),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckEsaMemberExists(n, email string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		parts := strings.SplitN(rs.Primary.ID, ":", 2)
		email := parts[1]

		api := testAccProvider.Meta().(*Api)
		invitations, _, err := api.PendingInvitations()
		if err != nil {
			return err
		}

		for _, invitation := range invitations.Invitations {
			if invitation.Email == email {
				return nil
			}
		}

		return fmt.Errorf("Not Found a invitation for %s", email)
	}
}

func testAccCheckEsaMemberDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "esa_member" {
			continue
		}

		parts := strings.SplitN(rs.Primary.ID, ":", 2)
		email := parts[1]

		api := testAccProvider.Meta().(*Api)
		invitations, resp, err := api.PendingInvitations()

		if resp.StatusCode == 404 {
			return nil
		}
		if err != nil {
			return err
		}

		for _, invitation := range invitations.Invitations {
			if invitation.Email == email {
				return fmt.Errorf("invitation %s still exists", email)
			}
		}

		return nil
	}

	return nil
}

func testAccEsaMemberConfig(email string) string {
	return fmt.Sprintf(`
resource "esa_member" "test" {
  email = "%s"
}
`, email)
}
