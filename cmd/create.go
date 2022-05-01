package cmd

import (
	"errors"
	"fmt"
	"folder_creator/pattern"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var (
	ErrValidate = errors.New("validation error")
)

var createCmd = &cobra.Command{
	Use:   "create pattern",
	Short: "The command creates folders according to the selected pattern and settings.",
	Long: "The command creates folders according to the selected pattern and settings.\n\n" +
		"pattern is a string that can contain regular text and special verbs:\n" +
		"%d - number",
	Example: "create -c 100 -p 2 \"%d. Example folder\"",
	Args:    cobra.ExactArgs(1),
	Run:     CreateExecute,
}

var (
	count  int
	number = pattern.NewNumberOptions()
)

func init() {
	createCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of iterations to create folders. Default = 1")
	createCmd.Flags().IntVarP(&number.Padding, "padding", "p", 0, "Number of digits in the number (missing digits are filled with 0). Default = 0")
	createCmd.Flags().IntVarP(&number.Increment, "increment", "i", 1, "By how much will the number increase with one iteration. Default = 1")
	createCmd.Flags().IntVarP(&number.Start, "start", "s", 1, "Initial number. Default = 1")
}

func validateCreateCommand(pattern string) error {
	switch {
	case count < 1:
		return fmt.Errorf("%w: count can not be smaller than 1", ErrValidate)
	case strings.TrimSpace(pattern) == "":
		return fmt.Errorf("%w: pattern must have at least one non-space character", ErrValidate)
	}

	return nil
}

func CreateExecute(_ *cobra.Command, args []string) {
	p := args[0]

	if err := validateCreateCommand(p); err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < count; i++ {
		currentNumber := number.Iterate()

		name := strings.Replace(p, pattern.NumberPattern, currentNumber, -1)

		err := os.Mkdir(name, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
