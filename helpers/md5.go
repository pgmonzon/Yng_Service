package helpers

import (
	"crypto/md5"
  "encoding/binary"
	"bytes"
)

func HashearMD5(pass string) (uint64) {
  var a uint64
  md5pass := md5.New()																					// Hay un field que no es Pass, sino que es PassMD.
  md5pass.Write([]byte(pass))       														// Tomamos usuario.Pass y lo hasheamos
  buf := bytes.NewBuffer(md5pass.Sum(nil))											// y por ultimo el checksum lo convertimos a uint64
  binary.Read(buf, binary.LittleEndian, &a)              // <- y lo pasamos a PassMD en el ultimo parametro
  //TODO: Err handler aca
  return a
}
