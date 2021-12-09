package conf

// 全局配置
type LogTransfer struct {
	KafkaConf `ini:"kafka"`
	EsConf `ini:"es"`
}

type KafkaConf struct {
	Address []string `ini:"address"`
	Topic string `ini:"topic"`
}


type EsConf struct {
	Address string `ini:"address"`
	MaxSize int64`ini:"maxSize"`
}
