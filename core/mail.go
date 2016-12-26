package core

import (
    "log"
    "net/mail"
    "net/smtp"
    "math/rand"
    "time"

    "github.com/pgmonzon/Yng_Servicios/cfg"

    "github.com/scorredoira/email"
    "gopkg.in/mgo.v2/bson"
)

var src = rand.NewSource(time.Now().UnixNano())
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func EnviarMailDeVerificacion(mail_usuario string, id_usuario bson.ObjectId){
    // Crea un codigo de verificacion nuevo, ENVIA este codigo a la persona, y lo modifica en la base de datos
    codigo_verificacion := CrearCodigoDeVerificacion(6)
    EnviarMail(mail_usuario, codigo_verificacion)
    updatearCodigoDeVerificacionAlUsuario(codigo_verificacion, id_usuario)
}

func EnviarMail(mail_usuario string, body string) {
    m := email.NewMessage("Codigo de verificacion", body)
    m.From = mail.Address{Name: "Yangee API", Address: cfg.EmailVerificacion}
    m.To = []string{mail_usuario}

    //Adjuntos:
    /*if err := m.Attach("email.go"); err != nil {
        log.Fatal(err)
    }*/

    auth := smtp.PlainAuth("", cfg.EmailVerificacion, cfg.EmailPwd, "smtp.gmail.com")
    if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
        log.Fatal(err)
    }
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CrearCodigoDeVerificacion(n int) string { //n es el numero que determina la longitud de la string que van a devolver
    b := make([]byte, n)
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

func updatearCodigoDeVerificacionAlUsuario(codigo string, id_usuario bson.ObjectId) {
	//var usuario models.Usuario

	creacion := time.Now()
	session := Session.Copy()
	defer session.Close()
	collection := session.DB(Dbname).C("usuarios")
	collection.Update(bson.M{"_id": id_usuario},
		bson.M{"$set": bson.M{"codigo": codigo, "creacion": creacion}})
}
