package flutter

type VersionInfo struct {
	Version  string `json:"version"`
	Channel  string `json:"channel"`
	Dart     string `json:"dart"`
	DevTools string `json:"dev_tools"`
}
