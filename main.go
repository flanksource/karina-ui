package main

import (
	"fmt"
	"net/http"
	"os"
	//"os/exec"
	//"strings"
	//"errors"

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
			logger.Infof("Starting %s", version)
		},
	}

	logger.BindFlags(root.PersistentFlags())

	if len(commit) > 8 {
		version = fmt.Sprintf("%v, commit %v, built at %v", version, commit[0:8], date)
	}

	/*root.AddCommand(
		pkg.GetCanary,
	)*/

	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version of karina-ui",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	})

	root.AddCommand(&cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			http.Handle("/", http.FileServer(http.Dir("./dist/")))
			http.HandleFunc("/api", pkg.Serve)
			logger.Infof("Listening on %s", ":8080")

			if err := http.ListenAndServe(":8080", nil); err != nil {
				logger.Fatalf("%v", err)
			}
		},
	})

	root.SetUsageTemplate(root.UsageTemplate() + fmt.Sprintf("\nversion: %s\n ", version))
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
