package models

import (
	"ASOS/pkg/flags"
)

type Example struct {
	Name     string     `json:"Name"`
	Messages []*Message `json:"Messages"`
}

var ExampleList []*Example

func GetExample(name string) *Example {
	for _, v := range ExampleList {
		if v.Name == name {
			return v
		} else {
			example := new(Example)
			example.Name = name

			message := new(Message)
			message.Text = "Hello World"
			message.Type = "Info"
			example.Messages = append(example.Messages, message)

			AddExample(example)
			return example
		}
	}
	return nil
}

func AddExample(example *Example) {
	if example == nil {
		return
	}
	ExampleList = append(ExampleList, example)
}

func Setup() {
	example := new(Example)
	example.Name = "Test"

	message := new(Message)
	message.Text = "Hello World"
	message.Type = "Info"
	example.Messages = append(example.Messages, message)

	AddExample(example)
}

func GetMessage(example *Example, message string) *Message {
	if example == nil {
		return nil
	}

	if example.Messages != nil {
		for _, v := range example.Messages {
			if v.Text == message {
				return v
			}
		}
	}
	return nil
}

func AddMessage(example *Example, message *Message) int {
	if example == nil || message == nil {
		return flags.INVALID_PARAMS
	}

	if GetMessage(example, message.Text) != nil {
		return flags.ERROR_EXIST_MESSAGE
	}

	example.Messages = append(example.Messages, message)
	return flags.SUCCESS
}
