terraform {
  backend "pg" {
  }

  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "~> 5.0"
    }
  }
}

resource "heroku_app" "gopin" {
  name   = "gopin"
  region = "us"
  stack  = "container"
}

resource "heroku_addon" "jawsdb" {
  app_id = heroku_app.gopin.id
  plan   = "jawsdb:kitefin"
}

output "gopin_url" {
  value = "https://${heroku_app.gopin.name}.herokuapp.com"
}
