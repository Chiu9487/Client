package handler

import (
	"chatdemo/commont/message"
	"chatdemo/commont/utils"
	"encoding/json"
	"fmt"
	"net"
)

func PublicChat(conn *net.Conn){
	fmt.Println("---------------------------------群聊频道-----------------------------")
	for{
		var content string
		fmt.Println("输入聊天内容(-1:退出群聊)：")
		fmt.Scanf("%s\n",&content)
		if content == "-1"{
			return
		}
		err := sendChatReq(conn,2,0,content)
		if err != nil{
			continue
		}

		datastring := utils.GetTime()
		fmt.Printf("你在%s向所有人发送了消息：%s\n",datastring,content)
	}
}
//私聊
func PrivateChat(conn *net.Conn){
	fmt.Println("---------------------------------私聊频道-----------------------------")
	for {
		var toId int
		fmt.Println("输入接收聊天的的id")
		fmt.Scanf("%d\n",&toId)
		for{
			var content string
			fmt.Println("输入聊天内容(-1:重新选择联系人 -2:退出私聊)：")
			fmt.Scanf("%s\n",&content)
			if content == "-1"{
				break
			}
			if content == "-2"{
				return
			}

			err := sendChatReq(conn,1,toId,content)
			if err != nil{
				continue
			}

			datastring := utils.GetTime()
			fmt.Printf("你在%s向%d发送了消息：%s\n",datastring,toId,content)
		}
	}
}



func sendChatReq(conn *net.Conn ,chatType ,toId int,content string)(err error){
	msg := message.CreateNewMsg(chatType,toId,content)
	s ,err := json.Marshal(msg)
	if err != nil{
		fmt.Println("封消息实体失败，err ：" ,err)
		return
	}

	req := message.CreateReqdMessage(2,string(s))

	reqStr ,err := json.Marshal(req)
	if err != nil{
		fmt.Println("封消息包失败，err ：" ,err)
		return
	}


	_ ,err = (*conn).Write(reqStr)
	if err != nil{
		fmt.Println("发送聊天消息给服务器失败， err:" ,err)
	}
	return
}

