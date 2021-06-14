package commonType

type Error map[string]string

type Response struct {
	Res    interface{} `json:"res"`
	Status string      `json:"status"`
	Err    Error       `json:"error"`
}

func (e Error) Error() string {
	return e["message"]
}
