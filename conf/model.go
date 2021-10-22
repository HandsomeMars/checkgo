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
	Index    int        `json:"_"`
	Row      int        `json:"_"`
	Col      int        `json:"_"`
}

type Comment struct {
	Name      string      `json:"name"`
	Code      string      `json:"code"`
	Validator string      `json:"validator"`
	Format    string      `json:"format"`
	Layout    layoutType  `json:"layout"`
	Type      commentType `json:"type"`
	Filed     []*Filed    `json:"fileds"`
	Index     int         `json:"_"`
	Row       int         `json:"_"`
	Col       int         `json:"_"`
}

type Filed struct {
	Name      string      `json:"name"`
	Code      string      `json:"code"`
	Validator string      `json:"validator"`
	Format    string      `json:"format"`
	Type      fieldType   `json:"type"`
	Example   interface{} `json:"example"`
	Index     int         `json:"_"`
	Row       int         `json:"_"`
	Col       int         `json:"_"`
}

func (c *Conf) check() error {
	return nil
}
