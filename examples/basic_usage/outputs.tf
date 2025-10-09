output "deployment_name" {
  description = "Nombre del deployment creado"
  value       = module.nginx_app.deployment_name
}

output "service_name" {
  description = "Nombre del servicio creado"
  value       = module.nginx_app.service_name
}

output "service_ip" {
  description = "IP externa del servicio LoadBalancer"
  value       = module.nginx_app.service_ip
}
