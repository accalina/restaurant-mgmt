package common

type DataMessageValue struct {
	Message string
}

type DataArrayValue struct {
	ArrMessage []string
}

func (d DataMessageValue) GetDataMessage() string {
	return d.Message
}

func (d DataArrayValue) GetDataArray() []string {
	return d.ArrMessage
}
