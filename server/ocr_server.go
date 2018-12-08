package main

import (
	"context"

	"grpc-go/ocr"
	"net"
	"google.golang.org/grpc"
	//"github.com/otiai10/gosseract"
	"mime/multipart"
	"fmt"
	"io/ioutil"
	"bytes"
	"net/http"
)

type OcrService struct {

}

func (o* OcrService) GetResult(ctx context.Context, in *ocr.File)(*ocr.OcrResult,error) {

	result:=&ocr.OcrResult{}

	buff := bytes.NewBuffer(in.Bytes)
	writer := multipart.NewWriter(buff)

	writer.WriteField("field", "this is a field")
	w, _ := writer.CreateFormFile("file", "t.png")
	w.Write([]byte(in.Bytes))
	writer.Close()
	var httpClient http.Client
	resp, err := httpClient.Post("http://127.0.0.1:8080/file", writer.FormDataContentType(), buff)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	result.Code=200
	result.Message="success"
	result.Data=string(data)

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