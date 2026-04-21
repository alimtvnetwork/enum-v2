package creationtests

import (
	"gitlab.com/auk-go/core/coreinterface/enuminf"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/issetter"
	"gitlab.com/auk-go/core/reqtype"
	"gitlab.com/auk-go/enum/accesstype"
	"gitlab.com/auk-go/enum/completionstate"
	"gitlab.com/auk-go/enum/dbaction"
	"gitlab.com/auk-go/enum/dbexposetype"
	"gitlab.com/auk-go/enum/eventtype"
	"gitlab.com/auk-go/enum/instructiontype"
	"gitlab.com/auk-go/enum/leveltype"
	"gitlab.com/auk-go/enum/licensetype"
	"gitlab.com/auk-go/enum/linuxservicestate"
	"gitlab.com/auk-go/enum/linuxtype"
	"gitlab.com/auk-go/enum/logtype"
	"gitlab.com/auk-go/enum/onofftype"
	"gitlab.com/auk-go/enum/osgroupexecution"
	"gitlab.com/auk-go/enum/overwritetype"
	"gitlab.com/auk-go/enum/pathpatterntype"
	"gitlab.com/auk-go/enum/resauthtype"
	"gitlab.com/auk-go/enum/revokereason"
	"gitlab.com/auk-go/enum/scripttype"
	"gitlab.com/auk-go/enum/servicestate"
	"gitlab.com/auk-go/enum/sqljointype"
	"gitlab.com/auk-go/enum/taskcategory"
	"gitlab.com/auk-go/enum/taskpriority"
)

var (
	bytesEnumContractsCollection = []enuminf.BasicEnumContractsBinder{
		reqtype.Invalid.AsBasicEnumContractsBinder(),
		stringcompareas.Invalid.AsBasicEnumContractsBinder(),
		accesstype.Invalid.AsBasicEnumContractsBinder(),
		completionstate.Invalid.AsBasicEnumContractsBinder(),
		dbaction.Invalid.AsBasicEnumContractsBinder(),
		dbexposetype.Invalid.AsBasicEnumContractsBinder(),
		eventtype.Invalid.AsBasicEnumContractsBinder(),
		instructiontype.Invalid.AsBasicEnumContractsBinder(),
		leveltype.Invalid.AsBasicEnumContractsBinder(),
		licensetype.Invalid.AsBasicEnumContractsBinder(),
		linuxservicestate.Invalid.AsBasicEnumContractsBinder(),
		linuxtype.Invalid.AsBasicEnumContractsBinder(),
		logtype.Invalid.AsBasicEnumContractsBinder(),
		onofftype.Invalid.AsBasicEnumContractsBinder(),
		osgroupexecution.Invalid.AsBasicEnumContractsBinder(),
		overwritetype.Invalid.AsBasicEnumContractsBinder(),
		pathpatterntype.Invalid.AsBasicEnumContractsBinder(),
		resauthtype.Invalid.AsBasicEnumContractsBinder(),
		revokereason.Unspecified.AsBasicEnumContractsBinder(),
		scripttype.Invalid.AsBasicEnumContractsBinder(),
		servicestate.Invalid.AsBasicEnumContractsBinder(),
		sqljointype.Invalid.AsBasicEnumContractsBinder(),
		taskcategory.Invalid.AsBasicEnumContractsBinder(),
		taskpriority.Invalid.AsBasicEnumContractsBinder(),
	}

	defaultOsScriptType = scripttype.OsDefaultScriptType()
	shellScriptType     = scripttype.Shell
	bashScriptType      = scripttype.Bash

	allScriptCreationTestCases = map[string]scripttype.Variant{
		"":                defaultOsScriptType,
		"def":             defaultOsScriptType,
		"default":         defaultOsScriptType,
		"Default":         defaultOsScriptType,
		"s":               shellScriptType,
		"sh":              shellScriptType,
		"Sh":              shellScriptType,
		"shell":           shellScriptType,
		"Shell":           shellScriptType,
		"/shell":          shellScriptType,
		"b":               bashScriptType,
		"bash":            bashScriptType,
		"Bash":            bashScriptType,
		"bh":              bashScriptType,
		"/bh":             bashScriptType,
		"Perl":            scripttype.Perl,
		"perl":            scripttype.Perl,
		"pl":              scripttype.Perl,
		"py":              scripttype.Python,
		"py2":             scripttype.Python2,
		"py3":             scripttype.Python3,
		"gcc":             scripttype.CLang,
		"gcc++":           scripttype.CLang,
		"c++":             scripttype.CLang,
		"CLang":           scripttype.CLang,
		"c":               scripttype.CLang,
		"Make":            scripttype.MakeScript,
		"MakeScript":      scripttype.MakeScript,
		"make":            scripttype.MakeScript,
		"m":               scripttype.MakeScript,
		"pw":              scripttype.Powershell,
		"pwsh":            scripttype.Powershell,
		"pwsh.exe":        scripttype.Powershell,
		"power":           scripttype.Powershell,
		"powershell":      scripttype.Powershell,
		"/powershell.exe": scripttype.Powershell,
		"Powershell":      scripttype.Powershell,
		"PowerShell":      scripttype.Powershell,
		"pshell":          scripttype.Powershell,
		"cmd":             scripttype.Cmd,
		"Cmd":             scripttype.Cmd,
		"dos":             scripttype.Cmd,
	}

	setterInvalid = issetter.Uninitialized
)
