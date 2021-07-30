package handler

import (
	"chatdemo/commont/message"
	"chatdemo/commont/utils"
	"encoding/json"
	"fmt"
	"net"
)

// 登陆
func Login(conn *net.Conn)bool{
	var userId int
	var pwd string
	for{
		fmt.Println("输入登陆的id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("输入登陆的密码")
		fmt.Scanf("%s\n",&pwd)

		//消息结构体
		snedMess := message.ReqMessage{
			Type : 1,
			MsgData : fmt.Sprintf("%d-%s",userId,pwd),
		}

		loginMsg,err := json.Marshal(snedMess) //登录消息
		if err != nil{
			fmt.Println("登陆消息序列化失败，，err：" ,err)
			continue
		}
		_,err = (*conn).Write([]byte(loginMsg))
		if err != nil{
			fmt.Println("登录失败  ，err :",err)
			continue;
		}
		var msg *message.ResMessage
		msg = utils.GetMessage(conn)

		fmt.Println(msg.MsgData)
		if msg.Type == 1 && (msg.Code == 1 || msg.Code == 2){
			return true
		}
		continue
	}
	return false;
}

