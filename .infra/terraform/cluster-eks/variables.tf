variable "account_id" {}

variable "region" {
    default = "us-east-1"
}

variable "vpcId" {
    default = "vpc-03e2dfdc264a07455"
}

variable "subnetA" {
    default = "subnet-02dd8eac5450c11c6"
}

variable "subnetB" {
    default = "subnet-0bc7e328b4f6d07a0"
}

variable "sgId" {
    default = "sg-02d4aa982423a7f7c"
}

variable "nodeName" {
    default = "ng-tech-challenge"
}

variable "accessConfig" {
    default = "API_AND_CONFIG_MAP"
}

variable "policyArn" {
    default = "arn:aws:eks::aws:cluster-access-policy/AmazonEKSAdminPolicy"
}