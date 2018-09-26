package di

import (
	"errors"
	"sync"
	"strings"
)

var (
	ErrFactoryNotFound = errors.New("factory not found")
)

type factory = func() (interface{}, error)

type Container struct {
	sync.Mutex
	singletons map[string]interface{}
	factories map[string]factory
}
/*初始化工厂*/
func initContainer()*Container{
	return &Container{
		singletons:make(map[string]interface{}),
		factories:make(map[string]factory),
	}
}
/*reg单例*/
func (p *Container)setSingletons(name string,singleton interface{})  {
	p.Lock()
	p.singletons[name]=singleton
	p.Unlock()
}
/*get单例*/
func (p *Container)getSingletons(name string)(interface{},error){
	factory,ok:=p.factories[name]
	if !ok{
		return nil,ErrFactoryNotFound
	}
	return factory,nil

}

/*check 单例*/
func (p *Container)checSingleton(tag string)bool{
	tags:=strings.Split(tag,",")
	for _,name:= range tags{
		if name=="prototype"{
			return false
		}
	}
	return true
}

/*reg Prototype*/
func (p *Container)setPrototype(name string,factory factory)  {
	p.Lock()
	p.factories[name]=factory
	p.Unlock()
}
/*get Prototype*/
func (p *Container)getPrototype(name string)interface{}{
	return p.singletons[name]
}

/*check Prototype*/
func (p *Container)checkPrototype(tag string)bool{
	tags:=strings.Split(tag,",")
	for _,name:= range tags{
		if name=="prototype"{
			return false
		}
	}
	return true
}

// 获取需要注入的依赖名称
func (p *Container) injectName(tag string) string {
	tags := strings.Split(tag, ",")
	if len(tags) == 0 {
		return ""
	}
	return tags[0]
}

func (p *Container) Inject(instance interface{}) error {
	return  nil
}