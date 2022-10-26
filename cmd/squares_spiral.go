package cmd

import (
	"github.com/rexposadas/gen2D/models"
	"github.com/rexposadas/gen2D/util/config"
	"github.com/rexposadas/gen2D/util/require"
	"github.com/spf13/cobra"
)

var squaresSpiralCmd = &cobra.Command{
	Use:   "spiral",
	Short: "squares making a spiral",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		total := require.Count(count)

		cfg := config.New(file)
		for i := 0; i < total; i++ {
			models.SpiralSquare(cfg)

		}
	},
}

func init() {
	squaresCmd.AddCommand(squaresSpiralCmd)
}
