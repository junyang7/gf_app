package api

type Conf struct {
	Ip     string   `json:"ip"`
	Port   string   `json:"port"`
	Origin []string `json:"origin"`
}
