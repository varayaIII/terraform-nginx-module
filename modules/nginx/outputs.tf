output "deployment_name" {
  description = "Nombre del deployment creado"
  value       = kubernetes_deployment.nginx.metadata[0].name
}

output "service_name" {
  description = "Nombre del servicio creado"
  value       = kubernetes_service.nginx_svc.metadata[0].name
}

output "service_ip" {
  description = "Dirección IP externa (LoadBalancer) del servicio"
  value = try(
    kubernetes_service.nginx_svc.status[0].load_balancer[0].ingress[0].ip,
    kubernetes_service.nginx_svc.status[0].load_balancer[0].ingress[0].hostname,
    "pending"
  )
}
