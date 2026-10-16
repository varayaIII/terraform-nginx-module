# Terraform NGINX Module

[![CI](https://github.com/varayaIII/terraform-nginx-module/actions/workflows/ci.yml/badge.svg)](https://github.com/varayaIII/terraform-nginx-module/actions/workflows/ci.yml)
[![E2E](https://github.com/varayaIII/terraform-nginx-module/actions/workflows/e2e.yml/badge.svg)](https://github.com/varayaIII/terraform-nginx-module/actions/workflows/e2e.yml)

A reusable Terraform module for deploying NGINX to Kubernetes with configurable replicas, health checks, resource limits, and a `LoadBalancer` service.

The project also demonstrates infrastructure testing with **Terratest**, automated validation with **GitHub Actions**, and generated Terraform documentation.

## Features

- Reusable Terraform module for Kubernetes.
- Configurable application name, image, namespace, port, and replicas.
- Kubernetes Deployment and LoadBalancer Service.
- Liveness and readiness probes.
- Configurable container resource limits.
- Automated Terraform validation and formatting checks.
- End-to-end tests with Terratest.
- Documentation generated with terraform-docs.
- Suitable as a building block for GitOps-oriented workflows.

## Requirements

- Terraform `>= 1.5.0`.
- Kubernetes cluster.
- `kubectl` configured for the target cluster.
- Go `1.22+` for running Terratest locally.
- Kind or another local Kubernetes cluster for end-to-end testing.

## Usage

```hcl
provider "kubernetes" {
  config_path = "~/.kube/config"
}

module "nginx_app" {
  source    = "github.com/varayaIII/terraform-nginx-module//modules/nginx"
  app_name  = "my-nginx"
  replicas  = 3
  namespace = "production"
}

output "service_ip" {
  value = module.nginx_app.service_ip
}
```

Initialize and apply the module from an example configuration:

```bash
terraform init
terraform plan
terraform apply
```

## Module interface

| Input | Description | Default |
|---|---|---:|
| `app_name` | Application and resource name. | Required |
| `app_image` | Container image for NGINX. | `nginx:1.25.1` |
| `container_port` | Container port. | `80` |
| `namespace` | Kubernetes namespace. | `default` |
| `replicas` | Number of Deployment replicas. | `2` |

| Output | Description |
|---|---|
| `deployment_name` | Name of the Kubernetes Deployment. |
| `service_name` | Name of the Kubernetes Service. |
| `service_ip` | External address exposed by the LoadBalancer. |

For the generated and complete interface documentation, see [`modules/nginx/README.md`](modules/nginx/README.md).

## Repository structure

```text
.
├── .github/workflows/
│   ├── ci.yml                # Format, validate, lint, and documentation
│   └── e2e.yml               # End-to-end tests
├── modules/nginx/
│   ├── main.tf
│   ├── variables.tf
│   ├── outputs.tf
│   ├── versions.tf
│   └── README.md
├── examples/
│   └── basic_usage/
└── test/
    └── nginx_test.go
```

## Testing locally

Create a local Kubernetes cluster with Kind:

```bash
kind create cluster
```

Run Terraform checks:

```bash
terraform fmt -check -recursive
terraform init -backend=false
terraform validate
```

Run Terratest:

```bash
cd test
go test -v -timeout 30m
```

## CI/CD

The repository includes two workflows:

- **CI:** formatting, validation, linting, and documentation updates.
- **E2E:** integration tests on pull requests and pushes to the default branch.

## Contributing

1. Create a feature branch.
2. Format and validate Terraform code.
3. Update module documentation when inputs or outputs change.
4. Run the Terratest suite when a Kubernetes cluster is available.
5. Open a pull request with a clear description.

## License

This project is available under the MIT License.

## Author

[VarayaIII](https://github.com/varayaIII)
