/*
   @Time : 2020-01-17 15:00
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : logger
   @Software: GoLand
*/

package logger

import (
	"QCloudDDNS/src/configer"
	"fmt"
	"os"
)

// 错误检查
func Check(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
		if configer.Config.Log != "" {
			writeStringToFile(configer.Config.Log, fmt.Sprintf("Error : %v", err))
		}
	}
}

func CheckWithoutSave(err error) {
	if err != nil {
		fmt.Println("Error : ", err)
	}
}

func Log(log string) {
	fmt.Println("Log : ", log)
	if configer.Config.Log != "" {
		writeStringToFile(configer.Config.Log, "Log : "+log)
	}
}

func Debug(debug string) {
	if configer.Config.Debug {
		fmt.Println("Debug : ", debug)
		if configer.Config.Log != "" {
			writeStringToFile(configer.Config.Log, "Debug : "+debug)
		}
	}
}

//写入文件
func writeStringToFile(filepath, content string) {
	//打开文件，没有则创建，有则append内容
	w1, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	CheckWithoutSave(err)

	_, err1 := w1.Write([]byte(content + "\r\n")) // \r\n 为Windows换行符， 其他系统为 \n
	CheckWithoutSave(err1)

	err2 := w1.Close()
	CheckWithoutSave(err2)
}
