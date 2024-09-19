terraform {

  backend "s3" {
    bucket         = "**REPLACE**"
    key            = "**REPLACE**"
    region         = "**REPLACE**"
    dynamodb_table = "**REPLACE**"
  }

  required_providers {
    acme = {
      source  = "vancluever/acme"
      version = "~> 2.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
  }
}

provider "acme" {
  server_url = "https://acme-v02.api.letsencrypt.org/directory"
}

provider "aws" {
  region = "us-east-1"
}

provider "github" {}
