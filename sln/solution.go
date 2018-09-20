package sln

import (
	"github.com/atakanozceviz/vsdep/graph"
)

// Open sln file(s), parse projects, create and return dependency graph and all parsed projects.
func Open(path ...string) ([]*Project, *graph.Graph, error) {
	allProjects := make([]*Project, 0, 10)
	for _, pth := range path {
		projects, err := parseSln(pth)
		if err != nil {
			return nil, nil, err
		}

		err = createProjectGraph(projects...)
		if err != nil {
			return nil, nil, err
		}
		allProjects = append(allProjects, projects...)
	}
	return allProjects, g, nil
}

// createProjectGraph add projects to graph
func createProjectGraph(projects ...*Project) error {
	for _, project := range projects {
		if err := project.createGraph(); err != nil {
			return err
		}
	}
	return nil
}
