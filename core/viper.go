package core

import (
	"BookRecSystem/config"
	"BookRecSystem/global"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var cfg string
	if len(path) == 0 {
		flag.StringVar(&cfg, "c", "", "choose cfg file.")
		flag.Parse()
		if cfg == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(config.ConfigEnv); configEnv == "" {
				cfg = config.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config.ConfigFile)
			} else {
				cfg = configEnv
				fmt.Printf("您正在使用GSD_CONFIG"+
					""+
					"环境变量,config的路径为%v\n", cfg)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", cfg)
		}
	} else {
		cfg = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", cfg)
	}

	v := viper.New()
	v.SetConfigFile(cfg)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error cfg file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("cfg file changed:", e.Name)
		if err := v.Unmarshal(&global.GSD_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GSD_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.GSD_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
