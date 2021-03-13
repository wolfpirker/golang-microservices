package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

// lesson from: https://tutorialedge.net/golang/go-protocol-buffer-tutorial/
// 1. follow this tutorial, plus
// > https://jbrandhorst.com/post/go-protobuf-tips/
// 2. revise code using newer google.golang.org/protobuf module
// > using https://developers.google.com/protocol-buffers/docs/gotutorial
// > https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

func main() {
	fmt.Println("Hello World")

	wolfgang := &Person{
		Name: "Wolfgang",
		Age:  34,
		SocialFollowers: &SocialFollowers{
			Twitter: 10,
			Youtube: 20,
		},
	}

	data, err := proto.Marshal(wolfgang)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	fmt.Println(data)

	newWolfgang := &Person{}
	err = proto.Unmarshal(data, newWolfgang)
	if err != nil {
		log.Fatal("unmarshalling error: ", err)
	}
	fmt.Println("Name: ", newWolfgang.Name)
	fmt.Println("Age: ", newWolfgang.Age)
	fmt.Println("YT: ", newWolfgang.SocialFollowers.GetYoutube())
	fmt.Println("Twitter: ", newWolfgang.SocialFollowers.GetTwitter())
}
