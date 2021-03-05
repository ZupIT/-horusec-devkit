// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ZupIT/horusec-admin/internal/logger"
)

const (
	Addr            = ":3000"
	ShutdownTimeout = 5 * time.Second
)

type (
	server struct {
		*http.Server
	}
	Interface interface {
		Start() Interface
		GracefullyShutdown() error
	}
)

func New(handler http.Handler) Interface {
	return &server{Server: &http.Server{Addr: Addr, Handler: handler}}
}

func (s *server) Start() Interface {
	go func() {
		log := logger.WithPrefix("server")
		log.WithField("addr", Addr).Info("listening")
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Fatal("listen error")
		}
	}()

	return s
}

func (s *server) GracefullyShutdown() error {
	logger.WithPrefix("server").Warn("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to gracefully shuts down the server: %w", err)
	}

	return nil
}
