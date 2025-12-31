variable "aws_region" {
  default     = "ap-south-1"
  description = "AWS region"
}

variable "cluster_name" {
  default     = "Harshal_Joshi-eks"
  description = "EKS cluster name"
}

variable "kubernetes_version" {
  default     = "1.30"
  description = "Kubernetes version"
}

variable "vpc_cidr" {
  default     = "10.0.0.0/16"
  description = "VPC CIDR"
}
