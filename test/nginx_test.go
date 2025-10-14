package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNginxModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic_usage",
		NoColor:      true,
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	deploymentName := terraform.Output(t, terraformOptions, "deployment_name")
	serviceName := terraform.Output(t, terraformOptions, "service_name")

	kubectlOptions := k8s.NewKubectlOptions("", "", "default")

	t.Run("DeploymentExists", func(t *testing.T) {
		deployment := k8s.GetDeployment(t, kubectlOptions, deploymentName)
		assert.NotNil(t, deployment)
	})

	t.Run("ServiceExists", func(t *testing.T) {
		service := k8s.GetService(t, kubectlOptions, serviceName)
		assert.NotNil(t, service)
	})

	t.Run("DeploymentReady", func(t *testing.T) {
		k8s.WaitUntilDeploymentAvailable(t, kubectlOptions, deploymentName, 60, 2)
	})
}
