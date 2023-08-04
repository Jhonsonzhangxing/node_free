package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/sanzhang007/webgin/models"
)

var line int
var file string

func init() {
	flag.IntVar(&line, "l", 100, "line nums")
	flag.StringVar(&file, "f", "v2ray.txt", "file")
	flag.Parse()
}

func main() {
	Template()
}

func Template() {
	t, _ := template.ParseFiles("README.tmpl.md")
	s := readLine(line, file)
	var count int64
	var availableCount int64
	var availableCount1 int64
	var availableCount5 int64
	models.DB.Model(&models.Nodes{}).Count(&count)
	models.DB.Model(&models.Nodes{}).Where("avg_speed > ?", 1).Count(&availableCount)
	models.DB.Model(&models.Nodes{}).Where("avg_speed > ?", 1024*1024).Count(&availableCount1)
	models.DB.Model(&models.Nodes{}).Where("avg_speed > ?", 1024*1024*5).Count(&availableCount5)
	f, err3 := os.OpenFile("README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err3 != nil {
		fmt.Printf("err3: %v\n", err3)
		return
	}
	//加载v2ray
	t.Execute(f, map[string]any{
		"v2ray":           s,
		"count":           count,
		"availableCount":  availableCount,
		"availableCount1": availableCount1,
		"availableCount5": availableCount5,
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
