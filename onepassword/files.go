package onepassword

type File struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Size        int      `json:"size"`
	Content     string `json:"content"`
	ContentPath string   `json:"content_path"`
}
