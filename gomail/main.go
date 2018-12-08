package main

func main() {
	tos := make([]string, 2)
	tos[0] = "yanrui@zhong-ying.com"
	tos[1] = "kurui_alarm@163.com"
	for i := range tos {
		Title := "测试邮件"
		Content := "<html><head><title>邮件标题</title></head><body>AAAAAA<br>BBBBB</body></html>"
		TTankMail(tos[i], Title, Content)
	}

}
