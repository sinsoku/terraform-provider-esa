# Terraform Provider for esa

![Test](https://github.com/sinsoku/terraform-provider-esa/workflows/Test/badge.svg)

## Installation

Download the latest version from [GitHub Releases](https://github.com/sinsoku/terraform-provider-esa/releases).

Unzip the file, and then manually installed into `~/.terraform.d/plugins`.

## Usage Example

```hcl
provider "esa" {
  token = "${var.esa_token}"
  team  = "${var.esa_team}"
}

resource "esa_member" "foo" {
  email = "mail@example.com"
}

data "esa_user" "current" {}

output "current_user_name" {
  value = data.esa_user.current.name
}
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/sinsoku/terraform-provider-esa/.
