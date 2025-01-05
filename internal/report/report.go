package report

import (
	"sync"
	"time"
)

type IPStatus struct {
	Status     string    `json:"status"`
	LastUpdate time.Time `json:"last_update"`
}

type Report struct {
	mu   sync.Mutex
	data map[string]IPStatus
}

func NewReport() *Report {
	return &Report{data: make(map[string]IPStatus)}
}

// Обновить данные IP
func (r *Report) Update(ip, status string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[ip] = IPStatus{
		Status:     status,
		LastUpdate: time.Now(),
	}
}

// Получить все данные в отчёт
func (r *Report) GetAll() map[string]IPStatus {
	r.mu.Lock()
	defer r.mu.Unlock()
	copy := make(map[string]IPStatus)
	for k, v := range r.data {
		copy[k] = v
	}
	return copy
}

// Получить все IP-адреса
func (r *Report) GetAllIPs() []string {
	r.mu.Lock()
	defer r.mu.Unlock()

	ips := make([]string, 0, len(r.data))
	for ip := range r.data {
		ips = append(ips, ip)
	}
	return ips
}

// Фильтрация отчёта по условию
func (r *Report) Filter(condition func(string) bool) map[string]IPStatus {
	r.mu.Lock()         // Блокировка
	defer r.mu.Unlock() // Разблокировка после обработки

	filtered := make(map[string]IPStatus)
	for ip, statusData := range r.data {
		if condition(statusData.Status) {
			filtered[ip] = statusData
		}
	}
	return filtered
}
