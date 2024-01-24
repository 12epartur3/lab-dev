package main

import(
	"gopkg.in/gomail.v2"
	//"crypto/tls"
	"fmt"
	"time"
)

func main() {
	body := "nihao"
	err := Send([]string{"yuanye@shopee.com"},
	     []string{},
	     "yytest",
	     body,
	     "")
	fmt.Printf("send over err=%v\n", err)
}

func Send(recipientsTo, recipientsCc []string, emailSubject, emailBody, filePath string) error {
   const (
      maxRetries  = 10
      retryPeriod = 200 * time.Millisecond
      //emailFrom   = "yuanshan@shopee.com"
      emailFrom   = "yuanye@shopee.com"
      emailHost   = "smtp.shopeemobile.com"
      emailPort   = 587
      //emailHost   = "smtp.shopee.io"
      //emailPort   = 2525
   )

   m := gomail.NewMessage()
   m.SetHeader("From", emailFrom)
   m.SetHeader("To", recipientsTo...)
   m.SetHeader("Subject", emailSubject)
   m.SetBody("text/html", emailBody)
   m.SetHeader("Cc", recipientsCc...)
   //m.Attach(filePath, gomail.Rename(filePath+".csv"))
   dialer := gomail.NewDialer(emailHost, emailPort, emailFrom, "")
   //dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
   fmt.Printf("from = %v, to = %v\n", emailFrom, recipientsTo)

   err := dialer.DialAndSend(m)
   //ulog.DefaultLogger().Withs("recipientsTo", recipientsTo).WithTypedFields(ulog.String("email_subject", emailSubject)).Info("email_info")
   if err == nil {
      return nil
   }
   for i := 0; err != nil && i < maxRetries; i++ {
      err = dialer.DialAndSend(m)
      time.Sleep(retryPeriod)
   }

   // try send without attachment
   m = gomail.NewMessage()
   m.SetHeader("From", emailFrom)
   m.SetHeader("To", recipientsTo...)
   m.SetHeader("Subject", emailSubject)
   m.SetBody("text/html", emailBody)
   m.SetHeader("Cc", recipientsCc...)
   err = dialer.DialAndSend(m)
   for i := 0; err != nil && i < maxRetries; i++ {
      err = dialer.DialAndSend(m)
      time.Sleep(retryPeriod)
   }
   return err
}
