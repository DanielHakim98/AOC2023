/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/DanielHakim98/aoc/cmd"
	"github.com/DanielHakim98/aoc/day5"
	"github.com/DanielHakim98/aoc/utils"
)

func main() {
	cmd.Execute()
	d := day5.Day5{Dnum: 5}
	fmt.Println("Lowest location number: ", d.PartTwo("example_1", utils.GetInput))
}
