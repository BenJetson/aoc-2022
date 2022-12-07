package day07

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/BenJetson/aoc-2022/aoc"
)

type FileSize int

type Directory struct {
	Name string

	Parent   *Directory
	Children map[string]*Directory

	Files map[string]FileSize
}

func NewRootDirectory() *Directory {
	return &Directory{
		Children: make(map[string]*Directory),
		Files:    make(map[string]FileSize),
	}
}

func (d *Directory) TotalSize() (total FileSize) {
	for _, fs := range d.Files {
		total += fs
	}
	for _, sd := range d.Children {
		total += sd.TotalSize()
	}
	return
}

func (d *Directory) Mkdir(name string) (new *Directory) {
	new = &Directory{
		Name: name,

		Parent:   d,
		Children: make(map[string]*Directory),

		Files: make(map[string]FileSize),
	}

	d.Children[name] = new

	return
}

func (d *Directory) AddFile(name string, size FileSize) {
	d.Files[name] = size
}

func (d *Directory) FullyQualifiedName() string {
	if d.Parent == nil {
		return "/"
	}

	return d.Parent.FullyQualifiedName() + "/" + d.Name
}

func (d *Directory) ContentListing(level int) (out string) {
	for _, child := range d.Children {
		out += fmt.Sprintf("%s- %s (dir)\n",
			strings.Repeat("  ", level),
			child.Name,
		)
		out += child.ContentListing(level + 1)
	}

	for fileName, fileSize := range d.Files {
		out += fmt.Sprintf("%s- %s (file, size=%d)\n",
			strings.Repeat("  ", level),
			fileName,
			fileSize,
		)
	}
	return
}

func (d *Directory) String() string {
	return d.ContentListing(0)
}

func ReadShellSession(input aoc.Input) (root *Directory, err error) {
	root = NewRootDirectory()
	pwd := root

	for _, line := range input {
		tokens := strings.Split(line, " ")
		if len(tokens) < 2 {
			continue
		}

		switch tokens[0] {
		case "$":
			cmd := tokens[1]
			if cmd == "cd" && len(tokens) > 2 {
				target := tokens[2]

				switch target {
				case "/":
					pwd = root
				case "..":
					if pwd.Parent == nil {
						err = errors.New("cannnot walk up from root directory")
						return
					}

					pwd = pwd.Parent
				default:
					var ok bool
					if pwd, ok = pwd.Children[target]; !ok {
						err = fmt.Errorf("no such directory %s found in %s",
							target, pwd.FullyQualifiedName())
						return
					}
				}

			}
		case "dir":
			dirName := tokens[1]
			pwd.Mkdir(dirName)
		default:
			sizeStr := tokens[0]
			fileName := tokens[1]

			var size int
			if size, err = strconv.Atoi(sizeStr); err != nil {
				err = fmt.Errorf("invalid size for file %s in directory %s",
					fileName, pwd.FullyQualifiedName())
				return
			}

			pwd.AddFile(fileName, FileSize(size))
		}
	}

	return
}

func FindDirectoriesSmallerThan(
	threshold FileSize,
	pwd *Directory,
) (out []*Directory) {

	for _, child := range pwd.Children {
		if child.TotalSize() <= threshold {
			out = append(out, child)
		}
		out = append(out, FindDirectoriesSmallerThan(threshold, child)...)
	}

	return
}

type FileSystem struct {
	Root *Directory

	Capacity    FileSize
	DesiredFree FileSize
}

func (fs *FileSystem) AmountUsed() FileSize {
	return fs.Root.TotalSize()
}

func (fs *FileSystem) AmountFree() FileSize {
	return fs.Capacity - fs.AmountUsed()
}

func (fs *FileSystem) AmountToFreeForDesired() FileSize {
	if fs.AmountFree() < fs.DesiredFree {
		return fs.DesiredFree - fs.AmountFree()
	}
	return 0
}

func (fs *FileSystem) FindDeletionCandidate() (target *Directory) {
	minSize := fs.AmountToFreeForDesired()
	allDirectories := FindDirectoriesSmallerThan(fs.Capacity, fs.Root)
	sort.Slice(allDirectories, func(i, j int) bool {
		return allDirectories[i].TotalSize() < allDirectories[j].TotalSize()
	})

	for _, pwd := range allDirectories {
		if pwd.TotalSize() >= minSize {
			target = pwd
			break
		}
	}

	return
}

func SolvePuzzle(input aoc.Input) (s aoc.Solution, err error) {
	root, err := ReadShellSession(input)
	if err != nil {
		err = fmt.Errorf("could not read shell session: %w", err)
		return
	}

	smallDirectories := FindDirectoriesSmallerThan(100000, root)
	var total FileSize
	for _, pwd := range smallDirectories {
		total += pwd.TotalSize()
	}

	s.Part1.SaveIntAnswer(int(total))

	fs := FileSystem{Root: root, Capacity: 70000000, DesiredFree: 30000000}

	dirToDelete := fs.FindDeletionCandidate()
	if dirToDelete == nil {
		err = errors.New("no suitable candidate for deletion found")
		return
	}

	s.Part2.SaveIntAnswer(int(dirToDelete.TotalSize()))

	return
}
