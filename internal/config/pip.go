package config

type Pip struct {
	IndexURL     string `mapstructure:"index-url" json:"indexURL"`
	SearchURL    string `mapstructure:"search-url" json:"searchURL"`
	ListPageSize int    `mapstructure:"list-page-size" json:"listPageSize"`
	Timeout      int    `mapstructure:"timeout" json:"timeout"`
}
