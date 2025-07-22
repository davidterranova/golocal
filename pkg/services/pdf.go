package services

import (
	"context"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type PDFSvc struct{}

func NewPDFSvc() *PDFSvc {
	return &PDFSvc{}
}

func (svc *PDFSvc) RemovePages(ctx context.Context, r io.ReadSeeker, w io.Writer, selectedPages []string) error {
	return api.RemovePages(r, w, selectedPages, nil)
}
