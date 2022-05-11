package front

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/chzyer/readline"
	"github.com/yimtun/remote/controller/host"
	"github.com/yimtun/remote/controller/strpro"
	"github.com/yimtun/remote/ssh"
)

/*

rs add  172.16.101.102  // default user root  default add way key  default key default-key
// rs add -u root -k default-key 172.16.101.102

rs add -p passname 172.16.101.102  default user root
rs add -k keyname 172.16.101.102   default user root

rs add -u root 172.16.101.102      default add way key use default-key
rs add -p pwdname -u root 172.16.101.102
rs add -k keyname -u username 172.16.101.102




*/

//terminal controller start
func FrontStart(user string) {
	//terminal
	fmt.Println("hello:", user)
	//check user privileges
	fmt.Println("欢迎使用 终端堡垒机 获取帮助请安 help")
	r1, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer r1.Close()

	for {

		input, err := r1.Readline()
		if err != nil { // io.EOF
			break
		}
		fmt.Println("输入:", input)

		inputSlice, n := strpro.StrFartory(input)

		if n != 0 && inputSlice[0] != "" {
			firstWord := inputSlice[0]
			switch firstWord {
			case "exit":
				if n == 1 {
					os.Exit(0)
				} else {
					fmt.Println("exit to quit ")
				}

			case "rs":
				args := argsParse(inputSlice)

				//fmt.Println(len(args))
				rsAdd(args)
				//sshclient.NewSessionByPasswd("root", "toshiba", "172.16.101.102", "22")
			//	sshclient.ConnectByPassword("172.16.101.102", "22", "root", "toshiba")

			//fmt.Println(argsParse(inputSlice)["arg1"])

			case "help":
				fmt.Println("remote usage .....")
			//		default:
			//			fmt.Println("press help for help")
			case "cobra":
				//fmt.Println("cobra")
				//cmd := exec.Command("/gopath/src/github.com/yimtun/rs/rs", "list")
				cmd := exec.Command("/gopath/src/github.com/stromland/cobra-prompt/example/example")
				//cmd := exec.Command("top") ok
				//fmt.Println("xxx")
				cmd.Stdout = os.Stdout //
				cmd.Stdin = os.Stdin
				cmd.Run()
			default:
				fmt.Println("rs usage .......  ")
			}

		} else {
			//fmt.Println("press help for help")

		}
		if n != 0 && inputSlice[0] != "rs" {
			///

		}

	}
	for { //  待优化
		time.Sleep(time.Second)
	}
}

func a1(input string) {

	inputSlice, n := strpro.StrFartory(input)

	// rs add host
	if inputSlice[0] == "rs" && n == 3 {
		if inputSlice[1] == "add" && strpro.IsIP(inputSlice[2]) && inputSlice[2] != "-k" {
			//rs add 172.16.11.99

			fmt.Println("合法测输入", inputSlice)
			//add host by current key  缺省的  the new create key is current key
			ok, error := host.TestConnect(inputSlice[2])
			if ok && error == nil {
				fmt.Println("test passed")
				//addHost()

			} else {
				fmt.Println("test connect failure:", error)
			}

		}

		if inputSlice[1] == "add" && !strpro.IsIP(inputSlice[2]) && inputSlice[2] != "-k" {
			fmt.Println("默认key 无效的地址")

		}

	}

	if inputSlice[0] == "rs" && n == 5 {
		if inputSlice[1] == "add" && inputSlice[2] == "-k" && inputSlice[3] != "" && strpro.IsIP(inputSlice[4]) {

			fmt.Println("通过指定key 添加host", inputSlice)
			//add host by default key

		}

		if inputSlice[1] == "add" && !strpro.IsIP(inputSlice[4]) {
			fmt.Println("自定义可以 添加 无效的地址")

		}

	}
	if inputSlice[0] == "rs" && n == 2 && strpro.IsIP(inputSlice[1]) {
		//checkout regiinputSlicey_type

		ssh.ConnectByPassword(inputSlice[1], "22", "root", "toshiba")
		// connect ssh by

	}
	if inputSlice[0] == "rs" && n == 2 && !strpro.IsIP(inputSlice[1]) {
		//checkout regiinputSlicey_type
		//ip = findipbyname()
		//ssh.ConnectByPassword(ip, "22", "root", "toshiba")
		// connect ssh by

	}
	if inputSlice[0] == "rs" && n == 3 && inputSlice[1] == "addkey" && inputSlice[2] != "" {
		keyName := inputSlice[2]
		cmd := exec.Command("vim", "-Z", "-c", "set paste", "./keys/"+keyName)
		//vim -Z  -c "set nu"  a.txt
		// of secure  vim -Z  for useful set paste
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()

	} else {
		//fmt.Println(inputSlice)
		if inputSlice[0] == "rs" && strpro.IsIP(inputSlice[1]) {

		} else {
			// bug [rs 172.16.101.102]
			fmt.Println("rs usage :rs addkey ")

		}
	}
	if inputSlice[0] == "rs" && n == 1 {
		fmt.Println("rs usage ")
	}
}

func SignalTrigger() {
	var notifySignals []os.Signal
	c := make(chan os.Signal, 10)
	notifySignals = append(notifySignals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)
	signal.Notify(c, notifySignals...)
	go func() {
		for sig := range c {
			if sig == syscall.SIGINT {
				fmt.Println("usr outlogin")
				// unregistry session before logout
				//return
				os.Exit(130)
				//signal.Ignore(sig)

			}

			if sig == syscall.SIGTSTP {
				//	os.Exit(0)
				//fmt.Println(os.Getpid())
				fmt.Println("usr outlogin")
				// unregistry session before logout
				//			ok = "xx"
				//			fmt.Println(ok)
				//return
				os.Exit(0)
				//			signal.Ignore(sig)

			}

			/*
				fmt.Printf("received signal: %v\n", sig)
								signal.Ignore(sig)
												os.Exit(0)
			*/
		}
	}()
}
