provider "aws" {
  region                  = "${var.aws_region}"
  shared_credentials_file = "${var.aws_credentials_file}"
  profile                 = "${lookup(var.aws_profile, var.aws_region)}"
}

data "aws_region" "current" {
  current = true
}

######################
## AWS - Slave Pool ##
######################

data "aws_ami" "jenkins-slave-coreos" {
  most_recent = true

  filter {
    name   = "name"
    values = ["CoreOS-stable-1235.6.0-hvm"]
  }

  owners = ["self"]
}

resource "aws_instance" "web" {
  count                       = "3"
  ami                         = "${data.aws_ami.jenkins-slave-coreos.id}"
  instance_type               = "m4.xlarge"
  key_name                    = "test-account"
  associate_public_ip_address = true

  tags {
    Name = "TestCount"
  }
}
