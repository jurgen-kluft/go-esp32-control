package wifi

// @see: https://www.arduino.cc/en/Reference/WiFi

// IPAddress ...
type IPAddress struct {
	A int
	B int
	C int
	D int
}

const (
	EncryptionTypeAuto    = 8
	EncryptionTypeCCMP    = 4
	EncryptionTypeNone    = 7
	EncryptionTypeTKIP    = 2
	EncryptionTypeWEP     = 5
	MaxSocketNum          = 4096
	StatusConnected       = 3
	StatusConnectionLost  = 4
	StatusConnectFailed   = 5
	StatusDisconnected    = 6
	StatusIdle            = 0
	StatusNoShield        = 255
	StatusNoSSIDAvailable = 1
	StatusScanCompleted   = 2
)

var (
	CurrentBSSID          = []int{0, 0, 0, 0, 0, 0, 0, 0}
	CurrentDNS            = []int{0, 0, 0, 0}
	CurrentEncryptionType = 7
	CurrentGateway        = &IPAddress{127, 0, 0, 255}
	CurrentLocalIP        = &IPAddress{127, 0, 0, 1}
	CurrentNetworks       = 0
	CurrentRSSI           = -1
	CurrentSSID           = ""
	CurrentStatus         = StatusIdle
	SocketPort            = make(map[int]int, MaxSocketNum)
	SocketState           = make(map[int]int, MaxSocketNum)
)

// BSSID gets the MAC address of the routher you are connected to.
// @see: https://www.arduino.cc/en/Reference/WiFiBSSID
func BSSID() []int {
	return CurrentBSSID
}

// Begin initializes the WiFi library's network settings and provides the current status.
// @see: https://www.arduino.cc/en/Reference/WiFiBegin
func Begin(ssid string) {
	CurrentRSSI = 0
	CurrentSSID = ssid
	CurrentStatus = StatusConnected
}

// BeginEncrypted initializes the WiFi library's network settings and provides the current status.
// @see: https://www.arduino.cc/en/Reference/WiFiBegin
func BeginEncrypted(ssid, passphrase string) {
	CurrentRSSI = 0
	CurrentSSID = ssid
	CurrentStatus = StatusConnected
}

// Disconnect disconnects the WiFi shield from the current network.
// @see: https://www.arduino.cc/en/Reference/WiFiDisconnect
func Disconnect() {
	CurrentStatus = StatusIdle
}

// EncryptionType gets the encryption type of the current network.
// @see: https://www.arduino.cc/en/Reference/WiFiEncryptionType
func EncryptionType() int {
	return CurrentEncryptionType
}

// HostByName ...
func HostByName(hostname string, addr string) int {
	return 0
}

// LocalIP gets the WiFi shield's IP address.
// @see: https://www.arduino.cc/en/Reference/WiFiLocalIP
func LocalIP() *IPAddress {
	return CurrentLocalIP
}

// RSSI gets the signal strength of the connection to the router.
// @see: https://www.arduino.cc/en/Reference/WiFiRSSI
func RSSI() int {
	return CurrentRSSI
}

// ScanNetworks scans for available WiFi networks and returns the discovered number.
// @see: https://www.arduino.cc/en/Reference/WiFiScanNetworks
func ScanNetworks() int {
	return CurrentNetworks
}

// SetDNS allows you to configure the DNS (Domain Name System) server.
// @see: https://www.arduino.cc/en/Reference/WiFiSetDns
func SetDNS(dns []int) {
	CurrentDNS = dns
}

// Status returns the connection status.
// @see: https://www.arduino.cc/en/Reference/WiFiStatus
func Status() int {
	return CurrentStatus
}

// SSID gets the SSID of the current network.
// @see: https://www.arduino.cc/en/Reference/WiFiSSID
func SSID() string {
	return CurrentSSID
}
