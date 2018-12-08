package main

type TestMailObj struct {
	Name string
	URL  string
}

func main() {
	mail := &TestMailObj{
		Name: "RiboseYim",
		URL:  "https://riboseyim.github.io",
	}
	tos := make([]string, 2)
	tos[0] = "yanrui@zhong-ying.com"
	tos[1] = "kurui_alarm@163.com"
	for i := range tos {
		if tos[i] != "" {
			r := NewMailRequest(tos[i], "Hello World!")
			r.ParseTemplate("./template/Test.html", mail)
			r.SendEmail()
		}

	}

}
