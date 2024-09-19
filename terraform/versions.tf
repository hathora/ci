terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    acme = {
      source = "vancluever/acme"
    }
    google = {
      source = "hashicorp/google"
    }
  }
}
