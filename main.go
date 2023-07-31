package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	Template()
}

func Template() {
	t, _ := template.ParseFiles("./README.tmpl.md")

	// f, err3 := os.OpenFile("./README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	// if err3 != nil {
	// 	fmt.Printf("err3: %v\n", err3)
	// 	return
	// }
	// t.Execute(f, c)
	// b, _ := os.ReadFile("./v2ray.txt")
	// // fmt.Printf("b: %v\n", b)
	// s := strings.Split(string(b), "\n")
	// var builder strings.Builder
	// for i := 1; i < 50; i = i + 2 {
	// 	builder.WriteString(s[i])
	// 	builder.WriteString("\n")
	// }

	s := readLine(50, "./v2ray.txt")

	f, err3 := os.OpenFile("README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err3 != nil {
		fmt.Printf("err3: %v\n", err3)
		return
	}

	//加载v2ray
	t.Execute(f, map[string]any{
		"v2ray": s,
	})
}

func readLine(count int, fileName string) string {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return ""
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	var builder strings.Builder
	for i := 0; i < count; i++ {
		a, _, _ := br.ReadLine()
		if i%2 == 1 {
			builder.WriteString(string(a))
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
