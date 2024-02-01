resource "aws_instance" "testnet_1" {
  ami           = "ami-0c7217cdde317cfec"
  instance_type = "t2.xlarge"
  subnet_id     = "subnet-0510c361d200223ae"

  tags = {
    Name = "Sonr Testnet"
  }

  root_block_device {
    volume_size = 30
  }
  provisioner "remote-exec" {
    inline = [
      "apt-get update",
      "apt-get install -y wget",
      "wget https://dist.ipfs.tech/kubo/v0.26.0/kubo_v0.26.0_linux-amd64.tar.gz",
      "tar -xvf kubo_v0.26.0_linux-amd64.tar.gz",
      "cd kubo && bash ./install.sh"
    ]
  }
}

resource "cloudflare_record" "testnet_dns_record" {
  zone_id = "acaa372320d4e317fd7c49815fbb6f7d"
  name    = "@"
  type    = "A"
  value   = aws_instance.testnet_1.public_ip
  proxied = true
}
