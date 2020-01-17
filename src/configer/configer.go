/*
   @Time : 2020-01-17 14:59
   @Author : Rebeta
   @Email : master@rebeta.cn
   @File : configer
   @Software: GoLand
*/

package configer

import (
	"fmt"
	"github.com/aWildProgrammer/fconf"
	"strings"
)

// 配置结构体
type config struct {
	IPv4  bool
	IPv6  bool
	Round int
	Log   string
	Debug bool

	QCloudSecretId  string
	QCloudSecretKey string

	IPv4Domain    string
	IPv4SubDomain string

	IPv6Domain    string
	IPv6SubDomain string
}

var Config = config{}

func init() {
	if config, err := fconf.NewFileConf("./QCloudDDNS.ini"); err != nil {
		panic(err)
	} else {
		if strings.ToLower(config.String("Common.IPv4")) == "true" {
			Config.IPv4 = true
		} else {
			Config.IPv4 = false
		}
		if strings.ToLower(config.String("Common.IPv6")) == "true" {
			Config.IPv6 = true
		} else {
			Config.IPv6 = false
		}
		if round, err := config.Int("Common.Round"); err != nil {
			// logger.Debug(fmt.Sprint(err)) // import cycle not allowed
			Config.Round = 120
		} else {
			Config.Round = round
		}
		Config.Log = config.String("Common.Log")
		if strings.ToLower(config.String("Common.Debug")) == "true" {
			Config.Debug = true
		} else {
			Config.Debug = false
		}
		Config.QCloudSecretId = config.String("QCloud.SecretId")
		if Config.QCloudSecretId == "" {
			panic(fmt.Errorf("未配置 QCloud.SecretId"))
		}
		Config.QCloudSecretKey = config.String("QCloud.SecretKey")
		if Config.QCloudSecretKey == "" {
			panic(fmt.Errorf("未配置 QCloud.SecretKey"))
		}
		Config.IPv4Domain = config.String("Domain.IPv4Domain")
		if Config.IPv4 && Config.IPv4Domain == "" {
			panic(fmt.Errorf("未配置 Domain.IPv4Domain"))
		}
		Config.IPv4SubDomain = config.String("Domain.IPv4SubDomain")
		if Config.IPv4 && Config.IPv4SubDomain == "" {
			panic(fmt.Errorf("未配置 Domain.IPv4SubDomain"))
		}
		Config.IPv6Domain = config.String("Domain.IPv6Domain")
		if Config.IPv6 && Config.IPv6Domain == "" {
			panic(fmt.Errorf("未配置 Domain.IPv6Domain"))
		}
		Config.IPv6SubDomain = config.String("Domain.IPv6SubDomain")
		if Config.IPv6 && Config.IPv6SubDomain == "" {
			panic(fmt.Errorf("未配置 Domain.IPv6SubDomain"))
		}
	}
}
