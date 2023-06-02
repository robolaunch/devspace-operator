package node

import (
	"errors"
	"io"
	"net/http"
	"path/filepath"
	"reflect"
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
	Organization string               `yaml:"organization"`
	Repository   string               `yaml:"repository"`
	Domains      map[string][]Element `yaml:"domains"`
}

type Element struct {
	Application   Application   `yaml:"application"`
	DevSpaceImage DevSpaceImage `yaml:"devspace"`
}

type Application struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type DevSpaceImage struct {
	UbuntuDistro string `yaml:"ubuntuDistro"`
	Desktop      string `yaml:"desktop"`
	Version      string `yaml:"version"`
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

		chosenElement := Element{}
		if devspace.Spec.Environment.Domain == "plain" {

		} else {
			if domain, ok := imageProps.Domains[devspace.Spec.Environment.Domain]; ok {
				for _, element := range domain {
					if element.Application.Name != devspace.Spec.Environment.Application.Name {
						continue
					}

					if element.Application.Version != devspace.Spec.Environment.Application.Version {
						continue
					}

					if element.DevSpaceImage.UbuntuDistro != devspace.Spec.Environment.DevSpaceImage.UbuntuDistro {
						continue
					}

					if element.DevSpaceImage.Desktop != devspace.Spec.Environment.DevSpaceImage.Desktop {
						continue
					}

					if element.DevSpaceImage.Version != devspace.Spec.Environment.DevSpaceImage.Version {
						continue
					}

					chosenElement = element
					repository += "-" + devspace.Spec.Environment.Domain
					tagBuilder.WriteString(chosenElement.Application.Name + "-")
					tagBuilder.WriteString(chosenElement.Application.Version + "-")
					break
				}

				if reflect.DeepEqual(chosenElement, Element{}) {
					return "", errors.New("environment is not supported")
				}

			} else {
				return "", errors.New("domain is not supported")
			}
		}

		tagBuilder.WriteString(chosenElement.DevSpaceImage.UbuntuDistro + "-")
		tagBuilder.WriteString(chosenElement.DevSpaceImage.Desktop + "-")
		tagBuilder.WriteString(chosenElement.DevSpaceImage.Version)

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
