package internal

import (
	"context"
	"io"

	"github.com/davidterranova/golocal/pkg/services"
)

type PDFSvc interface {
	RemovePages(ctx context.Context, r io.ReadSeeker, w io.Writer, selectedPages []string) error
}

type App struct {
	PDFSvc
}

func New() PDFSvc {
	pdfSvc := services.NewPDFSvc()

	return &App{
		PDFSvc: pdfSvc,
	}
}
