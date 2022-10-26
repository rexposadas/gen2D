package cmd

import (
	"github.com/rexposadas/gen2D/models"
	"github.com/rexposadas/gen2D/util/config"
	"github.com/rexposadas/gen2D/util/require"
	"github.com/spf13/cobra"
	"sync"
)

// holeCmd represents the hole command
var holeCmd = &cobra.Command{
	Use:   "hole",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := config.New(file)
		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				models.Hole(cfg)
			}()

		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(holeCmd)
}
