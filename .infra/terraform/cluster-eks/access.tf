resource "aws_eks_access_entry" "eks_access" {
  cluster_name = aws_eks_cluster.tech_challenge.name
  principal_arn = "arn:aws:iam::${var.account_id}:role/voclabs"
  kubernetes_groups = [ "fiap", "tech-challenge" ]
  type = "STANDARD"
}