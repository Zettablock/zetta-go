package internal

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	goModFile              = "go.mod"
	pipelineYml            = "pipeline.yml"
	projectYml             = "project.yml"
	blockHandlersFile      = "block_handlers.go"
	eventHandlersFile      = "event_handlers.go"
	examplePipelineDirName = "example-pipeline"
)

//go:embed templates/project.yml.tmpl
var projectYmlTemplate string

//go:embed templates/pipeline.yml.tmpl
var pipelineYmlTemplate string

//go:embed templates/block_handlers.go.tmpl
var blockHandlersTemplate string

//go:embed templates/event_handlers.go.tmpl
var eventHandlersTemplate string

//go:embed templates/go.mod.tmpl
var goModTemplate string

// Project contains name, license and paths to projects.
type Project struct {
	PkgName      string
	AbsolutePath string
	Viper        bool
	AppName      string
}

func (p *Project) Create() error {
	var err error
	var mode os.FileMode = 0755
	// check if AbsolutePath exists
	if _, err = os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err = os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	}

	dir := filepath.Base(p.AbsolutePath)
	projectYmlTemplate = strings.Replace(projectYmlTemplate, "[project-name]", dir, -1)

	// create project.yml
	projectFileName := fmt.Sprintf("%s/%s", p.AbsolutePath, projectYml)
	err = os.WriteFile(projectFileName, []byte(projectYmlTemplate), mode)
	if err != nil {
		return err
	}

	// create example-pipeline dir
	examplePipelineDir := fmt.Sprintf("%s/%s", p.AbsolutePath, examplePipelineDirName)
	if _, err = os.Stat(examplePipelineDir); os.IsNotExist(err) {
		// create directory
		if err = os.Mkdir(examplePipelineDir, 0755); err != nil {
			return err
		}
	}

	exDir := filepath.Base(examplePipelineDir)
	pipelineYmlTemplate = strings.Replace(pipelineYmlTemplate, "[pipeline-name]", exDir, -1)

	// create pipeline.yml
	configFileName := fmt.Sprintf("%s/%s", examplePipelineDir, pipelineYml)
	err = os.WriteFile(configFileName, []byte(pipelineYmlTemplate), mode)
	if err != nil {
		return err
	}

	// create block_handlers.go
	blockHandlersFileName := fmt.Sprintf("%s/%s", examplePipelineDir, blockHandlersFile)
	err = os.WriteFile(blockHandlersFileName, []byte(blockHandlersTemplate), mode)
	if err != nil {
		return err
	}

	// create event_handlers.go
	eventHandlersFileName := fmt.Sprintf("%s/%s", examplePipelineDir, eventHandlersFile)
	err = os.WriteFile(eventHandlersFileName, []byte(eventHandlersTemplate), mode)
	if err != nil {
		return err
	}

	// create go.mod
	goModFileName := fmt.Sprintf("%s/%s", p.AbsolutePath, goModFile)
	err = os.WriteFile(goModFileName, []byte(goModTemplate), mode)
	if err != nil {
		return err
	}

	return nil
}
