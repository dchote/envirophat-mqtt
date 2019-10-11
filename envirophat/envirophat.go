package envirophat

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/dchote/envirophat-mqtt/config"

	"github.com/dchote/go-envirophat"
	"github.com/dchote/go-envirophat/weather"

	"github.com/eclipse/paho.mqtt.golang"
)

var (
	i2c envirophat.Closer
)

// Init contains the generic read/publish loop
func Init() {
	i2c, err := weather.InitI2C()
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	opts := mqtt.NewClientOptions().AddBroker(config.ConnectionString()).SetClientID(config.ClientID)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var wg sync.WaitGroup

	for {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Read temperature in celsius degree
			t, err := weather.Temperature()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Temperature = %v*C\n", t)
			token := client.Publish(config.TopicPrefix+"/temperature", 0, false, fmt.Sprintf("%f", t))

			token.Wait()
			if token.Error() != nil {
				log.Fatal(token.Error())
			}

			// Read atmospheric pressure in pascal
			p, err := weather.Pressure()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Pressure = %v Pa\n", p)
			token = client.Publish(config.TopicPrefix+"/pressure", 0, false, fmt.Sprintf("%f", p))
			token.Wait()

			// Read atmospheric altitude in meters above sea level, if we assume
			// that pressure at see level is equal to 101325 Pa.
			a, err := weather.Altitude()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Altitude = %v m\n", a)
			token = client.Publish(config.TopicPrefix+"/altitude", 0, false, fmt.Sprintf("%f", a))
			token.Wait()
		}()

		time.Sleep(time.Duration(config.Interval) * time.Second)
	}

	wg.Wait()

	client.Disconnect(250)
	time.Sleep(1 * time.Second)
}
