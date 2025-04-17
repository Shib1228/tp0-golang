package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()
	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json") //LO guardo en mi variable global para poder usarla libremente
	log.Println("Soy un Log")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}

	// loggeamos el valor de la config
	log.Println(globals.ClientConfig.Mensaje)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje) //Va la primer letra en mayuscula porque no es como esta escrita en la config, sino como esta escrita en la globals, porque ya esta almacenado ahi el mensaje

	// leer de la consola el mensaje

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
}
