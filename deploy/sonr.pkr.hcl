packer {
  required_plugins {
    amazon = {
      version = ">= 1.2.8"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

source "amazon-ebs" "ubuntu" {
  ami_name      = "sonr-linux-aws"
  instance_type = "t2.micro"
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
  name = "sonr-testnet"
  sources = [
    "source.amazon-ebs.ubuntu"
  ]
  provisioner "shell" {
    environment_vars = [
      "FOO=hello world",
    ]
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y wget",
      "wget https://dist.ipfs.tech/kubo/v0.26.0/kubo_v0.26.0_linux-amd64.tar.gz",
      "tar -xvf kubo_v0.26.0_linux-amd64.tar.gz",
      "cd kubo && sudo bash ./install.sh"
    ]
  }
}
