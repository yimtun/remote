package front

import (
	"fmt"
	"strconv"

	//"github.com/yimtun/remote/controller/host"

	//"github.com/yimtun/remote/controller/host"
	"github.com/yimtun/remote/terminal-front/host"
	"github.com/yimtun/remote/terminal-front/httpclient"
	"github.com/yimtun/remote/terminal-front/secret"
	"github.com/yimtun/remote/terminal-front/sshclient"
	"github.com/yimtun/remote/terminal-front/utils"
)

func rsAdd(args map[string]string) {
	n := len(args)
	if len(args) != 0 {
		arg1 := args["arg1"]
		switch {
		case utils.IsIP(arg1) && n == 1: // like  rs 172.16.x.33
			ip := arg1
			fmt.Println(ip)
			//sshclient.ConnectByPassword(ip, "22", "root", "toshiba")
			sshclient.ConnectByKey(ip, "22", "root")
		case arg1 == "add": // rs add
			switch {
			case n == 2 || n == 3:
				arg2 := args["arg2"]
				fmt.Println(arg2)
				switch {

				case arg2 == "key":
					switch {
					case n == 2:
						fmt.Println("usage:", "rs add key keyname")
					case n == 3:

						/*
							err := os.Setenv("HOME", "/root") //临时设置 系统环境变量
							if err != nil {
								fmt.Println(err.Error())
							}
							keyname := args["arg3"]
							cmd := exec.Command("/usr/bin/vim", "-Z", "-c", "set paste", "./keys/"+keyname)
							cmd.Stdin = os.Stdin
							cmd.Stdout = os.Stdout
							cmd.Stderr = os.Stderr

							if err := cmd.Run(); err != nil {
								fmt.Println("Error: ", err)
							}
						*/
						kname := args["arg3"]
						//  isExist key ??
						ok, err := httpclient.KeyIsNotExist(kname)
						if ok {
							keystr := secret.AddKey(kname)
							httpclient.AddKey(kname, keystr)

						} else {
							fmt.Println(err)

						}

					}
				case arg2 == "passwd":
					switch {
					case n == 2:
						fmt.Println("usage:", "rs add passwd passwdname")
					case n == 3:
						//str := args["arg3"]
						//	secret.AddKey(str)

					}

				case utils.IsIP(args["arg2"]) && n == 2: // add host
					ip := args["arg2"]
					ok, error := host.TestConnect(ip)
					if ok && error == nil {
						fmt.Println("test passed")
						fmt.Println("rs add ip")
						host.Addhost(ip)
						fmt.Println(host.HostList)

					} else {
						fmt.Println("test connect failure:", error)
					}

				}

			}
		case arg1 == "list":
			switch {

			case n == 2:

				arg2 := args["arg2"]

				switch {
				case arg2 == "key":
					httpclient.ListKey()
					//		fmt.Println(secret.Keys)
				case arg2 == "passwd":
					fmt.Println("list passwds")
				case arg2 == "host":
					fmt.Println("[h1]-172.16.100.1 [h2]-162.16.100.2 \n[h3]-172.16.100.3 [h4]-172.16.100.4 \n[h5]-172.16.100.5 [h6]-172.16.100.6")

				}

			}
		case arg1 == "get":
			httpclient.GetKey("key1010")
		case arg1 == "test":
			httpclient.AddHost()
		case arg1 == "set":
			switch {
			case n == 2 || n == 3:
				arg2 := args["arg2"]
				switch {
				case arg2 == "default-key":
					fmt.Println("default-key")
				case arg2 == "default-passwd":

				}

			}

		}

	}
}

func argsParse(input []string) (args map[string]string) {

	//fmt.Println(defineArg(len(input)))

	//n := len(input)
	//var args map[string]string
	argMap := make(map[string]string)

	for n, _ := range input {

		if n != 0 {
			name := defineArg(n)
			//fmt.Println(name)
			name = input[n]
			//fmt.Println(defineArg(n), name)
			argMap[defineArg(n)] = name

		}
	}
	return argMap
	//	fmt.Println(args)

}

func defineArg(n int) string {
	return "arg" + strconv.Itoa(n)

}

func listhost() {
	fmt.Println("[h1]-172.16.100.1 [h2]-162.16.100.2 \n[h3]-172.16.100.3 [h4]-172.16.100.4 \n[h5]-172.16.100.5 [h6]-172.16.100.6")

}

// rs add
