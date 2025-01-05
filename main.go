package main

import (
	"log"
	"net/http"
	"time"

	"network-scanner/internal/config"
	"network-scanner/internal/handlers"
	"network-scanner/internal/ping"
	"network-scanner/internal/report"

	"github.com/natefinch/lumberjack"
)

func main() {
	// Настройка ротации логов
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/server.log",
		MaxSize:    5,    // Размер в мегабайтах
		MaxBackups: 3,    // Количество резервных копий
		MaxAge:     28,   // Срок хранения (в днях)
		Compress:   true, // Сжатие старых логов
	})

	cfg, err := config.LoadConfig("cfg/config.json")

	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
		return
	}

	startIP := cfg.IPRange[0]
	endIP := cfg.IPRange[1]
	threads := cfg.MaxThreads
	interval := time.Duration(cfg.ScanInterval) * time.Second

	rep := report.NewReport()

	// Запуск периодического обновления конфигурации
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			updateConfig, err := config.LoadConfig("cfg/config.json")
			if err != nil {
				log.Printf("Ошибка обновления конфигурации: %v\n", err)
			}
			cfg = updateConfig
			log.Println("Конфигурация обновлена")
		}
	}()

	// Запуск периодической проверки
	go func() {
		for {
			ips, _ := ping.GenerateIPs(startIP, endIP)
			results := ping.ParallelCheckWithPool(ips, threads)
			for ip, status := range results {
				rep.Update(ip, status)
			}

			time.Sleep(interval)
		}
	}()

	// Маршруты
	http.HandleFunc("/", handlers.HomeHandler())
	http.HandleFunc("/about", handlers.AboutHandler())
	http.HandleFunc("/report", handlers.GetFilteredReportHandler(rep))
	http.HandleFunc("/refresh", handlers.RefreshHandler(rep, cfg))
	http.HandleFunc("/reload-config", handlers.ReloadConfigHandler(cfg))

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
