resource "aws_instance" "testnet_1" {
  ami           = "ami-0bda289ea2d9fc54c"
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
      "earthly github.com/sonrhq/sonr+build",
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
