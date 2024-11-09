## Status Checker
Check server's status for every 5 seconds and expose to prometheus metrics.

### Monitoring

Prometheus metrics for all status code. If status code is true, metric value is 1 unless 0. 

For example, if status is 2xx the value is 1:
    
    check_status{status="2xx", server="manage.staging.mist-federal.com"} 1
    check_status{status="3xx", server="manage.staging.mist-federal.com"} 0
    check_status{status="4xx", server="manage.staging.mist-federal.com"} 0
    check_status{status="5xx", server="manage.staging.mist-federal.com"} 0

---

## Cert Checker

Issue with **self-signed** certificate is that to make a connection to server that uses self-signed certificate requires the "root cert" for cert verification.
If root cert was not provided and if you're trying to connect to the it would fail.

To check self-signed certificate of a server without using root cert, 
1. Establish a TSL connection to the server.
2. Get the peer certificate from the server. (**peer certificate)
3. Check the expiration date of the certificate.


More info about self-signed certificate and peer certificate: 

```txt
A peer certificate in the context of SSL/TLS connections refers to the certificate presented by the server to the client during the handshake process. For self-signed certificates, this has some specific implications:

    Self-signed nature:
    A self-signed certificate is one where the issuer and the subject are the same entity. This means the certificate is signed by its own private key rather than by a trusted Certificate Authority (CA).

    Verification challenges:
    When a client encounters a self-signed peer certificate, it typically cannot verify the certificate's authenticity through the normal chain of trust, as there's no higher authority to vouch for it.

    Trust issues:
    By default, most SSL/TLS implementations do not trust self-signed certificates. This often results in verification errors like "X509_V_ERR_DEPTH_ZERO_SELF_SIGNED_CERT" when attempting to establish a secure connection3.

    Manual trust establishment:
    To use self-signed certificates securely, clients often need to manually add the certificate to their trust store or implement custom verification logic2.

    Usage in development:
    Self-signed certificates are commonly used in development environments, internal networks, or scenarios where obtaining a CA-signed certificate is impractical.

    Security considerations:
    While self-signed certificates can provide encryption, they don't offer the same level of authentication as CA-signed certificates, since anyone can create a self-signed certificate claiming to be any entity.

    Expiration and renewal:
    Like any other certificate, self-signed certificates have an expiration date. When they expire, they need to be regenerated, and clients may need to update their trust stores accordingly.

When dealing with self-signed peer certificates in SSL/TLS connections, it's crucial to implement proper verification mechanisms to ensure the security of the communication, especially in production environments.