## Usage

Below, you can find a selection of the most used functionality of the Connect Go SDK. For more detailed information about the content of the SDK, please refer to the [GoDocs](https://pkg.go.dev/github.com/1Password/connect-sdk-go).

### Quickstart

Reading a secret:

```go
import "github.com/1Password/connect-sdk-go/connect"

func main () {
	client := connect.NewClient("<your_connect_host>", "<your_connect_token>")
	item, err := client.GetItem("<item-uuid>", "<vault-uuid>")
	if err != nil {
		log.Fatal(err)
	}
}
```

Writing a secret:

```go
import (
    "github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
)

func main () {
	client := connect.NewClient("<your_connect_host>", "<your_connect_token>")
	item := &onepassword.Item{
		Title:    "Secret String",
		Category: onepassword.Login,
		Fields: []*onepassword.ItemField{{
			Value: "mysecret",
			Type:  "STRING",
		}},
	}

	postedItem, err := client.CreateItem(item, "<vault-uuid>")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Creating an API Client

A 1Password Connect client (`connect.Client`) is required to make requests to the Connect server via the 1Password Go SDK.
The client is configured with a token and a hostname. Three constructor methods that allow for creating the 1Password Connect client are provided.

- `connect.NewClient` – Accepts a hostname and a token value.

  ```go
  package main

  import "github.com/1Password/connect-sdk-go/connect"

  func main () {
      client := connect.NewClient("<your_connect_host>", "<your_connect_token>")
  }
  ```

- `connect.NewClientFromEnvironment` – Fetches the hostname and token value from the environment, and expects these to be passed as environment variables (`OP_CONNECT_HOST` and `OP_CONNECT_TOKEN`, respectively).

Assuming that `OP_CONNECT_TOKEN` and `OP_CONNECT_HOST` have been set as environment variables, the `connect.NewClientFromEnvironment` can be invoked as such:

```go
package main

import "github.com/1Password/connect-sdk-go/connect"

func main () {
	client, err:= connect.NewClientFromEnvironment()
	if err != nil {
		panic(err)
	}
}
```

- `connect.NewClientWithUserAgent` – Accepts a hostname, a token value, and a custom User-Agent string for identifying the client to the 1Password Connect API:

```go
package main

import "github.com/1Password/connect-sdk-go/connect"

func main () {
	client := connect.NewClientWithUserAgent("<your_connect_host>", "<your_connect_token>", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4")
}
```

### Model Objects

The `onepassword.Item` model represents items and `onepassword.Vault` represents vaults, in 1Password.

### Item CRUD

The `connect.Client` supports methods for:

#### Retrieving a list of vaults that the Connect token has permission to read

```go
	vaults, err := client.GetVaults()
	if err != nil {
		log.Fatal(err)
	}
```

#### Retrieving a vault:

```go
	vault, err := client.GetVault("vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Retrieving all items in a vault

```go
	items, err := client.GetItems("vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Retrieving an item by title

To retrieve all items in a vault with a given title:

```go
	items, err := client.GetItemsByTitle("items-title", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

In case the item title is unique for a vault, another function is available as well, returning only one item, instead of a slice:

```go
	item, err := client.GetItemByTitle("item-title", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Retrieving items by vault UUID and item UUID

```go
	item, err := client.GetItem("item-uuid", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Creating an item in a vault

```go
	item := &onepassword.Item{
		Title:    "Secret String",
		Category: onepassword.Login,
		Tags:     []string{"1password-connect"},
		Fields: []*onepassword.ItemField{{
			Value: "mysecret",
			Type: "STRING",
		}},
	}

	postedItem, err := client.CreateItem(item, "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Updating an item

```go
	item, err := client.GetItem("item-uuid", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}

	item.Title = "new title"
	client.UpdateItem(item, "vault-uuid")
```

#### Deleting an item

```go
	item, err := client.GetItem("item-uuid", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}

	err = client.DeleteItem(item, "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Deleting an item by UUID

```go
	err := client.DeleteItemByID("item-uuid", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Retrieving a file from an item

```go
	file, err := client.GetFile("<file-uuid>", "item-uuid", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Retrieving the contents of a file from an item

```go
	file, err := client.GetFile("file-uuid", "item-uuid", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}

	content, err := client.GetFileContent(file)
	if err != nil {
		log.Fatal(err)
	}
```

#### Retrieving all files under an item

```go
	files, err := client.GetFiles("item-uuid", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}
```

#### Downloading a file

```go
	file, err := client.GetFile("file-uuid", "item-uuid", "vault-uuid")
	if err != nil {
		log.Fatal(err)
	}

	path, err := client.DownloadFile(file, "local/path/to/file", true)
	if err != nil {
		log.Fatal(err)
	}
```

### Unmarshalling into a Struct

Users can define tags on a struct and have the `connect.Client` unmarshall item data directly in them. Supported field tags are:

- `opvault` – The UUID of the vault the item should come from
- `opitem` – The title of the Item
- `opsection` - The section where the required field is located
- `opfield` – The item field whose value should be retrieved

All retrieved fields require at least the `opfield` and `opitem` tags, while all retrieved items require the `opitem` tag. Additionally, a custom vault can be specified by setting the `opvault` tag.
In case this is not set, the SDK will use the value of the `OP_VAULT` environment variable as the default UUID.
If a field is within a section, the `opsection` tag is required as well. Please note that one cannot retrieve a section in itself.

#### Example Struct

This example struct will retrieve 3 fields from one item and a whole item from another vault:

```go
package main

import (
	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
)

type Config struct {
	Username string           `opitem:"Demo TF Database" opfield:"username"`
	Password string           `opitem:"Demo TF Database" opfield:"password"`
    Host     string           `opitem:"Demo TF Database" opsection:"details" opfield:"hostname"`
	APIKey   onepassword.Item `opvault:"7vs66j55o6md5btwcph272mva4" opitem:"API Key"`
}

func main() {
	client, err := connect.NewClientFromEnvironment()
	if err != nil {
		panic(err)
	}
    	c := Config{}
	err = client.LoadStruct(&c)
}

```

Additionally, fields of the same item can be added to a struct at once, without needing to specify the `opitem` or `opvault` tags:

```go
package main

import "github.com/1Password/connect-sdk-go/connect"

type Config struct {
	Username string     `opfield:"username"`
	Password string     `opfield:"password"`
}

func main () {
	client, err := connect.NewClientFromEnvironment()
    	if err != nil {
		panic(err)
	}
	c := Config{}

    // retrieve using item title
	err = client.LoadStructFromItemByTitle(&c, "Demo TF Database", "7vs66j55o6md5btwcph272mva4")

    // retrieve using item uuid
    err = client.LoadStructFromItem(&c, "4bc73kao58g2usb582ndn3w4", "7vs66j55o6md5btwcph272mva4")
}
```

### Environment Variables

The Connect Go SDK makes use of the following environment variables:

- `OP_CONNECT_TOKEN`: the API token to be used to authenticate the client to your 1Password Connect instance. Used in order to authenticate via the `connect.NewClientFromEnvironment` function.
- `OP_CONNECT_HOST`: the hostname of your 1Password Connect instance. Used in order to authenticate via the `connect.NewClientFromEnvironment` function.
- `OP_VAULT`: a vault UUID. Used as default vault in the `LoadStruct`, `LoadStructFromItemByTitle` and `LoadStructFromItem` functions, for all fields where the `opvault` tag is not set.

### Errors

All errors returned by Connect API are unmarshalled into a `onepassword.Error` struct:

```go
type Error struct {
    StatusCode int    `json:"status"`
    Message    string `json:"message"`
}
```

Details of the errors can be accessed by using `errors.As()`:

```go
_, err := client.GetVaults()
if err != nil{
    var opErr *onepassword.Error
    if errors.As(err, &opErr){
        fmt.Printf("message=%s, status code=%d\n",
            opErr.Message,
            opErr.StatusCode,
        )
    }
}
```
