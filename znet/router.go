package znet

import "mzinx/ziface"

type BaseRouter struct {}

//处理conn之前的方法hook
func (r *BaseRouter) PreHandle(request ziface.IRequest)  {}

//处理conn方法
func (r *BaseRouter) Handle(request ziface.IRequest){}

//处理conn之后的方法hook
func (r *BaseRouter) PostHandle(request ziface.IRequest){}
