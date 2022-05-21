package config

// init文件所有字段都需要对应上
type LogConfig struct{
	*KafkaConfig `ini:"kafka"`
	*EtcdConfig `ini:"etcd"`
}

type KafkaConfig struct{
	Address string `ini:"address"`
	MaxSize int `ini:"chan_max_size"`
	//Topic string `ini:"topic"`
}

type EtcdConfig struct{
	Address []string `ini:"address"`
	Timeout int `ini:"timeout"`
	Key string `ini:"key"`
}

//type TailLogConfig struct{
//	Path string `ini:"path"`
//}
