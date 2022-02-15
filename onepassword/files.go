package onepassword

import (
	"encoding/json"
)

type File struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Section     *ItemSection `json:"section,omitempty"`
	Size        int          `json:"size"`
	ContentPath string       `json:"content_path"`
	Content     []byte       `json:"content,omitempty"`
}

func (f *File) UnmarshalJSON(data []byte) error {
	var jsonFile struct {
		ID          string       `json:"id"`
		Name        string       `json:"name"`
		Section     *ItemSection `json:"section,omitempty"`
		Size        int          `json:"size"`
		ContentPath string       `json:"content_path"`
		Content     []byte       `json:"content,omitempty"`
	}

	if err := json.Unmarshal(data, &jsonFile); err != nil {
		return err
	}
	f.ID = jsonFile.ID
	f.Name = jsonFile.Name
	f.Section = jsonFile.Section
	f.Size = jsonFile.Size
	f.ContentPath = jsonFile.ContentPath
	f.Content = jsonFile.Content
	return nil
}

// IsFetched returns true if the content of the file has been loaded and false if not.
// Use `client.GetFileContent(file *File)` to make sure the content is fetched automatically if not present.
func (f *File) IsFetched() bool {
	if f.Content == nil {
		return false
	}
	return true
}
