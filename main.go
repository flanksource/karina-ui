package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/flanksource/commons/logger"
	"github.com/flanksource/karina-ui/pkg"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {

	var root = &cobra.Command{
		Use: "karina-ui",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger.UseZap(cmd.Flags())
			logger.Infof("🏁 Starting %s", version)
		},
	}

	logger.BindFlags(root.PersistentFlags())

	if len(commit) > 8 {
		version = fmt.Sprintf("%v, commit %v, built at %v", version, commit[0:8], date)
	}

	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version of karina-ui",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	})

	serve := &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			file, _ := cmd.Flags().GetString("config")
			port, _ := cmd.Flags().GetUint("port")
			err := pkg.ParseConfiguration(file)
			if err != nil {
				logger.Fatalf("Failed to parse karina-ui configuration: %v", err)
			}
			http.Handle("/", http.FileServer(http.Dir("./dist/")))
			http.HandleFunc("/api", pkg.Serve)

			addr := fmt.Sprintf(":%d", port)
			logger.Infof("👂 Listening on %s", addr)
			if err := http.ListenAndServe(addr, nil); err != nil {
				logger.Fatalf("Failed to start http server: %v", err)
			}
		},
	}
	serve.Flags().UintP("port", "p", 8080, "Port to use for webserver")
	root.AddCommand(serve)
	root.SetUsageTemplate(root.UsageTemplate() + fmt.Sprintf("\nversion: %s\n ", version))
	root.PersistentFlags().String("config", "", "Specify a kubeconfig to use")

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
