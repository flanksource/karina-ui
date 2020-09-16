
package pkg

import (
	"os"
    "os/exec"

	"github.com/flanksource/commons/logger"
	"github.com/spf13/cobra"

)

var GetKube= &cobra.Command{
		Use: "get-kube",
		Short: "Get Kubernetes information",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof(`Retrieving Info...
				`)
			KubeVersion()	
			OSVersion()
			KernelVersion()	 
		},
	}

func KubeVersion() {

	cmd := exec.Command("kubectl", "version", "--short")

    outfile, err := os.Create("./pkg/api/getkube.txt")
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

func OSVersion() {

	cmd := exec.Command("cat", "/etc/os-release")

    outfile, err := os.Create("./pkg/api/getos.txt")
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


func KernelVersion() {

	cmd := exec.Command("uname", "-s", "-r")

    outfile, err := os.Create("./pkg/api/getkernel.txt")
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