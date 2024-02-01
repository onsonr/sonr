output "testnet_1_id" {
  description = "ID of the EC2 Testnet instance"
  value       = aws_instance.testnet_1.id
}

output "testnet_1_public_ip" {
  description = "Public IP address of the Testnet EC2 instance"
  value       = aws_instance.testnet_1.public_ip
}
