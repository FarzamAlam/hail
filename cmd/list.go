package cmd

import (
	"fmt"
	"hail/internal/hailconfig"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list is used to print all the alias and commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		hc := new(hailconfig.Hailconfig).WithLoader(hailconfig.DefaultLoader)
		defer hc.Close()
		err := hc.Parse()
		if err != nil {
			fmt.Println("error in list : ", err)
			os.Exit(2)
		}
		return hc.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
