package main

import(
	"gopkg.in/gomail.v2"
	//"crypto/tls"
	"fmt"
	"time"
)

func main() {
	body := `<!DOCTYPE html><html><head>    <meta charset="utf-8">    <title>Awesome go-echarts</title>    <script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"></script>    <script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts-gl.min.js"></script></head><body>    <style> .container {display: flex;justify-content: center;align-items: center;} .item {margin: auto;} </style><div class="container">    <div class="item" id="zwhXUTNTOaWM" style="width:900px;height:500px;"></div></div><script type="text/javascript">    "use strict";    let goecharts_zwhXUTNTOaWM = echarts.init(document.getElementById('zwhXUTNTOaWM'), "white");    let option_zwhXUTNTOaWM = {"color":["#5470c6","#91cc75","#fac858","#ee6666","#73c0de","#3ba272","#fc8452","#9a60b4","#ea7ccc"],"grid3D":{"show":true,"boxWidth":120,"boxDepth":60,"viewControl":{}},"legend":{"show":false,"type":""},"series":[{"name":"bar3d","type":"bar3D","showSymbol":false,"waveAnimation":false,"coordinateSystem":"cartesian3D","renderLabelForZeroData":false,"selectedMode":false,"animation":false,"data":[{"value":[4,3,99.7554]},{"value":[4,8,99.9361]},{"value":[4,10,99.9521]},{"value":[4,1,99.9512]},{"value":[4,6,99.9884]},{"value":[4,9,99.9886]},{"value":[4,0,98.7668]},{"value":[4,4,99.8173]},{"value":[4,5,99.9075]},{"value":[4,7,99.9946]},{"value":[4,2,99.9346]},{"value":[2,9,99.9838]},{"value":[2,8,99.9454]},{"value":[2,10,99.9539]},{"value":[2,6,99.9898]},{"value":[2,0,98.6864]},{"value":[2,1,99.9502]},{"value":[2,4,96.4241]},{"value":[2,5,99.9048]},{"value":[2,2,99.9367]},{"value":[2,3,99.7505]},{"value":[2,7,99.9976]},{"value":[11,5,1]},{"value":[11,6,1]},{"value":[11,2,1]},{"value":[11,4,1]},{"value":[11,7,1]},{"value":[11,10,1]},{"value":[11,1,1]},{"value":[11,3,1]},{"value":[11,8,1]},{"value":[11,0,1]},{"value":[11,9,1]},{"value":[10,8,99.9359]},{"value":[10,6,99.9907]},{"value":[10,7,99.9952]},{"value":[10,5,99.908]},{"value":[10,4,99.8175]},{"value":[10,0,98.5875]},{"value":[10,1,99.9491]},{"value":[10,2,99.935]},{"value":[10,9,99.9879]},{"value":[10,10,99.9565]},{"value":[10,3,99.7595]},{"value":[7,9,99.9882]},{"value":[7,8,99.9372]},{"value":[7,10,99.9424]},{"value":[7,0,98.6546]},{"value":[7,3,99.757]},{"value":[7,7,99.9996]},{"value":[7,1,99.9062]},{"value":[7,2,99.912]},{"value":[7,4,99.8055]},{"value":[7,5,99.9055]},{"value":[7,6,99.9892]},{"value":[1,9,99.9879]},{"value":[1,10,99.9533]},{"value":[1,5,99.9046]},{"value":[1,0,98.6918]},{"value":[1,6,99.9883]},{"value":[1,7,99.9949]},{"value":[1,8,99.9411]},{"value":[1,1,99.9544]},{"value":[1,2,99.9348]},{"value":[1,3,99.75]},{"value":[1,4,96.0439]},{"value":[9,2,99.9195]},{"value":[9,8,99.942]},{"value":[9,10,99.9548]},{"value":[9,1,99.9155]},{"value":[9,6,99.9895]},{"value":[9,7,99.9951]},{"value":[9,0,98.6217]},{"value":[9,5,99.9069]},{"value":[9,9,99.9884]},{"value":[9,3,99.7576]},{"value":[9,4,99.8008]},{"value":[6,2,99.9359]},{"value":[6,9,99.9867]},{"value":[6,10,99.9585]},{"value":[6,1,99.9488]},{"value":[6,3,99.758]},{"value":[6,4,99.8171]},{"value":[6,6,99.9881]},{"value":[6,5,99.9075]},{"value":[6,7,99.9942]},{"value":[6,8,99.9379]},{"value":[6,0,98.7225]},{"value":[5,8,99.9367]},{"value":[5,0,98.7269]},{"value":[5,1,99.9501]},{"value":[5,3,99.7582]},{"value":[5,2,99.9354]},{"value":[5,5,99.9074]},{"value":[5,7,99.9966]},{"value":[5,9,99.9898]},{"value":[5,4,99.8164]},{"value":[5,6,99.989]},{"value":[5,10,99.9533]},{"value":[13,4,1]},{"value":[13,10,1]},{"value":[13,0,1]},{"value":[13,3,1]},{"value":[13,6,1]},{"value":[13,2,1]},{"value":[13,7,1]},{"value":[13,9,1]},{"value":[13,1,1]},{"value":[13,5,1]},{"value":[13,8,1]},{"value":[12,1,1]},{"value":[12,7,1]},{"value":[12,8,1]},{"value":[12,9,1]},{"value":[12,6,1]},{"value":[12,0,1]},{"value":[12,4,1]},{"value":[12,2,1]},{"value":[12,3,1]},{"value":[12,5,1]},{"value":[12,10,1]},{"value":[8,5,99.9066]},{"value":[8,6,99.9889]},{"value":[8,7,99.9939]},{"value":[8,8,99.9388]},{"value":[8,10,99.957]},{"value":[8,0,98.6412]},{"value":[8,4,99.7949]},{"value":[8,9,99.9879]},{"value":[8,1,99.9101]},{"value":[8,2,99.9103]},{"value":[8,3,99.7585]},{"value":[3,1,99.9506]},{"value":[3,2,99.935]},{"value":[3,5,103.0041]},{"value":[3,7,104.7291]},{"value":[3,10,102.3163]},{"value":[3,0,98.8145]},{"value":[3,3,99.8715]},{"value":[3,6,101.5456]},{"value":[3,8,100.0287]},{"value":[3,4,104.007]},{"value":[3,9,99.9875]},{"value":[0,10,99.9539]},{"value":[0,1,99.8933]},{"value":[0,5,99.9042]},{"value":[0,7,99.9938]},{"value":[0,8,99.946]},{"value":[0,9,99.986]},{"value":[0,0,98.7276]},{"value":[0,3,99.7493]},{"value":[0,4,95.7632]},{"value":[0,6,99.9873]},{"value":[0,2,99.9541]}]}],"title":{"text":"success_rate"},"tooltip":{"show":false},"visualMap":[{"calculable":true,"min":98,"max":100,"range":[98,100],"inRange":{"color":["#313695","#4575b4","#74add1","#abd9e9","#e0f3f8","#fee090","#fdae61","#f46d43","#d73027","#a50026"]}}],"xAxis3D":{"data":["03-23 02:03","03-23 02:08","03-23 02:14","03-23 03:43","03-23 03:47","03-23 03:53","03-23 03:56","03-23 04:06","03-23 04:23","03-23 04:30","03-23 04:33","03-23 07:27","03-23 07:29","03-23 07:33"]},"yAxis3D":{"data":["br","cl","co","id","mx","my","ph","sg","th","tw","vn"]},"zAxis3D":{}};    goecharts_zwhXUTNTOaWM.setOption(option_zwhXUTNTOaWM);</script></body></html>`
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
      emailFrom   = "yuanshan@shopee.com"
      emailHost   = "smtp.shopeemobile.com"
      emailPort   = 587
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
