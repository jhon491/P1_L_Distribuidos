package componenteconexioncola

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"

	dtos "estadisticas/capaFachadaServices/DTOs"
	"estadisticas/capaFachadaServices/fachada"
)

type RabbitConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	fachada *fachada.FachadaEstadisticas
}

func NewRabbitConsumer(f *fachada.FachadaEstadisticas) (*RabbitConsumer, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error abriendo canal: %v", err)
	}

	q, err := ch.QueueDeclare(
		"reproducciones_audio",
		true, false, false, false, nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error declarando cola: %v", err)
	}

	return &RabbitConsumer{conn: conn, channel: ch, queue: q, fachada: f}, nil
}

// Escuchar consume mensajes de la cola y delega el procesamiento a la fachada
func (c *RabbitConsumer) Escuchar() error {
	mensajes, err := c.channel.Consume(
		c.queue.Name,
		"", true, false, false, false, nil,
	)
	if err != nil {
		return fmt.Errorf("error registrando consumidor: %v", err)
	}

	fmt.Println("Esperando mensajes de reproducción...")

	for d := range mensajes {
		fmt.Println("Mensaje consumido desde la cola 'reproducciones_audio'") // eco requerido

		var reproduccion dtos.ReproduccionAudioDTO
		if err := json.Unmarshal(d.Body, &reproduccion); err != nil {
			fmt.Println("Error leyendo mensaje:", err)
			continue
		}

		c.fachada.RegistrarReproduccion(reproduccion)
	}
	return nil
}

func (c *RabbitConsumer) Cerrar() {
	c.channel.Close()
	c.conn.Close()
}
