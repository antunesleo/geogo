/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/antunesleo/geogo/geo"
	"github.com/spf13/cobra"
)

// distanceCmd represents the distance command
var distanceCmd = &cobra.Command{
	Use:   "distance",
	Short: "Distance between two points on earth",
	Long: `Calculates distance between two points on earth.

geogo distance 45.55 45.99 66.99 67.88

where:

arg0 (45.55) is first point latitude
arg1 (45.99) is first point longitude
arg2 (66.99) is second point latitude
arg3 (57.99) is second point longitude
`,
	Run: func(cmd *cobra.Command, args []string) {
		distance(args)
	},
}

func init() {
	rootCmd.AddCommand(distanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// distanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// distanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func distance(args []string) {
	flat, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Skipping: First lat (arg 0) not provided")
		return
	}

	flong, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Skipping: First long (arg 1) not provided")
		return
	}

	slat, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		fmt.Println("Skipping: Second lat (arg 2) not provided")
		return
	}

	slong, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		fmt.Println("Skipping: Second long (arg 3) not provided")
		return
	}

	fp := geo.Point{flat, flong}
	sp := geo.Point{slat, slong}

	res := fp.DistanceBetween(sp)

	fmt.Println("distance is: %n", res)
}
