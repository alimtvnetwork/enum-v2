package creationtests

import (
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/reqtype"
	"gitlab.com/auk-go/enum/accesstype"
	"gitlab.com/auk-go/enum/brackets"
	"gitlab.com/auk-go/enum/certaction"
	"gitlab.com/auk-go/enum/cmdenumtypes/compresscmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/configcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/crontabscmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/decompresscmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/dnscmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/dockercmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/downloadcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/envpathcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/envvarscmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/ethernetcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/fail2bancmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/firewallcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/ftpcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/hostingplancmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/macrocmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/operatingsystemcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/packagecmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/rootcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/servicescmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/snapshotcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/sshcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/sslcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/sysgroupcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/toolingcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/usercmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/userrolecmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/webservercmdnames"
	"gitlab.com/auk-go/enum/completionstate"
	"gitlab.com/auk-go/enum/compressformats"
	"gitlab.com/auk-go/enum/compresslevels"
	"gitlab.com/auk-go/enum/configfilestate"
	"gitlab.com/auk-go/enum/conntrackstate"
	"gitlab.com/auk-go/enum/dbaction"
	"gitlab.com/auk-go/enum/dbdrivertype"
	"gitlab.com/auk-go/enum/dbexposetype"
	dbuserprivilegetype "gitlab.com/auk-go/enum/dbuserprivillegetype"
	"gitlab.com/auk-go/enum/eventtype"
	"gitlab.com/auk-go/enum/instructiontype"
	"gitlab.com/auk-go/enum/inttype"
	"gitlab.com/auk-go/enum/iptype"
	"gitlab.com/auk-go/enum/leveltype"
	"gitlab.com/auk-go/enum/licensetype"
	"gitlab.com/auk-go/enum/linescomparetype"
	"gitlab.com/auk-go/enum/linuxservicestate"
	"gitlab.com/auk-go/enum/linuxtype"
	"gitlab.com/auk-go/enum/linuxvendortype"
	"gitlab.com/auk-go/enum/logtype"
	"gitlab.com/auk-go/enum/nginxlogtype"
	"gitlab.com/auk-go/enum/onofftype"
	"gitlab.com/auk-go/enum/osarchs"
	"gitlab.com/auk-go/enum/osdetect"
	"gitlab.com/auk-go/enum/osgroupexecution"
	"gitlab.com/auk-go/enum/overwritetype"
	"gitlab.com/auk-go/enum/packageinstallmethod"
	"gitlab.com/auk-go/enum/pathpatterntype"
	"gitlab.com/auk-go/enum/promptclitype"
	"gitlab.com/auk-go/enum/protocoltype"
	"gitlab.com/auk-go/enum/querymethodtype"
	"gitlab.com/auk-go/enum/quotes"
	"gitlab.com/auk-go/enum/resauthtype"
	"gitlab.com/auk-go/enum/revokereason"
	"gitlab.com/auk-go/enum/runtype"
	"gitlab.com/auk-go/enum/scripttype"
	"gitlab.com/auk-go/enum/servicestate"
	"gitlab.com/auk-go/enum/sitestatetype"
	"gitlab.com/auk-go/enum/sqliteconnpathtype"
	"gitlab.com/auk-go/enum/sqljointype"
	"gitlab.com/auk-go/enum/strtype"
	"gitlab.com/auk-go/enum/taskcategory"
	"gitlab.com/auk-go/enum/taskpriority"
	"gitlab.com/auk-go/enum/timeunit"
	"gitlab.com/auk-go/enum/verifiertriggertype"
)

var allBasicEnumsCollection = [...]enuminf.BasicEnumer{
	&setterInvalid,
	reqtype.Invalid.AsBasicByteEnumContractsBinder(),
	stringcompareas.Invalid.AsBasicByteEnumContractsBinder(),
	accesstype.Invalid.AsBasicByteEnumContractsBinder(),
	brackets.Invalid.AsBasicByteEnumContractsBinder(),
	
	certaction.Invalid.AsBasicEnumContractsBinder(),
	
	completionstate.Invalid.AsBasicEnumContractsBinder(),
	
	compressformats.Invalid.AsBasicEnumContractsBinder(),
	compresslevels.Invalid.AsBasicEnumContractsBinder(),
	
	configfilestate.Invalid.AsBasicByteEnumContractsBinder(),
	conntrackstate.Invalid.AsBasicByteEnumContractsBinder(),
	
	dbaction.Invalid.AsBasicEnumContractsBinder(),
	dbexposetype.Invalid.AsBasicEnumContractsBinder(),
	dbdrivertype.Invalid.AsBasicByteEnumContractsBinder(),
	dbuserprivilegetype.Invalid.AsBasicEnumContractsBinder(),
	eventtype.Invalid.AsBasicByteEnumContractsBinder(),
	
	instructiontype.Invalid.AsBasicByteEnumContractsBinder(),
	inttype.Variant(0).AsBasicEnumer(),
	
	iptype.Invalid.AsBasicEnumContractsBinder(),
	
	leveltype.Invalid.AsBasicEnumContractsBinder(),
	licensetype.Invalid.AsBasicEnumContractsBinder(),
	linescomparetype.Invalid.AsBasicByteEnumContractsBinder(),
	linuxservicestate.Invalid.AsBasicByteEnumContractsBinder(),
	linuxtype.Invalid.AsBasicByteEnumContractsBinder(),
	linuxvendortype.Invalid.AsBasicByteEnumContractsBinder(),
	logtype.Invalid.AsBasicByteEnumContractsBinder(),
	
	nginxlogtype.Invalid.AsBasicEnumContractsBinder(),
	
	onofftype.Invalid.AsBasicEnumContractsBinder(),
	osarchs.Invalid.AsBasicByteEnumContractsBinder(),
	osgroupexecution.Invalid.AsBasicByteEnumContractsBinder(),
	osdetect.Invalid.AsBasicByteEnumContractsBinder(),
	overwritetype.Invalid.AsBasicByteEnumContractsBinder(),
	packageinstallmethod.Invalid.AsBasicEnumContractsBinder(),
	pathpatterntype.Invalid.AsBasicByteEnumContractsBinder(),
	promptclitype.Invalid.AsBasicByteEnumContractsBinder(),
	protocoltype.Invalid.AsBasicEnumContractsBinder(),
	querymethodtype.Invalid.AsBasicByteEnumContractsBinder(),
	quotes.Invalid.AsBasicByteEnumContractsBinder(),
	resauthtype.Invalid.AsBasicEnumContractsBinder(),
	revokereason.Unspecified.AsBasicEnumContractsBinder(),
	runtype.Invalid.AsBasicEnumContractsBinder(),
	scripttype.Invalid.AsBasicByteEnumContractsBinder(),
	
	servicestate.Invalid.AsBasicByteEnumContractsBinder(),
	
	sitestatetype.Invalid.AsBasicByteEnumContractsBinder(),
	
	sqliteconnpathtype.Invalid.AsBasicEnumContractsBinder(),
	sqljointype.Invalid.AsBasicEnumContractsBinder(),
	strtype.Variant("Invalid").AsBasicEnumer(),
	taskcategory.Invalid.AsBasicEnumContractsBinder(),
	taskpriority.Invalid.AsBasicEnumContractsBinder(),
	timeunit.Invalid.AsBasicEnumContractsBinder(),
	verifiertriggertype.Invalid.AsBasicEnumContractsBinder(),
	
	compresscmdnames.Invalid.AsBasicEnumContractsBinder(),
	configcmdnames.Invalid.AsBasicEnumContractsBinder(),
	crontabscmdnames.Invalid.AsBasicEnumContractsBinder(),
	decompresscmdnames.Invalid.AsBasicEnumContractsBinder(),
	dnscmdnames.Invalid.AsBasicEnumContractsBinder(),
	dockercmdnames.Invalid.AsBasicEnumContractsBinder(),
	envpathcmdnames.Invalid.AsBasicEnumContractsBinder(),
	envvarscmdnames.Invalid.AsBasicEnumContractsBinder(),
	ethernetcmdnames.Invalid.AsBasicEnumContractsBinder(),
	downloadcmdnames.Invalid.AsBasicEnumContractsBinder(),
	ethernetcmdnames.Invalid.AsBasicEnumContractsBinder(),
	fail2bancmdnames.Invalid.AsBasicEnumContractsBinder(),
	firewallcmdnames.Invalid.AsBasicEnumContractsBinder(),
	ftpcmdnames.Invalid.AsBasicEnumContractsBinder(),
	hostingplancmdnames.Invalid.AsBasicEnumContractsBinder(),
	macrocmdnames.Invalid.AsBasicEnumContractsBinder(),
	operatingsystemcmdnames.Invalid.AsBasicEnumContractsBinder(),
	packagecmdnames.Invalid.AsBasicEnumContractsBinder(),
	rootcmdnames.Invalid.AsBasicEnumContractsBinder(),
	servicescmdnames.Invalid.AsBasicEnumContractsBinder(),
	snapshotcmdnames.Invalid.AsBasicEnumContractsBinder(),
	sshcmdnames.Invalid.AsBasicEnumContractsBinder(),
	sslcmdnames.Invalid.AsBasicEnumContractsBinder(),
	sysgroupcmdnames.Invalid.AsBasicEnumContractsBinder(),
	toolingcmdnames.Invalid.AsBasicEnumContractsBinder(),
	usercmdnames.Invalid.AsBasicEnumContractsBinder(),
	userrolecmdnames.Invalid.AsBasicEnumContractsBinder(),
	webservercmdnames.Invalid.AsBasicEnumContractsBinder(),
}
