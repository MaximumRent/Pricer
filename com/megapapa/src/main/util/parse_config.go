package util

type ParseConfig struct {
	startPages []string
	index int
}

const EMPTY_SITE_NAME = ""

func (config *ParseConfig) GetStartPages() ([]string) {
	return config.startPages
}

func (config *ParseConfig) GetNextSite() (string) {
	site := config.startPages[config.index]
	config.index++
	if (config.index > (len(config.startPages) - 1)) {
		config.index = 0
	}
	return site
}

func (config *ParseConfig) GetSiteByName(name string) (string) {
	for _, page := range config.startPages {
		if (page == name) {
			return page;
		}
	}
	return EMPTY_SITE_NAME
}
