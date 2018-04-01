package entity

/*
	Type, which described configuration to concrete site
*/
type SiteConfig struct {
	SiteName 		string                  `json:"site_name"`
	SiteURL  		string                  `json:"site_url"`
	PagesToParsing 	int						`json:"page_to_parsing"`
	Tags     		map[string]ParseableTag `json:"tags"`
}
