package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xlog"
	"io/ioutil"
	"net/http"
)

var client *alipay.Client

func init() {
	// 初始化支付宝客户端
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	var err error
	client, err = alipay.NewClient("2016092300580007", "MIIEpAIBAAKCAQEAr3Bd9N1WAKwWoVQXZbK5/lfW01+OL46rQU2jgdKmfmnKrckg89LtRxBrkaDGeYaTvYN8w1R8iSeQnSuYpdUQV8ZYhJHxEyWWxdyw4XzqtV3vyhX2pKUcB7b+p5cmgFkqVxjE5ikS2Ehn15+gTqWNr31PL3GMkJlCSUsBxwtsAkOD1Ba4COSLhkMlXX360ypmCHOR0XsmWdpHW2zzPl/jABEFAJJyFpKXHDxUVr3kGJp2VeJib/K33ZZ6FtNjOp6lNsQtax+4W+/LAWP+LmvgUxBd9X+2VJ0E2B0e1B+dYMEx9da/jUOZZbp4R1tZNYNIT8J0roqd0Ta6ri5GubZ/JQIDAQABAoIBAG83Yri6H60pauc28eYsmTAd6zfwWajYXsR1PXQtYBK/ahcla9hmHLlyNV8KdwfPQpVfZ5ZEVvZzuH+KhaiLG2MpgC/6fMNQjvY7433uvLPhKwQHTGYvrUdOwLzpWyT+YZL5RslJvIKT1BCjRvLUjK91Vt+SsjNwlM7119ig6wW7Tn2O+HJZFTxuRR6mbtJAarpI5bCEC5YWqA76XDOwmfnsBIcsRvlGNIsozjjZQXXDA/EJZPfuMbjDS4cOlpGN4UX9kQysSrWthVt/tUcz6Flt4AGIPbMOMAy38gIUnjtnrV+wGb/wAlM3lcmaCLGqfNEc6FgqXOPgJgqt6mVbmvkCgYEA1oZSJacDKOjXo3D6UDRUjjegB7Fv4x75Ehsib8gbMVyacn1FT19vK0M4a/1XYXx9xKHJfbePXzVrfVmFC4iwPUwgo4nCiEu0SFPg80tHfRILH9tneTOxOdwqzKNBgQspxrHA+lJQMNeZn386j6q2Yy9YaM5+VBQw43esSJFdU58CgYEA0VuJZ96sSPXCWvaJRurm3kstoDmkVaQJSKN4M9acKD0NqHEU+cQ7Lra8Ge33ObJUKjdvVO5+ha6IcsYK9XyXceMg5q0dme+4nTqi3q001Iy8YEd4D2l70CRnv7+SdCDPRwnI74mKXBN4n0v3zya8r7fcnFn8jndTcYZqM5KtVrsCgYAzGZ3ZwDW9efFbyl6l2kFF+zQQgRCw3kDgMjx6ugTYNkHIAdUPuiDmAkYfWyKEYz4mJ4+ydZYeTMBFBJzviWjJwSeJLRh8lOPl0q0j5wupXDHwLbgxI8IUkUBUrJK+5WTXIqJgPllBuSmr14egjz6Q08dQd9eZuZGHrYsJYi8D6wKBgQCb2l0MKkRvfCTXurivRqffzG+G2+JurtMKdEdDr/xIKEtGnghTd7XnLLoTQ4fMN+lYFDSZMYKW2jtR1gTKg5nqen8/uphSxJbW1ZPx2Vd+89Y4hnvmXUyCwtdJc7//G32zEKG5UNbjfMtLKy6r09abqXBeJM9S75n65WmiQwsviQKBgQCWJ9KOHh56NxP8kMtEWqvId1Xn3psgnEjKl5VCQ9csEczu3XE8ygZEeoi/1aM1SZ2mueGyWlSI4TnZnYxk5qKNimg3CQwRUlWLyDT7SLygg1cs9myGqC9ou518Fm215l3WofchPKXvUvHAw7iNlZUS0U5WD4MrkERrtropdotNOw==", false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).                                   // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).                                  // 设置签名类型，不设置默认 RSA2
							SetReturnUrl("http://www.xuwuruoshui.club:8080/callback"). // 设置返回URL
							SetNotifyUrl("http://www.xuwuruoshui.club:8080/callback")  // 设置异步通知URL
	//SetAppAuthToken()                   // 设置第三方应用授权

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	fileData, _ := ioutil.ReadFile("./crt/alipayCertPublicKey_RSA2.crt")
	client.AutoVerifySign(fileData)
	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath("./crt/appCertPublicKey.crt", "./crt/alipayRootCert.crt", "./crt/alipayCertPublicKey_RSA2.crt")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	r := gin.Default()

	r.POST("/pay", func(context *gin.Context) {

		response, err := pay()
		if err != nil {
			fmt.Println(err)
		}
		context.JSON(http.StatusOK, response)
	})

	r.GET("/callback", func(c *gin.Context) {
		// 解析异步通知的参数
		//    req：*http.Request
		notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request) // c.Request 是 gin 框架的写法
		fmt.Println(notifyReq)
		if err != nil {
			xlog.Error(err)
			return
		}

		// 支付宝异步通知验签（公钥证书模式）
		//ok, err := alipay.VerifySignWithCert("./crt/alipayCertPublicKey_RSA2.crt", notifyReq)

		// 如果需要，可将 BodyMap 内数据，Unmarshal 到指定结构体指针 ptr
		//err = notifyReq.Unmarshal(ptr)

		// ====异步通知，返回支付宝平台的信息====
		//    文档：https://opendocs.alipay.com/open/203/105286
		//    程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）

		// 此写法是 gin 框架返回支付宝的写法
		//if !ok {
		//	c.String(http.StatusOK, "%s", "faild")
		//	return
		//}
		//
		//c.String(http.StatusOK, "%s", "success")
	})

	r.Run("127.0.0.1:8080")
}

func pay() (interface{}, error) {
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", "条码支付").
		Set("scene", "bar_code").
		Set("auth_code", util.RandomNumber(10)).
		Set("out_trade_no", "My_Pay_"+util.RandomString(10)).
		Set("total_amount", "100").
		Set("timeout_express", "2m")

	aliRsp, err := client.TradePrecreate(context.Background(), bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return "", err
		}
		xlog.Errorf("hello.TradePay(%+v),err:%+v", bm, err)
		return "", err
	}
	return aliRsp, nil
}
