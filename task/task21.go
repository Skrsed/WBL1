package task

// Реализовать паттерн «адаптер» на любом примере.

type JsonParser[T any] struct{}

type XMLParser[T any] struct{}

type Document[T any] struct{}

type XMLJSONAdapter[T any] struct {
	xmp XMLParser[T]
}

type IJsonParser[T any] interface {
	ParseJson(doc string, model *T)
}

type Pasport struct{}

func (dc Document[T]) ParseDocument(jp IJsonParser[T], model *T) {
	jp.ParseJson(dc.toString(), model)
}

func (dc Document[T]) toString() string {
	return ""
}

func (jp JsonParser[T]) ParseJson(doc string, model *T) {}

func (xp XMLParser[T]) ParseXML(doc string, model *T) {}

func (xjp XMLJSONAdapter[T]) ParseJson(doc string, model *T) {
	xjp.xmp.ParseXML(doc, model)
}

func Adapter() {
	doc := Document[Pasport]{}
	jsonParser := JsonParser[Pasport]{}
	xmlParser := XMLParser[Pasport]{}
	adapter := XMLJSONAdapter[Pasport]{
		xmp: xmlParser,
	}

	doc.ParseDocument(jsonParser, &Pasport{})
	doc.ParseDocument(adapter, &Pasport{})
}
