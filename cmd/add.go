package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Used for addition",
	Long: `This cli tool is used in addition of numbers.
	Numbers can be integer type or float type `,
	Run: func(cmd *cobra.Command, args []string) {
		// returning boolean value if flag float is used
		fstatus, _ := cmd.Flags().GetBool("float")

		if fstatus {
			addFloat(args)
		} else {
			AddInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	//BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

// Logic for adding integer value
func AddInt(args []string) int {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum = sum + itemp
	}

	fmt.Printf("Sum of integer numbers %s is %d\n", args, sum)
	return (sum)
}

// logic for adding float value
func addFloat(args []string) {
	var sum float64
	for _, fval := range args {
		//ParseFloat converts the string s to a floating-point number with the precision specified by bitSize
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum = sum + ftemp
	}

	fmt.Printf("Sum of floating numbers %s is %f\n", args, sum)
}
