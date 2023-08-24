openssl genpkey -algorithm RSA -out jwt_private_key.pem
openssl rsa -pubout -in jwt_private_key.pem -out jwt_public_key.pem