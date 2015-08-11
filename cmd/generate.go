package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files := getFiles("/home/gaspiman/Downloads/universal-pos-tags-master/", ".map")
	posMap := map[string]map[string]string{}
	for k := range files {
		extract(files[k], posMap)
	}
	res2B, err := json.Marshal(posMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(res2B))
	//fmt.Println(posMap)
}

func getFiles(folder string, fileSuffix string) []string {
	resp := []string{}
	visit := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, fileSuffix) {
			return nil
		}
		resp = append(resp, path)
		return nil
	}
	err := filepath.Walk(folder, visit)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func extract(path string, posMap map[string]map[string]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g := strings.Split(path, "/")
	tagSet := g[len(g)-1]
	tagSet = strings.TrimSuffix(tagSet, ".map")
	if _, ok := posMap[tagSet]; !ok {
		posMap[tagSet] = map[string]string{}
	}

	r := bufio.NewReaderSize(file, 500*1024)
	for {
		lineByte, isPrefix, err := r.ReadLine()
		if err != nil {
			break
		}
		if isPrefix != false {
			continue
		}
		line := string(lineByte)
		tokens := strings.Split(line, "	")
		if len(tokens) == 1 {
			tokens = strings.Split(tokens[0], " ")
		}
		posMap[tagSet][tokens[0]] = tokens[1]
	}
}
