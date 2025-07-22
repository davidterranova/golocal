package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/davidterranova/golocal/internal"
	ihttp "github.com/davidterranova/golocal/internal/adapters/http"
	"github.com/davidterranova/golocal/pkg/xhttp"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starts golocal http server",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := internal.New()

	go httpServer(ctx, app)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalCh:
		cancel()
	case <-ctx.Done():
	}
}

func httpServer(ctx context.Context, app internal.PDFSvc) {
	router := ihttp.New(
		app,
	)
	server := xhttp.NewServer(router, "", 8080)

	err := server.Serve(ctx)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("failed to start http server")
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
