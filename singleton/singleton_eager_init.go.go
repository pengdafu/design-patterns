package singleton

type EagerSingleton struct{}

var eagerInstance = &EagerSingleton{}

//var eagerInstance *EagerSingleton
//func init() {
//	eagerInstance = &EagerSingleton{}
//}

func GetEagerInstance() *EagerSingleton {
	return eagerInstance
}
