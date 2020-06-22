---
layout: "esa"
page_title: "esa: esa_member"
description: |-
  Provides an esa member resource.
---

# esa_member

This resource allows you to add/remove users from your team.

## Example Usage

```hcl
resource "esa_member" "example" {
  email = "mail@example.com"
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Required) The email to invite to the team.

## Attributes Reference

The following additional attributes are exported:

* `code` - The identifier of the invitation.

## Import

esa member can be imported using the `team:mail`, e.g.

```
$ terraform import esa_member.example hashicorp:mail@example.com
```
