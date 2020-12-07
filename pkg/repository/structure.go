package repository

import (
	"time"
)

type HarborRepository []struct {
	ArtifactCount int       `json:"artifact_count"`
	CreationTime  time.Time `json:"creation_time"`
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	ProjectID     int       `json:"project_id"`
	PullCount     int       `json:"pull_count"`
	UpdateTime    time.Time `json:"update_time"`
}

type HarborArtifact []struct {
	AdditionLink      map[string]interface{} `json:"addition_links,omitempty"`
	Digest            string                 `json:"digest"`
	ExtraAttrs        map[string]interface{} `json:"extra_attrs,omitempty"`
	Icon              string                 `json:"icon"`
	ID                int                    `json:"id"`
	Label             string                 `json:"labels,omitempty"`
	ManifestMediaType string                 `json:"manifest_media_type"`
	MediaType         string                 `json:"media_type"`
	ProjectID         int                    `json:"project_id"`
	PullTime          time.Time              `json:"pull_time"`
	PushTime          time.Time              `json:"push_time"`
	References        string                 `json:"references,omitempty"`
	RepositoryID      int                    `json:"repository_id"`
	Size              int                    `json:"size"`
	Tags              []TagArtifact          `json:"tags"`
	Type              string                 `json:"type"`
}

type TagArtifact struct {
	ArtifactID   int       `json:"artifact_id"`
	ID           int       `json:"id"`
	Immutable    string    `json:"immutable"`
	Name         string    `json:"name"`
	PullTime     time.Time `json:"pull_time"`
	PushTime     time.Time `json:"push_time"`
	RepositoryID int       `json:"repository_id"`
	Signed       string    `json:"signed"`
}
