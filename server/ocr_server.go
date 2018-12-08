package main

import (
	"context"

	"grpc-go/ocr"
	"net"
	"google.golang.org/grpc"

	"github.com/otiai10/gosseract"
)

type OcrService struct {

}

func (o* OcrService) GetResult(ctx context.Context, in *ocr.File)(*ocr.OcrResult,error) {

	result:=&ocr.OcrResult{}


	client:=gosseract.NewClient()

	defer client.Close()

	err:=client.SetImageFromBytes(in.Bytes)

	if err!=nil {

		result.Code=300
		result.Message="load image error"
		return result,nil
	}

	//获取到识别后的文字
	text,_:=client.Text()

	result.Code=200
	result.Message="success"
	result.Data=text

	return result,nil

}

func main() {
	serviceAddress:=":50052"
	ocrservice:=new(OcrService)

	ls,_:=net.Listen("tcp",serviceAddress)

	gs:=grpc.NewServer()

	ocr.RegisterOcrServiceServer(gs,ocrservice)

	gs.Serve(ls)
}