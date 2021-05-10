# 1Password Connect Go SDK

The 1Password Connect Go SDK provides access to the 1Password Connect API hosted on your infrastructure. The library is intended to be used by your applications, pipelines, and other automations to simplify accessing items stored in your 1Password vaults.

## Installation

```sh
$ go get github.com/1Password/connect-sdk-go
```

## Usage

### Environment Variables

| Variable           | Description | Feature |
|:-------------------|:------------|------:|
| `OP_CONNECT_TOKEN` | The API token to be used to authenticate the client to a 1Password Connect API. | [API Client](#/Creating-an-api-client) |
| `OP_CONNECT_HOST`  | The hostname of the 1Password Connect API | [API Client](#/Creating-an-api-client) |
| `OP_VAULT`         | If the `opvault` tag is not set the client will default to this vault UUID | [Unmarshalling](#/Unmarshalling-into-a-struct) |

### Creating an API Client

`connect.Client` instances require two pieces of configuration. A token and a hostname. There are three constructor methods provided by this library for creating your client.

- `connect.NewClient` – Accepts a hostname and a token value.
```go
client, err := connect.NewClient({connect_host}, {connect_token})
if err != nil {
	panic(err)
}
```
- `connect.NewClientFromEnvironment` – Fetches the hostname and token value from the environment
```go
client, err := connect.NewClientFromEnvironment()
if err != nil {
	panic(err)
}
```
- `connect.NewClientWithUserAgent` – Accepts a hostname, a token value, and a custom User-Agent string for identifying the client to the 1Password Connect API
```go
client, err := connect.NewClientWithUserAgent({connect_host}, {connect_token}, {user_agent})
if err != nil {
	panic(err)
}
```

### Unmarshalling into a Struct

Users can define tags on a struct and have the `connect.Client` unmarshall item data directly in them. Supported field tags are:

- `opvault` – The UUID of the vault the item should come from
- `opitem` – The title of the Item
- `opfield` – The item field whose value should be retrieved

#### Example Struct

This example struct will retrieve 3 fields from one Item and and a whole Item from another vault

```go
package main

import (
	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
)

type Config struct {
	Database string           `opitem:"Demo TF Database" opfield:".database"`
	Username string           `opitem:"Demo TF Database" opfield:".username"`
	Password string           `opitem:"Demo TF Database" opfield:".password"`
	APIKey   onepassword.Item `opvault:"7vs66j55o6md5btwcph272mva4" opitem:"API Key"`
}

var client connect.Client

func main() {
	client, err := connect.NewClientFromEnvironment()
	if err != nil {
		panic(err)
	}
	
	connect.Load(client, &c)
}

```

### Model Objects

The `onepassword.Item` model represents Items and `onepassword.Vault` represent Vaults in 1Password

### Item CRUD

The `connect.Client` also supports methods for:

- listing Vaults
```go
vaults, err := client.GetVaults()
if err != nil {
	panic(err)
}
```
- listing items in a Vault
```go
items, err := client.GetItems({vault_uuid})
if err != nil {
	panic(err)
}
```
- searching by Item Title
```go
items, err := client.GetItemsByTitle({item_title}, {vault_uuid})
if err != nil {
	panic(err)
}
```
- Retrieving Item by Vault and Item UUID
```go
item, err := client.GetItem({item_uuid}, {vault_uuid})
if err != nil {
	panic(err)
}
```
- Creating Items in a Vault
```go
item := &onepassword.Item{
	Fields: []*onepassword.ItemField{{
		Value: "mysecret",
		Type: "STRING",
	}},
	Tags:     []string{"1password-connect"},
	Category: onepassword.Login,
	Title:    "Secret String",
}

postedItem, err := client.CreateItem(item, {vault_uuid})
if err != nil {
	panic(err)
}
```
- Updating Items
```go
item2 := &onepassword.Item{
	Fields: []*onepassword.ItemField{{
		Value: "mysecret2",
		Type: "STRING",
	}},
	Tags:     []string{"1password-connect"},
	Category: onepassword.Login,
	Title:    "Second Secret String",
}

secondPostedItem, err := client.CreateItem(item2, {vault_uuid})
if err != nil {
	panic(err)
}
```
- Deleting Items
```go
err := client.DeleteItem({item_uuid}, {vault_uuid})
if err != nil {
	panic(err)
}
```

## Development

### Building

To build all packages run

```sh
$ go build ./...
```

### Running Tests

To run all tests and see test coverage run

```sh
$ go test -v ./... -cover
```

## Security

1Password requests you practice responsible disclosure if you discover a vulnerability.

Please file requests via [**BugCrowd**](https://bugcrowd.com/agilebits).

For information about security practices, please visit our [Security homepage](https://bugcrowd.com/agilebits).
