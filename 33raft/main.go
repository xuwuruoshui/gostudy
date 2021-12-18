package main

import (
	"fmt"
	"gostudy/33raft/utils"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 1.实现3节点选举
// 2.改造为分布式选举代码,加入RPC调用

// 定义3节点常量
const raftCount = 3

type Leader struct {

	// 任期
	Term int

	// LeaderId编号
	LeaderId int
}

type Raft struct {
	// 锁
	mu sync.Mutex
	// 节点编号
	me int
	// 当前任期
	currentTerm int
	// 为那个节点投票
	votedFor int
	// 0 follower 1 candidate 2leader
	state int
	// 发送最后一条数据的时间
	lastMessageTime int64
	// 当前leader是谁
	currentLeader int
	// 节点发送信息的通道
	message chan bool
	// 选举通道
	electCh chan bool
	// 心跳信号的通道
	heartBeat chan bool
	// 返回心跳信号的通道
	heartbeatRe chan bool
	// 超时时间
	timeout int
}

// 0还没上任 -1没有编号
var leader = &Leader{0,-1}

func NewRaft(me int) *Raft{
	// 随机种子
	rand.Seed(time.Now().UnixNano())
	r := &Raft{
		me: me,
		// -1代表谁都不投，表示此节点刚创建
		votedFor:-1,
		// 0 follower
		state:0,
		timeout: 0,
		currentLeader: -1,
		currentTerm: 0,
		message: make(chan bool),
		electCh: make(chan bool),
		heartBeat: make(chan bool),
		heartbeatRe: make(chan bool),
	}
	// 选举协程
	go r.election()

	// 心跳检测携程
	go r.sendLeaderHeartBeat()
	return r
}


func (r *Raft)election(){
	// 设置标记，判断是否选出了leader
	var flag bool

	for  {
		// 设置超时
		timeout := utils.RandRange(150,300)
		r.lastMessageTime = utils.Millisecond()
		select {
		// 延迟等待1毫秒
		case<-time.After(time.Duration(timeout)*time.Millisecond):
		fmt.Println("当前节点状态为:",r.state)
		default:
		}
		flag = false
		for !flag{
			// leader选举逻辑
			flag = r.election_one_round(leader)
		}
	}
}

// leader选举逻辑
func (r *Raft)election_one_round(leader *Leader)bool {
	// 定义超时
	var timeout int64 = 100

	//投票数量
	var vote int

	// 是否有心跳信号
	var triggerHeartbeat bool
	
	// 时间
	last := utils.Millisecond()

	// 是否当选
	isEleted := false

	// 节点变为cadidate
	r.mu.Lock()
	// 修改状态
	r.becomeCandidate()
	r.mu.Unlock()
	fmt.Println("start electing leader")
	for{
		// 遍历所有节点拉选票
		for i:=0;i<raftCount;i++{
			if i!=r.me{
				// 拉选票
				go func() {
					if leader.LeaderId<0{
						// 去拉票
						r.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量,自己成为候选人，本身就有1票
		vote = 1
		// 遍历
		for i:=0;i<raftCount;i++{
			// 计算投票数量
			select{
			case ok:=<-r.electCh:
				if ok{
					// 投票数量+1
					vote++
					isEleted = vote > raftCount/2
					if isEleted && !triggerHeartbeat{
						// 变为leader
						r.mu.Lock()
						r.becomeLeader()
						r.mu.Unlock()

						// 触发心跳检测
						triggerHeartbeat = true

						// leader向其他节点发送心跳信号
						r.heartBeat <- true
						fmt.Println(r.me,"号节点: 成为leader")
						fmt.Println(r.me,"leader发送心跳信号")

					}
				}
			}
		}

		// 做最后校验工作
		// 若不超时,且票数大于一半,则选举成功, break
		if timeout+last<utils.Millisecond() || vote>raftCount/2 || r.currentLeader>-1{
			break
		}else{
			// 等待操作
			select{
			case <-time.After(time.Duration(10)*time.Millisecond):
			}
		}
	}
	return isEleted
}

// leader节点发送心跳
// 顺便完成数据同步
// 看看folloer是否挂了
func(r *Raft) sendLeaderHeartBeat(){
	for{
		select {
		case <-r.heartBeat:
			r.sendAppendEntries()
		}
	}
}

// 用于返回给leader的确认信号
func (r *Raft)sendAppendEntries(){
	// 是leader就不确认了
	if r.currentLeader == r.me{
		// 记录确认信号的节点个数
		var success_count = 0
		// 设置确认信号
		for i:=0;i<raftCount;i++{
			if i!=r.me{
				go func() {
					r.heartbeatRe<-true
				}()
			}
		}

		// 计算返回确认信号个数
		for i:=0;i<raftCount;i++{
			select {
			case ok:=<-r.heartbeatRe:
				if ok{
					success_count++
					if success_count>raftCount/2{
						fmt.Println("投票选举结束")
						log.Fatalln("程序结束")
					}
				}
			}
		}
	}
}

// 变为候选人
func (r *Raft)becomeCandidate(){
	r.state=1
	r.currentTerm++
	r.votedFor = r.me
	r.currentLeader = -1
}

// 变为candidate
func (r *Raft)becomeLeader(){
	r.state=2
	r.currentLeader = r.me

}


func main(){
	// 3个节点，最初都是follower
	// 存在candidate状态，进行投票拉票
	// 选出leader

	// 创建3个节点
	for i:=0;i<raftCount;i++{
		// 创建3个raft节点
		NewRaft(i)
	}


	time.Sleep(time.Second*5)


}
