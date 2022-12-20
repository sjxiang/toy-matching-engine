// package conf

// import (


// 	"github.com/fsnotify/fsnotify"
// 	"github.com/spf13/viper"
// )


// var Cfg = Config{}

// type Config struct {
// 	Log    Log
// 	Server Server
// 	Redis  Redis
// }

// type Log struct {
// 	FileDir  string
// 	FileName string
// 	Prefix   string
// 	Level    string
// }

// type Server struct {
// 	Port  string
// }

// type Redis struct {
// 	Addr string
// }

// func Init() {
// 	viper := viper.New()
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath("./conf/")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}

// 	viper.Unmarshal(&Config)
// 	viper.WatchConfig()
// 	viper.OnConfigChange(func(e fsnotify.Event) {
// 		err = viper.ReadInConfig()
// 		if err == nil {
// 			viper.Unmarshal(&Config)
// 			log.Printf("config: %+v", Config)
// 		} else {
// 			log.Printf("ReadInConfig error: %s", err)
// 		}
// 	})
// }


package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)


var Cfg = Config{}

type Config struct {
	Log    Log      `mapstructure:"log"`
	Server Server   `mapstructure:"server"`
	Redis  Redis    `mapstructure:"redis"`
}

type Server struct {
	Port     string `mapstructure:"port"`
}

type Log struct {
	FileDir  string `mapstructure:"fileDir"`
	FileName string `mapstructure:"fileName"`
	Prefix   string `mapstructure:"prefix"`
	Level    string `mapstructure:"level"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
}

/* 

	*** 温馨小提示 tips

	这里推荐使用 mapstructure 作为序列化标签

	1. yaml 不支持 

		" AppSignExpire int64  `yaml:"app_sign_expire"` " 

		这种下划线的标签

	2. 使用 mapstructure 值得注意的地方是，只要标签中使用了下划线等连接符，":"后就不能有空格。
		比如： 
			AppSignExpire int64  `yaml:"app_sign_expire"`是可以被解析的
			AppSignExpire int64  `yaml: "app_sign_expire"` 不能被解析

*/


// 加载配置，失败直接 panic
func Load() {

	// 1. 初始化 viper 库的实例
	viper := viper.New()

	// 2. 设置配置文件路径
	viper.SetConfigFile("conf/config.yml")

	// 3. 配置读取
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 4. 将配置映射成结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}

	// 5. 监听配置文件变更，重新解析配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {  // 回调
		fmt.Println(e.Name)

		// Again，+1
		if err := viper.Unmarshal(&Cfg); err != nil {
			panic(err)
		}
	})
}
