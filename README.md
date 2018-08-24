# terraform-state-hash
Simple Terraform state hash code generator

## Purpose
I was playing around with a Terraform provider modification and needed a quick way to generate resource state hashes for acceptance testing. I didn't see a readily available solution in the test code base so I just wrote a little program that uses HashiCorp's `hashcode` helper to simulate the hashing functionality used in state file generation. Basically, you just modify the elements int the JSON file to include attribute values that need to be included in the hash. Currently only strings and bools are supported. Order has been implemented too as the order that you buffer the elements does in fact affect the hashing algorithm.


## NEWB ALERT
This is my first Go program and could probably use a lot of help. It was also developed for a very specific use case. Use with caution.
