package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type kafkaStruct struct{
	A int64
	B string
}


func main()  {
	syncProducer()
	//asyncProducer1(Address)
	//SaramaProducer()
}

// 同步消息模式
func syncProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()

	kafkaInstance := kafkaStruct{
		B: "xyz",
	}
	for i := 0; i<10; i++ {
		kafkaInstance.A = int64(i)
		pushData, err := json.Marshal(kafkaInstance)
		if err != nil {

		}
		msg := &sarama.ProducerMessage {
			Topic:"test_par",
			Value:sarama.ByteEncoder(pushData),
		}
		part, offset, err := p.SendMessage(msg)
		if err != nil {
			fmt.Println("fail", err)
		} else {
			fmt.Println("success", part, offset)
		}
		fmt.Println(msg.Partition, msg.Offset, msg.Value)
		time.Sleep(10 * time.Millisecond)
	}
}

// 异步消息模式
func SaramaProducer()  {

	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1

	fmt.Println("start make producer")
	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()

	//循环判断哪个通道发送过来数据.
	fmt.Println("start goroutine")
	go func(p sarama.AsyncProducer) {
		for{
			select {
			case  suc := <-p.Successes():
				fmt.Println("offset: ", suc.Offset,
					"partitions: ", suc.Partition,
					"key: ", suc.Key,
					"value: ", suc.Value)
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
			}
		}
	}(producer)

	var value string
	for i:=0; i<10; i++ {
		time11 := time.Now()
		value = "this is a message (test_abc) "+time11.Format("15:04:05")

		// 发送的消息,主题。
		// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
		msg := &sarama.ProducerMessage{
			Topic: "test_abc",
		}

		//将字符串转化为字节数组
		//msg.Value = sarama.ByteEncoder(value)
		//fmt.Println(value)

		// 这样传输也行
		msg.Value = sarama.StringEncoder(value)

		//使用通道发送
		producer.Input() <- msg
		fmt.Println("input end")
		//time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(10 * time.Millisecond)
}
