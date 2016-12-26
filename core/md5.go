package core

import (
	"crypto/md5"
	"encoding/binary"
	"bytes"
       )

func HashearMD5(pass string) (int32) {
	var a int32
	md5pass := md5.New()
	md5pass.Write([]byte(pass))
	buf := bytes.NewBuffer(md5pass.Sum(nil))
	binary.Read(buf, binary.LittleEndian, &a)
	return a
}
