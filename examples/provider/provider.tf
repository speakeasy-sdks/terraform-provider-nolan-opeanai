terraform {
  required_providers {
    openai = {
      source  = "nolans-big-test/openai"
      version = "1.0.0"
    }
  }
}

provider "openai" {
  # Configuration options
}