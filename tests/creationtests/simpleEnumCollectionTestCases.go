package creationtests

import (
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/issetter"
	"gitlab.com/auk-go/core/reqtype"
	"gitlab.com/auk-go/enum/accesstype"
	"gitlab.com/auk-go/enum/brackets"
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
	"gitlab.com/auk-go/enum/cmdenumtypes/toolingcmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/usercmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/userrolecmdnames"
	"gitlab.com/auk-go/enum/cmdenumtypes/webservercmdnames"
	"gitlab.com/auk-go/enum/completionstate"
	"gitlab.com/auk-go/enum/configfilestate"
	"gitlab.com/auk-go/enum/conntrackstate"
	"gitlab.com/auk-go/enum/dbaction"
	"gitlab.com/auk-go/enum/dbdrivertype"
	"gitlab.com/auk-go/enum/dbexposetype"
	"gitlab.com/auk-go/enum/dbuserprivillegetype"
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
	"gitlab.com/auk-go/enum/protocoltype"
	"gitlab.com/auk-go/enum/querymethodtype"
	"gitlab.com/auk-go/enum/quotes"
	"gitlab.com/auk-go/enum/resauthtype"
	"gitlab.com/auk-go/enum/revokereason"
	"gitlab.com/auk-go/enum/runtype"
	"gitlab.com/auk-go/enum/scripttype"
	"gitlab.com/auk-go/enum/servicestate"
	"gitlab.com/auk-go/enum/sqliteconnpathtype"
	"gitlab.com/auk-go/enum/sqljointype"
	"gitlab.com/auk-go/enum/strtype"
	"gitlab.com/auk-go/enum/taskcategory"
	"gitlab.com/auk-go/enum/taskpriority"
	"gitlab.com/auk-go/enum/timeunit"
	"gitlab.com/auk-go/enum/verifiertriggertype"
)

var simpleEnumCollectionTestCases = []enuminf.SimpleEnumer{
	issetter.Uninitialized,
	reqtype.Invalid,
	stringcompareas.Invalid.AsBasicEnumContractsBinder(),
	accesstype.Invalid,
	brackets.Invalid,
	
	completionstate.Invalid,
	
	configfilestate.Invalid,
	conntrackstate.Invalid,
	
	dbaction.Invalid,
	dbexposetype.Invalid,
	dbdrivertype.Invalid,
	
	dbuserprivilegetype.Invalid,
	eventtype.Invalid,
	
	instructiontype.Invalid,
	inttype.Invalid,
	iptype.Invalid,
	
	iptype.Invalid.AsBasicEnumContractsBinder(),
	
	leveltype.Invalid,
	licensetype.Invalid,
	linescomparetype.Invalid,
	linuxservicestate.Invalid,
	linuxtype.Invalid,
	linuxvendortype.Invalid,
	logtype.Invalid,
	
	nginxlogtype.Invalid.AsBasicEnumContractsBinder(),
	
	onofftype.Invalid,
	
	osarchs.Invalid,
	osgroupexecution.Invalid,
	osdetect.Invalid,
	overwritetype.Invalid,
	
	packageinstallmethod.Invalid,
	pathpatterntype.Invalid,
	protocoltype.Invalid,
	
	querymethodtype.Invalid,
	quotes.Invalid,
	
	resauthtype.Invalid,
	revokereason.Unspecified,
	runtype.Invalid,
	
	scripttype.Invalid,
	servicestate.Invalid,
	
	sqliteconnpathtype.Invalid,
	
	sqljointype.Invalid,
	strtype.Variant("Invalid"),
	taskcategory.Invalid,
	taskpriority.Invalid,
	
	timeunit.Invalid,
	verifiertriggertype.Invalid,
	
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
	toolingcmdnames.Invalid.AsBasicEnumContractsBinder(),
	usercmdnames.Invalid.AsBasicEnumContractsBinder(),
	userrolecmdnames.Invalid.AsBasicEnumContractsBinder(),
	webservercmdnames.Invalid.AsBasicEnumContractsBinder(),
}
