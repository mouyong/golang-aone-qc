package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"aone-qc/internal/handlers"
	"aone-qc/internal/initialization"
	"aone-qc/pkg/rabbitmq"
)

var cmd = &cobra.Command{
	Use:   "server",
	Short: "run aone-qc api server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Println(err)
			return
		}

		config := initialization.LoadConfig(cfg)
		go func() {
			fmt.Println("\n正在初始化")

			initialization.InitDatabaseConnection()
			rabbitmq.NewRabbitmq(initialization.AppConfig.MqHost, initialization.AppConfig.MqPort)

			fmt.Println("初始化完成")
		}()

		r := gin.Default()
		// 存储应用配置
		r.Use(func(c *gin.Context) {
			c.Keys = make(map[string]interface{})
			c.Keys["config"] = config
			c.Next()
		})

		// 开始/重新 批次质控
		r.POST("/api/qc/create", handlers.CreateOrRetryQcTask)
		r.POST("/api/qc/retry", handlers.CreateOrRetryQcTask)

		// 查看批次进度与说明
		r.GET("/api/qc/tasks/list", handlers.GetQcTaskListWithPage)
		r.GET("/api/qc/tasks/sample/list", handlers.GetQcTaskSampleListWithPage)
		// r.GET("/api/qc/notify/email", handlers.SendQcNotifyToEmail)

		// r.POST("/api/data/list", handlers.GetDataList)

		err = startHTTPServer(config, r)
		if err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	},
}

func startHTTPServer(config initialization.Config, r *gin.Engine) (err error) {
	fmt.Println("")
	log.Printf("http server started: http://localhost:%d/\n", config.HttpPort)
	err = r.Run(fmt.Sprintf("%s:%d", config.HttpHost, config.HttpPort))
	if err != nil {
		return fmt.Errorf("failed to start http server: %v", err)
	}
	return nil
}

func Register(rootCmd *cobra.Command) error {
	rootCmd.AddCommand(cmd)
	return nil
}
