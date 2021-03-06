package unitycloudbuild

import (
	"fmt"
	"time"
)

type BuildTarget struct {
	Name     string               `json:"name"`
	Platform string               `json:"platform"`
	Id       string               `json:"buildtargetid"`
	Enabled  bool                 `json:"enabled"`
	Builds   []Build              `json:"builds,omitempty"`
	Settings *BuildTargetSettings `json:"settings,omitempty"`
}

type BuildTargetSettings struct {
	AutoBuild      bool   `json:"autoBuild"`
	ExecutableName string `json:"executablename"`
	Scm            struct {
		Branch       string `json:"branch"`
		Subdirectory string `json:"subdirectory,omitempty"`
		Type         string `json:"type"`
	} `json:"scm"`
	UnityVersion string      `json:"unityVersion"`
	Advanced     interface{} `json:"advanced,omitempty"`
}

type Build struct {
	Number            int           `json:"build"`
	TargetId          string        `json:"buildTargetId"`
	TargetName        string        `json:"buildTargetName"`
	GUID              string        `json:"buildGUID,omitempty"` // NOTE: On unfinished builds this is empty!
	Created           time.Time     `json:"created"`
	Status            string        `json:"buildStatus"`
	Finished          time.Time     `json:"finished"`
	Platform          string        `json:"platform"`
	TotalTimeSeconds  float64       `json:"totalTimeInSeconds"`
	BuildTimeSeconds  float64       `json:"buildTimeInSeconds"`
	Links             Links         `json:"links"`
	ScmBranch         string        `json:"scmBranch"`
	LastBuiltRevision string        `json:"lastBuiltRevision,omitempty"`
	Changesets        []interface{} `json:"changeset,omitempty"`
	UnityVersion      string        `json:"unityVersion"`
}

func (b *Build) UniqueId() string {
	return fmt.Sprintf("%s-#%d", b.TargetId, b.Number)
}

type BuildAttempt struct {
	Build
	FailureDetails interface{} `json:"failureDetails,omitempty"`
	Error          string      `json:"error,omitempty"`
}

type Changeset struct {
}

type Links struct {
	Artifacts       []Artifact `json:"artifacts,omitempty"`
	Self            *Link      `json:"self,omitempty"`
	Log             *Link      `json:"log,omitempty"`
	AuditLog        *Link      `json:"auditlog,omitempty"`
	DownloadPrimary *Link      `json:"download_primary,omitempty"`
}

type Link struct {
	Method string `json:"method"`
	Href   string `json:"href"`
	Meta   struct {
		Type string `json:"type,omitempty"`
	} `json:"meta"`
}

type File struct {
	Filename string `json:"filename"`
	Href     string `json:"href"`
	Size     int64  `json:"size"`
}

type Artifact struct {
	Files []File `json:"files,omitempty"`
	Key   string `json:"key"`
	Name  string `json:"name"`
}

type GitCommit struct {
	Revision string `json:"revision"`
	Message  string `json:"message,omitempty"`
}
