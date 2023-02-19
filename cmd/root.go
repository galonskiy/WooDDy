package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wooddy",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			execute, _ := cmd.Flags().GetBool("execute")
			saveFilename, _ := cmd.Flags().GetString("save")

			readMds(args, execute, saveFilename)
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("save", "s", "", "Help message for save")
	rootCmd.Flags().BoolP("execute", "e", false, "Help message for execute")
}

func readMds(args []string, execute bool, saveFilename string) {

	buffer := bytes.Buffer{}
	buffer.WriteString("#!/usr/bin/env bash \n")
	buffer.WriteString("\n")

	for _, filename := range args {

		b, err := os.ReadFile(filename)

		if err != nil {
			log.Fatal(err)
		}

		matches := extractScriptsFromMd(b)

		buffer.WriteString("# ")
		buffer.WriteString(filename)
		buffer.WriteString("\n")

		for _, m := range matches {

			// TODO add log RUN command

			buffer.WriteString("\n")
			buffer.WriteString(strings.TrimSpace(m[1]))
			buffer.WriteString("\n\n")
		}

	}

	if saveFilename != "" {
		sf, err := os.Create(saveFilename) // in Go version older than 1.17 you can use ioutil.TempFile
		if err != nil {
			log.Fatal(err)
		}

		if _, err := sf.Write(buffer.Bytes()); err != nil {
			log.Fatal(err)
		}

		sf.Close()
	} else {
		fmt.Print(buffer.String())
	}

	if execute {
		// create and open a temporary file
		f, err := os.CreateTemp("", "tmpfile-") // in Go version older than 1.17 you can use ioutil.TempFile
		if err != nil {
			log.Fatal(err)
		}

		if _, err := f.Write(buffer.Bytes()); err != nil {
			log.Fatal(err)
		}

		f.Close()

		if err := os.Chmod(f.Name(), 0544); err != nil {
			log.Fatal(err)
		}

		execBashFile(f.Name())

		os.Remove(f.Name())
	}

}

func extractScriptsFromMd(b []byte) [][]string {

	fileText := string(b[:])

	left := "```"
	right := "\n```\n"

	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `[a-zA-Z0-9_\- ]{0,}\n` + `(.*?)` + regexp.QuoteMeta(right))

	matches := rx.FindAllStringSubmatch(fileText, -1)

	return matches
}

func execBashFile(filename string) {
	cmd := exec.Command(filename)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
