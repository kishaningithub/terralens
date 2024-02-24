# Terralens

See your Terraform state world with clarity and precision - TerraLens, your visual command center for infrastructure insight!

# Why?

I wanted to analyze statefiles across 100s of terraform roots present across mega mono repos and multiple poly repos to 
drive data driven decisions. On this journey, I did stumble upon [terrabord](https://github.com/camptocamp/terraboard) 
which even though was a great project did not meet my needs as it required a more involved setup and direct access to 
terraform state storage (S3 buckets and dynamodb tables). 

I just wanted a simple CLI which can run in my local machine on a single terraform root and did not require direct
access to terraform state storage (S3 buckets and dynamodb tables). Once that is in place, I can just use the full 
power of bash scripting to scale it across 100s of terraform roots to extract insights.