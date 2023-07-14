<!-- Image sourced from https://blog.1password.com/introducing-secrets-automation/ -->
<img alt="" role="img" src="https://blog.1password.com/posts/2021/secrets-automation-launch/header.svg"/>

<div align="center">
  <h1>1Password Connect SDK for Go</h1>
  <p>Access your 1Password items in your Go applications through your self-hosted <a href="https://developer.1password.com/docs/connect">1Password Connect server</a>.</p>
  <a href="#-quickstart">
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

2. Export the `OP_CONNECT_HOST` and `OP_CONNECT_TOKEN` environment variables:

   ```sh
   export OP_CONNECT_HOST=<your-connect-host> && \
       export OP_CONNECT_TOKEN=<your-connect-token>
   ```

3. Use it in your code

   - Read a secret:

     ```go
     import "github.com/1Password/connect-sdk-go/connect"

     func main () {
         client := connect.NewClientFromEnvironment()
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
         client := connect.NewClientFromEnvironment()
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
