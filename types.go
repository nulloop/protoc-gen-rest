package main

type PackageInfo struct {
	PackageName string
	FileName    string
}

type ImportValues struct {
	Name string
	Path string
}

type ProtoFile struct {
	PackageName string
	Messages    []*MessageValues
	Services    []*ServiceValues
	Enums       []*EnumValues
	Imports     map[string]*ImportValues
}

type ServiceValues struct {
	Name    string
	Methods []*ServiceMethodValues
}

type ServiceMethodValues struct {
	Name       string
	InputType  string
	OutputType string
}

type EnumKeyVal struct {
	Name  string
	Value int32
}

type EnumValues struct {
	Name   string
	Values []*EnumKeyVal
}

type MessageValues struct {
	Name   string
	Fields []*FieldValues
}

type FieldValues struct {
	Name       string
	Field      string
	Type       string
	IsEnum     bool
	IsRepeated bool
}
