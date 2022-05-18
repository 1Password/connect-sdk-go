[//]: # (START/LATEST)
# Latest

## Features
  * A user-friendly description of a new feature. {issue-number}

## Fixes
 * A user-friendly description of a fix. {issue-number}

## Security
 * A user-friendly description of a security fix. {issue-number}

---

[//]: # "START/v1.4.0"

# v1.4.0

## Features

- A field's `GeneratorRecipe` now supports a set of characters that should be excluded when generating a password. This is achieved with the `ExcludeCharacters` field of the `GeneratorRecipe` struct. (requires Connect `v1.4.0` or later) (#57)
- SDK functions now accept item/vault titles and UUIDs as parameters. (#55)
- A vault can now be fetched by either its title or UUID. (#52)
- SDK now supports 2 new Item categories: `MedicalRecord` and `SSHKey`. (#51)
- The SDK now enables to load item autofill URLs into structs. (#56)
- `ItemURL` struct now has an extra field which represent the label of the autofill URL. (#53)
- Readme now has more examples for using the SDK. (#32)

## Fixes

- `GetItemsByTitle` properly returns a list of items with all their detais instead of just their summaries. {#38}

---

[//]: # "START/v1.3.0"

# v1.3.0

## Features

- Added the ability to get the TOTP value for an item. {44}
- Added method to retrieve a list of files for an item {39}
- Added method to download a file {39}
- Added the ability to create tags for sections when loading a struct {37}
- Added method for deleting an item by id {33}
- Added item state field and deprecated trashed field {30}

## Fixes

- Added UUID validation {41}

---

[//]: # "START/v1.2.0"

# v1.2.0

## Features

- Files stored as attachments or as fields can now be retrieved (requires Connect `v1.3.0` or later).
- Details of Connect API errors can now be inspected by unwrapping a returned error into a `onepassword.Error` struct. {#17}

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
