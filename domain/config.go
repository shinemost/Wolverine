package domain

type Config struct {
	Age     string `json:"age"`
	Name    string `json:"name"`
	Address struct {
		Location string `json:"location"`
	} `json:"address"`
}

type Config2 struct {
	Age     int
	Name    string
	Address `mapstructure:"address"`
}

type Address struct {
	Location string
}
