<!-- Image sourced from https://blog.1password.com/introducing-secrets-automation/ -->
<img alt="" role="img" src="https://blog.1password.com/posts/2021/secrets-automation-launch/header.svg"/>

<div align="center">
  <h1>1Password Connect SDK for Go</h1>
  <p>Access your 1Password items in your Go applications through your self-hosted <a href="https://developer.1password.com/docs/connect">1Password Connect server</a>.</p>
  <a href="#✨-quickstart">
    <img alt="Get started" src="https://user-images.githubusercontent.com/45081667/226940040-16d3684b-60f4-4d95-adb2-5757a8f1bc15.png" height="37"/>
  </a>
</div>

---

The 1Password Connect Go SDK provides access to the [1Password Connect](https://developer.1password.com/docs/connect) API, to facilitate communication with the Connect server hosted on your infrastructure and 1Password. The library is intended to be used by your applications, pipelines, and other automations to simplify accessing items stored in your 1Password vaults.


## ✨ Quickstart

1. Download and install the 1Password Connect Go SDK:
   ```sh
   go get github.com/1Password/connect-sdk-go
   ```

2. Use it in your code
   - Read a secret:
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

   - Write a secret:
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

For more examples, check out [USAGE.md](USAGE.md). 

## 💙 Community & Support

- File an [issue](https://github.com/1Password/connect-sdk-go/issues) for bugs and feature requests.
- Join the [Developer Slack workspace](https://join.slack.com/t/1password-devs/shared_invite/zt-1halo11ps-6o9pEv96xZ3LtX_VE0fJQA).
- Subscribe to the [Developer Newsletter](https://1password.com/dev-subscribe/).

## 🔐 Security

1Password requests you practice responsible disclosure if you discover a vulnerability.

Please file requests via [**BugCrowd**](https://bugcrowd.com/agilebits).

For information about security practices, please visit the [1Password Bug Bounty Program](https://bugcrowd.com/agilebits).


<!-- # 1Password Connect Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/1Password/connect-sdk-go.svg)](https://pkg.go.dev/github.com/1Password/connect-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/1Password/connect-sdk-go)](https://goreportcard.com/report/github.com/1Password/connect-sdk-go)
[![Version](https://img.shields.io/github/release/1Password/connect-sdk-go.svg)](https://github.com/1Password/connect-sdk-go/releases/)

The 1Password Connect Go SDK provides access to the [1Password Connect](https://support.1password.com/secrets-automation/) API, to facilitate communication with the Connect server hosted on your infrastructure and 1Password. The library is intended to be used by your applications, pipelines, and other automations to simplify accessing items stored in your 1Password vaults.

<details>
  <summary>Table of Contents</summary>

  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
  * [Usage](#usage)
    + [Quickstart](#quickstart)
    + [Creating an API Client](#creating-an-api-client)
    + [Model Objects](#model-objects)
    + [Item CRUD](#item-crud)
      - [Retrieving list of vaults that the Connect token has permission to read](#retrieving-list-of-vaults-that-the-connect-token-has-permission-to-read)
      - [Retrieving all items in a vault](#retrieving-all-items-in-a-vault)
      - [Retrieving item by title](#retrieving-item-by-title)
      - [Retrieving items by vault and item UUID](#retrieving-items-by-vault-and-item-uuid)
      - [Creating items in a vault](#creating-items-in-a-vault)
      - [Update and Item](#update-and-item)
      - [Delete an item](#delete-an-item)
      - [Retrieving a file from an item](#retrieving-a-file-from-an-item)
      - [Retrieving the contents of a file from an item](#retrieving-the-contents-of-a-file-from-an-item)
    + [Unmarshalling into a Struct](#unmarshalling-into-a-struct)
      - [Example Struct](#example-struct)
    + [Environment Variables](#environment-variables)
    + [Errors](#errors)
  * [Development](#development)
    + [Building](#building)
    + [Running Tests](#running-tests)
  * [Security](#security)
</details>

## Prerequisites

- [1Password Connect](https://support.1password.com/secrets-automation/#step-2-deploy-a-1password-connect-server) deployed in your infrastructure

## Installation and Importing
To download and install the 1Password Connect Go SDK, as well as its dependencies:
```sh
go get github.com/1Password/connect-sdk-go
```

To import the 1Password Connect SDK in your Go project:
```go
import (
    "github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
)
```

## Development

### Building

To build all packages:

```sh
make build
```

### Running Tests

Run all tests:

```sh
make test
```

Run all tests with code coverage:

```sh
make test/coverage
```

## Security

1Password requests you practice responsible disclosure if you discover a vulnerability.

Please file requests via [**BugCrowd**](https://bugcrowd.com/agilebits).

For information about security practices, please visit our [Security homepage](https://bugcrowd.com/agilebits).

-->
