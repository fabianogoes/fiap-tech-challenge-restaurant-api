resource "aws_db_instance" "tech-challenge-db" {
    identifier = "tech-challenge-db"
    allocated_storage = 5
    engine = "postgres"
    engine_version = "16.1"
    parameter_group_name = "default.postgres16"
    instance_class = "db.t3.micro"
    db_name = "tech_challenge_db"
    username = "tech_challenge_usr"
    password = "tech_challenge_pwd"

    vpc_security_group_ids = [aws_security_group.instance.id]
    
    publicly_accessible = true
    skip_final_snapshot = true
}

resource "aws_security_group" "instance" {
    name = "tech-challenge-sg"
    ingress {
        from_port   = 5432
        to_port     = 5432
        protocol    = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
}