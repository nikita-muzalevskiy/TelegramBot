package main

import (
	pb "MangaService/protos"
)

var messageList = map[string]message{
	"menu": {
		"Новая манга или продолжить чтение?",
		[]button{},
		func(user string, manga string) ([]button, string) {
			return []button{
				{"Новая манга", user + "/MANGA/new-list/"},
				{"Начатая манга", user + "/MANGA/started-list/"},
			}, "Новая манга или продолжить чтение?"
		},
	},
	"started-list": {
		"",
		[]button{},
		func(user string, manga string) ([]button, string) {
			return GetStartedManga(user)
		},
	},
	"new-list": {
		"",
		[]button{},
		func(user string, manga string) ([]button, string) {
			return GetNewManga(user)
		},
	},
	"read": {
		"",
		[]button{},
		func(user string, manga string) ([]button, string) {
			return Read(user, manga)
		},
	},
}

type message struct {
	text     string
	keyboard []button
	f        func(user string, manga string) ([]button, string)
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
