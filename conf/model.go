package conf

type Info struct {
	Description string `json:"description"`
	Version     string `json:"version"`
	Title       string `json:"title"`
}

type Conf struct {
	Checkgo    string               `json:"checkgo"`
	Version    string               `json:"version"`
	Info       *Info                `json:"info"`
	Host       string               `json:"host"`
	Activities map[string]*Activity `json:"activities"`
}

type Activity struct {
	Code     string     `json:"code"`
	Name     string     `json:"name"`
	Version  string     `json:"version"`
	Comments []*Comment `json:"comments"`
}

type Comment struct {
	Name      string      `json:"name"`
	Code      string      `json:"code"`
	Validator string      `json:"validator"`
	Format    string      `json:"format"`
	Type      commentType `json:"type"`
	Comments  []*Comment  `json:"comments"`
	Example   interface{} `json:"example"`
}

func (c *Conf) check() error {
	return nil
}
