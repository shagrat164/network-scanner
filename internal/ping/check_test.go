package ping

/*
func TestCheckIP(t *testing.T) {
	originalDial := DialFunc
	defer func() { DialFunc = originalDial }()

	DialFunc = func(network, address string, timeout time.Duration) (net.Conn, error) {
		if address == "127.0.0.1" {
			return &net.TCPConn{}, nil
		}
		return &net.TCPConn{}, errors.New("unreacheble")
	}

	tests := []struct {
		ip       string
		expected bool
	}{
		{"127.0.0.1", true},
		{"192.0.2.1", false},
	}

	for _, test := range tests {
		result := ScanIP(test.ip)
		if result != test.expected {
			t.Errorf("CheckIP(%s): ожидалось %v, получено %v", test.ip, test.expected, result)
		}
	}
}

func TestGenerateIPs(t *testing.T) {
	start := "192.168.0.1"
	end := "192.168.0.3"

	expected := []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"}
	result, err := GenerateIPs(start, end)

	if err != nil {
		t.Fatalf("Ошибка генерации IP: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Ожидалось %d IP, получено %d", len(expected), len(result))
	}

	for i, ip := range result {
		if ip != expected[i] {
			t.Errorf("Ожидалось %s, получено %s", expected[i], ip)
		}
	}
}

func TestParallelCheckWithPool(t *testing.T) {
	originalDial := DialFunc
	defer func() { DialFunc = originalDial }()

	DialFunc = func(network, address string, timeout time.Duration) (net.Conn, error) {
		if address == "127.0.0.1" {
			return &net.TCPConn{}, nil
		}
		return &net.TCPConn{}, errors.New("unreacheble")
	}

	ips := []string{"127.0.0.1", "192.0.2.1"}
	results := ParallelCheckWithPool(ips, 2)

	if results["127.0.0.1"] != "online" {
		t.Errorf("Ожидалось, что 127.0.0.1 будет online")
	}

	if results["192.0.2.1"] != "offline" {
		t.Errorf("Ожидалось, что 192.0.2.1 будет offline")
	}
}
*/
