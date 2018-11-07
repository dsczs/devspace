package configure

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/covexo/devspace/pkg/devspace/builder/docker"
	"github.com/covexo/devspace/pkg/devspace/config/configutil"
	"github.com/covexo/devspace/pkg/devspace/config/v1"
	"github.com/covexo/devspace/pkg/util/log"
	"github.com/covexo/devspace/pkg/util/randutil"
	"github.com/covexo/devspace/pkg/util/stdinutil"
)

// Image configures the image name
func Image(dockerUsername string, skipQuestions bool) error {
	config := configutil.GetConfig()
	registryURL := "hub.docker.com"

	if !skipQuestions {
		registryURL = *stdinutil.GetFromStdin(&stdinutil.GetFromStdinParams{
			Question:               "Which registry do you want to push to? ('hub.docker.com' or URL)",
			DefaultValue:           "hub.docker.com",
			ValidationRegexPattern: "^.*$",
		})
	}

	if registryURL != "hub.docker.com" {
		imageBuilder, err := docker.NewBuilder(registryURL, "", "", false)
		if err == nil {
			log.StartWait("Checking Docker credentials")
			dockerAuthConfig, err := imageBuilder.Authenticate("", "", true)
			log.StopWait()

			if err != nil {
				return fmt.Errorf("Couldn't find credentials in credentials store. Make sure you login to the registry with: docker login %s", registryURL)
			}

			dockerUsername = dockerAuthConfig.Username
		}
	} else if dockerUsername == "" {
		return fmt.Errorf("Make sure you login to docker hub with: docker login")
	}

	googleRegistryRegex := regexp.MustCompile("^(.+\\.)?gcr.io$")
	isGoogleRegistry := googleRegistryRegex.Match([]byte(registryURL))
	isDockerHub := registryURL == "hub.docker.com"
	defaultImageName := ""
	createPullSecret := true

	if skipQuestions {
		defaultImageName = dockerUsername + "/devspace"
	} else {
		if isDockerHub {
			defaultImageName = *stdinutil.GetFromStdin(&stdinutil.GetFromStdinParams{
				Question:               "Which image name do you want to use on Docker Hub?",
				DefaultValue:           dockerUsername + "/devspace",
				ValidationRegexPattern: "^[a-zA-Z0-9/-]{4,60}$",
			})
		} else if isGoogleRegistry {
			project, err := exec.Command("gcloud", "config", "get-value", "project").Output()
			gcloudProject := "myGCloudProject"

			if err == nil {
				gcloudProject = strings.TrimSpace(string(project))
			}

			defaultImageName = *stdinutil.GetFromStdin(&stdinutil.GetFromStdinParams{
				Question:               "Which image name do you want to push to?",
				DefaultValue:           registryURL + "/" + gcloudProject + "/devspace",
				ValidationRegexPattern: "^.*$",
			})
		} else {
			defaultImageName = *stdinutil.GetFromStdin(&stdinutil.GetFromStdinParams{
				Question:               "Which image name do you want to push to?",
				DefaultValue:           registryURL + "/" + dockerUsername + "/devspace",
				ValidationRegexPattern: "^[a-zA-Z0-9\\./-]{4,90}$",
			})
		}

		createPullSecret = *stdinutil.GetFromStdin(&stdinutil.GetFromStdinParams{
			Question:               "Do you want to enable automatic creation of pull secrets for this image? (yes | no)",
			DefaultValue:           "yes",
			ValidationRegexPattern: "^(yes|no)$",
		}) == "yes"
	}

	imageMap := *config.Images
	imageMap["default"].Name = &defaultImageName
	imageMap["default"].CreatePullSecret = &createPullSecret

	return nil
}

// InternalRegistry configures the internal registry
func InternalRegistry() error {
	config := configutil.GetConfig()
	overwriteConfig := configutil.GetOverwriteConfig()

	imageMap := *config.Images
	defaultImageConf, defaultImageExists := imageMap["default"]
	if defaultImageExists {
		defaultImageConf.Registry = configutil.String("internal")
		defaultImageConf.CreatePullSecret = configutil.Bool(true)
	}

	overwriteRegistryMap := *overwriteConfig.Registries
	overwriteRegistryConfig, overwriteRegistryConfigFound := overwriteRegistryMap["internal"]
	if !overwriteRegistryConfigFound {
		overwriteRegistryConfig = &v1.RegistryConfig{
			Auth: &v1.RegistryAuth{},
		}
		overwriteRegistryMap["internal"] = overwriteRegistryConfig
	}

	registryAuth := overwriteRegistryConfig.Auth
	if registryAuth.Username == nil {
		randomUserSuffix, err := randutil.GenerateRandomString(5)
		if err != nil {
			return fmt.Errorf("Error creating random username: %s", err.Error())
		}

		registryAuth.Username = configutil.String("user-" + randomUserSuffix)
	}

	if registryAuth.Password == nil {
		randomPassword, err := randutil.GenerateRandomString(12)
		if err != nil {
			return fmt.Errorf("Error creating random password: %s", err.Error())
		}

		registryAuth.Password = &randomPassword
	}

	config.InternalRegistry = &v1.InternalRegistryConfig{
		Deploy: configutil.Bool(true),
	}
	config.Registries = &overwriteRegistryMap

	return nil
}
