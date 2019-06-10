// Package addstogo provides conversion of the xml file received from the server www.aviationweather.gov/dataserver into the go data structure
package addstogo

import (
	"encoding/xml"
	"io"
	"time"
)

type METARresponse struct {
	RequestIndex int `xml:"request_index"`
	DataSource   struct {
		Name string `xml:"name,attr"`
	} `xml:"data_source"`
	Request struct {
		Type string `xml:"type,attr"`
	} `xml:"request"`
	Errors      []string `xml:"errors>error"`
	Warnings    []string `xml:"warnings>warning"`
	TimeTakenMs int      `xml:"time_taken_ms"`
	Data        struct {
		METAR []struct {
			RawText             string    `xml:"raw_text"`
			StationID           string    `xml:"station_id"`
			ObservationTime     time.Time `xml:"observation_time"`
			Latitude            float32   `xml:"latitude"`
			Longitude           float32   `xml:"longitude"`
			TempC               float32   `xml:"temp_c"`
			DewpointC           float32   `xml:"dewpoint_c"`
			WindDirDegrees      int       `xml:"wind_dir_degrees"`
			WindSpeedKt         int       `xml:"wind_speed_kt"`
			WindGustKt          int       `xml:"wind_gust_kt"`
			VisibilityStatuteMi float32   `xml:"visibility_statute_mi"`
			AltimInHg           float32   `xml:"altim_in_hg"`
			SeaLevelPressureMb  float32   `xml:"sea_level_pressure_mb"`
			QualityControlFlags struct {
				Corrected               bool `xml:"corrected"`
				Auto                    bool `xml:"auto"`
				AutoStation             bool `xml:"auto_station"`
				MaintenanceIndicatorOn  bool `xml:"maintenance_indicator_on"`
				NoSignal                bool `xml:"no_signal"`
				LightningSensorOff      bool `xml:"lightning_sensor_off"`
				FreezingRainSensorOff   bool `xml:"freezing_rain_sensor_off"`
				PresentWeatherSensorOff bool `xml:"present_weather_sensor_off"`
			} `xml:"quality_control_flags"`
			WxString     string `xml:"wx_string"`
			SkyCondition []struct {
				SkyCover       string `xml:"sky_cover,attr"`
				CloudBaseFtAgl int    `xml:"cloud_base_ft_agl,attr"`
			} `xml:"sky_condition"`
			FlightCategory            string  `xml:"flight_category"`
			ThreeHrPressureTendencyMb float32 `xml:"three_hr_pressure_tendency_mb"`
			MaxTC                     float32 `xml:"maxT_c"`
			MinTC                     float32 `xml:"minT_c"`
			MaxT24HrC                 float32 `xml:"maxT24hr_c"`
			MinT24HrC                 float32 `xml:"minT24hr_c"`
			PrecipIn                  float32 `xml:"precip_in"`
			Pcp3HrIn                  float32 `xml:"pcp3hr_in"`
			Pcp6HrIn                  float32 `xml:"pcp6hr_in"`
			Pcp24HrIn                 float32 `xml:"pcp24hr_in"`
			SnowIn                    float32 `xml:"snow_in"`
			VertVisFt                 int     `xml:"vert_vis_ft"`
			MetarType                 string  `xml:"metar_type"`
			ElevationM                float32 `xml:"elevation_m"`
		} `xml:"METAR"`
		NumResults int `xml:"num_results,attr"`
	} `xml:"data"`
}

type TAFresponse struct {
	RequestIndex int `xml:"request_index"`
	DataSource   struct {
		Name string `xml:"name,attr"`
	} `xml:"data_source"`
	Request struct {
		Type string `xml:"type,attr"`
	} `xml:"request"`
	Errors      []string `xml:"errors>error"`
	Warnings    []string `xml:"warnings>warning"`
	TimeTakenMs int      `xml:"time_taken_ms"`
	Data        struct {
		TAF []struct {
			RawText       string    `xml:"raw_text"`
			StationID     string    `xml:"station_id"`
			IssueTime     time.Time `xml:"issue_time"`
			BulletinTime  time.Time `xml:"bulletin_time"`
			ValidTimeFrom time.Time `xml:"valid_time_from"`
			ValidTimeTo   time.Time `xml:"valid_time_to"`
			Remarks       string    `xml:"remarks"`
			Latitude      float32   `xml:"latitude"`
			Longitude     float32   `xml:"longitude"`
			ElevationM    float32   `xml:"elevation_m"`
			Forecast      []struct {
				FcstTimeFrom        time.Time `xml:"fcst_time_from"`
				FcstTimeTo          time.Time `xml:"fcst_time_to"`
				ChangeIndicator     string    `xml:"change_indicator"`
				TimeBecoming        time.Time `xml:"time_becoming"`
				Probability         string    `xml:"probability"`
				WindDirDegrees      int       `xml:"wind_dir_degrees"`
				WindSpeedKt         int       `xml:"wind_speed_kt"`
				WindGustKt          int       `xml:"wind_gust_kt"`
				WindShearHgtFtAgl   int       `xml:"wind_shear_hgt_ft_agl"`
				WindShearDirDegrees int       `xml:"wind_shear_dir_degrees"`
				WindShearSpeedKt    int       `xml:"wind_shear_speed_kt"`
				VisibilityStatuteMi float32   `xml:"visibility_statute_mi"`
				AltimInHg           float32   `xml:"altim_in_hg"`
				VertVisFt           int       `xml:"vert_vis_ft"`
				WxString            string    `xml:"wx_string"`
				NotDecoded          string    `xml:"not_decoded"`
				SkyCondition        []struct {
					SkyCover       string `xml:"sky_cover,attr"`
					CloudBaseFtAgl int    `xml:"cloud_base_ft_agl,attr"`
					CloudType      string `xml:"cloud_type,attr"`
				} `xml:"sky_condition"`
				TurbulenceCondition []struct {
					TurbulenceIntensity   string `xml:"turbulence_intensity,attr"`
					TurbulenceMinAltFtAgl int    `xml:"turbulence_min_alt_ft_agl,attr"`
					TurbulenceMaxAltFtAgl int    `xml:"turbulence_max_alt_ft_agl,attr"`
				} `xml:"turbulence_condition"`
				IcingCondition []struct {
					IcingIntensity   string `xml:"icing_intensity,attr"`
					IcingMinAltFtAgl int    `xml:"icing_min_alt_ft_agl,attr"`
					IcingMaxAltFtAgl int    `xml:"icing_max_alt_ft_agl,attr"`
				} `xml:"icing_condition"`
				Temperature []struct {
					ValidTime time.Time `xml:"valid_time,omitempty"`
					SfcTempC  float32   `xml:"sfc_temp_c,omitempty"`
					MaxTempC  string    `xml:"max_temp_c,omitempty"`
					MinTempC  string    `xml:"min_temp_c,omitempty"`
				} `xml:"temperature,omitempty"`
			} `xml:"forecast"`
		} `xml:"TAF"`
		NumResults int `xml:"num_results,attr"`
	} `xml:"data"`
}

type siteType struct {
	METAR        bool `xml:"METAR"`
	TAF          bool `xml:"TAF"`
	WFOoffice    bool `xml:"WFO_office"`
	NEXRAD       bool `xml:"NEXRAD"`
	Rawinsonde   bool `xml:"rawinsonde"`
	WindProfiler bool `xml:"wind_profiler"`
}

func (s *siteType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch tt := token.(type) {
		case xml.StartElement:
			switch tt.Name.Local {
			case "METAR":
				s.METAR = true
			case "TAF":
				s.TAF = true
			case "NEXRAD":
				s.NEXRAD = true
			case "WFO_office":
				s.WFOoffice = true
			case "rawinsonde":
				s.Rawinsonde = true
			case "wind_profiler":
				s.WindProfiler = true
			}
		}
	}
	return nil
}

type StationsInfoResponse struct {
	RequestIndex int `xml:"request_index"`
	DataSource   struct {
		Name string `xml:"name,attr"`
	} `xml:"data_source"`
	Request struct {
		Type string `xml:"type,attr"`
	} `xml:"request"`
	Errors      []string `xml:"errors>error"`
	Warnings    []string `xml:"warnings>warning"`
	NumResults  int      `xml:"num_results"`
	TimeTakenMs int      `xml:"time_taken_ms"`
	Data        struct {
		Station []struct {
			StationID  string   `xml:"station_id"`
			Latitude   float32  `xml:"latitude"`
			Longitude  float32  `xml:"longitude"`
			ElevationM float32  `xml:"elevation_m"`
			Site       string   `xml:"site"`
			Country    string   `xml:"country"`
			SiteType   siteType `xml:"site_type,omitempty"`
		} `xml:"Station"`
	} `xml:"data"`
}

func UnmarshalMetars(input []byte) (result *METARresponse, err error) {
	err = xml.Unmarshal(input, &result)
	return
}

func UnmarshalTafs(input []byte) (result *TAFresponse, err error) {
	err = xml.Unmarshal(input, &result)
	return
}

func UnmarshalStationsInfo(input []byte) (result *StationsInfoResponse, err error) {
	err = xml.Unmarshal(input, &result)
	return
}
