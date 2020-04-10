package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

// 订阅回调
func subCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
}

// 连接MQTT服务
func connMQTT(broker, user, passwd string) (bool, MQTT.Client) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(user)
	opts.SetPassword(passwd)

	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return false, mc
	}

	return true, mc
}

// 订阅消息
func Subscribe() {
	// sub的用户名和密码
	b, mc := connMQTT("tcp://127.0.0.1:1883", "admin", "admin")
	if !b {
		fmt.Println("sub connMQTT failed")
		return
	}
	mc.Subscribe("topic_tp", 0x00, subCallBackFunc)
}

// 发布消息
func Publish() {
	// pub的用户名和密码
	b, mc := connMQTT("tcp://127.0.0.1:1883", "admin", "admin")
	if !b {
		fmt.Println("pub connMQTT failed")
		return
	}

	for {
		mc.Publish("topic_tp", 0x00, true, "Hello, this is publisher")
		time.Sleep(time.Second)
	}
}

func main() {
	Subscribe()
	Publish()
}