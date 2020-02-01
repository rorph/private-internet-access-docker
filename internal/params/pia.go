package params

import (
	"fmt"

	libparams "github.com/qdm12/golibs/params"
	"github.com/qdm12/private-internet-access-docker/internal/models"
)

// GetUser obtains the user to use to connect to the VPN servers
func (p *paramsReader) GetUser() (s string, err error) {
	defer func() {
		unsetenvErr := p.unsetEnv("USER")
		if err == nil {
			err = unsetenvErr
		}
	}()
	s, err = p.envParams.GetEnv("USER")
	if err != nil {
		return "", err
	} else if len(s) == 0 {
		return s, fmt.Errorf("USER environment variable cannot be empty")
	}
	return s, nil
}

// GetPassword obtains the password to use to connect to the VPN servers
func (p *paramsReader) GetPassword() (s string, err error) {
	defer func() {
		unsetenvErr := p.unsetEnv("PASSWORD")
		if err == nil {
			err = unsetenvErr
		}
	}()
	s, err = p.envParams.GetEnv("PASSWORD")
	if err != nil {
		return "", err
	} else if len(s) == 0 {
		return s, fmt.Errorf("PASSWORD environment variable cannot be empty")
	}
	return s, nil
}

// GetPortForwarding obtains if port forwarding on the VPN provider server
// side is enabled or not from the environment variable PORT_FORWARDING
func (p *paramsReader) GetPortForwarding() (activated bool, err error) {
	s, err := p.envParams.GetEnv("PORT_FORWARDING", libparams.Default("off"))
	if err != nil {
		return false, err
	}
	// Custom for retro-compatibility
	if s == "false" || s == "off" {
		return false, nil
	} else if s == "true" || s == "on" {
		return true, nil
	}
	return false, fmt.Errorf("PORT_FORWARDING can only be \"on\" or \"off\"")
}

// GetPortForwardingStatusFilepath obtains the port forwarding status file path
// from the environment variable PORT_FORWARDING_STATUS_FILE
func (p *paramsReader) GetPortForwardingStatusFilepath() (filepath models.Filepath, err error) {
	filepathStr, err := p.envParams.GetPath("PORT_FORWARDING_STATUS_FILE", libparams.Default("/forwarded_port"))
	return models.Filepath(filepathStr), err
}

// GetPIAEncryption obtains the encryption level for the PIA connection
// from the environment variable ENCRYPTION
func (p *paramsReader) GetPIAEncryption() (models.PIAEncryption, error) {
	s, err := p.envParams.GetValueIfInside("ENCRYPTION", []string{"normal", "strong"}, libparams.Default("strong"))
	return models.PIAEncryption(s), err
}

// GetPIARegion obtains the region for the PIA server from the
// environment variable REGION
func (p *paramsReader) GetPIARegion() (models.PIARegion, error) {
	s, err := p.envParams.GetValueIfInside("REGION", []string{"Netherlands"}, libparams.Compulsory())
	return models.PIARegion(s), err
}
