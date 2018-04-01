package entity

/*
	Configuration with parse rules for sites
	index used for iterations
*/
type ParseConfig struct {
	Sites []SiteConfig `json:"sites"`
	index int
}

const EMPTY_SITE_NAME = ""

func (config *ParseConfig) GetSitesConfigs() ([]SiteConfig) {
	return config.Sites
}

func (config *ParseConfig) GetNextSiteURL() (string, bool) {
	siteURL := config.Sites[config.index].SiteURL
	var beRestarted = false
	config.index++
	if config.index > (len(config.Sites) - 1) {
		config.RestartCounting()
		beRestarted = true
	}
	return siteURL, beRestarted
}

func (config *ParseConfig) GetCurrentSiteConfiguration() (SiteConfig) {
	return config.Sites[config.index]
}

func (config *ParseConfig) RestartCounting() {
	config.index = 0;
}

func (config *ParseConfig) GetSiteByName(name string) (string) {
	for _, site := range config.Sites {
		if site.SiteURL == name {
			return site.SiteURL;
		}
	}
	return EMPTY_SITE_NAME
}
