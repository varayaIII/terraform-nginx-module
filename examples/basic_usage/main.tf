provider "kubernetes" {
  config_path = "~/.kube/config"
}

module "nginx_app" {
  source    = "../../modules/nginx"
  app_name  = "nginx-test-app"
  replicas  = 1
}
