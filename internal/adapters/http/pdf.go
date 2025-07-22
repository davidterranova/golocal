package http

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/davidterranova/golocal/pkg/xhttp"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type PDFSvc interface {
	RemovePages(ctx context.Context, r io.ReadSeeker, w io.Writer, selectedPages []string) error
}

type PDFHandler struct {
	app PDFSvc
}

func NewPDFHandler(app PDFSvc) *PDFHandler {
	return &PDFHandler{
		app: app,
	}
}

func (h *PDFHandler) RemovePages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pages := strings.Split(
		mux.Vars(r)[pdfPages],
		",",
	)
	if len(pages) == 0 {
		log.Ctx(ctx).Error().
			Str("pages", mux.Vars(r)[pdfPages]).
			Msg("PDFHandler:RemovePages invalid/empty provided pages")
		xhttp.WriteError(ctx, w, http.StatusBadRequest, "failed to get pages", errors.New("pages are required"))
		return
	}

	inMemoryFile, err := io.ReadAll(r.Body)
	if err != nil {
		log.Ctx(ctx).Error().
			Err(err).
			Msg("PDFHandler:RemovePages failed read file in memory")
		xhttp.WriteError(ctx, w, http.StatusInternalServerError, "failed to read file in memory", err)
		return
	}

	fileName := strings.Join(r.Header["X-Filename"], "")
	if fileName != "" {
		log.Info().
			Str("fileName", fileName).
			Msg("PDFHandler:RemovePages setting content disposition")
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	}

	err = h.app.RemovePages(ctx, bytes.NewReader(inMemoryFile), w, pages)
	if err != nil {
		log.Ctx(ctx).Error().
			Err(err).
			Msg("PDFHandler:RemovePages failed to delete pdf pages")
		xhttp.WriteError(ctx, w, http.StatusInternalServerError, "failed to delete pdf pages", err)
		return
	}

	log.Info().
		Str("deleted pages", mux.Vars(r)[pdfPages]).
		Msg("PDFHandler:RemovePages completed")
}
