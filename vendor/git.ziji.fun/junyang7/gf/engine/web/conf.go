package web

type Conf struct {
	Ip     string   `json:"ip"`
	Port   string   `json:"port"`
	Root   string   `json:"root"`
	Origin []string `json:"origin"`
}
