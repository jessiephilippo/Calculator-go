/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func newDivisionCmd() *cobra.Command {
	newDivisionCmd := &cobra.Command{
		Use:   "divide",
		Short: "divide your calculation",
		Run: func(cmd *cobra.Command, args []string) {
			fstatus, _ := cmd.Flags().GetBool("float")
			if fstatus { // if status is true, call divideFloat
				divideFloat(args)
			} else {
				divideInt(args)
			}
		},
	}

	newDivisionCmd.Flags().BoolP("float", "f", false, "Add floating numbers")
	return newDivisionCmd
}

func divideFloat(args []string) {
	var sum float64

	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
		}

		sum = sum / ftemp
	}

	fmt.Printf("Sum of floating numbers %s is %.2f\n", args, sum)
}

func divideInt(args []string) {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)
		if err != nil {
			fmt.Println(err)
		}

		sum = sum / itemp
	}

	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}
