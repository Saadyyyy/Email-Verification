package emails

import (
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

func SendWelcomeEmail(userEmail, name, verificationToken string) error {
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	sender := smtpUsername
	recipient := userEmail
	subject := "Laode Saady Website"
	verificationLink := "http://localhost:8080/verify?token=" + verificationToken
	emailBody := `
    <html>
    <head>
        <link href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
        <style>
            body {
    font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
    background-color: #f5f5f5;
    color: #333;
		}
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Website Laode saadyy</h1>
            <div class="message">
                <p>Hello, <strong>` + name + `</strong>,</p>
                <p>Terimakasih sudah register di website ini. Silahkan tekan tombol verify email untuk melanjutkan pendaftaran</p>
                <p><strong>Support Team:</strong> <a href="mailto:laodesaady12345@gmail.com">laodesaady12345@gmail.com</a></p>
                <a href="` + verificationLink + `" class="btn btn-verify-email">Verify Email</a>
            </div>
            <div class="footer">
                <p>&copy; 2024 laode saady. All rights reserved.</p>
            </div>
        </div>
    </body>
    </html>
    `

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", emailBody)

	d := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
