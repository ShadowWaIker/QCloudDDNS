/*
   @Time : 2020-01-17 15:00
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : qcloud
   @Software: GoLand
*/

package qcloud

import (
	"QCloudDDNS/src/configer"
	"QCloudDDNS/src/logger"
	"QCloudDDNS/src/utils"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func sign(method string, api string, params map[string]string, signType string) (dataPost string, err error) {
	timeStamp := time.Now().Unix()
	// 添加公共部分
	params["Timestamp"] = strconv.FormatInt(timeStamp, 10)
	rand.Seed(time.Now().UnixNano())
	params["Nonce"] = fmt.Sprintf("%v", rand.Int())
	params["SecretId"] = configer.Config.QCloudSecretId

	// 对参数的下标进行升序排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var dataParams string
	// 拼接参数
	for _, k := range keys {
		logger.Debug("Key -> " + k + "\tValue ->" + params[k])
		dataParams += k + "=" + params[k] + "&"
	}

	logger.Debug("修正前：" + dataParams)
	dataParams = dataParams[0 : len(dataParams)-1] // 去掉拼接后的参数最后的一个"&"
	logger.Debug("修正后：" + dataParams)

	var mac hash.Hash
	// 对字符串进行加密
	switch signType {
	case "HmacSHA1":
		mac = hmac.New(sha1.New, []byte(configer.Config.QCloudSecretKey))
	case "HmacSHA256":
		mac = hmac.New(sha256.New, []byte(configer.Config.QCloudSecretKey))
	default:
		return "", errors.New("加密参数设置有误") // 返回错误
	}

	mac.Write([]byte(strings.ToUpper(method) + api + ".api.qcloud.com/v2/index.php?" + dataParams))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	logger.Debug("Sign Before URLEncode -> " + sign)
	sign = url.QueryEscape(sign)
	logger.Debug("Sign After URLEncode -> " + sign)

	return dataParams + "&Signature" + "=" + sign, nil
}

/*
 * Cloud Name Service CNS 腾讯云解析 签名
 */
func SignCNS(params map[string]string) (response string, err error) {
	return sign("get", "cns", params, "HmacSHA1")
}

/*
 * Cloud Name Service CNS RecordList 获取域名的解析记录
 */
func CNSRecordList(domain, subDomain string) (map[string]interface{}, error) {
	apiURL := "https://cns.api.qcloud.com/v2/index.php?"
	params := map[string]string{
		"Action": "RecordList",
		"domain":    domain,
		"subDomain": subDomain,
	}
	if paramsStr, err := sign("get", "cns", params, "HmacSHA1"); err != nil {
		logger.Check(err)
		logger.Debug("paramsStr -> " + paramsStr)
		return nil, err
	} else {
		logger.Debug("paramsStr -> " + paramsStr)
		return utils.GetJson(apiURL + paramsStr)
	}
}

/*
 * Cloud Name Service CNS RecordModify 修改解析记录
 */
func CNSRecordModify(domain, recordId, subDomain, recordType, value, ttl string) (map[string]interface{}, error) {
	apiURL := "https://cns.api.qcloud.com/v2/index.php?"
	params := map[string]string{
		"Action":     "RecordModify",
		"domain":     domain,
		"recordId":   recordId,
		"subDomain":  subDomain,
		"recordType": recordType,
		"recordLine": "默认",
		"value":      value,
		"ttl":        ttl,
	}
	if paramsStr, err := sign("get", "cns", params, "HmacSHA1"); err != nil {
		logger.Check(err)
		logger.Debug("paramsStr -> " + paramsStr)
		return nil, err
	} else {
		logger.Debug("paramsStr -> " + paramsStr)
		return utils.GetJson(apiURL + paramsStr)
	}
}
