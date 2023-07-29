package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

/**
* @creator: xuwuruoshui
* @date: 2022-06-16 14:41:00
* @content:
 */

var client *wechat.ClientV3

func init() {

	PrivateKey, _ := ioutil.ReadFile("./crt/apiclient_key.pem")
	var err error
	client, err = wechat.NewClientV3("1588230631", "690F2A694D2B6EFA2E4AEFE76E652CFC61867524", "BSNDcbbL5r122gEdoCSd9tGNIPNvjSxx", string(PrivateKey))
	if err != nil {
		xlog.Error(err)
		return
	}

	// 设置微信平台API证书和序列号（如开启自动验签，请忽略此步骤）
	//hello.SetPlatformCert([]byte(""), "")

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		xlog.Error(err)
		return
	}

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn
}

func main() {
	r := gin.Default()

	// 1、获取code
	//r.GET("/wechat/auth", Auth)

	// 2、授权回调
	r.POST("/wechat/code/", CodeCallBack)

	// 1、购买
	r.POST("/wechat/pay", Pay)

	// 2、获取订单状态
	//r.POST("/wechat/status",)

	r.Run(":7902")
}

func Pay(ctx *gin.Context) {

	tradeNo := "sku3d_" + strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(tradeNo)

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", "wxb322c2d277046d02").
		Set("mchid", "1588230631").
		//Set("sp_appid", "wxb322c2d277046d02").
		//Set("sp_mchid", "sp_mchid").
		//Set("sub_mchid", "sub_mchid").
		Set("description", "测试Jsapi支付商品").
		Set("out_trade_no", tradeNo).
		Set("time_expire", time.Now().Add(10*time.Minute).Format(time.RFC3339)).
		Set("notify_url", "http://www.sku3d.com:8081/wechat/code/").
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", 1).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3TransactionNative(context.Background(), bm)
	if err != nil {
		xlog.Error(err)
		return
	}

	if wxRsp.Code == http.StatusOK {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)

	ctx.JSON(http.StatusOK, wxRsp)
}

func CodeCallBack(ctx *gin.Context) {
	notifyReq, err := wechat.V3ParseNotify(ctx.Request)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := notifyReq.DecryptCipherText("apiv3key")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.TradeState)
	fmt.Println(result.OutTradeNo)
}

func Auth(ctx *gin.Context) {
	apiUrl := "https://open.weixin.qq.com/connect/oauth2/authorize"

	data := url.Values{}
	data.Set("appid", "wxb322c2d277046d02")
	data.Set("redirect_uri", "www.sku3d.com/wechat/code")
	data.Set("response_type", "code")
	data.Set("scope", "snsapi_base")
	data.Set("state", "123#wechat_redirect")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())

	// resp, err := http.Get(u.String())
	// if err != nil {
	// 	fmt.Printf("post failed, err:%v\n", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// b, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("get resp failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Println(string(b))
	// ctx.String(http.StatusOK, "", string(b))
	ctx.String(http.StatusOK, u.String())
}
