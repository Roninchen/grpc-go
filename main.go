package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcObj struct {
	Id   int    `json:"id"` // struct标签， 如果指定，jsonrpc包会在序列化json时，将该聚合字段命名为指定的字符串
	Name string `json:"name"`
}

// 需要传输的对象
type ReplyObj struct {
	Ok  bool   `json:"ok"`
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

type ServerHandler struct{}

func (serverHandler ServerHandler) GetName(id int, returnObj *RpcObj) error {
	log.Println("server\t-", "recive GetName call, id:", id)
	returnObj.Id = id
	returnObj.Name = "名称1"
	return nil
}

func (serverHandler ServerHandler) SaveName(rpcObj RpcObj, returnObj *ReplyObj) error {
	log.Println("server\t-", "recive SaveName call, RpcObj:", rpcObj)
	returnObj.Ok = true
	returnObj.Id = rpcObj.Id
	returnObj.Msg = "存储成功"
	return nil
}

func main() {
	server := rpc.NewServer()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("server\t-", "listen error:", err.Error())
	}
	defer listener.Close()
	log.Println("server\t-", "start listion on port 8888")
	// 新建处理器
	serverHandler := &ServerHandler{}
	// 注册处理器
	server.Register(serverHandler)

	// 等待并处理链接
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err.Error())
			}

			// 在goroutine中处理请求
			// 绑定rpc的编码器，使用http connection新建一个jsonrpc编码器，并将该编码器绑定给http处理器
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}()

	//client
	client, err := net.DialTimeout("tcp", "localhost:8888", 1000*1000*1000*30) // 30秒超时时间
	if err != nil {
		log.Fatal("client\t-", err.Error())
	}
	defer client.Close()
	clientRpc := jsonrpc.NewClient(client)
	var rpcObj RpcObj
	// 请求数据，rpcObj对象会被填充
	clientRpc.Call("ServerHandler.GetName", 1, &rpcObj)
	fmt.Println("rpc",rpcObj)
	// 远程返回的对象
	var reply ReplyObj
	// 传给远程服务器的对象参数
	saveObj := RpcObj{2, "对象2"}
	// 请求数据
	clientRpc.Call("ServerHandler.SaveName", saveObj, &reply)
	fmt.Println(reply.Msg)

	// Asynchronous call

	divCall := clientRpc.Go("ServerHandler.SaveName", saveObj, &reply, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	fmt.Println(replyCall.Reply)
}