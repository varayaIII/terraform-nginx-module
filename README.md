![ci](https://github.com/varayaIII/terraform-nginx-module/actions/workflows/ci.yml/badge.svg)
![e2e](https://github.com/varayaIII/terraform-nginx-module/actions/workflows/e2e.yml/badge.svg)

# Terraform NGINX Module

Módulo Terraform reutilizable para desplegar Nginx en Kubernetes, demostrando prácticas avanzadas de IaC y validación con pruebas automatizadas (Terratest). Diseñado para flujos de GitOps.

## 📋 Características

- ✅ Despliegue automatizado de Nginx en Kubernetes
- ✅ Service tipo LoadBalancer configurado
- ✅ Health checks (liveness y readiness probes)
- ✅ Límites de recursos configurables
- ✅ Validación con Terratest
- ✅ CI/CD con GitHub Actions
- ✅ Documentación autogenerada con terraform-docs

## 🚀 Uso Rápido
```hcl
provider "kubernetes" {
  config_path = "~/.kube/config"
}

module "nginx_app" {
  source    = "github.com/varayaIII/terraform-nginx-module//modules/nginx"
  app_name  = "mi-nginx"
  replicas  = 3
  namespace = "production"
}

output "service_ip" {
  value = module.nginx_app.service_ip
}
