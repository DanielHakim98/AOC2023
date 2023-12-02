/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/DanielHakim98/aoc/day2"
	"github.com/spf13/cobra"
)

// dayTwoCmd represents the dayTwo command
var part int
var dayTwoCmd = &cobra.Command{
	Use:   "dayTwo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("No filename is passed in")
		}
		d := day2.Day2{}
		switch part {
		case 1:
			fmt.Println("sum of possible games ids: ", d.PartOne(args[0], day2.GetInput))
		case 2:
			fmt.Println("sum of power of minimum cube required from games: ", d.PartTwo(args[0], day2.GetInput))
		default:
			log.Fatal("Invalid 'part' flag")
		}

	},
}

func init() {
	rootCmd.AddCommand(dayTwoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dayTwoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	dayTwoCmd.Flags().IntVarP(&part, "part", "p", 0, "part to run")
}
