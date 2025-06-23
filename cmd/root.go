package cmd

import (
	"fmt"
	"net/http"
	"os"
	"task/config"
	"task/pkg/router"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ArgsPort     int
	ArgsLogLevel string
)

func init() {
	rootCmd.Flags().IntVar(&ArgsPort, "port", 8080, "Port to run the HTTP server on (overrides config)")
	rootCmd.Flags().StringVar(&ArgsLogLevel, "log-level", "debug", "Log level (debug, info, warn, error)")
}

var rootCmd = &cobra.Command{
	Use:   "taskd",
	Short: "Task API Gateway",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if ArgsPort != 0 {
			cfg.Port = ArgsPort
		}
		if ArgsLogLevel != "" {
			cfg.LogLevel = ArgsLogLevel
		}

		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(logLevelMapping[cfg.LogLevel])

		handler := router.SetupRouter()
		fmt.Printf("Server running at :%d\n", cfg.Port)
		return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), handler)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
