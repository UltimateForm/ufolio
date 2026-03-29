package githubapi

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	HTMLURL     string `json:"html_url"`
	Fork        bool   `json:"fork"`
	PushedAt    string `json:"pushed_at"`
}
