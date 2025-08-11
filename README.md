# 🛒 Cloud-Based E-commerce Application

A modern, cloud-native e-commerce application built with **GoLang** and **React**, following a microservices architecture and deployed on **Google Kubernetes Engine (GKE)**. This project demonstrates scalable cloud deployment, CI/CD readiness, and integration with Google Cloud services like **Cloud SQL** and **Cloud Run**.

---

## 🚀 Features

- ✅ Four microservices developed using **GoLang**
- ✅ Frontend application developed using **React + Vite**
- ✅ Deployment on **Google Kubernetes Engine (GKE)** for scalability and reliability
- ✅ Ingress configuration for unified access via a single external IP
- ✅ Integration with **Google Cloud SQL** for managing relational data
- ✅ Optional deployment of services using **Google Cloud Run** (serverless)
- ✅ Secure token-based communication between services (JWT)

---

## 📁 Microservices Overview

| Service          | Description                        | Deployment |
|------------------|------------------------------------|------------|
| `auth-service`   | Handles authentication & JWT       | GKE        |
| `product-service`| Manages product catalog             | GKE        |
| `order-service`  | Handles order processing            | GKE        |
| `payment-service`| Manages payment operations          | GKE        |

---

## 🖥️ Frontend

- Developed with **React** and **Vite**
- Interacts with backend services via REST APIs
- Deployed via **GKE** or optionally via **Cloud Run**
- Configurable environment for API endpoints

---

## ☁️ Google Cloud Integration

| Service           | Purpose                          |
|-------------------|----------------------------------|
| GKE               | Container orchestration          |
| Cloud SQL         | Relational database              |
| Cloud Run         | Serverless deployment (optional) |
| Artifact Registry | Container image hosting          |
| Ingress           | Unified routing for microservices|

---

## 🛠️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/cloud-ecommerce-app.git
cd cloud-ecommerce-app
