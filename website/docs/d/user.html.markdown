---
layout: "esa"
page_title: "esa: esa_user"
description: |-
  Get information about an authenticated user.
---

# esa\_user

Use this data source to retrieve information about an authenticated user.

## Example Usage

```hcl
data "esa_user" "current" {}
```

## Argument Reference

There are no arguments available for this data source.

## Attributes Reference

 * `name` - the user's full name.
