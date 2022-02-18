package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func Execute() {
	rootCmd := NewRootCommand()

	calculator :=
		`
		┌─┐┌─┐┬  ┌─┐┬ ┬┬  ┌─┐┌┬┐┌─┐┬─┐
		│  ├─┤│  │  │ ││  ├─┤ │ │ │├┬┘
		└─┘┴ ┴┴─┘└─┘└─┘┴─┘┴ ┴ ┴ └─┘┴└─																				  
     `
	for idx, s := range calculator {
		switch c := idx % 8; c {
		case 0:
			fmt.Print(Red(string(s)))
		case 1:
			fmt.Print(Yellow(string(s)))
		case 2:
			fmt.Print(Green(string(s)))
		case 3:
			fmt.Print(Purple(string(s)))
		case 4:
			fmt.Print(Magenta(string(s)))
		case 5:
			fmt.Print(Teal(string(s)))
		case 6:
			fmt.Print(White(string(s)))
		case 7:
			fmt.Print(Black(string(s)))
		}

	}

	if errExc := rootCmd.Execute(); errExc != nil {
		fmt.Printf("%v", errExc)
		os.Exit(1)
	}
}

func NewRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "calculator",
		Short: "Awesome calculator",
	}

	rootCmd.AddCommand(newAdditionCmd())
	rootCmd.AddCommand(newDivisionCmd())
	rootCmd.AddCommand(newSubtractionCmd())
	rootCmd.AddCommand(newMultiplicationCmd())

	return rootCmd
}

func init() {

}
