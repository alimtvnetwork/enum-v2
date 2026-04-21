package linuxtype

var invalidVersionsMap = map[Variant]bool{
	Invalid:            true,
	UbuntuServer:       true,
	Centos:             true,
	CentosStream:       true,
	DebianDesktop:      true,
	Docker:             true,
	DockerUbuntuServer: true,
	Android:            true,
	UbuntuDesktop:      true,
}
