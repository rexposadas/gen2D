package cmd

import (
	"sync"

	"github.com/rexposadas/gen2D/models"
	"github.com/rexposadas/gen2D/util/config"
	"github.com/rexposadas/gen2D/util/require"
	"github.com/spf13/cobra"
)

var circlesLoopCmd = &cobra.Command{
	Use:   "loop",
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
				models.CirclesLoop(cfg)
			}()
		}

		wg.Wait()
	},
}

func init() {
	circlesCmd.AddCommand(circlesLoopCmd)
}
