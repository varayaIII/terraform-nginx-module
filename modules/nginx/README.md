# Terraform Module: Nginx for Kubernetes

Este módulo despliega un servidor Nginx en un cluster Kubernetes usando Terraform.  
Crea automáticamente un **Deployment** y un **Service** tipo LoadBalancer.

## Uso
```hcl
provider "kubernetes" {
  config_path = "~/.kube/config"
}

module "nginx_app" {
  source    = "github.com/varayaIII/terraform-nginx-module//modules/nginx"
  app_name  = "nginx-demo"
  replicas  = 2
  namespace = "production"
}

output "deployment_name" {
  value = module.nginx_app.deployment_name
}

output "service_ip" {
  value = module.nginx_app.service_ip
}
