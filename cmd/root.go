package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    clipboard "github.com/atotto/clipboard"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    configFilePath = os.Getenv("HOME") + "/.zoom"
    configFileType = "yaml"
    configFileName = "config.yaml"
    myLinkKey      = "myLink"
)

func getLine(s string) string {
    fmt.Printf("%s: ", s)
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Info("Something went wrong getting input from stdin")
    }
    return strings.Replace(input, "\n", "", -1)
}

func NewCommand() *cobra.Command {
    var name string

    var command = &cobra.Command{
        Use:   "zoom",
        Short: "some useful commands for using Zoom",
        Run: func(cmd *cobra.Command, args []string) {
            if name != "" {
                link := viper.GetString(name)
                if link != "" {
                    clipboard.WriteAll(link)
                    fmt.Println("Copied meeting link for %s", name)
                } else {
                    log.Error("You haven't set that link yet!")
                }
            } else {
                myLink := viper.GetString(myLinkKey)
                clipboard.WriteAll(myLink)
                fmt.Println("Copied personal meeting link to clipboard")
            }
        },
    }

    command.AddCommand(NewInitCommand())
    command.AddCommand(NewAddCommand())
    command.AddCommand(NewResetCommand())

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

    command.Flags().StringVarP(&name, "name", "n", "", "name of zoom meeting to copy")

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
