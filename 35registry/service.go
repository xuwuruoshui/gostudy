package registry

// 服务节点
type Service struct{

	// 服务名
	Name string `json:"name"`

	// 节点列表
	Node []*Node `json:"nodes"`

}


// 单个服务节点
type Node struct{
	Id string	`json:"id"`
	Ip string `json:"ip"`
	Port int	`json:"port"`
	Weight int	`json:"weight"`
}
