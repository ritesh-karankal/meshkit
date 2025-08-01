package catalog

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/meshery/meshkit/models/catalog/v1alpha1"
)

func BuildArtifactHubPkg(name, downloadURL, user, version string, createdAt *time.Time, catalogData *v1alpha1.CatalogData) *ArtifactHubMetadata {
	var createdTime time.Time
	if createdAt != nil {
		createdTime = *createdAt
	} else {
		createdTime = time.Now()
	}

	artifacthubPkg := &ArtifactHubMetadata{
		Name:        toKebabCase(name),
		DisplayName: name,
		Description: valueOrElse(catalogData.PatternInfo, "A Meshery Design"),
		Provider: Provider{
			Name: user,
		},
		Links: []Link{
			{
				Name: "download",
				URL:  downloadURL, // this depends on where the design is stored by the user, we can give remote provider URL otherwise
			},
			{
				Name: "Meshery Catalog",
				URL:  "https://meshery.io/catalog",
			},
		},
		HomeURL:   "https://docs.meshery.io/concepts/logical/designs",
		Version:   valueOrElse(version, "0.0.1"),
		CreatedAt: createdTime.Format(time.RFC3339),
		License:   "Apache-2.0",
		LogoURL:   "https://raw.githubusercontent.com/meshery/meshery.io/0b8585231c6e2b3251d38f749259360491c9ee6b/assets/images/brand/meshery-logo.svg",
		Install:   "mesheryctl design import -f",
		Readme:    fmt.Sprintf("%s \n ##h4 Caveats and Consideration \n", catalogData.PatternCaveats),
	}

	if len(catalogData.SnapshotURL) > 0 {
		artifacthubPkg.Screenshots = append(artifacthubPkg.Screenshots, Screenshot{
			Title: "MeshMap Snapshot",
			URL:   catalogData.SnapshotURL[0],
		})

		if len(catalogData.SnapshotURL) > 1 {
			artifacthubPkg.Screenshots = append(artifacthubPkg.Screenshots, Screenshot{
				Title: "MeshMap Snapshot",
				URL:   catalogData.SnapshotURL[1],
			})
		}
	}

	artifacthubPkg.Screenshots = append(artifacthubPkg.Screenshots, Screenshot{
		Title: "Meshery Project",
		URL:   "https://raw.githubusercontent.com/meshery/meshery.io/master/assets/images/logos/meshery-gradient.png",
	})

	return artifacthubPkg
}

func valueOrElse(s string, fallback string) string {
	if len(s) == 0 {
		return fallback
	} else {
		return s
	}
}

func toKebabCase(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	s = strings.ReplaceAll(s, " ", "-")

	return s
}
