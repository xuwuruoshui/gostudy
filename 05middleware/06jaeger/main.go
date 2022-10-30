package main

import (
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

/**
* @creator: xuwuruoshui
* @date: 2022-10-30 15:39:17
* @content: jeager使用, 一个span就是一个模块
 */

func main() {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.0.132:6831",
		},
		ServiceName: "testMall",
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	defer closer.Close()
	if err != nil {
		panic(err)
	}

	orderSpan := tracer.StartSpan("order_srv")

	cartSpan := tracer.StartSpan("cart_srv", opentracing.ChildOf(orderSpan.Context()))
	time.Sleep(time.Second)
	cartSpan.Finish()

	productSpan := tracer.StartSpan("product_srv", opentracing.ChildOf(orderSpan.Context()))
	time.Sleep(2 * time.Second)
	productSpan.Finish()

	orderSpan.Finish()
}
