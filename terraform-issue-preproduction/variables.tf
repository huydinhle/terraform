variable "team_name" {}

############################
## AWS - General Settings ##
############################

variable "aws_credentials_file" {
  default = "/home/hle/.aws/credentials"
}

variable "aws_profile" {
  default = {
    "us-west-2" = "default"
  }
}

variable "aws_region" {
  description = "AWS region to launch servers."
  default     = "us-west-2"
}

variable "aws_availability_zone" {
  default = {
    "us-west-2" = "us-west-2a"
  }
}

variable "aws_route53_zone_id" {
  default = "Z26WMLA8RQFCN8"
}

##################
## AWS - Master ##
##################

variable "master_instance_type" {
  description = "Jenkins Master instance type"
  default     = "m4.large"
}

variable "master_aws_subnets" {
  default = {
    "us-west-2" = "subnet-d9dd7dbe"
  }
}

######################
## AWS - Slave Pool ##
######################

variable "slave_instance_type" {
  description = "Jenkins Slave instance type"
  default     = "m4.xlarge"
}

variable "slave_spot_price" {
  default = "0.1898"
}

variable "slave_aws_subnets" {
  default = {
    "us-west-2" = "subnet-d9dd7dbe"
  }
}

variable "slave_aws_sns_topic" {
  default = {
    "us-west-2" = "arn:aws:sns:us-west-2:314501373800:raptor-raptor-jenkins-slave-autoscaling-group-notifications"
  }
}
