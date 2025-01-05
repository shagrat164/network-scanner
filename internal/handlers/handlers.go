package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"network-scanner/internal/config"
	"network-scanner/internal/ping"
	"network-scanner/internal/report"
	"time"
)

// Обработчик главной страницы
func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	}
}

// Обработчик страницы "О нас"
func AboutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Страница 'О нас'. Тут будет описание.")
	}
}

// Обновить статус IP-адресов POST запрос
func RefreshHandler(rep *report.Report, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Обновить статус всех IP
		allIPs := rep.GetAllIPs()
		results := ping.ParallelCheckWithPool(allIPs, cfg.MaxThreads)

		// Обновить отчёт
		for ip, status := range results {
			rep.Update(ip, status)
		}

		// Обновление завершено
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "refresh complete",
		})
	}
}

// Обновление конфигурации POST запрос
func ReloadConfigHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		updateConfig, err := config.LoadConfig("cfg/config.json")
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка обновления конфигурации %v", err), http.StatusInternalServerError)
			return
		}
		cfg = updateConfig

		// Обновление завершено
		w.Write([]byte("Конфигурация обновлена"))
	}
}

// Обработчик report с фильтрацией
func GetFilteredReportHandler(rep *report.Report) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Чтение параметра filter из URL
		filter := r.URL.Query().Get("filter")

		var filtered map[string]report.IPStatus
		startTime := time.Now()
		switch filter {
		case "available": // Только доступные
			filtered = rep.Filter(func(status string) bool {
				return status == "online"
			})
		case "unavailable": // Только недоступные
			filtered = rep.Filter(func(status string) bool {
				return status == "offline"
			})
		default: // Весь список адресов
			filtered = rep.GetAll()
		}

		diffTime := time.Since(startTime)
		log.Printf("Данные /report выданы за %s", diffTime)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filtered)
	}
}
