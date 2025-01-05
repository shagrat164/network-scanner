package ping

import (
	"log"
	"net"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

// Проверка доступности IP
func CheckIP(ip string) bool {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		log.Fatalf("Ошибка создания пингера для %s: %v\n", ip, err)
		return false
	}

	pinger.Count = 2
	pinger.SetPrivileged(false)
	pinger.Timeout = 200 * time.Millisecond
	err = pinger.Run()
	if err != nil {
		log.Fatalf("Ошибка пинга %s: %v\n", ip, err)
		return false
	}

	return pinger.Statistics().PacketsRecv > 0
}

func ScanIP(ip string) bool {
	conn, err := net.DialTimeout("ip4:icmp", ip, 1*time.Second)
	if err != nil {
		log.Println(err)
		return false
	}
	conn.Close()
	return true
}

func ParallelCheckWithPool(ips []string, poolSize int) map[string]string {
	results := make(map[string]string, len(ips))
	var mu sync.Mutex
	var wg sync.WaitGroup

	ipChan := make(chan string, len(ips))

	startTime := time.Now()
	// Запуск воркеров
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ip := range ipChan {
				status := "offline"
				if CheckIP(ip) {
					status = "online"
				}
				mu.Lock()
				results[ip] = status
				mu.Unlock()
			}
		}()
	}

	go func() {
		for _, ip := range ips {
			ipChan <- ip
		}
		close(ipChan)
	}()

	wg.Wait()

	diffTime := time.Since(startTime)
	log.Printf("Цикл сканирования диапазона выполнен за %s", diffTime)

	return results
}
