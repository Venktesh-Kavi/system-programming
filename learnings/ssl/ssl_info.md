## TLS/SSL


## Should I call tls/ssl?

* TLS is the upgraded version SSL
* TLS (Transport layer security), SSL (Secure Sockets Layer)

## Single TLS or mTLS

* Most systems are TLS today, which means the client usually want to verify the servers certificate before a handshake is made.
* mTLS requires verification by both server and client of their respective certificates.

## Why are certificates required?

* Certificates are required by clients to trust a server which they are interacting with, as they might be sending some sensitive data.
* 90% of the certificates in the internet today are provided few large known CA (Certificate Issuing Authorities like LetsEncrypt, Sectigo, ZeroSSL, DigiCert etc..,)
* Most operating systems and browsers come pre-installed with default set of Root Certificates.

## How the Pre-installed Root CA Chain Works
* Root CAs: Root CA certificates are self-signed and form the top of the trust hierarchy.
* Intermediate CAs: Root CAs issue certificates to Intermediate CAs, which in turn issue certificates to websites or applications. These intermediate certificates chain back to the root certificate, forming the certificate chain.
* Server or Client Certificates: When you visit a secure website (for example), the server presents a certificate issued by an Intermediate CA. The browser or application will verify that this certificate ultimately traces back to one of the trusted Root CAs in the store. 

## What does SSL termination mean?

* Most applications terminate the SSL at the ALB layer, which means they do the following:
* SSL termination is the process of decrypting SSL/TLS traffic at a designated point in the network, allowing for secure communications between clients (such as web browsers) and a server. This point can be a load balancer, proxy server, or dedicated SSL termination device. Here's how it works and why it's used:

### How SSL Termination Works
- Client Connection: When a client initiates a connection to a server using HTTPS, it establishes a secure SSL/TLS handshake to negotiate the encryption parameters and exchange keys.

- Traffic Routing: Instead of sending the encrypted traffic directly to the application server, the traffic is routed to an SSL termination point (like a load balancer or reverse proxy).

Decryption: The SSL termination point decrypts the incoming SSL/TLS traffic, allowing it to inspect, manage, or log the content of the data.

- Forwarding: After decryption, the SSL termination device can forward the unencrypted traffic to the application server over an internal network using a standard protocol (like HTTP).

- Response Handling: The application server processes the request and sends the response back to the SSL termination device, which then re-encrypts it before sending it back to the client.

- **There can also cases were SSL terminate can be set as passthrough in ALB or backend ssl enabled can be set as true which mean even the communication from ALB to the servers are in SSL** (TBA add more details)

## Does server need some special keys to decrypt the encrpyted data in SSL?
1. Private Key
What It Is: The private key is part of the public/private key pair generated when creating an SSL/TLS certificate.
Role in Decryption: During the SSL handshake, the server uses its private key to decrypt the pre-master secret sent by the client, which is used to derive session keys for encrypting and decrypting the actual data exchanged during the session.
Where It Is Stored: The private key should be stored securely on the server or SSL termination device and must be kept confidential to maintain the security of the SSL/TLS connection.
2. SSL/TLS Certificate
What It Is: The SSL/TLS certificate contains the public key and other identifying information about the server.
Role in Handshake: During the SSL handshake, the server presents its certificate to the client. The client verifies the certificate against a trusted Certificate Authority (CA) to ensure it is legitimate and to obtain the public key.
Usage: While the certificate itself does not decrypt messages, it is essential for establishing trust and enabling the client to encrypt the session keys that are sent to the server.
3. Session Keys
What They Are: Session keys are symmetric keys generated during the SSL handshake. They are used to encrypt and decrypt the data sent during the session.
Role in Encryption/Decryption: After the handshake, the client and server use these session keys to encrypt and decrypt the messages they exchange. The session keys are derived from the pre-master secret, which is itself secured by the private key.
Summary of the Decryption Process
Handshake: The client initiates an SSL handshake with the server, during which it receives the server's SSL certificate (containing the public key).
Pre-Master Secret: The client generates a pre-master secret, encrypts it with the server's public key, and sends it to the server.
Session Key Derivation: The server uses its private key to decrypt the pre-master secret, allowing both parties to derive the session keys.
Encrypted Communication: The client and server use the session keys to encrypt and decrypt the data exchanged during the session.
Security Considerations
Protection of Private Keys: The security of SSL termination relies heavily on keeping the private key secure. If an attacker gains access to the private key, they can decrypt any traffic encrypted with the corresponding public key.
Certificate Management: It's important to manage SSL certificates and keys properly, including regular updates and revocation of compromised certificates.

## How does VPN in the middle affect SSL termination?

* VPN does not modify anything in SSL process
* It just acts as a secure tunnel for sending 

## Useful commands

1. Open SSL certificate file (pem/crt): `openssl x509 -in CredavenuePrivateLimited-40058RootCA.crt -text -noout`
2. Get Full Certificate Chain: `openssl s_client -showcerts -servername yubioci-uat-uls-gateway-api.go-yubi.in -connect yubioci-uat-uls-gateway-api.go-yubi.in:443 </dev/null`

## Real World Examples

* Yubi for its OCI servers uses Sectigo Certificate
* We had issues were sometimes only the leaf certificate (domain certificate) was presented to the clients
* Some clients like java require the intermediary certificate to climb up to the Root CA in the chain
* Depending on the CA issuer, they provide the full chain or domain + issuer certificate separately.

* Yubi in AWS uses AWS's native Certificate

What is X.509 certificates?
* An X.509 certificate is a standard format for public key certificates, which are used to verify that a public key belongs to a specific entity, such as an individual, organization, or device. 


References:

[* AWS SSL Termination](https://aws.amazon.com/blogs/aws/elastic-load-balancer-support-for-ssl-termination/)
* https://stackoverflow.com/questions/71155280/does-an-aws-application-load-balancer-always-terminate-https-connections-or-is
