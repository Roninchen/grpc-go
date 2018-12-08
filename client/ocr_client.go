package main

import (
	"google.golang.org/grpc"
	"grpc-go/ocr"
	"context"
	"io/ioutil"
	"fmt"
)

func main() {

	serviceAddress:="127.0.0.1:50052"

	coon,err:=grpc.Dial(serviceAddress,grpc.WithInsecure())

	if err !=nil {
		panic("connect error")
	}

	defer coon.Close()

	ocrClient:=ocr.NewOcrServiceClient(coon)

	bytes,err:=ioutil.ReadFile("C://myuser//hyperchain.png")

	if err!=nil {
		panic("open file error")
	}

	result,_:=ocrClient.GetResult(context.Background(),&ocr.File{Bytes:bytes})

	if result.Code==200 {

		fmt.Println(result.Data)

	}else {

		fmt.Println(result.Message)

	}


}
