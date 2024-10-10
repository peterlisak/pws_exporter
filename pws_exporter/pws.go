package pws_exporter

type PWSParams struct {
	// mandatory
	Id string `schema:"ID,required"`
	/* Password string `schema:"PASSWORD,required"`
	Dateutc  string `schema:"dateutc,required"` */
	// optionals
	/* Action         string
	Realtime       bool */
	Rtfreq         float32 `schema:"rtfreg,omitempty"`
	Baromin        float32 `schema:"baromin,omitempty"`
	Tempf          float32 `schema:"tempf,omitempty"`
	Dewptf         float32 `schema:"dewptf,omitempty"`
	Humidity       int     `schema:"humidity,omitempty"`
	Windspeedmph   float32 `schema:"windspeedmph,omitempty"`
	Windgustmph    float32 `schema:"windgustmph,omitempty"`
	Winddir        int     `schema:"winddir,omitempty"`
	Rainin         float32 `schema:"rainin,omitempty"`
	Dailyrainin    float32 `schema:"dailyrainin,omitempty"`
	Solarradiation float32 `schema:"solarradiation,omitempty"`
	UV             float32 `schema:"uv,omitempty"`
	Indoortempf    float32 `schema:"indoortempf,omitempty"`
	Indoorhumidity int     `schema:"indoorhumidity,omitempty"`
	Soiltempf      float32 `schema:"soiltempf,omitempty"`
	Soilmoisture   int     `schema:"soilmoisture,omitempty"`
	Soiltemp2f     float32 `schema:"soiltemp2f,omitempty"`
	Soilmoisture2  int     `schema:"soilmoisture2,omitempty"`
	Soiltemp3f     float32 `schema:"soiltemp3f,omitempty"`
	Soilmoisture3  int     `schema:"soilmoisture3,omitempty"`
	Soiltemp4f     float32 `schema:"soiltemp4f,omitempty"`
	Soilmoisture4  int     `schema:"soilmoisture4,omitempty"`
	Soiltemp5f     float32 `schema:"soiltemp5f,omitempty"`
	Soilmoisture5  int     `schema:"soilmoisture5,omitempty"`
	Soiltemp6f     float32 `schema:"soiltemp6f,omitempty"`
	Soilmoisture6  int     `schema:"soilmoisture6,omitempty"`
	Soiltemp7f     float32 `schema:"soiltemp7f,omitempty"`
	Soilmoisture7  int     `schema:"soilmoisture7,omitempty"`
	Soiltemp8f     float32 `schema:"soiltemp8f,omitempty"`
	Soilmoisture8  int     `schema:"soilmoisture8,omitempty"`
}
