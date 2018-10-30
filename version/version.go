package version

var (
	Version   string
	Revision  string
	BuildDate string
	GoVersion string
)

// Info returns map of version info
var Info = map[string]string{
	"version":   Version,
	"revision":  Revision,
	"buildDate": BuildDate,
	"goVersion": GoVersion,
}
