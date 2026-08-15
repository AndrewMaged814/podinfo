package http

import (
	"net/http"
	"runtime"

	"github.com/stefanprodan/podinfo/pkg/version"
)

// Version godoc
// @Summary Version
// @Description returns podinfo version, git commit hash, build time, and Go runtime
// @Tags HTTP API
// @Produce json
// @Router /version [get]
// @Success 200 {object} http.MapResponse
func (s *Server) versionHandler(w http.ResponseWriter, r *http.Request) {
	result := map[string]string{
		"version":   version.VERSION,
		"commit":    version.REVISION,
		"buildtime": version.BUILDTIME,
		"runtime":   runtime.Version(),
	}
	s.JSONResponse(w, r, result)
}
