package handlers

import (
	"net/http"
	"net/http/httptest"
	"network-scanner/internal/report"
	"strings"
	"testing"
)

func TestGetReportHandler(t *testing.T) {
	rep := report.NewReport()
	rep.Update("127.0.0.1", "online")

	req := httptest.NewRequest(http.MethodGet, "/report", nil)
	w := httptest.NewRecorder()

	handler := GetFilteredReportHandler(rep)
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("ожидался статус 200, получено %d", resp.StatusCode)
	}

	body := w.Body.String()
	if !strings.Contains(body, "127.0.0.1") {
		t.Fatalf("ответ не содержит ожидаемый IP: %s", body)
	}
}
