package main

import (
	"os"

	"github.com/yimtun/remote/terminal-front/front"
	//	"./front"
	//"github.com/yimtun/remote/terminal-front/front"
)

func main() {
	front.SignalTrigger()
	//fmt.Println("len:", len(os.Args))
	//fmt.Println(os.Args)
	user := os.Args[1]

	//fmt.Println("front get user:", user)
	front.FrontStart(user)

}
