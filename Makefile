.PHONY: help fmt validate lint test clean

help: ## Muestra este mensaje de ayuda
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

fmt: ## Formatea el código Terraform
	terraform -chdir=modules/nginx fmt -recursive
	terraform -chdir=examples/basic_usage fmt

validate: ## Valida la configuración de Terraform
	terraform -chdir=modules/nginx init -backend=false
	terraform -chdir=modules/nginx validate

lint: ## Ejecuta tflint
	tflint --init
	tflint

docs: ## Genera documentación del módulo
	terraform-docs markdown table --output-file README.md --output-mode inject modules/nginx

test: ## Ejecuta los tests de Terratest
	cd test && go test -v -timeout 30m

clean: ## Limpia archivos temporales
	find . -type d -name ".terraform" -exec rm -rf {} +
	find . -type f -name "terraform.tfstate*" -delete
	find . -type f -name ".terraform.lock.hcl" -delete

init-example: ## Inicializa el ejemplo básico
	terraform -chdir=examples/basic_usage init

apply-example: ## Aplica el ejemplo básico
	terraform -chdir=examples/basic_usage apply

destroy-example: ## Destruye el ejemplo básico
	terraform -chdir=examples/basic_usage destroy

all: fmt validate lint docs ## Ejecuta fmt, validate, lint y docs
