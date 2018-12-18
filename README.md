GRPC 微服务

# 安装 go-kit
* git clone https://github.com/go-kit/kit.git

# otiai10
* go get -t github.com/otiai10/gosseract
* docker run -it --rm otiai10/gosseract

# install tesseract(调用文件)
* https://github.com/tesseract-ocr/tesseract/wiki
* mac 下 brew install tesseract
* brew list tesseract


# 安装ocrserver(调用服务)
* https://github.com/otiai10/ocrserver
* % go get github.com/otiai10/ocrserver/...
* % PORT=8080 ocrserver
* open http://localhost:8080