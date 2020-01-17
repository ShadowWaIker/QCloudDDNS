/*
   @Time : 2020-01-17 14:03
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : main
   @Software: GoLand
*/

package main

import (
	"QCloudDDNS/src/configer"
	"QCloudDDNS/src/logger"
	"QCloudDDNS/src/qcloud"
	"QCloudDDNS/src/utils"
	"fmt"
	"strconv"
	"time"
)

func main() {
	logger.Log(time.Unix(time.Now().Unix(), 0).String() + " -> 开始运行")
	for range time.Tick(time.Millisecond * 1000 * time.Duration(configer.Config.Round)) {
		if configer.Config.IPv4 {
			handler("4")
		}
		if configer.Config.IPv6 {
			handler("6")
		}
	}
}

func handler(version string) {
	if ip := utils.GetGatewayIP(version); ip == "" {
		logger.Log(time.Unix(time.Now().Unix(), 0).String() + " -> 获取 IPv" + version + " 地址失败")
	} else {
		if CNSRecord, err := qcloud.CNSRecordList(configer.Config.IPv4Domain, configer.Config.IPv4SubDomain); err != nil {
			logger.Log(time.Unix(time.Now().Unix(), 0).String() + " -> 获取 IPv" + version + " 解析记录失败")
		} else {
			// 判断记录中的 ip 与当前ip 是否相符
			if ip == CNSRecord["data"].(map[string]interface{})["records"].([]interface{})[0].(map[string]interface{})["value"].(string) {
				// logger.Log(time.Unix(time.Now().Unix(), 0).String() + " -> IPv"+version+" 记录未发生变化")
			} else {
				recordType := "A"
				if version != "4" {
					recordType = "AAAA"
				}
				if response, err := qcloud.CNSRecordModify(configer.Config.IPv4Domain, strconv.FormatFloat(CNSRecord["data"].(map[string]interface{})["records"].([]interface{})[0].(map[string]interface{})["id"].(float64), 'f', -1, 64), configer.Config.IPv4SubDomain, recordType, ip, strconv.FormatFloat(CNSRecord["data"].(map[string]interface{})["domain"].(map[string]interface{})["min_ttl"].(float64), 'f', -1, 64)); err != nil {
					logger.Log(time.Unix(time.Now().Unix(), 0).String() + " -> 修改 IPv" + version + " 解析失败")
				} else {
					logger.Debug(fmt.Sprint(response))
					if response["code"] == 0 {
						logger.Log(time.Unix(time.Now().Unix(), 0).String() + " -> 修改 IPv" + version + " 解析失败 [ " + fmt.Sprint(response["code"]) + " ] ( " + fmt.Sprint(response["codeDesc"]) + " ) ")
					} else {
						logger.Log(time.Unix(time.Now().Unix(), 0).String() + " -> 修改 IPv" + version + " 解析成功 ( " + CNSRecord["data"].(map[string]interface{})["records"].([]interface{})[0].(map[string]interface{})["value"].(string) + " -> " + ip + " )")
					}
				}
			}
		}
	}
}
