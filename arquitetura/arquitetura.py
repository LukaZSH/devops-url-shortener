# arquitetura.py

from diagrams import Diagram, Cluster, Edge
from diagrams.onprem.vcs import Gitlab
from diagrams.onprem.ci import GitlabCI
from diagrams.onprem.container import Docker
from diagrams.onprem.database import Postgresql
from diagrams.onprem.inmemory import Redis
from diagrams.onprem.queue import Rabbitmq
from diagrams.k8s.compute import Deployment, Pod
from diagrams.k8s.network import Service
from diagrams.onprem.gitops import Argocd
from diagrams.digitalocean.compute import K8SCluster as DOKS
from diagrams.digitalocean.network import LoadBalancer
from diagrams.onprem.client import User
from diagrams.onprem.iac import Terraform

graph_attr = {
    "fontsize": "20",
    "bgcolor": "white",
    "splines": "curved",
    "ranksep": "2.0",
    "nodesep": "1.0",
}

with Diagram("Arquitetura GitOps - URL Shortener", show=False, graph_attr=graph_attr, direction="LR"):

    usuario_final = User("Usuário Final")

    with Cluster("Ambiente de Desenvolvimento & IaC"):
        terraform = Terraform("Terraform\n(Infra as Code)")

        with Cluster("GitLab.com"):
            repo_app = Gitlab("Repo: url-shortener-app")
            repo_manifests = Gitlab("Repo: url-shortener-manifests")
            gitlab_ci = GitlabCI("GitLab CI/CD Pipeline")
            container_registry = Docker("GitLab Container Registry")

    with Cluster("Cloud Provider: DigitalOcean"):

        with Cluster("DOKS - Cluster Kubernetes"):
            argo_cd = Argocd("ArgoCD\n(Agente GitOps)")

            with Cluster("Namespace: 'default'"):
                # Ponto de entrada
                ingress_service = Service("Frontend Service (LB)")

                # Aplicação
                frontend = Deployment("Frontend")
                api_gateway = Deployment("API Gateway")
                analytics_worker = Deployment("Analytics Worker")

                # Backend / Infra
                redis = Redis("Redis")
                rabbitmq = Rabbitmq("RabbitMQ")
                postgres = Postgresql("PostgreSQL")

    # --- FLUXOS ---

    # Fluxo de CI
    repo_app >> Edge(label="git push") >> gitlab_ci
    gitlab_ci >> Edge(label="build & push") >> container_registry
    gitlab_ci >> Edge(label="atualiza .yaml") >> repo_manifests

    # Fluxo de GitOps
    repo_manifests >> Edge(label="PULL: observa o repo") >> argo_cd
    argo_cd >> Edge(label="SYNC: aplica no cluster") >> [frontend, api_gateway, analytics_worker]

    # Fluxo da Aplicação
    usuario_final >> ingress_service >> frontend >> api_gateway
    api_gateway >> [redis, rabbitmq]
    rabbitmq >> analytics_worker >> postgres

    # Fluxo de Provisionamento (IaC)
    terraform >> Edge(label="terraform apply", style="dashed") >> argo_cd