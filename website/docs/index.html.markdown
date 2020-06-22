---
layout: "esa"
page_title: "Provider: esa"
description: |-
  The esa provider is used to interact with esa resources.
---

# esa Provider

The esa provider is used to interact with esa resources.

The provider allows you to manage your esa team's resources.
It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the esa Provider
provider "esa" {
  token = "my-oauth-token"
  team  = "my-team"
}

# Add a member to the team
resource "esa_member" "example" {
  email = "mail@example.com"
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - (Required) This is the OAuth token. It must be provided, but it can also be sourced from
  the `ESA_ACCESS_TOKEN` environment variable.

* `team` - (Required) This is the team name. It must be provided, but it can also be sourced from
  the `ESA_TEAM` environment variable.

* `api_endpoint` - (Optional) This is the API endpoint. It is optional to provide this value and
  it can also be sourced from the `ESA_API_ENDPOINT` environment variable.  The value must end with a slash.
