# 🚀 Projeto de DevOps: Encurtador de URL com Microserviços e GitOps

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white) ![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white) ![GitLab CI](https://img.shields.io/badge/gitlab%20ci-%23181717.svg?style=for-the-badge&logo=gitlab&logoColor=B95A20) ![ArgoCD](https://img.shields.io/badge/Argo%20CD-FFFFFF?style=for-the-badge&logo=argo&logoColor=black) ![RabbitMQ](https://img.shields.io/badge/Rabbitmq-FF6600?style=for-the-badge&logo=rabbitmq&logoColor=white)

## 📄 Visão Geral do Projeto

Este repositório documenta a criação e implantação de um **Encurtador de URLs** utilizando uma arquitetura moderna de microserviços. O projeto demonstra um fluxo de trabalho DevOps completo, desde o provisionamento da infraestrutura como código até o deploy contínuo automatizado com GitOps, servindo como um case prático de habilidades prontas para produção.

A aplicação permite encurtar URLs e rastrear os cliques de forma assíncrona, garantindo alta performance e resiliência.

## 🛠️ Tecnologias e Pilares DevOps Demonstrados

| Pilar DevOps | Ferramentas e Conceitos Aplicados |
| :--- | :--- |
| **Infraestrutura como Código (IaC)** | **Terraform** para provisionar e gerenciar de forma automatizada o cluster **Kubernetes (DOKS)** na **DigitalOcean**. |
| **Containerização** | **Docker** para empacotar cada microserviço (`Frontend`, `API Gateway`, `Worker`) em imagens otimizadas, utilizando **builds multi-stage**. |
| **CI (Integração Contínua)** | **GitLab CI/CD** para automatizar os processos de `linting` (Hadolint, Go), testes, `build` e `push` das imagens para o **GitLab Container Registry**. |
| **CD (Deploy Contínuo) & GitOps** | **ArgoCD** para implementar o deploy contínuo. O ArgoCD monitora um repositório de manifestos e sincroniza automaticamente o estado do cluster com o declarado no Git, seguindo o paradigma **pull-based**. |
| **Arquitetura de Microserviços** | A aplicação é desacoplada em serviços independentes que se comunicam via APIs e mensageria, aumentando a resiliência e a escalabilidade. |
| **Mensageria e Cache** | **RabbitMQ** como message broker para processamento assíncrono de analytics, e **Redis** como banco de dados de alta velocidade para mapeamento das URLs. |
| **Orquestração** | **Kubernetes** para orquestrar todos os contêineres, gerenciando o deploy, a rede (`Services`) e a persistência de dados (`PersistentVolumeClaim`) para o **PostgreSQL**. |

---

## 🏛️ Arquitetura da Solução

O diagrama abaixo, gerado com a abordagem de "Diagrams as Code" (Python), ilustra o fluxo completo da arquitetura implementada.

![Diagrama da Arquitetura GitOps](arquitetura/arquitetura_gitops_-_url_shortener.png)

O fluxo de trabalho funciona da seguinte forma:
1.  **Desenvolvimento:** O desenvolvedor envia o código para o repositório da aplicação no **GitLab**.
2.  **CI Pipeline:** O `push` aciona a pipeline no **GitLab CI/CD**, que executa:
    * Testes e linting.
    * Build das imagens Docker e push para o **GitLab Container Registry**.
    * Um commit automático no **repositório de manifestos**, atualizando a tag da imagem.
3.  **Deploy com GitOps:** O **ArgoCD**, rodando no cluster **Kubernetes**, detecta a alteração no repositório de manifestos.
4.  **Sincronização:** O ArgoCD "puxa" os novos manifestos e aplica as alterações ao cluster, atualizando a aplicação em produção sem intervenção manual.
5.  **Acesso:** O usuário final acessa a aplicação através de um **Load Balancer** da DigitalOcean, que direciona o tráfego para o pod do **Frontend**.

---

## ✨ Showcase do Projeto

### 🚀 Aplicação em Produção

<img width="2488" height="1200" alt="Pós deploy no ArgoCD" src="https://github.com/user-attachments/assets/85d8c50c-6cc6-4eef-8fb9-55d298f5b2d6" />

### 🔄 Pipeline de CI/CD (GitLab)

<img width="2086" height="810" alt="Pipeline final" src="https://github.com/user-attachments/assets/a2709720-5ad6-43d6-8c7d-d4da5478dcac" />


###  ArgoCD em Ação (GitOps)

<img width="2537" height="1431" alt="ArgoCD full healthy" src="https://github.com/user-attachments/assets/175c0c33-61f2-42be-a039-532ae85f372c" />


---
