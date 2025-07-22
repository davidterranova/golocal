package cmd

import (
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var pdfCmd = &cobra.Command{
	Use:   "pdf",
	Short: "pdf operations",
}

var pdfDeletePageCmd = &cobra.Command{
	Use:   "delpage",
	Short: "delete pdf pages",
	Run:   runPDFDeletePages,
}

var (
	sourceFile string
	dstFile    string

	pdfDeletePageNumbers string
)

func runPDFDeletePages(cmd *cobra.Command, args []string) {
	pages := strings.Split(pdfDeletePageNumbers, ",")

	if len(pages) == 0 {
		log.Info().
			Str("src_file", sourceFile).
			Str("dst_file", dstFile).
			Msg("no pages to delete")
		return
	}

	log.Info().
		Str("src_file", sourceFile).
		Str("dst_file", dstFile).
		Any("pages", pages).
		Msg("deleting file page")

	err := api.RemovePagesFile(sourceFile, dstFile, pages, nil)
	if err != nil {
		log.Error().Err(err).Msg("error while deleting pdf file page")
	}
}

func init() {
	rootCmd.AddCommand(pdfCmd)
	pdfCmd.AddCommand(pdfDeletePageCmd)

	pdfDeletePageCmd.PersistentFlags().StringVarP(&sourceFile, "src", "s", "", "local path of the source file to read from")
	pdfDeletePageCmd.PersistentFlags().StringVarP(&dstFile, "dst", "d", "", "local path of the file to write")
	pdfDeletePageCmd.PersistentFlags().StringVarP(&pdfDeletePageNumbers, "pages", "p", "", "pages to delete, coma separated")
}
