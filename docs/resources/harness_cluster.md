---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "imagetest_harness_cluster Resource - terraform-provider-imagetest"
subcategory: ""
description: |-
  A harness that provisions a remote cluster in a cloud environment and runs steps as requested
---

# imagetest_harness_cluster (Resource)

A harness that provisions a remote cluster in a cloud environment and runs steps as requested



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `inventory` (Attributes) The inventory this harness belongs to. This is received as a direct input from a data.imagetest_inventory data source. (see [below for nested schema](#nestedatt--inventory))
- `name` (String) The name of the harness. This must be unique within the scope of the provided inventory.

### Optional

- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `id` (String) The unique identifier for the harness. This is generated from the inventory seed and harness name.
- `skipped` (Boolean) Whether or not to skip creating the harness based on runtime inputs and the dependent features within this inventory.

<a id="nestedatt--inventory"></a>
### Nested Schema for `inventory`

Required:

- `seed` (String)


<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) The maximum time to wait for the k3s harness to be created.
