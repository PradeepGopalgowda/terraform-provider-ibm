---
layout: "ibm"
page_title: "IBM : resource_group"
sidebar_current: "docs-ibm-resource-resource-group"
description: |-
  Manages IBM Resource Group.
---

# ibm\_resource_group

Provides a resource group resource. This allows resource groups to be created, and updated and deleted.

## Example Usage

```hcl

resource "ibm_resource_group" "resourceGroup" {
  name     = "prod"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string)The name of the resource group.
* `quota_id` - (Removed, string) The id of the quota.You can [refer to a quota by name using a data source](../d/resource_quota.html).
* `tags` - (Optional, array of strings) Tags associated with the resource group instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new resource group.
* `default` - Specifies whether its default resource group or not.
* `state` - State of the resource group.


## Import

ibm_resource_group can be imported using resource group id, eg

```
$ terraform import ibm_resource_group.example 5ffda12064634723b079acdb018ef308
```