package cmd

import (
	"fmt"
	
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGetCommand() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "print a meeting link to STDOUT",
		Run: func(cmd *cobra.Command, args []string) {
			name := GetArg(args, 0, defaultKey)
			link := viper.GetString(name)
			if link != "" {
				fmt.Println(link)
			} else {
				log.Error(LinkNotSet)
			}
		},
	}

	return getCmd
}
