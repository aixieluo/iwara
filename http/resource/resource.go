package resource

type Res struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func (res *Res) SetData(data interface{}) *Res {
	res.Data = data
	return res
}

func (res *Res) SetMeta(meta interface{}) *Res {
	res.Meta = meta
	return res
}

func Factory(data interface{}) *Res {
	return &Res{
		Data: data,
		Meta: nil,
	}
}
