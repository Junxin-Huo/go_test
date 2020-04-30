package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"unsafe"
	"bytes"
)

func httpPost() {
	resp, err := http.Post("http://www.01happy.com/demo/accept.php",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost2() {
	//生成client 参数为默认
	client := &http.Client{}

	//生成要访问的url
	url := "http://www.baidu.com"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)
}

type picCaptchaResponse struct {
	Code     int32
	Msg      string
}

func httpPost3() {
	tmp := make(map[string]interface{})
	tmp["captcha"] = "12346"
	tmp["captchaid"] = "xjPVif8YbYFkCidm7hMy"
	bytesData, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://127.0.0.1:8082/picCaptchaVerify"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)

	var response picCaptchaResponse
	err = json.Unmarshal(respBytes, &response)
	if err == nil {
		fmt.Println(response)
	}
}

type mailCaptchaResponse struct {
	Code     int32
	Msg      string
}

func httpPost4() {
	tmp := make(map[string]interface{})
	tmp["to"] = "liuxiaojun.333@gmail.com"
	tmp["lang"] = "zh_CN"
	tmp["biz_key"] = "testbiz"
	tmp["expiry"] = 3600
	bytesData, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://47.97.167.221:8087/captcha"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var response mailCaptchaResponse
	err = json.Unmarshal(respBytes, &response)
	if err == nil {
		fmt.Println(response)
	}
}

func main() {
	httpPost4()
}