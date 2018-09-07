package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

const inputDir = "./input"
const outputDir = "./output"

func main() {
	p, _ := os.Getwd()
	fmt.Println(p)

	garbages, _ := ioutil.ReadDir("./")
	for _, v := range garbages {
		if v.IsDir() {
			continue
		}
		name := v.Name()
		if strings.Contains(name, ".txt") {
			os.Remove(name)
		}
	}

	fmt.Println("ストリームKEYを入力してください")

	var key string
	fmt.Scan(&key)

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf("%s/%s.mp4", outputDir, key)
	input := key + ".txt"

	inputFile, err := os.Create(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("-------------------------")
	for _, v := range files {
		name := v.Name()
		if !strings.Contains(name, ".mp4") {
			continue
		}
		path := inputDir + "/" + name
		inputFile.Write([]byte(fmt.Sprintf("file '%s'\n", path)))
		fmt.Printf("%s\n", name)
	}
	inputFile.Close()
	fmt.Println("-------------------------")

	fmt.Printf("この順番で連結し %s を作成します、よろしいですか？ (yes/no)\n", output)
	var ok string
	fmt.Scan(&ok)
	if ok != "yes" {
		os.Remove(input)
		fmt.Println("作業を中止しました")
		return
	}
	fmt.Println("しばらく時間がかかります、少々お待ちください")

	cmd := exec.Command(
		"ffmpeg", "-y", "-f", "concat", "-safe", "-1", "-i", input, "-c", "copy", "-bsf:a", "aac_adtstoasc", output,
	)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	os.Remove(input)
	fmt.Println("正常に結合が完了しました")
}
