package flutter

type IFlutter interface {
	GetVersion() (*VersionInfo, error)
	RunDoctor() (string, error)
}
type Flutter struct{}
