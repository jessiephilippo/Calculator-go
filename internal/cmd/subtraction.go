/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func newSubtractionCmd() *cobra.Command {
	subtractionCmd := &cobra.Command{
		Use:   "subtraction",
		Short: "subtract your calculation",
		Run: func(cmd *cobra.Command, args []string) {
			fstatus, _ := cmd.Flags().GetBool("float")
			if fstatus { // if status is true, call subtractFloat
				subtractFloat(args)
			} else {
				subtractInt(args)
			}
		},
	}

	return subtractionCmd
}

func subtractFloat(args []string) {
	var sum float64

	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
		}

		sum = sum + ftemp
	}

	fmt.Printf("Sum of floating numbers %s is %.2f\n", args, sum)
}

func subtractInt(args []string) {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}

		sum = sum + itemp
	}

	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}
