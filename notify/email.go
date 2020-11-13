package notify

import (
	"log"
	"os"

	model "../model"
)

var (
	emailUsername = ""
	emailPass     = ""
	emailHost     = ""
	emailPort     = ""
)

func init() {
	emailUsername = os.Getenv("EMAIL_USERNAME")
	emailPass = os.Getenv("EMAIL_PASS")
	emailHost = os.Getenv("EMAIL_HOST")
	emailPort = os.Getenv("EMAIL_PORT")
}

func email(msg string, client *model.User) (string, error) {
	log.Println("email methond chosen")
	// Set up authentication information.
	// auth := smtp.PlainAuth("", emailUsername, emailPass, emailHost)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	// to := []string{client.email}
	// body := "To: " + client.email + "\r\n" +
	// 	"Subject: Purchase Info!\r\n" +
	// 	"\r\n" +
	// 	msg + "\r\n"

	// // Send
	// err := smtp.SendMail(fmt.Sprintf("%s:%s", emailHost+emailPort), auth, emailUsername, to, []byte(body))
	// if err != nil {
	// 	log.Fatal("Email method error: ", err)
	// 	return "", err
	// }
	return "Email Sent", nil
}
