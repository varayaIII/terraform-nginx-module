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

📦 Requisitos

Terraform >= 1.5.0
Kubernetes cluster activo
kubectl configurado
(Para desarrollo) Go 1.22+ para ejecutar tests

🏗️ Estructura del Proyecto
.
├── .github/workflows/     # CI/CD pipelines
│   ├── ci.yml            # Lint, validación y docs
│   └── e2e.yml           # Tests end-to-end
├── modules/nginx/        # Módulo principal
│   ├── main.tf
│   ├── variables.tf
│   ├── outputs.tf
│   └── versions.tf
├── examples/             # Ejemplos de uso
│   └── basic_usage/
└── test/                 # Tests con Terratest
    └── nginx_test.go
🧪 Testing
Ejecutar tests localmente
bash# Configurar cluster local (KinD)
kind create cluster

# Ejecutar tests
cd test
go test -v -timeout 30m
CI/CD
El proyecto incluye dos workflows:

CI: Valida formato, sintaxis, lint y actualiza documentación
E2E: Ejecuta tests de integración en cada PR/push a main

📚 Documentación del Módulo
Para documentación detallada del módulo (inputs, outputs, requirements), consulta:
modules/nginx/README.md
🤝 Contribuir

Fork el proyecto
Crea una rama para tu feature (git checkout -b feature/amazing-feature)
Commit tus cambios (git commit -m 'Add amazing feature')
Push a la rama (git push origin feature/amazing-feature)
Abre un Pull Request

Convenciones

Código Terraform debe seguir terraform fmt
Variables deben estar documentadas
Tests deben pasar antes de merge
Los commits deben seguir Conventional Commits

📄 Licencia
Este proyecto es de código abierto y está disponible bajo la MIT License.
👤 Autor: varayalabs
