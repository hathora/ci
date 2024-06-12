package shorthand

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hathora/ci/internal/sdk/models/shared"
)

var (
	containerPortShorthandRegex = regexp.MustCompile(`^(?:(?P<name>\S+)\:)?(?P<port>[0-9]+)(?:/(?P<transport>\S+))?$`)
	nameAllowedChars            = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	defaultTransportType        = shared.TransportTypeTCP
	maxPort                     = 65535
	minPort                     = 1
)

func ParseContainerPort(s string) (*shared.ContainerPort, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, fmt.Errorf("container port cannot be empty")
	}

	var port int
	transportType := shared.TransportTypeTCP
	var name string

	matches := containerPortShorthandRegex.FindStringSubmatch(s)
	if matches == nil {
		return nil, fmt.Errorf("invalid container port format: %s", s)
	}

	var err error
	for i, group := range containerPortShorthandRegex.SubexpNames() {
		if matches[i] == "" {
			continue
		}
		switch group {
		case "port":
			candidate := matches[i]
			foundPort, portErr := strconv.Atoi(candidate)
			port = foundPort
			err = errors.Join(err, portErr)
			if port < minPort || port > maxPort {
				err = errors.Join(err, fmt.Errorf("port outside of valid range (%d,%d): %d", minPort, maxPort, port))
			}
		case "transport":
			candidate := matches[i]
			switch candidate {
			case "tcp":
				transportType = defaultTransportType
			case "udp":
				transportType = shared.TransportTypeUDP
			case "tls":
				transportType = shared.TransportTypeTLS
			default:
				err = errors.Join(err, fmt.Errorf("invalid transport type: %s", candidate))
			}
		case "name":
			name = matches[i]
			if !nameAllowedChars.MatchString(name) {
				err = errors.Join(err, fmt.Errorf("invalid name: %s", name))
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return &shared.ContainerPort{
		Port:          port,
		TransportType: transportType,
		Name:          name,
	}, nil
}
