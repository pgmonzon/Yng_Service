package core

import (
  //"log"
	"crypto/md5"
  "encoding/binary"
	"bytes"
)

func HashearMD5(pass string) (int32) {
  var a int32
  md5pass := md5.New()																					// Hay un field que no es Pass, sino que es PassMD.
  md5pass.Write([]byte(pass))       														// Tomamos usuario.Pass y lo hasheamos
  buf := bytes.NewBuffer(md5pass.Sum(nil))											// y por ultimo el checksum lo convertimos a uint64
  binary.Read(buf, binary.LittleEndian, &a)              // <- y lo pasamos a PassMD en el ultimo parametro
  //log.Printf("Se acaba de crear una contraseña de string: %s a MD5 %d", pass, a)
  return a
}
