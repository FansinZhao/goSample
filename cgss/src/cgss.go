package main

import (
	"bufio"
	"cg"
	"fmt"
	"ipc"
	"log"
	"os"
	"strconv"
	"strings"
)

var centerClient *cg.CenterClient

func startCenterClient() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}
	return nil
}

func Help(args []string) int {
	fmt.Println(`
	Commands:
	login <username><level><exp>
	logout <username>
	send <message>
	listplayer
	quit(q)
	help(h)
	`)
	return 0
}

func Quit(args []string) int {
	return 1
}

func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
		return 0
	}
	centerClient.RemovePlayer(args[1])
	return 0
}

func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("USAGE: login <username><level><exp>")
		return 0
	}

	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("level参数应该为整数!")
	}

	exp, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("exp参数应该为整数!")
	}

	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp
	err = centerClient.AddPlayer(player)

	if err != nil {
		fmt.Println("增加玩家失败")
	}

	return 0
}

func ListPlayer(args []string) int {
	ps, err := centerClient.ListPlayer("")
	if err != nil {
		fmt.Println("查询玩家失败", err)
	} else {
		for i, v := range ps {
			fmt.Println(i+1, ":", v)
		}
	}
	return 0
}

func Send(args []string) int {
	message := strings.Join(args[1:], " ")
	err := centerClient.BroadCast(message)
	if err != nil {
		fmt.Println("Failed.", err)
	}
	return 0
}

func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"help":       Help,
		"h":          Help,
		"quit":       Quit,
		"q":          Quit,
		"login":      Login,
		"logout":     Logout,
		"listplayer": ListPlayer,
		"send":       Send,
	}
}

func main() {
	logfile, err := os.OpenFile("../logs/cgss.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer logfile.Close()
	logger := log.New(os.Stdout, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	//	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	logger.Println("开始游戏")
	startCenterClient()
	Help(nil)
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()

	for {
		fmt.Print("commad> ")
		b, _, _ := r.ReadLine()
		lines := string(b)
		tokens := strings.Split(lines, " ")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("未知命令")
		}
	}
}
