/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/IamDaedalus/yyt/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		// ignore the error message, cobra already prints it
		os.Exit(1)
	}
}
