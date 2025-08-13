# Full-DevOps-Stack Template

**Production-ready Go REST-service wrapped into IaC, CI/CD, Monitoring & Logging**

---

## 📌 What it is

- **Micro-service** written in Go (Gin) with PostgreSQL and Redis  
- **Docker → Kubernetes → GitOps** flow:  
  `develop` (docker-compose) → `dev` (Skaffold) → `prod` (Ansible + Helm)  
- **Observability**: Prometheus, Grafana, Alertmanager.
- **Logs**: EFK / Fluent/Filebeat.
- **IaC**: Terraform (cloud) + Ansible (OS + cluster)  
- **CI/CD**: GitHub Actions → Docker → Helm → Argo CD  
- **QA**: Molecule tests, pre-commit, Renovate, CVE scanners  
