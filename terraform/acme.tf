data "aws_region" "current" {}

resource "acme_registration" "this" {
  email_address = "info@geode.io"
}

resource "acme_certificate" "certificate" {
  account_key_pem           = acme_registration.this.account_key_pem
  common_name               = "hathora.dev"

  dns_challenge {
    provider = "route53"
    config = {
      AWS_REGION          = "**REPLACE**"
      AWS_HOSTED_ZONE_ID  = "**REPLACE**"
    }
  }
}
