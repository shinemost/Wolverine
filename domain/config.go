package domain

type Config struct {
	Age     string `json:"age"`
	Name    string `json:"name"`
	Address struct {
		Location string `json:"location"`
	} `json:"address"`
}
