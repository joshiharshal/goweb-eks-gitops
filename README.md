# goweb-

# Go Web App on EKS with GitOps

This repository contains the source code and configuration for deploying a simple Go web application to an Amazon EKS cluster. The infrastructure is provisioned using Terraform, and the application is deployed using standard Kubernetes manifests and a Helm chart. This project serves as a demonstration of Infrastructure as Code (IaC) and containerized application deployment on Kubernetes.


# Go Web Application

This is a simple website written in Golang. It uses the `net/http` package to serve HTTP requests.

## Running the server

To run the server, execute the following command:

```bash
go run main.go


http://localhost:8080/home
http://localhost:8080/services
http://localhost:8080/contact
http://localhost:8080/about 
```

The server will start on port 8080. You can access it by navigating to `http://localhost:8080/home` in your web browser.

## Project Structure

```
.
├── Dockerfile                  
├── main.go                     
├── go.mod              
├── helm/                       
│   └── goweb-eks-gitops-chart/
├── k8s/                        
├── nginx-ingress-con
├── static/
└── terraform/                  
```

## Core Components

*   **Go Web Application**: A simple HTTP server built in Go that serves static HTML pages from the `/static` directory.
*   **Docker**: A multi-stage `Dockerfile` is used to build a lightweight, distroless container image for the application.
*   **Terraform**: The `terraform/` directory contains all the necessary code to provision a VPC and an EKS cluster on AWS. It uses the official AWS modules for VPC and EKS.
*   **Kubernetes Manifests**: The `k8s/` directory provides basic YAML files for a Deployment, a ClusterIP Service, and an Ingress to deploy the application.
*   **Helm Chart**: For more complex and configurable deployments, a Helm chart is available in the `helm/` directory.
*   **NGINX Ingress Controller**: The project includes manifests and instructions to set up the NGINX Ingress Controller, which manages external access to the services in the EKS cluster.

## Getting Started

### Prerequisites

*   [AWS CLI](https://aws.amazon.com/cli/) configured with appropriate credentials.
*   [Terraform](https://www.terraform.io/downloads.html) (>= 1.3.0)
*   [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
*   [Docker](https://docs.docker.com/get-docker/)
*   [Helm](https://helm.sh/docs/intro/install/)

### 1. Provision Infrastructure with Terraform

First, create the AWS EKS cluster and the underlying VPC network.

1.  Navigate to the `terraform` directory:
    ```sh
    cd terraform
    ```

2.  Initialize Terraform to download the required providers and modules:
    ```sh
    terraform init
    ```

3.  Create an execution plan to preview the resources that will be created:
    ```sh
    terraform plan
    ```

4.  Apply the plan to provision the resources on AWS. This may take 15-20 minutes.
    ```sh
    terraform apply --auto-approve
    ```

5.  Once the EKS cluster is ready, configure `kubectl` to communicate with it:
    ```sh
    aws eks --region $(terraform output -raw region) update-kubeconfig --name $(terraform output -raw cluster_name)
    ```

### 2. Install NGINX Ingress Controller

The Ingress Controller is required to expose your application to the internet.

```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml
```
Wait a few minutes for the Network Load Balancer to be created and provisioned.

### 3. Build and Push the Docker Image

1.  Navigate to the root directory of the repository.

2.  Build the Docker image. Replace `your-dockerhub-username` with your Docker Hub username or another container registry path.
    ```sh
    docker build -t your-dockerhub-username/go-eks-gitops:latest .
    ```

3.  Push the image to your container registry:
    ```sh
    docker push your-dockerhub-username/go-eks-gitops:latest
    ```

### 4. Deploy the Application

You can deploy the application using either raw Kubernetes manifests or the Helm chart.

#### Option A: Using Kubernetes Manifests

1.  Open `k8s/deployment.yaml` and update the `image` field to match the image you just pushed.

    ```yaml
    # k8s/deployment.yaml
    ...
          containers:
          - name: go-eks-gitops
            image: your-dockerhub-username/go-eks-gitops:latest # <-- UPDATE THIS LINE
    ...
    ```

2.  Apply the manifests to your cluster:
    ```sh
    kubectl apply -f k8s/
    ```

#### Option B: Using the Helm Chart

1.  Open `helm/goweb-eks-gitops-chart/values.yaml` and update the `image.repository` and `image.tag` fields.

    ```yaml
    # helm/goweb-eks-gitops-chart/values.yaml
    image:
      repository: your-dockerhub-username/go-eks-gitops
      pullPolicy: IfNotPresent
      tag: "latest"
    ```

2.  Install the Helm chart from the chart's directory:
    ```sh
    helm install goweb-app helm/goweb-eks-gitops-chart/
    ```

### 5. Access the Application

The Ingress is configured to respond to the host `gitops.joshiharshal.cloud`. You will need to either configure DNS for this domain to point to the NGINX Ingress Load Balancer's public address or modify your local `/etc/hosts` file for testing.

The application serves the following pages:
*   `/home`
*   `/about`
*   `/services`
*   `/contact`

## Cleanup

To avoid ongoing charges for the AWS resources, you can destroy the infrastructure created by Terraform.

1.  Navigate to the `terraform` directory:
    ```sh
    cd terraform
    ```

2.  Destroy all resources:
    ```sh
    terraform destroy --auto-approve