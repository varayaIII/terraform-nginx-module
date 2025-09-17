variable "namespace" {
  description = "El namespace de Kubernetes donde se desplegará Nginx."
  type        = string
  default     = "default"
}

variable "app_name" {
  description = "El nombre para la aplicación y las etiquetas."
  type        = string
}

variable "app_image" {
  description = "La imagen de Docker a utilizar para el contenedor de Nginx."
  type        = string
  default     = "nginx:1.25.1"
}

variable "replicas" {
  description = "El número de réplicas para el despliegue de Nginx."
  type        = number
  default     = 2
}

variable "container_port" {
  description = "El puerto que expone el contenedor."
  type        = number
  default     = 80
}
