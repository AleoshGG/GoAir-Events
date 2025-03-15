package infrastructure

type FCM struct {
}

func NewFCM() *FCM {
	return &FCM{}
}

func (e *FCM) StartVentilation()    {}
func (e *FCM) StopVentilation()     {}
func (e *FCM) DurationVentilation() {}