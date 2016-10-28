package core

import (
    "log"
    "net/mail"
    "net/smtp"

    "github.com/scorredoira/email"
)

func EnviarMail() {
    // compose the message
    m := email.NewMessage("Hi", "this is the body")
    m.From = mail.Address{Name: "From", Address: "yangeeapp@gmail.com"}
    m.To = []string{"adrian.diasdacostalima@gmail.com"}

    // add attachments
    /*if err := m.Attach("email.go"); err != nil {
        log.Fatal(err)
    }*/

    // send it
    auth := smtp.PlainAuth("", "yangeeapp@gmail.com", "la1962ser", "smtp.gmail.com")
    if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
        log.Fatal(err)
    }
}
