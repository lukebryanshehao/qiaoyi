package model

type ResultBean struct {
	Status bool
	Msg    string
	Code    string
	Data   interface{}
	Token  string
}
type ResultPage struct {
	Status    bool
	Msg       string
	PageCount int
	Data      interface{}
}

/*
d为通用返回数据，e为可选参数，e不为空时返回失败result
 */
func NewResultBean(d interface{}, e ... string) *ResultBean {
	if len(e) > 0 {
		return &ResultBean{false, e[0],"200", nil, ""}
	}
	return &ResultBean{true, "success", "200",d, ""}
}

/*
d为通用返回数据，e为可选参数，e不为空时返回失败result
 */
func NewResultPage(d interface{}, pageCount int, e ... string) *ResultPage {
	if len(e) > 0 {
		return &ResultPage{false, e[0], 0, ""}
	}
	return &ResultPage{true, "success", pageCount, d}
}
