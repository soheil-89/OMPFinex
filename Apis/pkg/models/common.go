package models

type Maker interface {
	handler(params interface{})
}

func Make(m Maker, params interface{}) {
	m.handler(params)
}
