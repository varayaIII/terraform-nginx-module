provider "kubernetes" {
  # Terraform asumirá la configuración de tu `kubeconfig` local
  # que apunta a tu clúster (Minikube/Kind).
}

resource "kubernetes_deployment" "nginx" {
  metadata {
    name      = var.app_name
    namespace = var.namespace
    labels = {
      app = var.app_name
    }
  }

  spec {
    replicas = var.replicas

    selector {
      match_labels = {
        app = var.app_name
      }
    }

    template {
      metadata {
        labels = {
          app = var.app_name
        }
      }

      spec {
        container {
          image = var.app_image
          name  = var.app_name
          port {
            container_port = var.container_port
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "nginx" {
  metadata {
    name      = "${var.app_name}-service"
    namespace = var.namespace
  }
  spec {
    selector = {
      app = kubernetes_deployment.nginx.metadata[0].labels.app
    }
    port {
      port        = 80
      target_port = var.container_port
    }
    type = "NodePort" # Usamos NodePort para que sea fácil de probar localmente
  }
}
