package cmd

import (
	"fmt"
	
	clipboard "github.com/atotto/clipboard"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCopyCommand() *cobra.Command {
	var copyCmd = &cobra.Command{
		Use:   "copy",
		Short: "copy a meeting link to clipboard",
		Run: func(cmd *cobra.Command, args []string) {
			name := GetArg(args, 0, defaultKey)
			link := viper.GetString(name)
			if link != "" {
				clipboard.WriteAll(link)
				fmt.Printf("Copied link for %s\n", name)
			} else {
				log.Error(LinkNotSet)
			}
		},
	}

	return copyCmd
}
