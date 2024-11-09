package app

import(
    "crypto/tls"
    "log"
    "time"
)

func checkExpiration(url string,){

	// 1. Establish a TLS connection to the server
	conn, err := tls.Dial("tcp", url, &tls.Config{
		InsecureSkipVerify: true, // ignore verify certificate (to avoid getting "certificate is not trusted" error with self signed certificate)
	})
	if err != nil {
		log.Printf("[ERROR] Failed to connect to %s: %v\n", url, err)
		// log.Fatal("Failed to connect to %s: %v", url, err)
		// log.Fatal() calls os.Exit(1) after logging the message, which immediately terminates the program without running deferred functions or allowing cleanup.
	}
	defer conn.Close() 

	// 2. Get the peer certificate from the server
    certs := conn.ConnectionState().PeerCertificates

	// 3. Check the expiration date of the certificate
	if len(certs) > 0 {
        cert := certs[0]
        expirationTime := cert.NotAfter
        currentTime := time.Now()
        
        if currentTime.After(expirationTime) {
            log.Printf("[INFO] %s - Certificate has expired on %v\n", url, expirationTime)
        } else {
            daysUntilExpiration := expirationTime.Sub(currentTime).Hours() / 24
            log.Printf("[INFO] %s - Certificate expires on %v (expires in %d days)\n", url, expirationTime, int(daysUntilExpiration))
        }
    } else {
        log.Printf("[INFO] No certificates found\n")
    }
}

func CheckCert(domain string, servers []string){
	for i := 0; i < len(servers); i++ {
		checkExpiration(servers[i]+domain+":443")
	}
}