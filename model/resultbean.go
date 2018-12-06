package model

type ResultBean struct {
	Status bool
	Msg    string
	Data   interface{}
}

func CreateResultWithData(data interface{}) ResultBean {
	return ResultBean{true, "SUCCCESS", data}
}

func CreateResultWithMsg(msg string) ResultBean {
	return ResultBean{false, msg, nil}
}

