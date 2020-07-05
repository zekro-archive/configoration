# configoration

[![GoDoc](https://godoc.org/github.com/zekroTJA/configoration?status.svg)](https://pkg.go.dev/github.com/zekroTJA/configoration) [![Actions](https://github.com/zekroTJA/configoration/workflows/Go/badge.svg)](https://github.com/zekroTJA/configoration/actions) [![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/zekroTJA/configoration)](https://github.com/zekroTJA/configoration/releases)

> go get github.com/zekroTJA/configoration

## Details

confi**go**ration *(yeah, it's no typo)* is another configuration interface library which provides a similar API to the [Configuration Extension](https://docs.microsoft.com/en-us/dotnet/api/microsoft.extensions.configuration?view=dotnet-plat-ext-3.1) of .NET Platform Extensions 3.1 providing a [`ConfigurationBuilder`](https://docs.microsoft.com/en-us/dotnet/api/microsoft.extensions.configuration.configurationbuilder?view=dotnet-plat-ext-3.1) like builder struct which produces a Section instance which is inspired by the [`IConfiguration`](https://docs.microsoft.com/en-us/dotnet/api/microsoft.extensions.configuration.iconfiguration?view=dotnet-plat-ext-3.1) interface of .NET.

## Example

```go
c, err := configoration.NewBuilder().
	SetBasePath("./testdata").
	AddJsonFile("config.json", true).
	AddJsonFile("config.Development.json", true).
	AddYamlFile("config.yaml", true).
	AddEnvironmentVariables("APP_", false).
	Build()

if err != nil {
	panic(err)
}

tls := c.GetSection("WebServer:TLS")
enable, ok := tls.GetBool("enable")
```
