/*
   @Time : 2020-01-17 16:08
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : utils
   @Software: GoLand
*/

package utils

import (
	"QCloudDDNS/src/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetJson(URL string) (responseJsonMap map[string]interface{}, err error) {
	if res, err := http.Get(URL) ; err != nil {
		logger.Check(err)
		return nil, err
	} else {
		//defer logger.Check(res.Body.Close())
		if body, err := ioutil.ReadAll(res.Body); err != nil {
			logger.Check(err)
			return nil, err
		} else {
			logger.Debug(string(body))
			// 解码 body
			var responseJsonMap map[string]interface{}
			if err := json.Unmarshal(body, &responseJsonMap); err != nil {
				logger.Check(err)
				return nil, err
			} else {
				// 返回请求回来的 Json 的 Map
				return responseJsonMap, nil
			}
		}
	}
}

func GetGatewayIP(version string) string {
	if version != "4" {
		version = "6"
	}
	if responseJsonMap, err := GetJson("http://ipv"+version+".lookup.test-ipv6.com/ip/"); err != nil {
		logger.Check(err)
		 return ""
	} else {
		if responseJsonMap["ip"] != nil && len(responseJsonMap["ip"].(string)) > 1 {
			return responseJsonMap["ip"].(string)
		} else {
			return ""
		}
	}
}
