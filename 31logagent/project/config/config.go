package config

// init文件所有字段都需要对应上
type LogConfig struct{
	*KafkaConfig `ini:"kafka"`
	*TailLogConfig `ini:"taillog"`
}

type KafkaConfig struct{
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type TailLogConfig struct{
	Path string `ini:"path"`
}
