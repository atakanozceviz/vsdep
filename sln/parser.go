package sln

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

var projectPattern = regexp.MustCompile(`(?m)Project\(\"(.*?)\"\)\s+=\s+\"(.*?)\",\s+"(.*?)\",\s+\"(.*?)\"`)

// parseSln open sln file, parse information from that file and return projects.
func parseSln(pth string) ([]*Project, error) {
	var projects []*Project

	slnFile, err := os.Open(pth)
	if err != nil {
		return nil, err
	}
	defer slnFile.Close()

	sln, err := ioutil.ReadAll(slnFile)
	if err != nil {
		return nil, err
	}

	for _, v := range projectPattern.FindAllStringSubmatch(string(sln), -1) {
		if len(v) >= 4 {
			if !strings.Contains(v[3], ".csproj") {
				continue
			}
			csprojPath := path.Clean(strings.Replace(path.Join(path.Dir(pth), v[3]), "\\", "/", -1))

			csproj, err := parseCsproj(csprojPath)
			if err != nil {
				return nil, err
			}

			project := &Project{
				Name:   v[2],
				Path:   csprojPath,
				Sln:    pth,
				Csproj: csproj,
			}

			projects = append(projects, project)
		}
	}
	return projects, nil
}

// parseCsproj open csproj file, unmarshal and return.
func parseCsproj(pth string) (*Csproj, error) {
	csproj := &Csproj{}

	csprojFile, err := os.Open(pth)
	if err != nil {
		return csproj, err
	}
	defer csprojFile.Close()

	csprojBytes, err := ioutil.ReadAll(csprojFile)
	if err != nil {
		return csproj, err
	}

	if err := xml.Unmarshal(csprojBytes, csproj); err != nil {
		return csproj, err
	}
	return csproj, nil
}
