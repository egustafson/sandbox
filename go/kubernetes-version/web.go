package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

type WebServer interface {
	// Start starts the web server and blocks until it is stopped.
	Start() error
}

type webServerImpl struct {
	srv     *http.Server
	apiCtrl KubeAPIController
}

func InitWebServer(apiCtrl KubeAPIController) (WebServer, error) {
	srv := &http.Server{
		Addr: ":8080",
	}
	return &webServerImpl{
		srv:     srv,
		apiCtrl: apiCtrl,
	}, nil
}

func (s *webServerImpl) Start() error {
	// Set up HTTP handlers
	http.HandleFunc("/version", s.handleVersion)

	// Start the server
	slog.Info(fmt.Sprintf("Starting web server on %s", s.srv.Addr))
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start web server: %w", err)
	}
	return nil
}

func (s *webServerImpl) handleVersion(w http.ResponseWriter, _ *http.Request) {
	versionInfo, err := s.apiCtrl.ServerVersion()
	if err != nil {
		http.Error(w, "Failed to get Kubernetes API Server version", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Kubernetes API Server Version: %s", versionInfo.GitVersion)
}
