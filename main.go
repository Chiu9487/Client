package main

import (
	"chatdemo/client/handler"
	"chatdemo/commont/message"
	"chatdemo/commont/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main(){
	//连接服务器
	conn ,err := net.Dial("tcp" ,"127.0.0.1:8181")
	if(err != nil){
		fmt.Println("客户端连接服务器失败, err:" ,err)
		return
	}

	loginState := handler.Login(&conn)
	if loginState {
		go clientHandler(&conn)

		for{
			//var msg *message.ResMessage
			msg := utils.GetMessage(&conn)
			//fmt.Println(msg.MsgData)
			if msg.Type == 3 && (msg.Code == 1){ //退出监听
				(conn).Close()
				os.Exit(0)
				break
			}

			if msg.Type == 2 { //聊天消息
				if msg.Code == 1{ //发送聊天返回

				} else { //被动接收聊天消息
					str := utils.UnMarshalChatData(msg.MsgData)
					fmt.Println(str)
				}
			}
		}
	}
	fmt.Println("客户端结束关闭" )
}



func clientHandler(conn *net.Conn){
	//登录成功,发送聊天
	var choose int //选择
	//监听活跃
	go keepActive(conn)
	var flag bool
	for{
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("-----------------------输入以下的数字，进行下一步操作-----------------------")
		fmt.Println("`-----------------------1.进行私聊-------------------------------------")
		fmt.Println("`-----------------------2.进行群聊-------------------------------------")
		fmt.Println("`-----------------------3.退出-------------------------------------")
		fmt.Scanf("%d\n",&choose)
		switch(choose){
		case 1:
			handler.PrivateChat(conn)
		case 2:
			handler.PublicChat(conn)
		case 3:
			exitCha(*conn)
			flag = true

		default:
			fmt.Println("选择错误~~重新选择")
		}
		if(flag){
			break
		}
	}
}

//保持监听，循环接收服务器发来的消息
func keepActive(conn *net.Conn){

}


func exitCha(conn net.Conn){
	sendMsg := message.ReqMessage{
		Type : 3,
	}
	data ,err := json.Marshal(sendMsg)
	if err != nil{
		fmt.Println("序列话退出信息失败 ,err:" ,err)
		return
	}
	conn.Write(data)

}
