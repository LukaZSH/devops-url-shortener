<p align="right">
  <a href="README-en.md"><img src="https://raw.githubusercontent.com/twitter/twemoji/master/assets/svg/1f1fa-1f1f8.svg" width="24" alt="English" /> English</a> •
  <a href="README.md"><img src="https://raw.githubusercontent.com/twitter/twemoji/master/assets/svg/1f1e7-1f1f7.svg" width="24" alt="Português" /> Português</a>
</p>

# 🚀 DevOps Project: URL Shortener with Microservices and GitOps

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white) ![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white) ![GitLab CI](https://img.shields.io/badge/gitlab%20ci-%23181717.svg?style=for-the-badge&logo=gitlab&logoColor=B95A20) ![ArgoCD](https://img.shields.io/badge/Argo%20CD-FFFFFF?style=for-the-badge&logo=argo&logoColor=black) ![RabbitMQ](https://img.shields.io/badge/Rabbitmq-FF6600?style=for-the-badge&logo=rabbitmq&logoColor=white)

## 📄 Project Overview

This repository documents the creation and deployment of a **URL Shortener** using a modern microservices architecture. The project demonstrates a complete DevOps workflow, from provisioning infrastructure as code to automated continuous deployment with GitOps, serving as a practical, production-ready skills showcase.

The application allows users to shorten URLs and track clicks asynchronously, ensuring high performance and resilience.

## 🛠️ Technologies and DevOps Pillars Demonstrated

| DevOps Pillar | Tools and Concepts Applied |
| :--- | :--- |
| **Infrastructure as Code (IaC)** | **Terraform** to automatically provision and manage the **Kubernetes (DOKS)** cluster on **DigitalOcean**. |
| **Containerization** | **Docker** to package each microservice (`Frontend`, `API Gateway`, `Worker`) into optimized images using **multi-stage builds**. **Docker Hub** as the image registry. |
| **CI (Continuous Integration)** | **GitLab CI/CD** to automate the processes of `linting` (Hadolint), testing, `building`, vulnerability scanning with **Trivy**, and `pushing` images to the **GitLab Container Registry**. |
| **CD (Continuous Deployment) & GitOps** | **ArgoCD** to implement continuous deployment. ArgoCD monitors a manifest repository and automatically synchronizes the cluster's state with the state declared in Git, following the **pull-based** paradigm. |
| **Microservices Architecture** | The application is decoupled into independent services that communicate via APIs and messaging, increasing resilience and scalability. |
| **Messaging and Caching** | **RabbitMQ** as a message broker for asynchronous analytics processing, and **Redis** as a high-speed database for URL mapping. |
| **Orchestration** | **Kubernetes** to orchestrate all containers, managing deployments, networking (`Services`), and data persistence (`PersistentVolumeClaim`) for **PostgreSQL**. |

---

## 🏛️ Solution Architecture

The diagram below, generated with a "Diagrams as Code" approach (Python), illustrates the complete architectural flow implemented in this project.

![GitOps Architecture Diagram](architecture/Arquitetura_GitOps_-_URL_Shortener.png)

The workflow operates as follows:
1.  **Development:** The developer pushes code to the application repository on **GitLab**.
2.  **CI Pipeline:** The `push` triggers the pipeline in **GitLab CI/CD**, which executes tests, linting, builds, vulnerability scans, and pushes the Docker images to the **GitLab Container Registry**, in addition to updating the image tag in the **manifests repository**.
3.  **GitOps Deployment:** **ArgoCD**, running in the **Kubernetes** cluster, detects the change in the manifests repository and "pulls" the new manifests, updating the application in production without manual intervention.
4.  **Access:** The end-user accesses the application through a **Load Balancer** provided by DigitalOcean.

---

## ✨ Project Showcase

### 🚀 Application in Production
*The URL Shortener application, after a successful deployment via ArgoCD, is publicly accessible.*

<img width="2488" height="1200" alt="Post deploy in ArgoCD" src="https://github.com/user-attachments/assets/85d8c50c-6cc6-4eef-8fb9-55d298f5b2d6" />

### 🔄 CI/CD Pipeline (GitLab)
*The GitLab Actions workflow showing the successful execution of all stages (Quality, Test, Build & Scan, Deploy).*

<img width="2086" height="810" alt="Final pipeline" src="https://github.com/user-attachments/assets/a2709720-5ad6-43d6-8c7d-d4da5478dcac" />

### 🛡️ Vulnerability Scanning in Action (Trivy)
*Evidence of the DevSecOps step, where the pipeline failed upon detecting `HIGH` severity vulnerabilities in the Nginx base image, blocking the deployment. The issue was resolved by updating the image tag in the Dockerfile.*

`[INSERIR SEU PRINT DO TRIVY AQUI]`

### 📨 Asynchronous Messaging (RabbitMQ)
*The RabbitMQ management dashboard running in the cluster, showing the "clicks" queue ready to receive and process events in a decoupled and resilient manner.*

`[INSERIR SEU PRINT DO RABBITMQ AQUI]`

### 🤖 Continuous Deployment with GitOps (ArgoCD)
*A view of ArgoCD with all application resources synced and healthy, demonstrating that the cluster state faithfully mirrors the manifests repository.*

<img width="2537" height="1431" alt="ArgoCD full healthy" src="https://github.com/user-attachments/assets/175c0c33-61f2-42be-a039-532ae85f372c" />

---

## 🎓 Conclusion

This project served as a practical, in-depth immersion into the modern DevOps ecosystem, transforming a microservices application into a fully automated system, from infrastructure as code with **Terraform** to continuous deployment with **GitLab CI** and **ArgoCD**.

The main outcomes were the creation of a resilient pipeline that ensures code quality and security, and the implementation of a **GitOps** workflow that makes deployments safer and more traceable.

### 🚀 Next Steps
* **Implement Load Testing:** Use tools like K6 to analyze performance under stress.
* **Add Full Observability:** Integrate the **Logging** (with the EFK/Loki stack) and **Tracing** (with Jaeger/OpenTelemetry) pillars for a 360º view of the system's behavior.
* **Cost Optimization:** Explore using **KEDA** (Kubernetes Event-driven Autoscaling) to scale the analytics workers from zero to `N` based on the number of messages in the RabbitMQ queue.