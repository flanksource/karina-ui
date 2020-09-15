package pkg

import (
	"os"
    "os/exec"

	"github.com/flanksource/commons/logger"
	"github.com/spf13/cobra"

)

var GetCanary = &cobra.Command{
		Use: "get-canary",
		Short: "Get canary information",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof(`Retrieving Canary Info...
				`)
			Retrieva()			 
		},
	}

func Retrieva() {

	cmd := exec.Command("kubectl", "get", "canary")

    outfile, err := os.Create("./pkg/api/getcanary.txt")
    if err != nil {
        panic(err)
    }
    defer outfile.Close()
    cmd.Stdout = outfile

    err = cmd.Start(); if err != nil {
        panic(err)
    }
    cmd.Wait()
    	
}
