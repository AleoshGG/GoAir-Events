package infrastructure

var FMC *FCM

func GoDependences() {
	FMC = NewFCM()
}

func GetFCM() *FCM {
	return FMC
}