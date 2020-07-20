package cmd

import (
    "fmt"
    "os"

    clipboard "github.com/atotto/clipboard"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    configFilePath = os.Getenv("HOME") + "/.zoom"
    configFileType = "yaml"
    configFileName = "config.yaml"
    myLinkKey      = "myLink"
)

func NewCommand() *cobra.Command {
    var command = &cobra.Command{
        Use:   "zoom",
        Short: "argo is the command line interface to Argo",
        Run: func(cmd *cobra.Command, args []string) {
            if ok := viper.IsSet(myLinkKey); ok {
                myLink := viper.GetString(myLinkKey)
                clipboard.WriteAll(myLink)
                fmt.Println("Copied personal meeting link to clipboard")
            } else {
                cmd.HelpFunc()(cmd, args)
            }
        },
    }

    command.AddCommand(NewInitCommand())

    command.PersistentPreRun = func(cmd *cobra.Command, args []string) {
        viper.AddConfigPath(configFilePath)
        viper.SetConfigName(configFileName)
        viper.SetConfigType(configFileType)
        if err := viper.ReadInConfig(); err != nil {
            if _, ok := err.(viper.ConfigFileNotFoundError); ok {
                fmt.Println("Configuration not found. Let's get started!")
                NewInitCommand().Execute()
            }
        }
    }
    return command
}

var rootCmd = &cobra.Command{
    Use:   "zoom",
    Short: "some useful commands for using Zoom",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
