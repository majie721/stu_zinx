package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mzinx/ziface"
)

type GlobalObj struct {
	TCPServer ziface.IServer //全局Server 对象

	Host string //监听的主机dIP

	TCPPort int //服务器监听的端口

	Name string //当前服务器的名称

	/**
	Zinx
	*/
	Version string //zinx 版本号

	MaxConn int //zinx 最大连接数

	MaxPackageSize uint32 //zinx数据包得最大值

}

func (g *GlobalObj) Reload() {
	path := "C:\\Users\\LENOVO\\go\\src\\mzinx\\conf\\zinx.json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("reload config file [%s]error,error:%#v", path, err)
		return
	}

	if err = json.Unmarshal(data, g); err != nil {
		panic(err)
	}

	fmt.Printf(" json:%#v \n", g)
}

var GlobalObject *GlobalObj

//提供一个init方法,初始化当前GlobalObject的对象
func init() {
	GlobalObject = &GlobalObj{
		Name:           "zinxApp",
		Version:        "0.3",
		TCPPort:        8888,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	fmt.Println("INIT ")
	GlobalObject.Reload()
}
