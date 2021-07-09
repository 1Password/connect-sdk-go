[//]: # (START/LATEST)
# Latest

## Features
  * A user-friendly description of a new feature. {issue-number}

## Fixes
 * A user-friendly description of a fix. {issue-number}

## Security
 * A user-friendly description of a security fix. {issue-number}

---

[//]: # (START/v1.2.0)
# v1.2.0

## Features
 * Files stored as attachments or as fields can now be retrieved (requires Connect `v1.3.0` or later).
 * Details of Connect API errors can now be inspected by unwrapping a returned error into a `onepassword.Error` struct. {#17}

---

[//]: # "START/v1.1.0"

# v1.1.0

## Features

- Vaults can be retrieved by their UUID with `GetVault(vaultUUID)`.

## Fixes

- The `API_CREDENTIAL` category is now supported. {#14}

---

[//]: # "START/v1.0.1"

# v1.0.1

## Fixes

- Includes the correct version number in the user agent when making requests

---

[//]: # "START/v1.0.0"

# v1.0.0

Use the 1Password Connect SDK to leverage 1Password Secrets Automation in your Go applications.

## Features:

- Add GetValue method for accessing the value of a field on an item by its label

---

[//]: # "START/v0.0.2"

# v0.0.2

## Features:

- Add method for retrieving all items or vaults with a given title (#1)

---

[//]: # "START/v0.0.1"

# v0.0.1

## Features:

- 1Password Connect API Client
- 1Password Connect API Models
- Automatic Item field unmarshalling through struct tags
- Integrates with existing open tracing implementations

---
