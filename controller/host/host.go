package host

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func TestConnect(ip string) (bool, error) {

	ok, error := TestConnectByKey(ip, "root", "22", "./keys/defaultKey")
	fmt.Println(ok, error)
	return ok, error

}

func TestConnectByKey(host, user, port, keyPath string) (bool, error) {
	pemBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return false, err
	}

	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		return false, err
		//		log.Fatalf("parse key failed:%v", err)
	}
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return false, err
		//		log.Fatalf("dial failed:%v", err)
	}
	defer conn.Close()
	session, err := conn.NewSession()
	if err != nil {
		//		log.Fatalf("session failed:%v", err)
		return false, err
	}
	defer session.Close()

	b, err := session.CombinedOutput("echo ok")
	if err != nil {
		return false, err

	}

	outputBuf := string(b)
	//fmt.Println(outputBuf)
	return outputBuf == "ok\n", nil

}

type Host struct {
	ip            string
	port          string
	user          string
	registry_tyte string
	id            string
	key           string
	passwd        string
}

var hostlist map[Host]string

func Addhost() {
	/*
	   rs add -p  172.16.100.2
	   rs add -k  172.16.100.2
	   rs add    172.16.100.2
	*/
}

/*

func RegistryHostByKey(ip, port, user, registry_tyte, id string) error {
	var host Host
	if ip == "" {
		var err error = errors.New("ip cant't empty")
		return err
	} else {
		host.ip = ip
	}
	if port == "" {
		host.port = "22"
	} else {
		host.port = port
	}
	if user == "" {
		host.user = "root"
	} else {
		host.user = user
	}
	if registry_tyte == "" && key == "" {
		host.registry_tyte = "key"
		host.key = "defaultKey"
	}
	if registry_tyte == "" && key != "" {
		host.registry_tyte = "key"
		host.key = key
	}
	if registry_tyte != "" && key == "" {
		host.registry_tyte = "key"
		host.key = "defaultKey"
	}

	return host
}
*/
