terraform {
  required_version = ">= 1.5"
}

resource "local_file" "dummy" {
  filename = "${path.module}/.terraform-ok"
  content  = "Terraform stub"
}