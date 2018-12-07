package model

type ResultBean struct {
	Status bool
	Count	int
	Msg    string
	Data   interface{}
}

func CreateResultWithData(data interface{}) ResultBean {
	return ResultBean{true,0, "SUCCCESS", data}
}

func CreateResultWithCountAndData(count int,data interface{}) ResultBean {
	return ResultBean{true,count, "SUCCCESS", data}
}

func CreateResultWithMsg(msg string) ResultBean {
	return ResultBean{false,0, msg, nil}
}

