package main

import (
	"errors"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/alibaba/sentinel-golang/util"
	"log"
	"math/rand"
	"time"
)

func main(){
	// AAA()
	// BBB()


}

// 1秒10个请求,多了直接扔掉
func AAA(){
	conf := config.NewDefaultConfig()
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err!=nil{
		panic(err)
	}

	var rules []*flow.Rule
	rule := &flow.Rule{
		// 当前流量控制器的Token计算策略。Direct表示直接使用字段 Threshold 作为阈值；WarmUp表示使用预热方式计算Token的阈值
		Resource:               "test",
		TokenCalculateStrategy: flow.Direct,
		// Reject表示超过阈值直接拒绝，Throttling表示匀速排队。
		ControlBehavior:        flow.Reject,
		// 1000ms内只能有10个
		Threshold:              10,
		StatIntervalInMs:       1000,
	}

	rules = append(rules,rule)
	_, err = flow.LoadRules(rules)
	if err!=nil{
		panic(err)
	}

	for i := 0; i < 20; i++ {
		entry,blockError := sentinel.Entry("test",sentinel.WithTrafficType(base.Inbound))
		if blockError!=nil{
			fmt.Println("流量太大了,开启限流")
		}else{
			fmt.Println("限流通过")
			entry.Exit()
		}
	}
}

// 30s预热达到每秒1000个峰值
func BBB(){
	err := sentinel.InitDefault()
	if err!=nil{
		panic(err)
	}

	var all,through,block int

	var rules []*flow.Rule
	rule := &flow.Rule{
		Resource:               "test",
		TokenCalculateStrategy: flow.WarmUp,
		// Reject表示超过阈值直接拒绝，Throttling表示匀速排队。
		ControlBehavior:        flow.Reject,
		// 30s预热达到每秒1000个峰值
		Threshold:              1000,
		WarmUpPeriodSec:       30,
		// 最大等待时间
		MaxQueueingTimeMs: 500,
	}

	rules = append(rules, rule)
	_, err = flow.LoadRules(rules)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		go func() {

			for {
				all++
				entry,blockError := sentinel.Entry("test",sentinel.WithTrafficType(base.Inbound))
				if blockError!=nil{
					//fmt.Println("流量太大了,开启限流")
					block++
					time.Sleep(time.Second)
				}else{
					//fmt.Println("限流通过")
					through++
					entry.Exit()
				}
			}
		}()
	}



	go func() {
		var oldAll,oldThrough,oldBlock int
		for{
			a := all-oldAll
			oldAll = all

			t := through-oldThrough
			oldThrough = through

			b := block-oldBlock
			oldBlock = block
			time.Sleep(time.Second)
			fmt.Println(a,"-",t,"-",b)
		}
	}()

	select {

	}
}


type stateChangeTestListener struct {
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Closed, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	fmt.Printf("rule.steategy: %+v, From %s to Open, snapshot: %.2f, time: %d\n", rule.Strategy, prev.String(), snapshot, util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}
func CCC(){
	conf := config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan struct{})
	// Register a state change listener so that we could observer the state change of the internal circuit breaker.
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})

	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		// Statistic time span=5s, recoveryTimeout=3s, maxErrorRatio=40%
		{
			Resource:                     "abc",
			// 按错误比例
			Strategy:                     circuitbreaker.ErrorRatio,
			// 熔断持续时间
			RetryTimeoutMs:               3000,
			// 触发熔断的最小请求数目，若当前统计周期内的请求数小于此值，即使达到熔断条件规则也不会触发
			MinRequestAmount:             10,
			// 统计的时间窗口长度（单位为 ms）
			StatIntervalMs:               5000,
			//
			StatSlidingWindowBucketCount: 10,
			// 当前资源的慢调用比例如果高于Threshold，那么熔断器就会断开；否则保持闭合状态
			Threshold:                    0.4,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	logging.Info("[CircuitBreaker ErrorRatio] Sentinel Go circuit breaking demo is running. You may see the pass/block metric in the metric log.")
	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g1 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				if rand.Uint64()%20 > 6 {
					// Record current invocation as error.
					sentinel.TraceError(e, errors.New("biz error"))
				}
				// g1 passed
				time.Sleep(time.Duration(rand.Uint64()%80+20) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	go func() {
		for {
			e, b := sentinel.Entry("abc")
			if b != nil {
				// g2 blocked
				time.Sleep(time.Duration(rand.Uint64()%20) * time.Millisecond)
			} else {
				// g2 passed
				time.Sleep(time.Duration(rand.Uint64()%80+40) * time.Millisecond)
				e.Exit()
			}
		}
	}()
	<-ch
}