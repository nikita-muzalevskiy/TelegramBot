package main

import (
	pb "NovelService/protos"
)

var messageList = map[string]message{
	"menu": {
		"Новая новелла или продолжить чтение?",
		[]button{},
		func(user string, novel string) ([]button, string) {
			return []button{
				{"Новая новелла", user + "/NOVEL/new-list/"},
				{"Начатая новелла", user + "/NOVEL/started-list/"},
			}, "Новая новелла или продолжить чтение?"
		},
	},
	"started-list": {
		"",
		[]button{},
		func(user string, novel string) ([]button, string) {
			return GetStartedNovel(user)
		},
	},
	"new-list": {
		"",
		[]button{},
		func(user string, novel string) ([]button, string) {
			return GetNewNovel(user)
		},
	},
	"read": {
		"",
		[]button{},
		func(user string, novel string) ([]button, string) {
			return Read(user, novel)
		},
	},
}

type message struct {
	text     string
	keyboard []button
	f        func(user string, novel string) ([]button, string)
}

type button struct {
	text     string
	callback string
}

//type getButtons func(user string) []button

func MessageToCallback(user string, param string, mes message) *pb.CallbackReply {
	mes.keyboard, mes.text = mes.f(user, param)
	var bs = make([]*pb.Button, 0)
	for _, b := range mes.keyboard {
		bs = append(bs, &pb.Button{Text: b.text, Data: b.callback})
	}
	return &pb.CallbackReply{
		Text:    mes.text,
		Buttons: bs,
	}
}
