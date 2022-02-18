/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newMinusCmd() *cobra.Command {
	minusCmd := &cobra.Command{
		Use:   "minus",
		Short: "minus your calculation",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("minus called")
		},
	}

	return minusCmd
}
