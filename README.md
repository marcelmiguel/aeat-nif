## Library to check Spanish NIF/CIF against government web service

[![Actions Status](https://github.com/marcelmiguel/aeat-nif/workflows/test/badge.svg)](https://github.com/marcelmiguel/aeat-nif/actions) 

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

This lib works but requires still some important job to do (for security reasons)

# Install

``` sh
go get github.com/marcelmiguel/aeat-nif
```

Get pfx file

Convert pfc encrypted to pem unencrypted

``` sh
openssl pkcs12 -in cert.pfx -nodes -out cert.crt
```

# Usage for tests

Put a cert.crt file in the project directory
Change aeat-nif_test.go and set password CERTPWD constant (DO NOT UPDATE TO GIT THE PASSWORD !!).
Launch tests.

# TODO

- Use Environment variables for file and password of certificate
- Use secrets in BASE64 to store cert and password for use in Kubernetes
- Verify caducity of certifcate
- Enable read encrypted pem
- Enable read encrypted pfx
