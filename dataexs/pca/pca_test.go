package pca

import (
	"os"
	"testing"
)

func TestExportJSON(t *testing.T) {
	_ = os.Setenv("REDIS_ADDR", "139.9.119.21:56379")
	ExportJSON()
}
