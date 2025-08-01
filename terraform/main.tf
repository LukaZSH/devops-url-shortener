terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {}

resource "digitalocean_kubernetes_cluster" "url_shortener_cluster" {
  name    = "url-shortener-cluster"
  region  = "nyc1"                 
  version = "1.33.1-do.2"           

  node_pool {
    name       = "url-shortener-pool"
    size       = "s-2vcpu-2gb" 
    node_count = 2
  }
}

# Como dar o nome: TIPO_DO_RECURSO.NOME_QUE_VOCE_DEU.ATRIBUTO
output "cluster_name" {
  value = digitalocean_kubernetes_cluster.url_shortener_cluster.name
  description = "Nome do cluster Kubernetes criado na DigitalOcean"
}