package test

import (
	"fmt"
	"testing"
	"time"

	// Librerías de Terratest para Terraform, Kubernetes y peticiones HTTP
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Nombre de la función de prueba
func TestTerraformNginxModule(t *testing.T) {
	t.Parallel()

	// Opciones de Terraform, apuntando al directorio
	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/basic_usage",
	}

	// 'defer' asegura que 'terraform destroy' se ejecute al final de la prueba
	defer terraform.Destroy(t, terraformOptions)

	// Ejecuta 'terraform init' y 'terraform apply'. Si falla, la prueba se detiene.
	terraform.InitAndApply(t, terraformOptions)

	// Opciones para conectarse a nuestro clúster de Kubernetes.
	kubectlOptions := k8s.NewKubectlOptions("", "", "default")

	// Terratest reintentará durante 5 minutos, cada 15 segundos.
	k8s.WaitUntilServiceAvailable(t, kubectlOptions, "nginx-test-app-service", 5, 15*time.Second)

	service := k8s.GetService(t, kubectlOptions, "nginx-test-app-service")

	// Obtenemos la URL para acceder al servicio.
	// Usamos un túnel local (forward) para que sea accesible desde nuestra máquina.
	url := k8s.GetServiceEndpoint(t, kubectlOptions, service, 80)

	// Hacemos una petición HTTP a la URL del servicio y verificamos que responde
	// con el código 200 (OK) y que el cuerpo del HTML contiene "Welcome to nginx!".
	// Terratest reintentará si hay fallos temporales.
	http_helper.HttpGetWithRetry(t, url, nil, 200, "Welcome to nginx!", 30, 5*time.Second)
}
