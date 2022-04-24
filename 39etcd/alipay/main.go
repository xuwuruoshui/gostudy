package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
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
	client.AutoVerifySign([]byte("alipayCertPublicKey_RSA2 bytes"))

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath("39etcd/crt/appCertPublicKey.crt", "39etcd/crt/alipayRootCert.crt", "39etcd/crt/alipayCertPublicKey_RSA2.crt")
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

	r.POST("/callback", func(c *gin.Context) {
		fmt.Println("1")
		fmt.Println(c)
		fmt.Println("1")
	})

	r.Run("127.0.0.1:8080")
}

func pay() (string, error) {
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", "条码支付").
		Set("scene", "bar_code").
		Set("auth_code", "286248566432274952").
		Set("out_trade_no", "GZ201909081745431345").
		Set("total_amount", "74").
		Set("timeout_express", "2m")

	aliRsp, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return "", err
		}
		xlog.Errorf("client.TradePay(%+v),err:%+v", bm, err)
		return "", err
	}
	return aliRsp, nil
}
