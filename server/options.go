package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"time"
)

var allowUnknownTopLevelField = int32(0)

type Logger struct {
	path string
	size string
	days uint8
}

// Options block for nats-server.
// NOTE: This structure is no longer used for monitoring endpoints
// and json tags are deprecated and may be removed in the future.
type Options struct {
	ConfigFile     string        `json:"-"`
	ServerName     string        `json:"server_name"`
	Host           string        `json:"addr"`
	Port           int           `json:"port"`
	Authorization  authorization `json:"-"`
	HTTPSPort      int           `json:"https_port"`
	AuthTimeout    float64       `json:"auth_timeout"`
	MaxControlLine int32         `json:"max_control_line"`
	MaxPayload     int32         `json:"max_payload"`
	MaxPending     int64         `json:"max_pending"`
	LogPath        string        `json:"-"`
	LogSizeLimit   int64         `json:"-"`
	TLS            bool          `json:"-"`
	TLSConfig      *tls.Config   `json:"-"`
	AllowNonTLS    bool          `json:"-"`
	inConfig       map[string]bool
	inCmdLine      map[string]bool
}

type authorization struct {
	user  string
	pass  string
	token string
}

// TLSConfigOpts holds the parsed tls config information,
// used with flag parsing
type TLSConfigOpts struct {
	CertFile          string
	KeyFile           string
	CaFile            string
	Verify            bool
	Insecure          bool
	Map               bool
	TLSCheckKnownURLs bool
	Timeout           float64
}

func ProcessConfigFile(configFile string) (*Options, error) {
	opts := &Options{}
	return opts, nil
}

func (o *Options) processConfigFileLine(k string, v interface{}, errors *[]error, warnings *[]error) {
	switch strings.ToLower(k) {
	case "listen":
		hp, err := parseListen(v)
		if err != nil {

			return
		}
		o.Host = hp.host
		o.Port = hp.port
	case "client_advertise":

	case "port":
		o.Port = int(v.(int64))
	case "server_name":
		o.ServerName = v.(string)
	case "host", "net":
		o.Host = v.(string)

	default:
		if au := atomic.LoadInt32(&allowUnknownTopLevelField); au == 0 && !tk.IsUsedVariable() {

		}
	}
}

type hostPort struct {
	host string
	port int
}

func parseListen(v interface{}) (*hostPort, error) {
	hp := &hostPort{}
	switch vv := v.(type) {
	// Only a port
	case int64:
		hp.port = int(vv)
	case string:
		host, port, err := net.SplitHostPort(vv)
		if err != nil {
			return nil, fmt.Errorf("could not parse address string %q", vv)
		}
		hp.port, err = strconv.Atoi(port)
		if err != nil {
			return nil, fmt.Errorf("could not parse port %q", port)
		}
		hp.host = host
	default:
		return nil, fmt.Errorf("expected port or host:port, got %T", vv)
	}
	return hp, nil
}

func parseURLs(a []interface{}, typ string) (urls []*url.URL, errors []error) {
	urls = make([]*url.URL, 0, len(a))

	for _, u := range a {
		sURL := u.(string)
		url, err := parseURL(sURL, typ)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		urls = append(urls, url)
	}
	return urls, errors
}

func parseURL(u string, typ string) (*url.URL, error) {
	urlStr := strings.TrimSpace(u)
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s url [%q]", typ, urlStr)
	}
	return url, nil
}

//将命令行的配置覆盖文件配置
func MergeOptions(fileOpts, flagOpts *Options) *Options {
	if fileOpts == nil {
		return flagOpts
	}
	if flagOpts == nil {
		return fileOpts
	}
	// Merge the two, flagOpts override
	opts := *fileOpts

	if flagOpts.Port != 0 {
		opts.Port = flagOpts.Port
	}
	if flagOpts.Host != "" {
		opts.Host = flagOpts.Host
	}
	return &opts
}

func isIPInList(list1 []net.IP, list2 []net.IP) bool {
	for _, ip1 := range list1 {
		for _, ip2 := range list2 {
			if ip1.Equal(ip2) {
				return true
			}
		}
	}
	return false
}

func getURLIP(ipStr string) ([]net.IP, error) {
	ipList := []net.IP{}

	ip := net.ParseIP(ipStr)
	if ip != nil {
		ipList = append(ipList, ip)
		return ipList, nil
	}

	hostAddr, err := net.LookupHost(ipStr)
	if err != nil {
		return nil, fmt.Errorf("Error looking up host with route hostname: %v", err)
	}
	for _, addr := range hostAddr {
		ip = net.ParseIP(addr)
		if ip != nil {
			ipList = append(ipList, ip)
		}
	}
	return ipList, nil
}

func getInterfaceIPs() ([]net.IP, error) {
	var localIPs []net.IP

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		return nil, fmt.Errorf("Error getting self referencing address: %v", err)
	}

	for i := 0; i < len(interfaceAddr); i++ {
		interfaceIP, _, _ := net.ParseCIDR(interfaceAddr[i].String())
		if net.ParseIP(interfaceIP.String()) != nil {
			localIPs = append(localIPs, interfaceIP)
		} else {
			return nil, fmt.Errorf("Error parsing self referencing address: %v", err)
		}
	}
	return localIPs, nil
}
func ProcessCommandLineArgs(cmd *flag.FlagSet) (showVersion bool, showHelp bool, err error) {
	if len(cmd.Args()) > 0 {
		arg := cmd.Args()[0]
		switch strings.ToLower(arg) {
		case "version":
			return true, false, nil
		case "help":
			return false, true, nil
		default:
			return false, false, fmt.Errorf("unrecognized command: %q", arg)
		}
	}
	return false, false, nil
}
func ConfigureOptions(fs *flag.FlagSet, args []string, printVersion func()) (*Options, error) {
	opts := &Options{}
	var (
		showVersion            bool
		showHelp               bool
		showTLSHelp            bool
		signal                 string
		configFile             string
		dbgAndTrace            bool
		trcAndVerboseTrc       bool
		dbgAndTrcAndVerboseTrc bool
		err                    error
	)
	fs.BoolVar(&showHelp, "h", false, "Show this message.")
	fs.BoolVar(&showHelp, "help", false, "Show this message.")
	fs.IntVar(&opts.Port, "port", 0, "Port to listen on.")
	fs.IntVar(&opts.Port, "p", 0, "Port to listen on.")
	fs.StringVar(&opts.ServerName, "n", "", "Server name.")
	fs.StringVar(&opts.ServerName, "name", "", "Server name.")
	fs.StringVar(&opts.ServerName, "server_name", "", "Server name.")
	fs.StringVar(&opts.Host, "addr", "", "Network host to listen on.")
	fs.StringVar(&opts.Host, "a", "", "Network host to listen on.")
	fs.StringVar(&opts.Host, "net", "", "Network host to listen on.")
	fs.BoolVar(&trcAndVerboseTrc, "VV", false, "Enable Verbose Trace logging. (Traces system account as well)")
	fs.BoolVar(&dbgAndTrace, "DV", false, "Enable Debug and Trace logging.")
	fs.BoolVar(&dbgAndTrcAndVerboseTrc, "DVV", false, "Enable Debug and Verbose Trace logging. (Traces system account as well)")
	fs.IntVar(&opts.HTTPPort, "m", 0, "HTTP Port for /varz, /connz endpoints.")
	fs.IntVar(&opts.HTTPPort, "http_port", 0, "HTTP Port for /varz, /connz endpoints.")
	fs.IntVar(&opts.HTTPSPort, "ms", 0, "HTTPS Port for /varz, /connz endpoints.")
	fs.IntVar(&opts.HTTPSPort, "https_port", 0, "HTTPS Port for /varz, /connz endpoints.")
	fs.StringVar(&configFile, "c", "", "Configuration file.")
	fs.StringVar(&configFile, "config", "", "Configuration file.")
	fs.StringVar(&signal, "sl", "", "Send signal to nats-server process (stop, quit, reopen, reload).")
	fs.StringVar(&signal, "signal", "", "Send signal to nats-server process (stop, quit, reopen, reload).")

	fs.Int64Var(&opts.LogSizeLimit, "log_size_limit", 0, "Logfile size limit being auto-rotated")

	fs.BoolVar(&showVersion, "version", false, "Print version information.")
	fs.BoolVar(&showVersion, "v", false, "Print version information.")

	fs.BoolVar(&showTLSHelp, "help_tls", false, "TLS help.")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	if showVersion {
		printVersion()
		return nil, nil
	}

	// Process args looking for non-flag options,
	// 'version' and 'help' only for now
	showVersion, showHelp, err = ProcessCommandLineArgs(fs)
	if err != nil {
		return nil, err
	} else if showVersion {
		printVersion()
		return nil, nil
	}

	// Process signal control.
	if signal != "" {
		if err := processSignal(signal); err != nil {
			return nil, err
		}
	}

	// Parse config if given
	if configFile != "" {
		// This will update the options with values from the config file.
		err := opts.ProcessConfigFile(configFile)
		if err != nil {

			fmt.Fprint(os.Stderr, err)
		}

		// Call this again to override config file options with options from command line.
		// Note: We don't need to check error here since if there was an error, it would
		// have been caught the first time this function was called (after setting up the
		// flags).
		fs.Parse(args)
	}

	// Special handling of some flags
	var (
		flagErr     error
		tlsDisabled bool
		tlsOverride bool
	)
	fs.Visit(func(f *flag.Flag) {
		// short-circuit if an error was encountered
		if flagErr != nil {
			return
		}
		if strings.HasPrefix(f.Name, "tls") {
			if f.Name == "tls" {
				if !opts.TLS {
					// User has specified "-tls=false", we need to disable TLS
					opts.TLSConfig = nil
					tlsDisabled = true
					tlsOverride = false
					return
				}
				tlsOverride = true
			} else if !tlsDisabled {
				tlsOverride = true
			}
		} else {
			switch f.Name {

			}
		}
	})
	if flagErr != nil {
		return nil, flagErr
	}

	return opts, nil
}

func normalizeBasePath(p string) string {
	if len(p) == 0 {
		return "/"
	}
	// add leading slash
	if p[0] != '/' {
		p = "/" + p
	}
	return path.Clean(p)
}

func ProcessSignal(command Command, pidStr string) error {
	var pid int
	if pidStr == "" {
		pids, err := resolvePids()
		if err != nil {
			return err
		}
		if len(pids) == 0 {
			return fmt.Errorf("no %s processes running", processName)
		}
		if len(pids) > 1 {
			errStr := fmt.Sprintf("multiple %s processes running:\n", processName)
			prefix := ""
			for _, p := range pids {
				errStr += fmt.Sprintf("%s%d", prefix, p)
				prefix = "\n"
			}
			return errors.New(errStr)
		}
		pid = pids[0]
	} else {
		p, err := strconv.Atoi(pidStr)
		if err != nil {
			return fmt.Errorf("invalid pid: %s", pidStr)
		}
		pid = p
	}

	var err error
	switch command {
	case CommandStop:
		err = kill(pid, syscall.SIGKILL)
	case CommandQuit:
		err = kill(pid, syscall.SIGINT)
	case CommandReopen:
		err = kill(pid, syscall.SIGUSR1)
	case CommandReload:
		err = kill(pid, syscall.SIGHUP)
	case commandLDMode:
		err = kill(pid, syscall.SIGUSR2)
	case commandTerm:
		err = kill(pid, syscall.SIGTERM)
	default:
		err = fmt.Errorf("unknown signal %q", command)
	}
	return err
}
func processSignal(signal string) error {
	var (
		pid           string
		commandAndPid = strings.Split(signal, "=")
	)
	if l := len(commandAndPid); l == 2 {
		pid = maybeReadPidFile(commandAndPid[1])
	} else if l > 2 {
		return fmt.Errorf("invalid signal parameters: %v", commandAndPid[2:])
	}
	if err := ProcessSignal(Command(commandAndPid[0]), pid); err != nil {
		return err
	}
	os.Exit(0)
	return nil
}

// maybeReadPidFile returns a PID or Windows service name obtained via the following method:
// 1. Try to open a file with path "pidStr" (absolute or relative).
// 2. If such a file exists and can be read, return its contents.
// 3. Otherwise, return the original "pidStr" string.
func maybeReadPidFile(pidStr string) string {
	if b, err := ioutil.ReadFile(pidStr); err == nil {
		return string(b)
	}
	return pidStr
}

func homeDir() (string, error) {
	if runtime.GOOS == "windows" {
		homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH")
		userProfile := os.Getenv("USERPROFILE")

		home := filepath.Join(homeDrive, homePath)
		if homeDrive == "" || homePath == "" {
			if userProfile == "" {
				return "", errors.New("nats: failed to get home dir, require %HOMEDRIVE% and %HOMEPATH% or %USERPROFILE%")
			}
			home = userProfile
		}

		return home, nil
	}

	home := os.Getenv("HOME")
	if home == "" {
		return "", errors.New("failed to get home dir, require $HOME")
	}
	return home, nil
}

func expandPath(p string) (string, error) {
	p = os.ExpandEnv(p)

	if !strings.HasPrefix(p, "~") {
		return p, nil
	}

	home, err := homeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, p[1:]), nil
}
