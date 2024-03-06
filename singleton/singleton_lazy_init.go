package singleton

import (
	"sync"
)

type LazySingleton struct{}

var lazyInstance *LazySingleton
var once sync.Once

func GetLazyInstance() *LazySingleton {
	once.Do(func() {
		lazyInstance = &LazySingleton{}
	})
	return lazyInstance
}
