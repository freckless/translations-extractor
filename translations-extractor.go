package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	open_tag := "__\\('"
	close_tag := "'\\)"
	save_file := "language.json"

	var _tmp_string string
	fmt.Print("Open tag [" + open_tag + "]: ")
	fmt.Scanf("%s", &_tmp_string)
	if len(_tmp_string) > 0 {
		open_tag = _tmp_string
	}
	fmt.Print("Close tag [" + close_tag + "]: ")
	fmt.Scanf("%s", &_tmp_string)
	if len(_tmp_string) > 0 {
		close_tag = _tmp_string
	}
	fmt.Print("Output file [" + save_file + "]: ")
	fmt.Scanf("%s", &_tmp_string)
	if len(_tmp_string) > 0 {
		save_file = _tmp_string
	}

	regexp_tag := regexp.MustCompile(open_tag + "(.*?)" + close_tag)

	if len(os.Args) > 1 {
		cwd, _ := os.Getwd()
		fmt.Println(cwd)
		files := os.Args[1:]
		translations := make(map[string]string)

		for _, file := range files {
			data, _ := ioutil.ReadFile(cwd + "/" + file)
			matches := regexp_tag.FindAll(data, -1)
			for _, match := range matches {
				key := regexp_tag.ReplaceAllString(string(match), "$1")
				translations[key] = key
			}
			json_data, _ := json.Marshal(translations)
			var out bytes.Buffer
			json.Indent(&out, json_data, "", "  ")
			f, _ := os.Create(save_file)
			w := bufio.NewWriter(f)
			out.WriteTo(w)
			w.Flush()
		}
	} else {
		fmt.Println("No files passed")
	}
}
