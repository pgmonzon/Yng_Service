package rbac

import (
	//"encoding/json"
	"net/http"
	"log"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"
)

func ParsearPermisosAJSON(r *http.Request) (models.RP){
	//Conseguimos los RP del usuario
	id := core.ExtraerClaim(r, "id")
	user, _ := core.ExtraerInfoUsuario(id.(string))
	RP, _ := core.ExtraerPermisosDelRol(user.Rol)
	log.Println(RP)
	core.BuscarLosPermisos(RP)
	return RP
}
