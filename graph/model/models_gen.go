// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CurrentConditions struct {
	Datetime       *string   `json:"datetime,omitempty"`
	DatetimeEpoch  *int32    `json:"datetimeEpoch,omitempty"`
	Temp           *float64  `json:"temp,omitempty"`
	Feelslike      *float64  `json:"feelslike,omitempty"`
	Humidity       *float64  `json:"humidity,omitempty"`
	Dew            *float64  `json:"dew,omitempty"`
	Precip         *float64  `json:"precip,omitempty"`
	Precipprob     *float64  `json:"precipprob,omitempty"`
	Snow           *float64  `json:"snow,omitempty"`
	Snowdepth      *float64  `json:"snowdepth,omitempty"`
	Preciptype     []*string `json:"preciptype,omitempty"`
	Windgust       *float64  `json:"windgust,omitempty"`
	Windspeed      *float64  `json:"windspeed,omitempty"`
	Winddir        *float64  `json:"winddir,omitempty"`
	Pressure       *float64  `json:"pressure,omitempty"`
	Visibility     *float64  `json:"visibility,omitempty"`
	Cloudcover     *float64  `json:"cloudcover,omitempty"`
	Solarradiation *float64  `json:"solarradiation,omitempty"`
	Solarenergy    *float64  `json:"solarenergy,omitempty"`
	Uvindex        *float64  `json:"uvindex,omitempty"`
	Conditions     *string   `json:"conditions,omitempty"`
	Icon           *string   `json:"icon,omitempty"`
	Stations       []*string `json:"stations,omitempty"`
	Source         *string   `json:"source,omitempty"`
	Sunrise        *string   `json:"sunrise,omitempty"`
	SunriseEpoch   *float64  `json:"sunriseEpoch,omitempty"`
	Sunset         *string   `json:"sunset,omitempty"`
	SunsetEpoch    *float64  `json:"sunsetEpoch,omitempty"`
	Moonphase      *float64  `json:"moonphase,omitempty"`
}

type Day struct {
	Datetime         *string   `json:"datetime,omitempty"`
	DatetimeEpoch    *float64  `json:"datetimeEpoch,omitempty"`
	Tempmax          *float64  `json:"tempmax,omitempty"`
	Tempmin          *float64  `json:"tempmin,omitempty"`
	Temp             *float64  `json:"temp,omitempty"`
	Feelslikemax     *float64  `json:"feelslikemax,omitempty"`
	Feelslikemin     *float64  `json:"feelslikemin,omitempty"`
	Feelslike        *float64  `json:"feelslike,omitempty"`
	Dew              *float64  `json:"dew,omitempty"`
	Humidity         *float64  `json:"humidity,omitempty"`
	Precip           *float64  `json:"precip,omitempty"`
	Precipprob       *float64  `json:"precipprob,omitempty"`
	Precipcover      *float64  `json:"precipcover,omitempty"`
	Preciptype       []*string `json:"preciptype,omitempty"`
	Snow             *float64  `json:"snow,omitempty"`
	Snowdepth        *float64  `json:"snowdepth,omitempty"`
	Windgust         *float64  `json:"windgust,omitempty"`
	Windspeed        *float64  `json:"windspeed,omitempty"`
	Winddir          *float64  `json:"winddir,omitempty"`
	Pressure         *float64  `json:"pressure,omitempty"`
	Cloudcover       *float64  `json:"cloudcover,omitempty"`
	Visibility       *float64  `json:"visibility,omitempty"`
	Solarradiation   *float64  `json:"solarradiation,omitempty"`
	Solarenergy      *float64  `json:"solarenergy,omitempty"`
	Uvindex          *float64  `json:"uvindex,omitempty"`
	Severerisk       *float64  `json:"severerisk,omitempty"`
	Sunrise          *string   `json:"sunrise,omitempty"`
	SunriseEpoch     *float64  `json:"sunriseEpoch,omitempty"`
	Sunset           *string   `json:"sunset,omitempty"`
	SunsetEpoch      *float64  `json:"sunsetEpoch,omitempty"`
	Moonphase        *float64  `json:"moonphase,omitempty"`
	Conditions       *string   `json:"conditions,omitempty"`
	Description      *string   `json:"description,omitempty"`
	Icon             *string   `json:"icon,omitempty"`
	Stations         []*string `json:"stations,omitempty"`
	Source           *string   `json:"source,omitempty"`
	Hours            []*Hour   `json:"hours,omitempty"`
	Precipsum        *float64  `json:"precipsum,omitempty"`
	Precipsumnormal  *float64  `json:"precipsumnormal,omitempty"`
	Snowsum          *float64  `json:"snowsum,omitempty"`
	DatetimeInstance *string   `json:"datetimeInstance,omitempty"`
}

type Hour struct {
	Datetime         *string   `json:"datetime,omitempty"`
	DatetimeEpoch    *int32    `json:"datetimeEpoch,omitempty"`
	Temp             *float64  `json:"temp,omitempty"`
	Feelslike        *float64  `json:"feelslike,omitempty"`
	Humidity         *float64  `json:"humidity,omitempty"`
	Dew              *float64  `json:"dew,omitempty"`
	Precip           *float64  `json:"precip,omitempty"`
	Precipprob       *float64  `json:"precipprob,omitempty"`
	Snow             *float64  `json:"snow,omitempty"`
	Snowdepth        *float64  `json:"snowdepth,omitempty"`
	Preciptype       []*string `json:"preciptype,omitempty"`
	Windgust         *float64  `json:"windgust,omitempty"`
	Windspeed        *float64  `json:"windspeed,omitempty"`
	Winddir          *float64  `json:"winddir,omitempty"`
	Pressure         *float64  `json:"pressure,omitempty"`
	Visibility       *float64  `json:"visibility,omitempty"`
	Cloudcover       *float64  `json:"cloudcover,omitempty"`
	Solarradiation   *float64  `json:"solarradiation,omitempty"`
	Solarenergy      *float64  `json:"solarenergy,omitempty"`
	Uvindex          *float64  `json:"uvindex,omitempty"`
	Severerisk       *float64  `json:"severerisk,omitempty"`
	Conditions       *string   `json:"conditions,omitempty"`
	Icon             *string   `json:"icon,omitempty"`
	Stations         []*string `json:"stations,omitempty"`
	Source           *string   `json:"source,omitempty"`
	DatetimeInstance *string   `json:"datetimeInstance,omitempty"`
}

type Query struct {
}

type WeatherResponse struct {
	QueryCost         *int32             `json:"queryCost,omitempty"`
	Latitude          *float64           `json:"latitude,omitempty"`
	Longitude         *float64           `json:"longitude,omitempty"`
	ResolvedAddress   *string            `json:"resolvedAddress,omitempty"`
	Address           *string            `json:"address,omitempty"`
	Timezone          *string            `json:"timezone,omitempty"`
	Tzoffset          *float64           `json:"tzoffset,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Days              []*Day             `json:"days,omitempty"`
	CurrentConditions *CurrentConditions `json:"currentConditions,omitempty"`
	Alerts            []*string          `json:"alerts,omitempty"`
	IsProcessed       *bool              `json:"isProcessed,omitempty"`
}
