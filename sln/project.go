package sln

import (
	"path"
	"strings"

	"github.com/atakanozceviz/vsdep/graph"
)

// Project hold information about a Visual Studio Project
type Project struct {
	Name string
	Path string
	Sln  string
	*Csproj
}

var g = graph.NewGraph()

// createGraph creates dependency graph for project.
func (project *Project) createGraph() error {
	g.AddNode(project.Name)
	for _, ig := range project.Csproj.ItemGroups {
		for _, pr := range ig.ProjectReferences {
			csprojFilePath := strings.Replace(path.Join(path.Dir(project.Path), pr.Include), "\\", "/", -1)
			csproj, err := parseCsproj(csprojFilePath)
			if err != nil {
				return err
			}

			fileName := path.Base(strings.Replace(pr.Include, "\\", "/", -1))
			fileNameNoExt := strings.Replace(fileName, path.Ext(fileName), "", -1)
			dep := &Project{
				Name:   fileNameNoExt,
				Path:   csprojFilePath,
				Csproj: csproj,
			}

			if project.Name == "" || dep.Name == "" {
				continue
			}

			g.AddEdge(project.Name, dep.Name)

			dep.createGraph()
		}
	}
	return nil
}

// IsTest return true if project is a test project, if not return false.
func (project *Project) IsTest() bool {
	for _, ig := range project.ItemGroups {
		for _, pr := range ig.PackageReferences {
			if pr.Include == "Microsoft.NET.Test.Sdk" {
				return true
			}
		}
	}
	return false
}
