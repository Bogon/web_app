package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序的全部配置
var Conf = new(WebAppConfig)

// WebAppConfig is a struct that contains the following fields: `AppConf`, `LogConf`, `MySQLConf`, and `RedisConf`.
// @property {AppConf}  - AppConf: Application configuration
// @property {LogConf}  - AppConf: Application configuration
// @property {MySQLConf}  - AppConf: Application configuration
// @property {RedisConf}  - AppConf: Application configuration
type WebAppConfig struct {
	*AppConf   `mapstructure:"app" json:"app"`
	*LogConf   `mapstructure:"log" json:"log"`
	*MySQLConf `mapstructure:"mysql" json:"mysql"`
	*RedisConf `mapstructure:"redis" json:"redis"`
}

// AppConf It's a struct with three fields, each of which is a string.
// @property {string} Name - The name of the application
// @property {string} Mode - The mode in which the application will run. It can be either debug or release.
// @property {string} Port - The port on which the application will run.
type AppConf struct {
	Name    string `mapstructure:"name" json:"name"`
	Mode    string `mapstructure:"mode" json:"mode"`
	Port    int    `mapstructure:"port" json:"port"`
	Verison string `mapstructure:"version" json:"verison"`
}

// LogConf It's a struct with a bunch of fields that are all strings.
// @property {string} Level - The log level, which can be debug, info, warn, error, fatal, or panic.
// @property {string} Filename - The name of the log file.
// @property {int} MaxSize - The maximum size of the log file in megabytes.
// @property {int} MaxPage - The maximum number of pages in the log file.
// @property {int} MaxBackups - The maximum number of old log files to retain.
type LogConf struct {
	Level      string `mapstructure:"level" json:"level"`
	Filename   string `mapstructure:"filename" json:"filename"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
}

// MySQLConf "MySQLConf is a struct that contains the configuration for a MySQL database."
//
// The `mapstructure` tags are used by the `mapstructure` package to decode the configuration file into the struct. The
// `json` tags are used by the `json` package to encode the struct into a JSON string.
//
// The `mapstructure` tags are used by the `mapstructure` package to decode the configuration file into the struct. The
// `json` tags are used by the `json` package to encode the struct into a JSON string.
//
// The `
// @property {string} Host - The hostname of the MySQL server.
// @property {int} Port - The port number of the MySQL server.
// @property {string} User - The user to connect as.
// @property {string} Password - The password for the user.
// @property {string} Dbname - The name of the database to connect to.
type MySQLConf struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	Dbname       string `mapstructure:"dbname" json:"dbname"`
	OpenMaxConns int    `mapstructure:"open_max_conns" json:"open_max_conns"`
	IdelMaxConns int    `mapstructure:"idel_max_conns" json:"idel_max_conns"`
}

// RedisConf is a struct that contains a string, an int, and another int.
// @property {string} Host - The hostname of the Redis server.
// @property {int} Port - The port on which the Redis server is running.
// @property {int} DB - The database number to connect to.
type RedisConf struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	DB       int    `mapstructure:"db" json:"db"`
	Password string `mapstructure:"password" json:"password"`
	PoolSize int    `mapstructure:"pool_szie" json:"pool_size"`
}

// Init It reads a configuration file, and then watches it for changes
func Init(filePath string) (err error) {
	// 读取文件的位置
	//viper.SetDefault("fileDir", "./")

	// 方式1：指定配置文件路径（可以是绝对路径、相对路径）
	// 相对路径：相对执行的可执行文件的相对路径
	// 绝对路径：系统中实际的文件路径
	// 读取配置文件
	//viper.SetConfigFile("./conf/config.yaml")

	// 方式2：指定配置文件名和配置文件位置，viper自行查找可用的配置文件
	// 配置文件名不需要后缀
	// 配置文件位置可配置多个
	// 配置文件目录中避免出现 重名的配置文件(读取配置文件出错)
	//viper.SetConfigName("config") // 配置文件名称(无文件后缀)
	//viper.AddConfigPath(".")      // 指定查找配置文件路径（这里使用相对路径）
	//viper.AddConfigPath("./conf") // 指定查找配置文件路径（这里使用相对路径）

	// 配合远程配置中心使用
	//viper.SetConfigType("yaml")   // 指定配置文件类型（专用于从远程获取配置信息时指定指定配置文件类型）

	// 从命令行获取到 配置文件
	viper.SetConfigFile(filePath)

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("viper.ReadInConfig() failed, error: ", err)
		return
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 把读取到的到的配置信息反序列化到模型中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal failed, error:", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("修改配置文件……  %s \n", in.Name)
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper.Unmarshal failed, error:", err)
		}
	})
	return
}
