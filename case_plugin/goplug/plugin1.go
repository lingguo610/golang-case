/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-07-28 16:02:41
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-07-28 16:30:48
 * @FilePath: \goproject\goplug\plugin1.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"log"

	"common"
)

func init() {
	log.Println("plugin1 init")

	common.F2()
}

func F() {
	fmt.Println("plugin1 : xxxx")
}

//go build -buildmode=plugin -o plugin1.so plugin1.go
