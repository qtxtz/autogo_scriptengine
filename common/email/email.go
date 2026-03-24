package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"mime/multipart"
	"net/mail"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

// Client 邮件客户端
type Client struct {
	smtpServer string
	username   string
	password   string
	useAuth    bool
}

// NewClient 创建邮件客户端
func NewClient(smtpServer, username, password string, useAuth bool) *Client {
	return &Client{
		smtpServer: smtpServer,
		username:   username,
		password:   password,
		useAuth:    useAuth,
	}
}

// Send 发送普通文本邮件
func (c *Client) Send(to, subject, body string) error {
	host := strings.Split(c.smtpServer, ":")[0]

	// 创建邮件
	from := mail.Address{Name: "", Address: c.username}
	toAddr := mail.Address{Name: "", Address: to}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = toAddr.String()
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	// 构建邮件内容
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 发送邮件
	var auth smtp.Auth
	if c.useAuth {
		auth = smtp.PlainAuth("", c.username, c.password, host)
	}

	err := smtp.SendMail(c.smtpServer, auth, c.username, []string{to}, []byte(message))
	if err != nil {
		// 如果 SSL/TLS 失败，尝试使用 TLS
		if strings.Contains(c.smtpServer, "465") || strings.Contains(c.smtpServer, "587") {
			return c.sendWithTLS(to, message)
		}
	}

	return err
}

// SendWithAttachment 发送带附件的邮件
func (c *Client) SendWithAttachment(to, subject, body string, files []string) error {
	host := strings.Split(c.smtpServer, ":")[0]

	// 创建邮件
	from := mail.Address{Name: "", Address: c.username}
	toAddr := mail.Address{Name: "", Address: to}

	// 构建邮件内容
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 添加邮件头
	fmt.Fprintf(&buf, "From: %s\r\n", from.String())
	fmt.Fprintf(&buf, "To: %s\r\n", toAddr.String())
	fmt.Fprintf(&buf, "Subject: %s\r\n", subject)
	fmt.Fprintf(&buf, "MIME-Version: 1.0\r\n")
	fmt.Fprintf(&buf, "Content-Type: multipart/mixed; boundary=%s\r\n\r\n", writer.Boundary())

	// 添加邮件正文
	part, _ := writer.CreatePart(map[string][]string{
		"Content-Type": []string{"text/plain; charset=utf-8"},
	})
	part.Write([]byte(body))

	// 添加附件
	for _, file := range files {
		fileData, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		attachment, _ := writer.CreatePart(map[string][]string{
			"Content-Type":              []string{"application/octet-stream"},
			"Content-Disposition":       []string{fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(file))},
			"Content-Transfer-Encoding": []string{"base64"},
		})
		attachment.Write(fileData)
	}

	writer.Close()

	// 发送邮件
	var auth smtp.Auth
	if c.useAuth {
		auth = smtp.PlainAuth("", c.username, c.password, host)
	}

	err := smtp.SendMail(c.smtpServer, auth, c.username, []string{to}, buf.Bytes())
	if err != nil {
		// 如果 SSL/TLS 失败，尝试使用 TLS
		if strings.Contains(c.smtpServer, "465") || strings.Contains(c.smtpServer, "587") {
			return c.sendWithTLS(to, buf.String())
		}
	}

	return err
}

// sendWithTLS 使用 TLS 发送邮件
func (c *Client) sendWithTLS(to, message string) error {
	host := strings.Split(c.smtpServer, ":")[0]

	// 创建 TLS 连接
	conn, err := tls.Dial("tcp", c.smtpServer, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	})
	if err != nil {
		return err
	}
	defer conn.Close()

	// 创建 SMTP 客户端
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer client.Quit()

	// 认证
	if c.useAuth {
		auth := smtp.PlainAuth("", c.username, c.password, host)
		if err := client.Auth(auth); err != nil {
			return err
		}
	}

	// 设置发件人和收件人
	if err := client.Mail(c.username); err != nil {
		return err
	}
	if err := client.Rcpt(to); err != nil {
		return err
	}

	// 发送邮件内容
	wc, err := client.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	_, err = wc.Write([]byte(message))
	return err
}
