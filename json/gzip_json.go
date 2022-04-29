package main

import (
	"compress/gzip"
	"io/ioutil"
	"os"
)

func main() {
	////文件打开
	//jsonFile, err := os.Open("wind (1).json")
	//if err != nil {
	//	fmt.Println("error opening json file")
	//	return
	//}
	//defer jsonFile.Close()
	////json数据读入
	//jsonData, err := ioutil.ReadAll(jsonFile)
	//if err != nil {
	//	fmt.Println("error reading json file")
	//	return
	//}
	////压缩json文件
	//data := MarshalToJsonWithGzip(jsonData)
	////写出压缩后的json文件
	//fileWriter(data)
	a("abc")
}

//func Encode(input []byte) ([]byte, error) {
//	// 创建一个新的 byte 输出流
//	var buf bytes.Buffer
//	// 创建一个新的 gzip 输出流
//	gzipWriter := gzip.NewWriter(&buf)
//	// 将 input byte 数组写入到此输出流中
//	_, err := gzipWriter.Write(input)
//	if err != nil {
//		_ = gzipWriter.Close()
//		return nil, err
//	}
//	if err := gzipWriter.Close(); err != nil {
//		return nil, err
//	}
//	// 返回压缩后的 bytes 数组
//	return buf.Bytes(), nil
//}
//
//func MarshalToJsonWithGzip(jsonData []byte) []byte {
//	//
//	dataAfterMarshal, _ := json.Marshal(jsonData)
//	dataAfterGzip, err := Encode(dataAfterMarshal)
//	if err != nil {
//		return nil
//	}
//	return dataAfterGzip
//}
//func fileWriter(str []byte) {
//	file, err := os.OpenFile("gzip.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		fmt.Println("open fail:err=", err)
//		return
//	}
//	defer file.Close()
//	file.Write([]byte(str))
//}

func a(b string) {
	buf, _ := ioutil.ReadFile("wind (1).json")
	fs, _ := os.Create("c")
	w := gzip.NewWriter(fs)
	w.Write(buf)
	w.Flush()
	fs.Close()
}
