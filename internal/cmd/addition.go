/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func newAdditionCmd() *cobra.Command {
	additionCmd := &cobra.Command{
		Use:   "addition",
		Short: "addition your calculation",
		Run: func(cmd *cobra.Command, args []string) {
			fstatus, _ := cmd.Flags().GetBool("float")
			if fstatus { // if status is true, call addFloat
				addFloat(args)
			} else {
				addInt(args)
			}
		},
	}

	additionCmd.Flags().BoolP("float", "f", false, "Add floating numbers")
	return additionCmd
}

func addFloat(args []string) {
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

func addInt(args []string) {
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
