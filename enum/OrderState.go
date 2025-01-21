package enum

import "log"

type OrderState interface {
	SetState(state string)
	GetState() string
}

const (
	Successful = "Successful"
	Failed     = "Failed"
)

type orderstate struct {
	value string
}

func validateState(state string) bool {
	return state == Successful || state == Failed
}

func NewOrderState(state string) *orderstate {
	if !validateState(state) {
		log.Panicf("Invalid order state: %s", state)
	}

	return &orderstate{value: state}
}

func (o *orderstate) SetState(state string) {
	if !validateState(state) {
		log.Panicf("Invalid order state: %s", state)
	}

	o.value = state
}

func (o *orderstate) GetState() string {
	return o.value
}
