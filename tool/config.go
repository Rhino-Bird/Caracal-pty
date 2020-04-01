package tool

import (
	"fmt"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	// Config returns an initialized Viper instance
	Config = viper.New()
)

// InitConfig initialize the configuration file
func InitConfig() {
	var err error

	Config.AddConfigPath(ConfDir)

	Config.SetConfigName("config")
	err = Config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file config: %s", err.Error()))
	}

	Config.SetConfigName("log")
	err = Config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file log: %s", err.Error()))
	}

	Config.SetConfigName("base")
	err = Config.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file log: %s", err.Error()))
	}

	watchON := Config.GetBool("base.notify_watch_config")
	if watchON {
		WatchConfig()
	}

	Log.Info("Initialize the configuration file")
}

// WatchConfig automatic discovery configuration
func WatchConfig() {
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			Log.Errorf("config watch init errror, err: %s", err.Error())
		}
		defer watcher.Close()

		confPath := "conf/yaml"
		watcher.Add(confPath)
		Log.Infof("viper start watch config notify, path: %s", confPath)

		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					filename := path.Base(event.Name)
					fileType := path.Ext(filename)
					filenameOnly := strings.TrimSuffix(filename, fileType)

					Config.SetConfigName(filenameOnly)
					err := Config.MergeInConfig() // 合并配置
					if err != nil {
						Log.Errorf("read config error, err: %s", err.Error())
					}
					Log.Infof("config file changed, file_name: %s", event.Name)
				}
			case err := <-watcher.Errors:
				if err != nil {
					Log.Errorf("watch error, err: %s", err.Error())
				}
			}
		}
	}()
}
