package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNginxModule(t *testing.T) {
	t.Parallel()

	// Configuración de Terraform
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic_usage",
		NoColor:      true,
	}

	// Limpieza al finalizar
	defer terraform.Destroy(t, terraformOptions)

	// Inicializar y aplicar Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Obtener outputs
	deploymentName := terraform.Output(t, terraformOptions, "deployment_name")
	serviceName := terraform.Output(t, terraformOptions, "service_name")

	// Configuración de kubectl
	kubectlOptions := k8s.NewKubectlOptions("", "", "default")

	// Verificar que el deployment existe
	t.Run("DeploymentExists", func(t *testing.T) {
		k8s.GetDeployment(t, kubectlOptions, deploymentName)
	})

	// Verificar que el service existe
	t.Run("ServiceExists", func(t *testing.T) {
		k8s.GetService(t, kubectlOptions, serviceName)
	})

	// Esperar a que el deployment esté listo
	t.Run("DeploymentReady", func(t *testing.T) {
		k8s.WaitUntilDeploymentAvailable(t, kubectlOptions, deploymentName, 60, 2*time.Second)
		deployment := k8s.GetDeployment(t, kubectlOptions, deploymentName)
		assert.Equal(t, int32(1), deployment.Status.ReadyReplicas)
	})

	// Verificar que los pods están corriendo
	t.Run("PodsRunning", func(t *testing.T) {
		filters := map[string]string{
			"app": "nginx-test-app",
		}
		pods := k8s.ListPods(t, kubectlOptions, filters)
		assert.True(t, len(pods) > 0, "Should have at least one pod")

		for _, pod := range pods {
			assert.Equal(t, "Running", string(pod.Status.Phase))
		}
	})

	// Verificar labels del deployment
	t.Run("DeploymentLabels", func(t *testing.T) {
		deployment := k8s.GetDeployment(t, kubectlOptions, deploymentName)
		assert.Equal(t, "nginx-test-app", deployment.Labels["app"])
	})

	// Listar todos los deployments (para debugging)
	t.Run("ListAllDeployments", func(t *testing.T) {
		output, err := k8s.RunKubectlAndGetOutputE(t, kubectlOptions, "get", "deploy", "-A")
		assert.NoError(t, err)
		fmt.Println("All deployments:")
		fmt.Println(output)
	})
}
