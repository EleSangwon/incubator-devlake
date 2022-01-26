package models

import (
	"github.com/merico-dev/lake/models/common"
)

// Please note that Issue Labels can also apply to Pull Requests.
// Pull Requests are considered Issues in GitHub.

type GithubIssueLabel struct {
	IssueId        int    `gorm:"primaryKey;autoIncrement:false"`
	IssueLabelName string `gorm:"primaryKey"`
	common.NoPKModel
}
