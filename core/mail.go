package core

import (
    "log"
    "net/mail"
    "net/smtp"
    "math/rand"
    "time"

    "github.com/scorredoira/email"
)

func EnviarMail() {
    // compose the message
    m := email.NewMessage("Codigo de verificacion", RandStringBytesMaskImprSrc(6))
    m.From = mail.Address{Name: "Yangee API", Address: "yangeeapp@gmail.com"}
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

var src = rand.NewSource(time.Now().UnixNano())
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)


func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringBytesMaskImprSrc(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}
