package main

import (
	"crypto/tls"
	"emailenvios/src/conexion"
	"fmt"
	"strconv"

	"gopkg.in/gomail.v2"
)

const nombres = "Playa Tuquillo"
const descripcion = "El balneario de Tuquillo, denominado como la piscina del Océano Pacífico, es considerado como una de las mejores playas del Perú. Está rodeado por casas y restaurantes de material recuperable a orillas de una ensenada."
const portada = "75150a9a-1b27-71a1-b62b-be3e37117d44.jpg"
const idzona = "387"

func main() {
	enviarEmails()
	enviarEmailsUserRegistrados()
	guardarEnvio()
}

func guardarEnvio() {
	query, err2 := conexion.Session.Prepare("INSERT INTO envios_emails (idzona, zona) VALUES ($1,$2)")
	catch(err2)

	query.Exec(idzona, nombres)
	fmt.Println("Guardado!!")
}

func enviarEmailsUserRegistrados() {
	var email string = ""
	var contador int = 0

	query := `select email from usuarios where email != ''`
	filas, err := conexion.Session2.Query(query)
	catch(err)

	for filas.Next() {
		contador++
		errsql := filas.Scan(&email)
		catch(errsql)
		sendEmailTienda(email, contador)
	}
}

func enviarEmails() {
	var d Emails
	var contador int = 0

	query := `select email from emails`
	filas, err := conexion.Session.Query(query)
	catch(err)

	for filas.Next() {
		contador++
		errsql := filas.Scan(&d.Email)
		catch(errsql)
		sendEmailTienda(d.Email, contador)
	}
}

func sendEmailTienda(email string, contador int) {
	m := gomail.NewMessage()
	m.SetHeader("From", "labradsoft@labradsoft.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Lugares Turísticos #"+idzona)
	htmlContent := `<!doctype html>
	<html lang="es">
	  <head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.1/dist/css/bootstrap.min.css" integrity="sha384-zCbKRCUGaJDkqS1kPbPd7TveP5iyJE0EjAuZQTgFLD2ylzuqKfdKlfG/eSrtxUkn" crossorigin="anonymous">
	  </head>
	  <body>
		<div class="container">
			<h2>` + nombres + `</h2>
			<p style="margin-bottom: 40px;">` + descripcion + `</p>
			<div class="text-center">
				<a  href="https://lugaresturisticos.app/#/zona-show/` + idzona + `" style="background-color: rgb(20, 56, 173); padding-top: 14px; padding-bottom: 14px; padding-left: 30px; padding-right: 30px; color: white; border-radius: 23px; -webkit-border-radius: 23px; margin-right: 30px;">
					Ver lugar turístico
				</a>
				<a  href="https://lugaresturisticos.app/apps" style="background-color: rgb(110, 7, 141); padding-top: 14px; padding-bottom: 14px; padding-left: 30px; padding-right: 30px; color: white; border-radius: 23px; -webkit-border-radius: 23px;">
					Descargar app
				</a>
			</div>
			
			<br>

<p style="color: rgb(63, 63, 121)"><b>Visítanos en las redes sociales:</b></p>
<a href="https://twitter.com/LugaresTuristc1" target="_blank"><img
src="https://lugaresturisticos.app:3000/static/iconoredesociales/twitter.png"
style="width:30px; margin-right: 7px;"></a>
<a href="https://www.facebook.com/profile.php?id=100067880418873" target="_blank"><img
src="https://lugaresturisticos.app:3000/static/iconoredesociales/facebook.png"
style="width:30px; margin-right: 7px;"></a>
<a href="https://www.instagram.com/lgturismo213/" target="_blank"><img
src="https://lugaresturisticos.app:3000/static/iconoredesociales/instagram.png"
style="width:30px; margin-right: 7px;"></a>
<a href="https://www.tiktok.com/@lugares_turisticos" target="_blank"><img
src="https://lugaresturisticos.app:3000/static/iconoredesociales/tiktok.png"
style="width:30px; margin-right: 7px;"></a>
<a href="https://www.pinterest.ca/lturistcos" target="_blank"><img
src="https://lugaresturisticos.app:3000/static/iconoredesociales/pinterest.png"
style="width:30px; margin-right: 7px;"></a>
<br>

			<img src="https://lugaresturisticos.app:3000/static/` + portada + `" class="w-25" style="margin-top: 20px">
		</div>
		
		<script src="https://cdn.jsdelivr.net/npm/jquery@3.5.1/dist/jquery.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.6.1/dist/js/bootstrap.bundle.min.js" integrity="sha384-fQybjgWLrvvRgtW6bFlB7jaZrFsaBXjsOMm/tB9LTS58ONXgqbR9W8oWht/amnpF" crossorigin="anonymous"></script>
	  </body>
	</html>`
	m.SetBody("text/html", htmlContent)

	d := gomail.NewPlainDialer("mail.labradsoft.com", 465, "_mainaccount@labradsoft.com", "gETUMeTD^z4Z")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Enviado " + strconv.Itoa(contador))
	}
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

type Emails struct {
	Idemail int    `json:"idemail"`
	Email   string `json:"email"`
}
