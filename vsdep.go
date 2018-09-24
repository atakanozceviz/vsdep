package vsdep

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/atakanozceviz/vsdep/changes"
	"github.com/atakanozceviz/vsdep/graph"
	"github.com/atakanozceviz/vsdep/sln"
)

var ch = make(chan string)
var wg = &sync.WaitGroup{}

// FindOut which Visual Studio projects needs to build and
// which tests needs to run by checking differences between
// a git commit id and HEAD.
func FindOut(lastcommit string, walkpath ...string) (*Result, error) {
	var wp string
	switch {
	case len(walkpath) > 1:
		return nil, fmt.Errorf("walkpath cannot be more than one")
	case len(walkpath) == 1:
		wp = walkpath[0]
	default:
		wp = "."
	}

	affectedProjects, err := changes.Get(lastcommit)
	if err != nil {
		return nil, err
	}

	var paths []string
	if err := filepath.Walk(wp, func(pth string, info os.FileInfo, err error) error {
		if path.Ext(info.Name()) == ".sln" {
			paths = append(paths, pth)
		}
		return err
	}); err != nil {
		return nil, err
	}

	projects, g, err := sln.Open(paths...)
	if err != nil {
		return nil, err
	}

	for id := range affectedProjects {
		go findDeps(g, id)
		wg.Add(1)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	deps := make(map[string]struct{})
	testsNeedsToRun := make(map[string]string)
	solutionsNeedsToBuild := make(map[string]string)

	for dep := range ch {
		deps[dep] = struct{}{}
	}

	for projectName := range deps {
		for _, project := range projects {
			if project.Name == projectName {
				if project.IsTest() {
					testsNeedsToRun[project.Name] = project.Path
					continue
				}
				solutionsNeedsToBuild[path.Base(project.Sln)] = project.Sln
			}
		}
	}

	for projectName := range affectedProjects {
		for _, project := range projects {
			if project.Name == projectName {
				solutionsNeedsToBuild[path.Base(project.Sln)] = project.Sln
			}
		}
	}

	result := &Result{
		Solutions: solutionsNeedsToBuild,
		Tests:     testsNeedsToRun,
	}

	return result, nil
}

func findDeps(g *graph.Graph, id string) {
	defer wg.Done()
	for u, m := range g.Edges {
		for v := range m {
			if v == id {
				ch <- u
				wg.Add(1)
				go findDeps(g, u)
			}
		}
	}
}
