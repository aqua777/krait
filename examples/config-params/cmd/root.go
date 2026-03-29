package cmd

import (
    "fmt"
    "os"
	"time"
    "github.com/spf13/cobra"
)

var versionCobraCmd = &cobra.Command{
    Use:   "version",
    Short: "version short",
    Long:  "version long",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Executing: version Cobra style")
    },
}

var rootCobraCmd = &cobra.Command{
    Use:   "app",
    Short: "Cobra app short",
    Long:  "Cobra app long",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Executing: app Cobra style")
    },
}

func Execute() {
    err := rootCobraCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCobraCmd.Flags().BoolP("bool", "b", false, "bool param description")
    rootCobraCmd.Flags().IntP("int", "i", 123, "int param description")
    rootCobraCmd.Flags().StringP("string", "s", "defaultStr", "string param description")
    rootCobraCmd.Flags().Float64P("float64", "f", 123.456, "float64 param description")
    rootCobraCmd.Flags().DurationP("duration", "d", time.Second*5, "duration param description")
    rootCobraCmd.AddCommand(versionCobraCmd)
}
