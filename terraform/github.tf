resource "github_actions_secret" "certificate_pfx" {
  repository       = "hathora/ci"
  secret_name      = "CERTIFICATE_PFX"
  encrypted_value  = acme_certificate.certificate.certificate_p12
}
