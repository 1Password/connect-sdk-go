package connect

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/1Password/connect-sdk-go/onepassword"
)

// Tests in this file exercise the section/URL name lookup helpers in
// config_helper.go. The lookups must be case-insensitive: a struct tag like
// `opsection:"Details"` should match a section whose label is "Details",
// "details", or any other casing. Prior to the fix, comparisons used
// strings.ToLower on one side only, which silently dropped fields whenever
// the tag's case did not match the lowercased label.

var testSections = []*onepassword.ItemSection{
	{ID: "section-id-1", Label: "Details"},
	{ID: "section-id-2", Label: "API Credentials"},
	{ID: "section-id-3", Label: "lowercase"},
}

var testURLs = []onepassword.ItemURL{
	{Label: "Primary", URL: "https://example.com", Primary: true},
	{Label: "Backup", URL: "https://backup.example.com", Primary: false},
}

func TestSectionIDForName(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		want string
	}{
		{"exact case", "Details", "section-id-1"},
		{"lowercase tag against mixed-case label", "details", "section-id-1"},
		{"uppercase tag against mixed-case label", "DETAILS", "section-id-1"},
		{"mixed-case tag against mixed-case label", "DeTaIlS", "section-id-1"},
		{"label with space, lowercase tag", "api credentials", "section-id-2"},
		{"label with space, exact case", "API Credentials", "section-id-2"},
		{"lowercase label, lowercase tag", "lowercase", "section-id-3"},
		{"lowercase label, uppercase tag", "LOWERCASE", "section-id-3"},
		{"no match", "nonexistent", ""},
		{"empty tag", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sectionIDForName(tt.tag, testSections))
		})
	}
}

func TestSectionIDForName_NilSections(t *testing.T) {
	assert.Equal(t, "", sectionIDForName("Details", nil))
}

func TestSectionIDForName_EmptySections(t *testing.T) {
	assert.Equal(t, "", sectionIDForName("Details", []*onepassword.ItemSection{}))
}

func TestSectionLabelForName(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		want string
	}{
		{"exact case returns stored label", "Details", "Details"},
		{"lowercase tag returns stored mixed-case label", "details", "Details"},
		{"uppercase tag returns stored mixed-case label", "DETAILS", "Details"},
		{"mixed-case tag returns stored mixed-case label", "DeTaIlS", "Details"},
		{"label with space, lowercase tag", "api credentials", "API Credentials"},
		{"no match", "nonexistent", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sectionLabelForName(tt.tag, testSections))
		})
	}
}

func TestSectionLabelForName_NilSections(t *testing.T) {
	assert.Equal(t, "", sectionLabelForName("Details", nil))
}

func TestURLPrimaryForName(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		want bool
	}{
		{"exact case primary URL", "Primary", true},
		{"lowercase tag against mixed-case label, primary", "primary", true},
		{"uppercase tag against mixed-case label, primary", "PRIMARY", true},
		{"exact case non-primary URL", "Backup", false},
		{"lowercase tag against mixed-case label, non-primary", "backup", false},
		{"no match returns false", "nonexistent", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, urlPrimaryForName(tt.tag, testURLs))
		})
	}
}

func TestURLPrimaryForName_NilURLs(t *testing.T) {
	assert.Equal(t, false, urlPrimaryForName("Primary", nil))
}

func TestURLLabelForName(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		want string
	}{
		{"exact case", "Primary", "Primary"},
		{"lowercase tag returns stored mixed-case label", "primary", "Primary"},
		{"uppercase tag returns stored mixed-case label", "PRIMARY", "Primary"},
		{"mixed-case tag returns stored mixed-case label", "PrImArY", "Primary"},
		{"no match", "nonexistent", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, urlLabelForName(tt.tag, testURLs))
		})
	}
}

func TestURLLabelForName_NilURLs(t *testing.T) {
	assert.Equal(t, "", urlLabelForName("Primary", nil))
}

func TestURLURLForName(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		want string
	}{
		{"exact case", "Primary", "https://example.com"},
		{"lowercase tag against mixed-case label", "primary", "https://example.com"},
		{"uppercase tag against mixed-case label", "PRIMARY", "https://example.com"},
		{"non-primary URL, lowercase tag", "backup", "https://backup.example.com"},
		{"no match", "nonexistent", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, urlURLForName(tt.tag, testURLs))
		})
	}
}

func TestURLURLForName_NilURLs(t *testing.T) {
	assert.Equal(t, "", urlURLForName("Primary", nil))
}
