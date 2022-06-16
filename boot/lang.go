package boot

var langText map[string]string

type LangText struct{}

// 中文
func (LangText) Zh() {

	langText = map[string]string{
		"url not found": "地址未找到",
	}
}

// 英文
func (LangText) En() {

	langText = map[string]string{
		// 无需定义内容
	}
}
