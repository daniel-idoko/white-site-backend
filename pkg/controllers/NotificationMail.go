package controller

import (
	"fmt"
	"net/http"
	"os"
	"puppyspot-backend/pkg/utils"
	"time"

	"github.com/go-mail/mail"
)




func HandleNotification(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(w, r)
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// Get form values
		documentID := r.FormValue("documentID")
		message := r.FormValue("message")


		var emailAdd = os.Getenv("EMAIL")
		var emailPassword = os.Getenv("APP_PASSWORD")
		var emailHost = os.Getenv("EMAIL_HOST")

		// Create a new mailer
		m := mail.NewMessage()
		m.SetHeader("From", emailAdd)
		m.SetHeader("To", "info.idoko@gmail.com")
		m.SetAddressHeader("Cc", emailAdd, "Puppy Spot")
		m.SetHeader("Subject", "Notification!!!")

		m.SetBody("text/html", "Someone just triggered something: <br/> <p>ID: "+documentID+"</p> <br/> <p>Message: "+message+"</p>")
	
		
		// Send email
		d := mail.NewDialer(emailHost, 465, emailAdd, emailPassword)
		d.Timeout = 120 * time.Second
		d.StartTLSPolicy = mail.MandatoryStartTLS		

	}
}