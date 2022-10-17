package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	// from file
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))

	//from Flags
	fmt.Println(os.Args)
	pflag.String("mq.ip", "127.0.0.1", "Redis port to connect")
	pflag.Int("mq.port", 8381, "Redis port to connect")
	pflag.Parse()
	// 绑定命令行
	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("mq.ip: ", viper.Get("mq.ip"))
	fmt.Println("mq.port: ", viper.Get("mq.port"))

	// from ENV
	viper.AutomaticEnv()
	fmt.Println("GOPATH: ", viper.Get("GOPATH"))
	fmt.Println("MYSQL_IP: ", viper.Get("MYSQL_IP"))
	fmt.Println("MYSQL_PORT: ", viper.Get("MYSQL_PORT"))

	// single bind
	viper.BindEnv("redis.port", "redis_port")
	viper.BindEnv("go.path", "GOPATH")

	fmt.Println("redis.port: ", viper.Get("redis.port"))
	fmt.Println("go.path: ", viper.Get("go.path"))
}
