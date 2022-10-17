Terraform Provider
==================

- Website: https://www.terraform.io

Maintainers
-----------

This provider plugin is maintained by [Elio Severo Junior](https://github.com/elioseverojunior).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.13.x
-	[Go](https://golang.org/doc/install) 1.18 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/elioseverojunior/terraform-provider-string`

```sh
$ mkdir -p $GOPATH/src/github.com/elioseverojunior; cd $GOPATH/src/github.com/elioseverojunior
$ git clone git@github.com:elioseverojunior/terraform-provider-strcase
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/elioseverojunior/terraform-provider-strcase
$ make build
```

Using the provider
----------------------
## Fill in for each provider

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.18+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-strcase
...
```


Usage sample:

```hcl
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
```
