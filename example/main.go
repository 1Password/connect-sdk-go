package main

import (
	"bufio"
	"fmt"
	"github.com/1Password/connect-sdk-go/connect"
	"github.com/1Password/connect-sdk-go/onepassword"
	"os"
	"time"
)

func main() {
	vault := os.Getenv("OP_VAULT")
	secret := os.Getenv("SECRET_STRING")

	fmt.Print(steps["intro"])

	client, err := connect.NewClientFromEnvironment()

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(steps["step1"])

	item := &onepassword.Item{
		Fields: []*onepassword.ItemField{{
			Value: secret,
			Type: "STRING",
		}},
		Tags:     []string{"1password-connect"},
		Category: onepassword.Login,
		Title:    "Secret String",
	}

	fmt.Print(steps["step2"])

	postedItem, err := client.CreateItem(item, vault)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(steps["step3"])
	time.Sleep(10 * time.Second)

	retrievedItem, err := client.GetItem(postedItem.ID, vault)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(steps["step4"])
	fmt.Print(steps["confirmation"])

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	for (char != 'y') && (char != 'n') {
		fmt.Print(steps["confirmation2"])
		char, _, err = reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}
	}


	if char == 'y' {
		_ = client.DeleteItem(retrievedItem, vault)
		fmt.Print(steps["step5"])
	}

	fmt.Print(steps["outro"])

}


var steps = map[string]string{
	"intro":         "\nHello from 1Password! In order to exemplify the end-to-end process of creating, posting, retrieving and, eventually, deleting an item, the following steps are taken: \n",
	"step1":         "1. The SDK has contacted the Connect Server, and a client has been created, based on the provided OP_CONNECT_TOKEN.\n",
	"step2":         "2. An item containing the secret string has been successfully created.\n",
	"step3":         "3. The item containing the secret string has been successfully added in the default vault.\n",
	"step4":         "4. The item containing the secret string has been successfully retrieved from the default vault.\n",
	"confirmation":  "Would you like to delete the newly created item from your vault? (y/n)",
	"confirmation2": "Your answer should be either 'y' or 'n'. Would you like to delete the newly created item from your vault? (y/n)",
	"step5":         "\n5. The item containing the secret string has been successfully deleted from the default vault.\n",
	"outro":         "All done!\n",
}