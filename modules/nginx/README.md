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

Recursos Creados
Este módulo crea los siguientes recursos en Kubernetes:

Deployment: Despliega los pods de Nginx con las réplicas especificadas
Service (LoadBalancer): Expone la aplicación externamente

Características

✅ Health checks configurados (liveness y readiness probes)
✅ Límites de recursos (CPU y memoria)
✅ Soporte para múltiples réplicas
✅ Namespace personalizable
✅ Imagen de Nginx configurable

Requirements
NameVersionterraform>= 1.5.0kubernetes~> 2.25
Providers
NameVersionkubernetes~> 2.25
Inputs
NameDescriptionTypeDefaultRequiredapp_nameEl nombre para la aplicación y las etiquetasstringn/ayesnamespaceEl namespace de Kubernetes donde se desplegará Nginxstring"default"noapp_imageLa imagen de Docker a utilizar para el contenedor de Nginxstring"nginx:1.25.1"noreplicasEl número de réplicas para el despliegue de Nginxnumber2nocontainer_portEl puerto que expone el contenedornumber80no
Outputs
NameDescriptiondeployment_nameNombre del deployment creadoservice_nameNombre del servicio creadoservice_ipDirección IP externa (LoadBalancer) del servicio
Ejemplos
Despliegue básico
hclmodule "nginx_basic" {
  source   = "github.com/varayaIII/terraform-nginx-module//modules/nginx"
  app_name = "nginx-basic"
}
Despliegue en producción
hclmodule "nginx_prod" {
  source    = "github.com/varayaIII/terraform-nginx-module//modules/nginx"
  app_name  = "nginx-production"
  namespace = "production"
  replicas  = 5
  app_image = "nginx:1.26.0-alpine"
}
Múltiples instancias
hclmodule "nginx_dev" {
  source    = "github.com/varayaIII/terraform-nginx-module//modules/nginx"
  app_name  = "nginx-dev"
  namespace = "development"
  replicas  = 1
}

module "nginx_staging" {
  source    = "github.com/varayaIII/terraform-nginx-module//modules/nginx"
  app_name  = "nginx-staging"
  namespace = "staging"
  replicas  = 3
}
Notas

El Service LoadBalancer puede tardar unos minutos en obtener una IP externa dependiendo de tu proveedor cloud
Los health checks están configurados con un delay inicial de 5 segundos
Los límites de recursos están optimizados para cargas ligeras-medias. Ajusta según tus necesidades

License
MIT
