package rbac

import (
    "net/http"
    "log"

    "github.com/pgmonzon/Yng_Servicios/models"
    "github.com/pgmonzon/Yng_Servicios/core"
)

func ConseguirRP(r *http.Request) (models.RP){
  //Conseguimos los RP del usuario
  id := core.ExtraerClaim(r, "id")
  user, _ := core.ExtraerInfoUsuario(id.(string))
  RP, _ := core.ExtraerPermisosDelRol(user.Rol)
  ParsearPermisosAJSON(RP)
  return RP
}

func ParsearPermisosAJSON(permisos models.RP) (){
  log.Println(permisos)
}
