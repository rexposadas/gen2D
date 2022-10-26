package cmd

import (
	"github.com/rexposadas/gen2D/models"
	"github.com/rexposadas/gen2D/util/config"
	"github.com/rexposadas/gen2D/util/require"
	"sync"

	"github.com/spf13/cobra"
)

// dotCmd represents the dot command
var dotCmd = &cobra.Command{
	Use:   "dot",
	Short: "generate dot images",
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
		total := require.Count(count)
		cfg := config.New(file)

		dot(cfg, total)
	},
}

func dot(cfg config.Config, count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			models.Dot(cfg)
		}()
	}
	wg.Wait()
}

func init() {
	rootCmd.AddCommand(dotCmd)
}
