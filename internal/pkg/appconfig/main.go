package appconfig

// Server is a structure containing a single server configuration
type Server struct {
	URL      string
	Login    string
	Password string
	BaseJQL  string
	CSVOut   string
}

// RequestParams is a structure
type RequestParams struct {
}

// Config is a container for all stored configuration. This structure can either be prepared manually or read from a configuration file
type Config struct {
	Servers   []Server
	TargetDir string
}

// GetDummy is a temporal function to fill the Config with pre-hardcoded data
func GetDummy() *Config {
	//TODO: Add reading from the actual config file
	//https://www.thepolyglotdeveloper.com/2017/04/load-json-configuration-file-golang-application/

	var server1 Server
	server1.URL = "https://tandbergdata.atlassian.net"
	server1.Login = "e.lavnikevich@sam-solutions.com"
	server1.Password = "Js6Us47UtB78qcsY9[qP"
	server1.BaseJQL = "PROJECT in (RDX, VTX2U, VTX1U) AND worklogAuthor in (a.hrytsevich, v.redzhepov, vshakhov) AND timespent is not EMPTY"

	var config Config
	config.Servers = append(config.Servers, server1)
	config.TargetDir = "/Users/lava/Desktop"

	return &config
}
