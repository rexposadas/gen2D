package cmd

import (
	"github.com/spf13/cobra"
	"sync"

	"github.com/rexposadas/gen2D/models"
	"github.com/rexposadas/gen2D/util/config"
	"github.com/rexposadas/gen2D/util/require"
)

var circlesRegularCmd = &cobra.Command{
	Use:   "regular",
	Short: "regular",
	Long:  `regular circles`,
	Run: func(cmd *cobra.Command, args []string) {
		total := require.Count(count)

		cfg := config.New(file)
		circlesRegular(cfg, total)
	},
}

func circlesRegular(cfg config.Config, count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			models.Circles(cfg)
		}()
	}

	wg.Wait()
}

func init() {
	circlesCmd.AddCommand(circlesRegularCmd)
}
