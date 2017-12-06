package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/afanti-com/mobile"
	"os"
	"strings"
)

func main() {
	dataFile := flag.String("d", "../data/mobile.dat", "data file path")
	if err := mobile.Init(*dataFile); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("请输入手机号：")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		no := strings.TrimSpace(scanner.Text())
		mi, err := mobile.Search(no)
		if err != nil {
			fmt.Println(no, err.Error())
			continue
		}
		fmt.Printf("手机号：%s\t省份：%s\t城市：%s\t城市区号：%s\t城市邮编：%s\t运营商：%s\n",
			no, mi.Province, mi.City, mi.CityAreaCode, mi.CityZipCode, mi.Company)
	}
}
