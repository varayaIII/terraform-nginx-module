<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5.0 |
| <a name="requirement_kubernetes"></a> [kubernetes](#requirement\_kubernetes) | ~> 2.25 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_kubernetes"></a> [kubernetes](#provider\_kubernetes) | 2.38.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [kubernetes_deployment.nginx](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/deployment) | resource |
| [kubernetes_service.nginx_svc](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_app_image"></a> [app\_image](#input\_app\_image) | La imagen de Docker a utilizar para el contenedor de Nginx. | `string` | `"nginx:1.25.1"` | no |
| <a name="input_app_name"></a> [app\_name](#input\_app\_name) | El nombre para la aplicación y las etiquetas. | `string` | n/a | yes |
| <a name="input_container_port"></a> [container\_port](#input\_container\_port) | El puerto que expone el contenedor. | `number` | `80` | no |
| <a name="input_namespace"></a> [namespace](#input\_namespace) | El namespace de Kubernetes donde se desplegará Nginx. | `string` | `"default"` | no |
| <a name="input_replicas"></a> [replicas](#input\_replicas) | El número de réplicas para el despliegue de Nginx. | `number` | `2` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_deployment_name"></a> [deployment\_name](#output\_deployment\_name) | Nombre del deployment creado |
| <a name="output_service_ip"></a> [service\_ip](#output\_service\_ip) | Dirección IP externa (LoadBalancer) del servicio |
| <a name="output_service_name"></a> [service\_name](#output\_service\_name) | Nombre del servicio creado |
<!-- END_TF_DOCS -->