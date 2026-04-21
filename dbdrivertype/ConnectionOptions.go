package dbdrivertype

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreutils/stringutil"
)

type ConnectionOptions struct {
	Host, Port         string
	User, Password     string
	Options            string
	DbName             string
	IsSpecificDatabase bool
}

func (it ConnectionOptions) CreateMap() map[string]string {
	return map[string]string{
		"{db}":       it.DbName,
		"{ip}":       it.Host,
		"{port}":     it.Port,
		"{user}":     it.User,
		"{password}": it.User,
		"{?options}": it.Options,
	}
}

func (it ConnectionOptions) CreateMapUsingParams(
	host, port, dbName string,
	user, password, options string,
) map[string]string {
	return map[string]string{
		"{db}":       dbName,
		"{ip}":       host,
		"{port}":     port,
		"{user}":     user,
		"{password}": password,
		"{?options}": options,
	}
}

func (it ConnectionOptions) CreateMapUsingParamsNoOptions(
	host, port, dbName string,
	user, password string,
) map[string]string {
	return map[string]string{
		"{db}":       dbName,
		"{ip}":       host,
		"{port}":     port,
		"{user}":     user,
		"{password}": password,
		"{?options}": constants.EmptyString,
	}
}

func (it ConnectionOptions) CompileUsingConnectionFormat(
	format string,
) string {
	createdMap := it.CreateMap()

	return stringutil.ReplaceTemplate.DirectKeyUsingMap(
		format,
		createdMap)
}
