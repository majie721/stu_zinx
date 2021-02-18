package znet

import "errors"

type Message struct {
	ID uint32

	DataLen uint32

	Data []byte
}


func (m *Message)GetMsgId() uint32{
	return  m.ID
}

func (m *Message)GetMsgLen() uint32{
	return  m.DataLen
}

func (m *Message)GetData() []byte{
	return  m.Data
}

func(m *Message)SetMsgId(id uint32){
	m.ID = id
}

func(m *Message)SetMsgLen(len uint32 ){
	m.DataLen = len
}

func(m *Message)SetData(data []byte) error{
	m.Data = data

}