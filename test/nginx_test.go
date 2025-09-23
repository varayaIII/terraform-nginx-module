package test

import (
  "testing"
  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/gruntwork-io/terratest/modules/k8s"
)

func TestNginxModule(t *testing.T) {
  kube := k8s.NewKubectlOptions("", "", "default")
  tf := &terraform.Options{ TerraformDir: "../examples/basic_usage" }
  defer terraform.Destroy(t, tf)
  terraform.InitAndApply(t, tf)
  k8s.RunKubectl(t, kube, "get", "deploy", "-A")
}
