packer {
  required_plugins {
    amazon = {
      version = ">= 1.2.8"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

source "amazon-ebs" "ubuntu" {
  ami_name      = "sonr-ubuntu-aws"
  instance_type = "t2.xlarge"
  region        = "us-east-1"
  subnet_id     = "subnet-0510c361d200223ae"
  source_ami_filter {
    filters = {
      name                = "ubuntu/images/*ubuntu-jammy-22.04-amd64-server-*"
      root-device-type    = "ebs"
      virtualization-type = "hvm"
    }
    most_recent = true
    owners      = ["099720109477"]
  }
  ssh_username = "ubuntu"
}


build {
  name = "sonr-base-vm"
  sources = [
    "source.amazon-ebs.ubuntu"
  ]

  # Install Docker
  provisioner "shell" {
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y wget ca-certificates curl",
      "sudo curl -fsSL https://get.docker.com -o get-docker.sh",
      "sudo sh get-docker.sh",
      "rm get-docker.sh",
    ]
  }

  # Install Earthly
  provisioner "shell" {
    inline = [
      "sudo wget https://github.com/earthly/earthly/releases/latest/download/earthly-linux-amd64 -O /usr/local/bin/earthly",
      "sudo chmod +x /usr/local/bin/earthly",
      "sudo /usr/local/bin/earthly bootstrap",
    ]
  }
}
