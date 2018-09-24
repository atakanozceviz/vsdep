package changes

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

var commitRe = regexp.MustCompile(`(?m)(.+/)(.+)`)

// Get which projects changed after "lastcommit" (using git diff)
func Get(lastcommit string) (map[string]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	pth := ""

	if len(commitRe.FindAllStringSubmatch(lastcommit, -1)) >= 1 {
		pth = commitRe.FindAllStringSubmatch(lastcommit, -1)[0][1]
		lastcommit = commitRe.FindAllStringSubmatch(lastcommit, -1)[0][2]
	}
	wd = path.Join(wd, pth)

	cmd := exec.Command("git", "diff", "--name-only", lastcommit, "HEAD")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Dir = wd

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	dirs := make(map[string]struct{})

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		dirs[path.Dir(scanner.Text())] = struct{}{}
	}

	projects := make(map[string]string)

dirs:
	for p := range dirs {
		for p != "." {
			files, err := ioutil.ReadDir(path.Join(wd, p))
			if err != nil && os.IsExist(err) {
				return nil, err
			}

			for _, f := range files {
				if path.Ext(f.Name()) == ".csproj" {
					pn := strings.Split(path.Clean(p), "/")
					projects[pn[len(pn)-1]] = path.Clean(p)

					continue dirs
				}
			}

			p = path.Clean(p + "/..")

		}
	}
	return projects, nil
}
