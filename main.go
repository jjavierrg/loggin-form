package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

var tpl = template.Must(template.New("login").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Login</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background: #f4f4f4;
            height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            margin: 0;
        }
        .box {
            background: white;
            padding: 30px;
            width: 320px;
            border-radius: 10px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
        }
        input {
            width: 100%;
            margin-bottom: 12px;
            padding: 10px;
            font-size: 14px;
            border-radius: 6px;
            border: 1px solid #ccc;
        }
        button {
            width: 100%;
            padding: 10px;
            background: #0078ff;
            color: white;
            border: none;
            border-radius: 6px;
            font-size: 15px;
            cursor: pointer;
        }
    </style>
</head>
<body>
    <div class="box">
        <form method="POST">
            <h2 style="text-align:center">Login</h2>
            <input type="text" name="username" placeholder="Usuario" required />
            <input type="password" name="password" placeholder="Contraseña" required />
            <button type="submit">Entrar</button>
        </form>
    </div>
</body>
</html>
`))

func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user := r.FormValue("username")
		pass := r.FormValue("password")
		ip := getIP(r)

		log.Printf("LOGIN --> user='%s' password='%s' ip='%s'\n", user, pass, ip)
		w.Write([]byte("<h3>Datos recibidos. Revisa los logs del contenedor.</h3>"))
		return
	}
	tpl.Execute(w, nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func main() {
	// Healthcheck CLI: loginapp health
	if len(os.Args) > 1 && os.Args[1] == "health" {
		resp, err := http.Get("http://localhost/health")
		if err != nil || resp.StatusCode != 200 {
			os.Exit(1)
		}
		os.Exit(0)
	}

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/health", healthHandler)

	// Servidores HTTP y HTTPS en goroutines
	go func() {
		log.Println("HTTP escuchando en :80")
		log.Println(http.ListenAndServe(":80", nil))
	}()

	log.Println("HTTPS escuchando en :443")
	log.Fatal(http.ListenAndServeTLS(":443", "/server.crt", "/server.key", nil))
}
