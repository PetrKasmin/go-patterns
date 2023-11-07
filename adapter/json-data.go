package adapter

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.jsonDocument.ConvertToXml()
	println("Отправка xml документа!")
}
