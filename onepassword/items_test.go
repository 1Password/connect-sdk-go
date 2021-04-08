package onepassword

import (
	"testing"

	"github.com/google/uuid"
)

var testUsernmae = "user123"
var testPassword = "Password2!"
var testOther = "Test Value"

func TestItemGetValuePassword(t *testing.T) {
	login := testLogin()

	value := login.GetValue("password")
	if value != testPassword {
		t.Logf("Expected password %q, found %q", testPassword, value)
		t.FailNow()
	}

	value = login.GetValue(".password")
	if value != testPassword {
		t.Logf("Expected password %q, found %q", testPassword, value)
		t.FailNow()
	}
}

func TestItemGetValueOther(t *testing.T) {
	login := testLogin()

	value := login.GetValue("other")
	if value != testOther {
		t.Logf("Expected other value %q, found %q", testOther, value)
		t.FailNow()
	}

	value = login.GetValue("example.other")
	if value != testOther {
		t.Logf("Expected other value %q, found %q", testOther, value)
		t.FailNow()
	}
}

func TestItemGetValueMissingField(t *testing.T) {
	login := testLogin()

	value := login.GetValue("missing")
	if value != "" {
		t.Logf("Expected other value %q, found %q", "", value)
		t.FailNow()
	}
}

func testLogin() *Item {
	sectionUUID := uuid.New().String()
	return &Item{
		ID:    uuid.New().String(),
		Title: "Example Login",
		URLs: []ItemURL{
			{
				URL:     "example.com",
				Primary: true,
			},
		},
		Category: Login,
		Sections: []*ItemSection{
			{
				ID:    sectionUUID,
				Label: "example",
			},
		},
		Fields: []*ItemField{
			{
				ID:      "username",
				Type:    "STRING",
				Purpose: "USERNAME",
				Label:   "username",
				Value:   testUsernmae,
			},
			{
				ID:      "password",
				Type:    "CONCEALED",
				Purpose: "PASSWORD",
				Label:   "password",
				Value:   testPassword,
			},
			{
				ID: uuid.New().String(),
				Section: &ItemSection{
					ID: sectionUUID,
				},
				Type:  "STRING",
				Label: "other",
				Value: testOther,
			},
		},
	}
}
