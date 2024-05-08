resource "aws_eks_cluster" "tech_challenge" {
  name     = "tech_challenge_eks_cluster"
  role_arn = "arn:aws:iam::${var.account_id}:role/LabRole"
  version = 1.29

  vpc_config {
    subnet_ids = ["${var.subnetA}", "${var.subnetB}"]
    security_group_ids = ["${var.sgId}"]
    endpoint_private_access = true
    endpoint_public_access  = true

  }

  access_config {
    authentication_mode = var.accessConfig
  }  
  
}
