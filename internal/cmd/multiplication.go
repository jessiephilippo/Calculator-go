/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func newMultiplicationCmd() *cobra.Command {
	multiplicationCmd := &cobra.Command{
		Use:   "multiplication",
		Short: "multiply your calculation",
		Run: func(cmd *cobra.Command, args []string) {
			fstatus, _ := cmd.Flags().GetBool("float")
			if fstatus { // if status is true, call multiplyFloat
				multiplyFloat(args)
			} else {
				multiplyInt(args)
			}
		},
	}

	multiplicationCmd.Flags().BoolP("float", "f", false, "Add floating numbers")
	return multiplicationCmd
}

func multiplyFloat(args []string) {
	var sum float64

	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
		}

		sum = sum * ftemp
	}

	fmt.Printf("Sum of floating numbers %s is %.2f\n", args, sum)
}

func multiplyInt(args []string) {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}

		sum = sum * itemp
	}

	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}
