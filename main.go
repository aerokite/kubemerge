package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"time"

	"github.com/fatih/color"
	yaml "gopkg.in/yaml.v2"
	"magic.pathao.com/platform/kubeconfig/pkg"
)

func main() {

	if len(os.Args) != 2 {
		color.New(color.FgRed).Printf("ArgsError: ")
		fmt.Println("must pass new kubeconfig filepath to merge")
		os.Exit(0)
	}

	usr, err := user.Current()
	if err != nil {
		color.New(color.FgRed).Printf("UserError: ")
		fmt.Println(err)
		os.Exit(0)
	}

	cfn := fmt.Sprintf("%s/.kube/config", usr.HomeDir)
	dat, err := ioutil.ReadFile(cfn)
	if err != nil {
		color.New(color.FgRed).Printf("FileError: ")
		fmt.Println(err)
		os.Exit(0)
	}

	suffix := fmt.Sprintf("%s", time.Now().Format("20060102T150405"))
	cfnb := fmt.Sprintf("%s-%s", cfn, suffix)
	if err := ioutil.WriteFile(cfnb, dat, 0600); err != nil {
		color.New(color.FgRed).Printf("FileError: ")
		fmt.Println(err)
		os.Exit(0)
	}

	config := new(pkg.Config)
	if err := yaml.Unmarshal(dat, config); err != nil {
		color.New(color.FgRed).Printf("UnmarshalError: ")
		fmt.Println(err)
		os.Exit(0)
	}
	config.Rename()

	newConfigPath := os.Args[1]
	newDat, err := ioutil.ReadFile(newConfigPath)
	if err != nil {
		color.New(color.FgRed).Printf("FileError: ")
		fmt.Println(err)
		os.Exit(0)
	}
	newConfig := new(pkg.Config)
	if err := yaml.Unmarshal(newDat, &newConfig); err != nil {
		color.New(color.FgRed).Printf("UnmarshalError: ")
		fmt.Println(err)
		os.Exit(0)
	}
	newConfig.Rename()

	mergedConfig, err := pkg.Merge(config, newConfig)
	if err != nil {
		color.New(color.FgRed).Printf("MergeConflict: ")
		fmt.Println(err)
		os.Exit(0)
	}

	data, err := yaml.Marshal(mergedConfig)
	if err != nil {
		color.New(color.FgRed).Printf("MarshalError: ")
		fmt.Println(err)
		os.Exit(0)
	}

	if err := ioutil.WriteFile(cfn, data, 0600); err != nil {
		color.New(color.FgRed).Printf("FileError: ")
		fmt.Println(err)
		os.Exit(0)
	}

	color.New(color.FgGreen).Printf("Success: ")
	fmt.Println("merge completed")
	color.New(color.FgBlue).Printf("--> ")
	fmt.Println("To restore your previous kubeconfig")
	fmt.Println()
	color.New(color.FgYellow).Printf("\tcp ~/.kube/config-%s ~/.kube/config\n", suffix)
	fmt.Println()

}
