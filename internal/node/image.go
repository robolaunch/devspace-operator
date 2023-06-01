package node

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/robolaunch/devspace-operator/internal"
	devv1alpha1 "github.com/robolaunch/devspace-operator/pkg/api/roboscale.io/v1alpha1"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
)

type Platform struct {
	Versions []Version `yaml:"versions"`
}

type Version struct {
	Date     string   `yaml:"date"`
	Version  string   `yaml:"version"`
	DevCloud DevCloud `yaml:"devCloud"`
}

type DevCloud struct {
	Kubernetes Kubernetes `yaml:"kubernetes"`
}

type Kubernetes struct {
	Operators Operators `yaml:"operators"`
}

type Operators struct {
	DevSpaceOperator DevSpaceOperator `yaml:"devspace"`
}

type DevSpaceOperator struct {
	Images Images `yaml:"images"`
}

type Images struct {
	Organization string   `yaml:"organization"`
	Repository   string   `yaml:"repository"`
	Tags         []string `yaml:"tags"`
}

// Not used in devspace manifest, needed for internal use.
type ReadyDevSpaceProperties struct {
	Enabled bool
	Image   string
}

func GetReadyDevSpaceProperties(devspace devv1alpha1.DevSpace) ReadyDevSpaceProperties {
	labels := devspace.GetLabels()

	if user, hasUser := labels[internal.DEVSPACE_IMAGE_USER]; hasUser {
		if repository, hasRepository := labels[internal.DEVSPACE_IMAGE_REPOSITORY]; hasRepository {
			if tag, hasTag := labels[internal.DEVSPACE_IMAGE_TAG]; hasTag {
				return ReadyDevSpaceProperties{
					Enabled: true,
					Image:   user + "/" + repository + ":" + tag,
				}
			}
		}
	}

	return ReadyDevSpaceProperties{
		Enabled: false,
	}
}

func GetImage(node corev1.Node, devspace devv1alpha1.DevSpace) (string, error) {

	var imageBuilder strings.Builder
	var tagBuilder strings.Builder

	readyDevSpace := GetReadyDevSpaceProperties(devspace)

	if readyDevSpace.Enabled {

		imageBuilder.WriteString(readyDevSpace.Image)

	} else {

		platformVersion := GetPlatformVersion(node)
		imageProps, err := getImageProps(platformVersion)
		if err != nil {
			return "", err
		}

		organization := imageProps.Organization
		repository := imageProps.Repository

		tagBuilder.WriteString(string(devspace.Spec.UbuntuDistro))

		hasGPU := HasGPU(node)

		if hasGPU {
			tagBuilder.WriteString("-xfce") // TODO: make desktop selectable

		} else {
			tagBuilder.WriteString("-xfce") // TODO: make desktop selectable
		}

		// get latest tag
		tagBuilder.WriteString("-" + imageProps.Tags[0])

		imageBuilder.WriteString(filepath.Join(organization, repository) + ":")
		imageBuilder.WriteString(tagBuilder.String())

	}

	return imageBuilder.String(), nil

}

func getImageProps(platformVersion string) (Images, error) {

	resp, err := http.Get("https://raw.githubusercontent.com/robolaunch/robolaunch/main/platform.yaml")
	if err != nil {
		return Images{}, err
	}

	defer resp.Body.Close()

	var yamlFile []byte
	if resp.StatusCode == http.StatusOK {
		yamlFile, err = io.ReadAll(resp.Body)
		if err != nil {
			return Images{}, err
		}
	}

	var platform Platform
	err = yaml.Unmarshal(yamlFile, &platform)
	if err != nil {
		return Images{}, err
	}

	var imageProps Images
	for _, v := range platform.Versions {
		if v.Version == platformVersion {
			imageProps = v.DevCloud.Kubernetes.Operators.DevSpaceOperator.Images
		}
	}

	return imageProps, nil
}
