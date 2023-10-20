package server

import (
	"log"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/gin-gonic/gin"

	"aone-qc/internal/initialization"
	"aone-qc/internal/handlers"
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
		// initialization.Migrate()
		initialization.InitDatabaseConnection()

		r := gin.Default()
		// 存储应用配置
		r.Use(func (c *gin.Context) {
			c.Keys = make(map[string]interface{})
			c.Keys["config"] = config
			c.Next()
		})

        api := r.Group("/api")
        {
            api.POST("/qc/create", handlers.CreateOrRetryQcTask)
            api.POST("/qc/retry", handlers.CreateOrRetryQcTask)

            // api.GET("/qc/batch/list", handlers.GetQcBatchList)
            // api.GET("/qc/notify/email", handlers.SendQcNotifyToEmail)
            // api.GET("/qc/batch/detail/list", handlers.GetQcBatchDetailList)

            // api.POST("/data/list", handlers.GetDataList)
        }

		err = startHTTPServer(config, r)
		if err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	},
}

func startHTTPServer(config initialization.Config, r *gin.Engine) (err error) {
	fmt.Println("")
	log.Printf("http server started: http://localhost:%d/\n", config.HttpPort)
	err = r.Run(fmt.Sprintf(":%d", config.HttpPort))
	if err != nil {
		return fmt.Errorf("failed to start http server: %v", err)
	}
	return nil
}

func Register(rootCmd *cobra.Command) error {
	rootCmd.AddCommand(cmd)
	return nil
}