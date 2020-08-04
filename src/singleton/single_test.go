package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)


type Singleton struct {

}

var singleInstance *Singleton
var once sync.Once

//只执行一次
func GetSingletonObj() *Singleton  {
	if singleInstance != nil {
		return singleInstance
	}
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup //携程安全
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	//等待是由协程运行完成
	wg.Wait()
}
