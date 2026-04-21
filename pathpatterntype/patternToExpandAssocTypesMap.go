package pathpatterntype

var patternToExpandAssocTypesMap = map[Variant][]Variant{
	Invalid: {
		Invalid,
	},
	Root: {
		Root,
	},
	VariableDir: {
		VariableDir,
	},
	App: {
		App,
	},
	AppCore: {
		AppCore,
	},
	Id: {
		Id,
	},
	File: {
		File,
	},
	VarApp: {
		VariableDir,
		App,
	},
	AppConfigStore: {
		App,
		ConfigStore,
	},
	AppLog: {
		App,
		Log,
	},
	AppDb: {
		App,
		Database,
	},
	AppTest: {
		App,
		Test,
	},
	AppTemp: {
		App,
		Temp,
	},
	VarAppTemp: {
		VariableDir,
		App,
		Temp,
	},
	Relative: {
		Relative,
	},
	RelativeId: {
		Relative,
		Id,
	},
	RelativeIdFile: {
		Relative,
		Id,
		File,
	},
	AppRelative: {
		App,
		Relative,
	},
	AppRelativeId: {
		App,
		Relative,
		Id,
	},
	AppRelativeIdFile: {
		App,
		Relative,
		Id,
		File,
	},
	PrefixApp: {
		Prefix,
		App,
	},
	PrefixAppRelative: {
		Prefix,
		App,
		Relative,
	},
	PrefixAppRelativeId: {
		Prefix,
		App,
		Relative,
		Id,
	},
	PrefixAppRelativeIdFile: {
		Prefix,
		App,
		Relative,
		Id,
		File,
	},
	TempRoot: {
		TempRoot,
	},
	AppInstalled: {
		AppInstalled,
	},
	AppDownloads: {
		App,
		Downloads,
	},
	Home: {
		Home,
	},
	User: {
		User,
	},
	HomeUser: {
		Home,
		User,
	},
	HomeUserApp: {
		Home,
		User,
		App,
	},
	UsersRoot: {
		UsersRoot,
	},
	WebServerRoot: {
		WebServerRoot,
	},
	WebServerConfigsRoot: {
		WebServerConfigsRoot,
	},
	WebServerConfigsUsersRoot: {
		WebServerConfigsUsersRoot,
	},
	Packages: {
		Packages,
	},
	Instructions: {
		Instructions,
	},
	VarAppPackages: {
		VariableDir,
		App,
		Packages,
	},
	VarAppDownloads: {
		VariableDir,
		App,
		Downloads,
	},
	VarAppInstructions: {
		VariableDir,
		App,
		Instructions,
	},
	VarAppLog: {
		VariableDir,
		App,
		Log,
	},
	VarAppLogId: {
		VariableDir,
		App,
		Log,
		Id,
	},
	IdFile: {
		Id,
		File,
	},
	LogFile: {
		Log,
		File,
	},
	LogId: {
		Log,
		Id,
	},
	DbId: {
		Database,
		Id,
	},
	DbIdFile: {
		Database,
		Id,
		File,
	},
	ConfigStore: {
		ConfigStore,
	},
	Log: {
		Log,
	},
	Database: {
		Database,
	},
	Test: {
		Test,
	},
	Temp: {
		Temp,
	},
	Prefix: {
		Prefix,
	},
	Downloads: {
		Downloads,
	},
	Extension: {
		Extension,
	},
	FileWithExtension: {
		File,
		Extension,
	},
	Webserver: {
		Webserver,
	},
	PrefixAppFile: {
		Prefix,
		App,
		File,
	},
	AppFileWithExtension: {
		App,
		File,
		Extension,
	},
	PrefixAppFileWithExtension: {
		Prefix,
		App,
		File,
		Extension,
	},
	PrefixRelativeAppFileWithExtension: {
		Prefix,
		Relative,
		App,
		File,
		Extension,
	},
	UserTemp: {
		User,
		Temp,
	},
	Audit: {
		Audit,
	},
	LogDb: {
		Log,
		Database,
	},
	LogDbFile: {
		Log,
		Database,
		File,
	},
	LogAppDb: {
		Log,
		App,
		Database,
	},
	LogAppDbFile: {
		Log,
		App,
		Database,
		File,
	},
	TempUser: {
		Temp,
		User,
	},
	TempApp: {
		Temp,
		App,
	},
	TempAudit: {
		Temp,
		Audit,
	},
	LogApp: {
		Log,
		App,
	},
	LogAppFile: {
		Log,
		App,
		File,
	},
	Random: {
		Random,
	},
	RandomUuid: {
		RandomUuid,
	},
	RandomNumber: {
		RandomNumber,
	},
	AppDbFile: {
		App,
		Database,
		File,
	},
	AppDbRelativeFile: {
		App,
		Database,
		Relative,
		File,
	},
	AppDbRandom: {
		App,
		Database,
		Random,
	},
	AppDbRandomRelative: {
		App,
		Database,
		Random,
		Relative,
	},
	AppDbRandomRelativeFile: {
		App,
		Database,
		Random,
		Relative,
		File,
	},
	LogRandom: {
		Log,
		Random,
	},
	AppDbRandomFile: {
		App,
		Database,
		Random,
		File,
	},
	LogRandomFile: {
		Log,
		Random,
		File,
	},
	VarAppRandom: {
		VariableDir,
		App,
		Random,
	},
	VarAppRandomFile: {
		VariableDir,
		App,
		Random,
		File,
	},
	VarAppRandomRelative: {
		VariableDir,
		App,
		Random,
		Relative,
	},
	VarAppRandomRelativeFile: {
		VariableDir,
		App,
		Random,
		Relative,
		File,
	},
	Specific: {
		Specific,
	},
	Backup: {
		Backup,
	},
	HomeBackup: {
		Home,
		Backup,
	},
	BackupSpecific: {
		Backup,
		Specific,
	},
	BackupRelative: {
		Backup,
		Relative,
	},
	Ssl: {
		Ssl,
	},
	WebServerSsl: {
		Webserver,
		Ssl,
	},
	WebServerConfigUsers: {
		Webserver,
		Config,
		Users,
	},
	WebServerConfigUsersSpecific: {
		Webserver,
		Config,
		Users,
		Specific,
	},
	WebServerConfigSsl: {
		Webserver,
		Config,
		Ssl,
	},
	RelativeSsl: {
		Relative,
		Ssl,
	},
	VarAppBackup: {
		VariableDir,
		App,
		Backup,
	},
	VarAppBackupFile: {
		VariableDir,
		App,
		Backup,
		File,
	},
	VarAppBackupRandom: {
		VariableDir,
		App,
		Backup,
		Random,
	},
	VarAppBackupRandomFile: {
		VariableDir,
		App,
		Backup,
		Random,
		File,
	},
	VarAppBackupRandomRelative: {
		VariableDir,
		App,
		Backup,
		Random,
		Relative,
	},
	VarAppBackupRandomRelativeFile: {
		VariableDir,
		App,
		Backup,
		Random,
		Relative,
		File,
	},
	Config: {
		Config,
	},
	Users: {
		Users,
	},
	BackupFile: {
		Backup,
		File,
	},
	BackupRelativeFile: {
		Backup,
		Relative,
		File,
	},
	PrefixBackup: {
		Prefix,
		Backup,
	},
	PrefixBackupFile: {
		Prefix,
		Backup,
		File,
	},
	PrefixBackupRelativeFile: {
		Prefix,
		Backup,
		Relative,
		File,
	},
	PrefixBackupRandomRelativeFile: {
		Prefix,
		Backup,
		Random,
		Relative,
		File,
	},
}
