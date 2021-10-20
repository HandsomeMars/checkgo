package main

type JsonConf struct {
	Pages   []*Page `json:"pages"`
	Version string   `json:"version"`
}

type Page struct {
	Key      string     `json:"key"`
	Comments []*Comment `json:"comments"`
}

type Comment struct {
	Key   string `json:"key"`
	Index int    `json:"index"`
	Type  string `json:"type"`
	Api   *Api   `json:"api"`
}

type Api struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ApiType string `json:"api_type"`
	Params  *Model `json:"params"`
	Result  *Model `json:"model"`
}

type Model struct {
	Checker []string `json:"checker"`
	Fields  []*Field `json:"fields"`
}

type Field struct {
	Name       string   `json:"name"`
	FieldType  string      `json:"field_type"`
	Value      interface{} `json:"value"`
	IsKey      bool        `json:"is_key"`
	IsRequired bool     `json:"is_required"`
	IsShow     bool     `json:"is_show"`
	Checker    []string `json:"checker"`
}

//func main() {
//	jsonconf := JsonConf{
//		Pages:   []*Page{&Page{
//			Key:      "list_page",
//			Comments: []*Comment{&Comment{
//				Key:   "list",
//				Index: 0,
//				Type:  "table",
//				Api:   &Api{
//					Url:     "http://127.0.0.1:8090/hello",
//					Method:  "post",
//					ApiType: "query",
//					Params:  &Model{
//						Checker: nil,
//						Fields:  []*Field{
//							&Field{
//								Name:       "name",
//								IsKey:      false,
//								FieldType:  "string",
//								IsRequired: false,
//								IsShow:     true,
//								Checker:    nil,
//							},
//							&Field{
//								Name:       "id",
//								IsKey:      false,
//								FieldType:  "int",
//								IsRequired: false,
//								IsShow:     true,
//								Checker:    nil,
//							},
//						},
//					},
//					Result:  nil,
//				},
//			}},
//		}},
//		Version: "v1.0.0",
//	}
//	data,err:=json.Marshal(jsonconf);
//	if err!=nil{
//		log.Fatal(err)
//	}
//	log.Print(string(data))
//}