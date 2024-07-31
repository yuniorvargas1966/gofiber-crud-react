package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

var status = 301

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasena := ""
	Nombre := "taller"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasena+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", c.Handler(handler)))
	// Routes

	http.HandleFunc("/datos", Inicio)

	http.HandleFunc("/crear", Crear)

	http.HandleFunc("/insertar", Insertar)

	http.HandleFunc("/borrar", Borrar)

	http.HandleFunc("/editar", Editar)

	http.HandleFunc("/actualizar", Actualizar)

	//var port = ":8080"
	//fmt.Println("Servidor corriendo en el puerto", port, "entre a http://localhost:8080/")
	//log.Fatal(http.ListenAndServe(port, nil))

}

type Servicio struct {
	Id       int
	Nombre   string
	Equipo   string
	Trabajo  string
	Telefono string
	Correo   string
	Fecha    string
}

// Inicio
func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := conexionBD()

	registros, err := conexionEstablecida.Query("SELECT * FROM servicios")

	if err != nil {
		panic(err.Error())
	}
	servicio := Servicio{}
	arregloServicio := []Servicio{}

	for registros.Next() {
		var id int
		var nombre, equipo, trabajo, telefono, correo, fecha string
		err = registros.Scan(&id, &nombre, &equipo, &trabajo, &telefono, &correo, &fecha)
		if err != nil {
			panic(err.Error())
		}
		servicio.Id = id
		servicio.Nombre = nombre
		servicio.Equipo = equipo
		servicio.Trabajo = trabajo
		servicio.Telefono = telefono
		servicio.Correo = correo
		servicio.Fecha = fecha

		arregloServicio = append(arregloServicio, servicio)

	}
	// fmt.Println(arregloServicio)
	plantillas.ExecuteTemplate(w, "inicio", arregloServicio)
}

// Borrar
func Borrar(w http.ResponseWriter, r *http.Request) {
	idServicio := r.URL.Query().Get("id")
	fmt.Println(idServicio)

	conexionEstablecida := conexionBD()

	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM servicios WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	borrarRegistro.Exec(idServicio)
	http.Redirect(w, r, "/datos", status)
}

// Editar
func Editar(w http.ResponseWriter, r *http.Request) {
	idServicio := r.URL.Query().Get("id")
	fmt.Println(idServicio)

	conexionEstablecida := conexionBD()

	registro, err := conexionEstablecida.Query("SELECT * FROM servicios WHERE id=?", idServicio)

	servicio := Servicio{}
	for registro.Next() {
		var id int
		var nombre, equipo, trabajo, telefono, correo, fecha string
		err = registro.Scan(&id, &nombre, &equipo, &trabajo, &telefono, &correo, &fecha)
		if err != nil {
			panic(err.Error())
		}
		servicio.Id = id
		servicio.Nombre = nombre
		servicio.Equipo = equipo
		servicio.Trabajo = trabajo
		servicio.Telefono = telefono
		servicio.Correo = correo
		servicio.Fecha = fecha
	}

	fmt.Println(servicio)
	plantillas.ExecuteTemplate(w, "editar", servicio)
}

// Crear
func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

// Insertar
func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		nombre := r.FormValue("nombre")
		equipo := r.FormValue("equipo")
		trabajo := r.FormValue("trabajo")
		telefono := r.FormValue("telefono")
		correo := r.FormValue("correo")
		fecha := r.FormValue("fecha")

		conexionEstablecida := conexionBD()

		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO servicios(nombre,equipo,trabajo,telefono,correo,fecha) VALUES(?, ?, ?, ?, ?, ?) ")

		if err != nil {
			panic(err.Error())
		}

		insertarRegistros.Exec(nombre, equipo, trabajo, telefono, correo, fecha)

		http.Redirect(w, r, "/datos", status)

	}

}

// Actualizar
func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		equipo := r.FormValue("equipo")
		trabajo := r.FormValue("trabajo")
		telefono := r.FormValue("telefono")
		correo := r.FormValue("correo")
		fecha := r.FormValue("fecha")

		conexionEstablecida := conexionBD()

		modificarRegistros, err := conexionEstablecida.Prepare("UPDATE servicios SET nombre=?,equipo=?,trabajo=?,telefono=?,correo=?,fecha=? WHERE id=? ")

		if err != nil {
			panic(err.Error())
		}

		modificarRegistros.Exec(nombre, equipo, trabajo, telefono, correo, fecha, id)

		http.Redirect(w, r, "/datos", status)

	}

}
