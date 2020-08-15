#!/usr/bin/expect

# generate private key and enter pass phrase
set PASS_PHRASE "1234567890"
puts "\n********** Generate RSA PrivateKey with '$PASS_PHRASE' passphrase **********\n"
spawn openssl genrsa -des3 -out private_key.pem 2048
expect "Enter pass phrase"
send "$PASS_PHRASE\r"
expect "Verifying - Enter pass phrase"
send "$PASS_PHRASE\r"
interact
puts ""

