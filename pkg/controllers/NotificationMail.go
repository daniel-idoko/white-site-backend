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
}


func HandleUserNotificationEmail(w http.ResponseWriter, r *http.Request){
		utils.EnableCors(w, r)
	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// Get form values
		documentID := r.FormValue("documentID")
		email := r.FormValue("email")
		puppyName := r.FormValue("puppyName")
		breed := r.FormValue("breed")
		

		var emailAdd = os.Getenv("EMAIL")
		var emailPassword = os.Getenv("APP_PASSWORD")
		var emailHost = os.Getenv("EMAIL_HOST")

		// Create a new mailer
		m := mail.NewMessage()
		m.SetHeader("From", emailAdd)
		m.SetHeader("To", email)
		m.SetAddressHeader("Cc", emailAdd, "Puppy Spot Adoption")
		m.SetHeader("Subject", "Your Puppy Adoption Application is Being Reviewed!")

		m.SetBody("text/html",
		    "Thank you for applying to adopt " + puppyName + "! 🐾 We’re excited to process your application and will review it shortly.<br><br>" +
		    "<strong>Application Details:</strong><br>" +
		    "Puppy Name: " + puppyName + "<br>" +
		    "Breed: " + breed + "<br>" +
		    "Tracking ID: " + documentID + "<br>" +
		    "Application Status: <strong>Under Review</strong><br><br>" +
		    "You can track your application status anytime using your Tracking ID at:<br>" +
		    "<a href='https://puppyspotadoption.shop/shop/puppy-tracker/" + documentID + "'>🔗 Check Status</a><br><br>" +
		    "We will notify you as soon as your application is approved or if we need any further information. If you have any questions, feel free to reach out.<br><br>" +
		    "Thanks for choosing us to help you find your new furry friend! 🐶❤️<br><br>" +
		    "<br>" +
		    "Puppy Spot Adoption<br>" +
		   	"📧 info.puppyspotadoption@gmail.com")
	
		
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

}
