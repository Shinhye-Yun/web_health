package app

import(
	"crypto/tls"
	"log"
	"time"
	"math"
)


func ValidateCert(url string, port string){

	// 1: Check website SSL cert by establishing a TLS connection.
	conn, err := tls.Dial("tcp", url+":"+port, nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}

	// 2. Check if SSL cert has correct hostname.
	err = conn.VerifyHostname(url) // match common domain name with url
	if err != nil {
		panic("Hostname doesn't match with certificate: " + err.Error())
	}

	// 3. Check cert expiration date.
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	log.Println("Issuer:", conn.ConnectionState().PeerCertificates[0].Issuer, "Expiry:", expiry.Format(time.RFC850))
	now := time.Now().UTC()
	diff := expiry.Sub(now).Hours()
	diff = math.Floor(diff / 24.0)
	if diff < 300 {
		log.Println(diff, "days")
	} else {
		log.Println(diff, "days left") //TODO: send notification
	}
}