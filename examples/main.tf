terraform {
  required_providers {
    strcase = {
      source = "elioseverojunior/strcase"
      version = "1.0.0"
    }
  }
}

provider "strcase" {}

data "strcase" "this" {
  string = "Terraform Provider Strcase"
  delimiter = "."
}

output "original_string" {
  value = data.strcase.this.string
}

output "strcase" {
  value = data.strcase.this
}

output "strcase_to_kebab" {
  value = data.strcase.this.to_kebab
}

output "strcase_to_snake" {
  value = data.strcase.this.to_snake
}