package util

type SiteConfig struct {
	SiteName string			`json:"site_name"`
	SiteURL string			`json:"site_url"`
	ParseableTags []string	`json:"tags"`
}
