package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	mail "github.com/go-mail/mail/v2"
)

func HandleNotification(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r) // Ensure CORS is enabled properly

	// Log request headers to debug the content type
	fmt.Println("Received Headers:", r.Header)

	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB max size
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		fmt.Println("Error parsing multipart form:", err)
		return
	}

	// Get form values
	documentID := r.FormValue("documentID")
	message := r.FormValue("message")

	fmt.Println("Received documentID:", documentID)
	fmt.Println("Received message:", message)

	// Ensure environment variables are set
	emailAdd := os.Getenv("EMAIL")
	emailPassword := os.Getenv("APP_PASSWORD")
	emailHost := os.Getenv("EMAIL_HOST")

	if emailAdd == "" || emailPassword == "" || emailHost == "" {
		http.Error(w, "Email configuration is missing", http.StatusInternalServerError)
		return
	}

	// Create a new mailer
	m := mail.NewMessage()
	m.SetHeader("From", emailAdd)
	m.SetHeader("To", "info.idoko@gmail.com")
	m.SetAddressHeader("Cc", emailAdd, "Puppy Spot")
	m.SetHeader("Subject", "Notification!!!")
	m.SetBody("text/html", fmt.Sprintf("Someone just triggered something:<br/><p>ID: %s</p><br/><p>Message: %s</p>", documentID, message))

	// Send email
	d := mail.NewDialer(emailHost, 465, emailAdd, emailPassword)
	d.Timeout = 120 * time.Second
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Attempt to send email
	if err := d.DialAndSend(m); err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		fmt.Println("Error sending email:", err)
		return
	}

	// Send a success response to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification sent successfully"))
}
