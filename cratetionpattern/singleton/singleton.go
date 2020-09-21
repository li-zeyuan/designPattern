package singleton

import "sync"

// 单例类
type Singleton struct {

}

var (
	singleton *Singleton
	once sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
