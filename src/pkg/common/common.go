package common

import (
	"context"
	"fmt"
	"os"

	"github.com/defenseunicorns/lula/src/pkg/domains/api"
	kube "github.com/defenseunicorns/lula/src/pkg/domains/kubernetes"
	"github.com/defenseunicorns/lula/src/pkg/message"
	"github.com/defenseunicorns/lula/src/pkg/providers/kyverno"
	"github.com/defenseunicorns/lula/src/pkg/providers/opa"
	"github.com/defenseunicorns/lula/src/types"
	goversion "github.com/hashicorp/go-version"

	"sigs.k8s.io/yaml"
)

func ReadFileToBytes(path string) ([]byte, error) {
	var data []byte
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return data, fmt.Errorf("Path: %v does not exist - unable to digest document", path)
	}
	data, err = os.ReadFile(path)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Returns version validity
func IsVersionValid(versionConstraint string, version string) (bool, error) {
	if version == "unset" {
		// Default cli version is "unset", enabling users to run directly from source code
		// This is not a valid version, but we want to allow it for development purposes
		return true, nil
	}

	currentVersion, err := goversion.NewVersion(version)
	if err != nil {
		return false, err
	}
	constraints, err := goversion.NewConstraint(versionConstraint)
	if err != nil {
		return false, err
	}
	if constraints.Check(currentVersion) {
		return true, nil
	}
	return false, nil
}

// Get the domain and providers
func GetDomain(domain Domain, ctx context.Context) types.Domain {
	switch domain.Type {
	case "kubernetes":
		return kube.KubernetesDomain{
			Context: ctx,
			Spec:    domain.KubernetesSpec,
		}
	case "api":
		return api.ApiDomain{
			Spec: domain.ApiSpec,
		}
	default:
		return nil
	}
}

func GetProvider(provider Provider, ctx context.Context) types.Provider {
	switch provider.Type {
	case "opa":
		return opa.OpaProvider{
			Context: ctx,
			Spec:    provider.OpaSpec,
		}
	case "kyverno":
		return kyverno.KyvernoProvider{
			Context: ctx,
			Spec:    provider.KyvernoSpec,
		}
	default:
		return nil
	}
}

// Converts a raw string to a Validation object (string -> common.Validation -> types.Validation)
func ValidationFromString(raw string) (validation types.LulaValidation, err error) {
	if raw == "" {
		return validation, fmt.Errorf("validation string is empty")
	}

	var validationData Validation

	err = yaml.Unmarshal([]byte(raw), &validationData)
	if err != nil {
		message.Fatalf(err, "error unmarshalling yaml: %s", err.Error())
		return validation, err
	}

	validation, err = validationData.ToLulaValidation()
	if err != nil {
		return validation, err
	}

	return validation, nil
}
