package sln

type Csproj struct {
	ItemGroups []ItemGroup `xml:"ItemGroup,omitempty"`
}

type ItemGroup struct {
	ProjectReferences []ProjectReference `xml:"ProjectReference,omitempty"`
	PackageReferences []PackageReference `xml:"PackageReference"`
}

type ProjectReference struct {
	Include string `xml:"Include,attr,omitempty"`
}
type PackageReference struct {
	Text    string `xml:",chardata"`
	Include string `xml:"Include,attr"`
	Version string `xml:"Version,attr"`
}
