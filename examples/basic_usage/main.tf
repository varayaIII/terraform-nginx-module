module "nginx_app" {
  # La ruta apunta al módulo local
  source = "../../../modules/nginx"

  # Asignamos un nombre
  app_name = "nginx-test-app"
  
  # Se usa solo una réplica para que la prueba sea más rápida
  replicas = 1 
}
