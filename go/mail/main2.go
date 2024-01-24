package main
import (
    "fmt"
    "log"
    "strings"
    sasl "github.com/emersion/go-sasl"
    smtp "github.com/emersion/go-smtp"
)
var (
    server        = "smtp.shopee.io:2525"
    username      = "search_official"
    password      = "CkWSlUby2Qrf3VIzkYzk4Hd0tX4VWYc8"
    sender        = "yuanye@shopee.com"           // must be a white listed sender address, e.g. devops@shopee.com
    recipients    = []string{"yuanye@shopee.com"} // receiver email address
    subject       = `Test Email Subject 1`        // your email subject
    email_content = `Some email content`          // your email content
)

func SendSMTPWithAuth() {
    auth := sasl.NewPlainClient("", username, password)                                              // support plain method
    msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", recipients[0], subject, email_content) // Your email message here
    c, err := smtp.Dial(server)
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()
    if err = c.Auth(auth); err != nil {
        log.Fatal(err)
    }
    if c.SendMail(sender, recipients, strings.NewReader(msg)); err != nil {
        log.Fatal(err)
    }
}

func main() {
	SendSMTPWithAuth()
}
