package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
)

const (
	SMTP_SERVER = "smtp.qq.com"
	SMTP_PORT   = "25"
)

type Mail struct {
	From    string
	To      string
	Message string
}

type User struct {
	Name   string
	Passwd string
}

type SendMail func(m *Mail) error

func SendWithoutAuth() SendMail {
	conn, err := smtp.Dial(SMTP_SERVER)
	if err != nil {
		panic(err)
	}

	return func(m *Mail) error {
		conn.Mail(m.From)
		conn.Rcpt(m.To)
		data, err := conn.Data()
		if err != nil {
			panic(err)
		}
		defer data.Close()

		buf := bytes.NewBufferString(m.Message)
		if _, err = buf.WriteTo(data); err != nil {
			return err
		}
		return nil
	}

}

func SendWithAuth(user *User) SendMail {
	auth := smtp.PlainAuth(
		"",
		user.Name,
		user.Passwd,
		SMTP_SERVER,
	)

	return func(m *Mail) error {
		err := smtp.SendMail(
			SMTP_SERVER+":"+SMTP_PORT,
			auth,
			m.From,
			[]string{m.To},
			[]byte(m.Message),
		)

		return err
	}

}

func main() {

	mail := &Mail{
		From:    "12345678@qq.com",
		To:      "iloveu@sina.com",
		Message: "subject:欢迎使用go语言\n\n这是一封使用go smtp发送的测试邮件",  //Subject：xxxx \n\n 这个代表这封邮件的标题
	}

	user := &User{Name: "12345678@qq.com", Passwd: "your password"}
	
	//sendMail := SendWithoutAuth() //无需验证
	sendMail := SendWithAuth(user)  //需要验证
	err := sendMail(mail)
	if err != nil {
		panic(err)
		os.Exit(-1)
	}
	fmt.Println("Send Mail Successfully!")
	os.Exit(0)
}
