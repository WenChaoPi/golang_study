package main

import (
	"fmt"
	"time"
	
	"errors"
)
//定义可重用对象
type ReusableObj struct{
}
//定义对象池
type ObjPool struct{
	bufChan chan *ReusableObj	//用于缓存可重用对象
}
//创建对象池
func NewObjPool(numOfObj int) *ObjPool{
	objPool := ObjPool{}
	// objPool := new(ObjPool)
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i:=0; i<numOfObj; i++{
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}
//获取对象
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error){
	select{
	case ret := <- p.bufChan:
		return ret, nil
	case <- time.After(timeout):
		return nil, errors.New("time out")
	}
}
//释放对象
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error{
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
//
func TestObjPool(){
	pool := NewObjPool(10)
	for i:=0; i<11; i++{
		if v, err := pool.GetObj(time.Second * 1); err != nil{
			fmt.Println(err)
		} else {
			//使用获得的对象 v 处理任务， 处理完后释放对象到对象池
			fmt.Printf("%T, %v\n", v, v)
			if err = pool.ReleaseObj(v); err != nil{
				fmt.Println(err)
			}
		}
	}
	fmt.Println("Done.")
}

func main()  {
	TestObjPool()
}