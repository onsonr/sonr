resource "aws_instance" "internal_dxi" {
  ami           = "ami-0c7217cdde317cfec"
  instance_type = "t2.xlarge"
  subnet_id     = "subnet-0510c361d200223ae"

  tags = {
    Name = "Internal DXI"
  }
}

resource "aws_instance" "testnet_1" {
  ami           = "ami-0c7217cdde317cfec"
  instance_type = "t2.xlarge"
  subnet_id     = "subnet-0510c361d200223ae"

  tags = {
    Name = "Sonr Testnet"
  }
}

resource "cloudflare_record" "internal_dxi_dns_record" {
  zone_id = "4b0faeeadab13c19609a949a0db2439b"
  name    = "@"
  type = "A"
  value = aws_instance.internal_dxi.public_ip
  proxied = true
}
