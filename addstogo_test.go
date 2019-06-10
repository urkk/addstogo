package addstogo

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnmarshalStationsInfo(t *testing.T) {
	Convey("Unmarshal stations info should work correctly", t, func() {
		input := []byte(`<response xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XML-Schema-instance" version="1.0" xsi:noNamespaceSchemaLocation="http://weather.aero/schema/station1_0.xsd"><request_index>84789653</request_index><data_source name="stations"/><request type="retrieve"/><errors/><warnings/><time_taken_ms>5</time_taken_ms><data num_results="3"><Station><station_id>KDEN</station_id><wmo_id>72565</wmo_id><latitude>39.85</latitude><longitude>-104.65</longitude><elevation_m>1640.0</elevation_m><site>DENVER (DIA)</site><state>CO</state><country>US</country><site_type><METAR/></site_type></Station><Station><station_id>KSEA</station_id><wmo_id>72793</wmo_id><latitude>47.45</latitude><longitude>-122.32</longitude><elevation_m>136.0</elevation_m><site>SEATTLE/METRO</site><state>WA</state><country>US</country><site_type><METAR/><TAF/></site_type></Station><Station><station_id>PHNL</station_id><wmo_id>91182</wmo_id><latitude>21.33</latitude><longitude>-157.92</longitude><elevation_m>4.0</elevation_m><site>HONOLULU</site><state>HI</state><country>US</country><site_type><METAR/><TAF/></site_type></Station><Station><station_id>KABR</station_id><wmo_id>72659</wmo_id><latitude>45.45</latitude><longitude>-98.42</longitude><elevation_m>397.0</elevation_m><site>ABERDEEN</site><state>SD</state><country>US</country><site_type><METAR/><NEXRAD/><rawinsonde/><WFO_office/><TAF/></site_type></Station></data></response>`)
		expected := &StationsInfoResponse{RequestIndex: 84789653, DataSource: struct {
			Name string "xml:\"name,attr\""
		}{Name: "stations"}, Request: struct {
			Type string "xml:\"type,attr\""
		}{Type: "retrieve"}, Errors: []string(nil), Warnings: []string(nil), NumResults: 0, TimeTakenMs: 5, Data: struct {
			Station []struct {
				StationID  string   `xml:"station_id"`
				Latitude   float32  `xml:"latitude"`
				Longitude  float32  `xml:"longitude"`
				ElevationM float32  `xml:"elevation_m"`
				Site       string   `xml:"site"`
				Country    string   `xml:"country"`
				SiteType   siteType `xml:"site_type,omitempty"`
			} "xml:\"Station\""
		}{Station: []struct {
			StationID  string   `xml:"station_id"`
			Latitude   float32  `xml:"latitude"`
			Longitude  float32  `xml:"longitude"`
			ElevationM float32  `xml:"elevation_m"`
			Site       string   `xml:"site"`
			Country    string   `xml:"country"`
			SiteType   siteType `xml:"site_type,omitempty"`
		}{struct {
			StationID  string   `xml:"station_id"`
			Latitude   float32  `xml:"latitude"`
			Longitude  float32  `xml:"longitude"`
			ElevationM float32  `xml:"elevation_m"`
			Site       string   `xml:"site"`
			Country    string   `xml:"country"`
			SiteType   siteType `xml:"site_type,omitempty"`
		}{StationID: "KDEN", Latitude: 39.85, Longitude: -104.65, ElevationM: 1640, Site: "DENVER (DIA)", Country: "US", SiteType: siteType{METAR: true, TAF: false, WFOoffice: false, NEXRAD: false, Rawinsonde: false, WindProfiler: false}}, struct {
			StationID  string   `xml:"station_id"`
			Latitude   float32  `xml:"latitude"`
			Longitude  float32  `xml:"longitude"`
			ElevationM float32  `xml:"elevation_m"`
			Site       string   `xml:"site"`
			Country    string   `xml:"country"`
			SiteType   siteType `xml:"site_type,omitempty"`
		}{StationID: "KSEA", Latitude: 47.45, Longitude: -122.32, ElevationM: 136, Site: "SEATTLE/METRO", Country: "US", SiteType: siteType{METAR: true, TAF: true, WFOoffice: false, NEXRAD: false, Rawinsonde: false, WindProfiler: false}}, struct {
			StationID  string   `xml:"station_id"`
			Latitude   float32  `xml:"latitude"`
			Longitude  float32  `xml:"longitude"`
			ElevationM float32  `xml:"elevation_m"`
			Site       string   `xml:"site"`
			Country    string   `xml:"country"`
			SiteType   siteType `xml:"site_type,omitempty"`
		}{StationID: "PHNL", Latitude: 21.33, Longitude: -157.92, ElevationM: 4, Site: "HONOLULU", Country: "US", SiteType: siteType{METAR: true, TAF: true, WFOoffice: false, NEXRAD: false, Rawinsonde: false, WindProfiler: false}}, struct {
			StationID  string   `xml:"station_id"`
			Latitude   float32  `xml:"latitude"`
			Longitude  float32  `xml:"longitude"`
			ElevationM float32  `xml:"elevation_m"`
			Site       string   `xml:"site"`
			Country    string   `xml:"country"`
			SiteType   siteType `xml:"site_type,omitempty"`
		}{StationID: "KABR", Latitude: 45.45, Longitude: -98.42, ElevationM: 397, Site: "ABERDEEN", Country: "US", SiteType: siteType{METAR: true, TAF: true, WFOoffice: true, NEXRAD: true, Rawinsonde: true, WindProfiler: false}}}}}

		si, err := UnmarshalStationsInfo(input)
		Convey("struct should be builded correctly", func() {
			So(si, ShouldResemble, expected)
		})
		Convey("err must bi nil", func() {
			So(err, ShouldBeNil)
		})
	})
}

func TestUnmarshalTafs(t *testing.T) {
	Convey("Unmarshal TAF should work correctly", t, func() {
		input := []byte(`<response xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XML-Schema-instance" version="1.2" xsi:noNamespaceSchemaLocation="http://aviationweather.gov/adds/schema/taf1_2.xsd"><request_index>36724144</request_index><data_source name="tafs"/><request type="retrieve"/><errors/><warnings/><time_taken_ms>9</time_taken_ms><data num_results="1"><TAF><raw_text>TAF URSS 070456Z 0706/0806 23005MPS 9999 FEW040 BECMG 0708/0709 28006G11MPS SCT030CB TEMPO 0709/0717 -TSRA BECMG 0717/0718 05005MPS BKN011 TEMPO 0718/0806 VRB06G11MPS -TSRA BKN007 SCT030CB</raw_text><station_id>URSS</station_id><issue_time>2019-06-07T04:56:00Z</issue_time><bulletin_time>2019-06-07T05:00:00Z</bulletin_time><valid_time_from>2019-06-07T06:00:00Z</valid_time_from><valid_time_to>2019-06-08T06:00:00Z</valid_time_to><latitude>43.45</latitude><longitude>39.95</longitude><elevation_m>16.0</elevation_m><forecast><fcst_time_from>2019-06-07T06:00:00Z</fcst_time_from><fcst_time_to>2019-06-07T08:00:00Z</fcst_time_to><wind_dir_degrees>230</wind_dir_degrees><wind_speed_kt>10</wind_speed_kt><visibility_statute_mi>6.21</visibility_statute_mi><sky_condition sky_cover="FEW" cloud_base_ft_agl="4000"/></forecast><forecast><fcst_time_from>2019-06-07T08:00:00Z</fcst_time_from><fcst_time_to>2019-06-07T17:00:00Z</fcst_time_to><change_indicator>BECMG</change_indicator><time_becoming>2019-06-07T09:00:00Z</time_becoming><wind_dir_degrees>280</wind_dir_degrees><wind_speed_kt>12</wind_speed_kt><wind_gust_kt>21</wind_gust_kt><visibility_statute_mi>6.21</visibility_statute_mi><sky_condition sky_cover="SCT" cloud_base_ft_agl="3000" cloud_type="CB"/></forecast><forecast><fcst_time_from>2019-06-07T09:00:00Z</fcst_time_from><fcst_time_to>2019-06-07T17:00:00Z</fcst_time_to><change_indicator>TEMPO</change_indicator><wx_string>-TSRA</wx_string></forecast><forecast><fcst_time_from>2019-06-07T17:00:00Z</fcst_time_from><fcst_time_to>2019-06-08T06:00:00Z</fcst_time_to><change_indicator>BECMG</change_indicator><time_becoming>2019-06-07T18:00:00Z</time_becoming><wind_dir_degrees>50</wind_dir_degrees><wind_speed_kt>10</wind_speed_kt><visibility_statute_mi>6.21</visibility_statute_mi><sky_condition sky_cover="BKN" cloud_base_ft_agl="1100"/></forecast><forecast><fcst_time_from>2019-06-07T18:00:00Z</fcst_time_from><fcst_time_to>2019-06-08T06:00:00Z</fcst_time_to><change_indicator>TEMPO</change_indicator><wind_dir_degrees>0</wind_dir_degrees><wind_speed_kt>12</wind_speed_kt><wind_gust_kt>21</wind_gust_kt><wx_string>-TSRA</wx_string><sky_condition sky_cover="BKN" cloud_base_ft_agl="700"/><sky_condition sky_cover="SCT" cloud_base_ft_agl="3000" cloud_type="CB"/></forecast></TAF></data></response>`)
		expected := &TAFresponse{RequestIndex: 36724144, DataSource: struct {
			Name string "xml:\"name,attr\""
		}{Name: "tafs"}, Request: struct {
			Type string "xml:\"type,attr\""
		}{Type: "retrieve"}, Errors: []string(nil), Warnings: []string(nil), TimeTakenMs: 9, Data: struct {
			TAF []struct {
				RawText       string    "xml:\"raw_text\""
				StationID     string    "xml:\"station_id\""
				IssueTime     time.Time "xml:\"issue_time\""
				BulletinTime  time.Time "xml:\"bulletin_time\""
				ValidTimeFrom time.Time "xml:\"valid_time_from\""
				ValidTimeTo   time.Time "xml:\"valid_time_to\""
				Remarks       string    "xml:\"remarks\""
				Latitude      float32   "xml:\"latitude\""
				Longitude     float32   "xml:\"longitude\""
				ElevationM    float32   "xml:\"elevation_m\""
				Forecast      []struct {
					FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
					FcstTimeTo          time.Time "xml:\"fcst_time_to\""
					ChangeIndicator     string    "xml:\"change_indicator\""
					TimeBecoming        time.Time "xml:\"time_becoming\""
					Probability         string    "xml:\"probability\""
					WindDirDegrees      int       "xml:\"wind_dir_degrees\""
					WindSpeedKt         int       "xml:\"wind_speed_kt\""
					WindGustKt          int       "xml:\"wind_gust_kt\""
					WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
					WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
					WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
					VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
					AltimInHg           float32   "xml:\"altim_in_hg\""
					VertVisFt           int       "xml:\"vert_vis_ft\""
					WxString            string    "xml:\"wx_string\""
					NotDecoded          string    "xml:\"not_decoded\""
					SkyCondition        []struct {
						SkyCover       string "xml:\"sky_cover,attr\""
						CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
						CloudType      string "xml:\"cloud_type,attr\""
					} "xml:\"sky_condition\""
					TurbulenceCondition []struct {
						TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
						TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
						TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
					} "xml:\"turbulence_condition\""
					IcingCondition []struct {
						IcingIntensity   string "xml:\"icing_intensity,attr\""
						IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
						IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
					} "xml:\"icing_condition\""
					Temperature []struct {
						ValidTime time.Time "xml:\"valid_time,omitempty\""
						SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
						MaxTempC  string    "xml:\"max_temp_c,omitempty\""
						MinTempC  string    "xml:\"min_temp_c,omitempty\""
					} "xml:\"temperature,omitempty\""
				} "xml:\"forecast\""
			} "xml:\"TAF\""
			NumResults int "xml:\"num_results,attr\""
		}{TAF: []struct {
			RawText       string    "xml:\"raw_text\""
			StationID     string    "xml:\"station_id\""
			IssueTime     time.Time "xml:\"issue_time\""
			BulletinTime  time.Time "xml:\"bulletin_time\""
			ValidTimeFrom time.Time "xml:\"valid_time_from\""
			ValidTimeTo   time.Time "xml:\"valid_time_to\""
			Remarks       string    "xml:\"remarks\""
			Latitude      float32   "xml:\"latitude\""
			Longitude     float32   "xml:\"longitude\""
			ElevationM    float32   "xml:\"elevation_m\""
			Forecast      []struct {
				FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
				FcstTimeTo          time.Time "xml:\"fcst_time_to\""
				ChangeIndicator     string    "xml:\"change_indicator\""
				TimeBecoming        time.Time "xml:\"time_becoming\""
				Probability         string    "xml:\"probability\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
				WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
				WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				VertVisFt           int       "xml:\"vert_vis_ft\""
				WxString            string    "xml:\"wx_string\""
				NotDecoded          string    "xml:\"not_decoded\""
				SkyCondition        []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				} "xml:\"sky_condition\""
				TurbulenceCondition []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				} "xml:\"turbulence_condition\""
				IcingCondition []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				} "xml:\"icing_condition\""
				Temperature []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				} "xml:\"temperature,omitempty\""
			} "xml:\"forecast\""
		}{struct {
			RawText       string    "xml:\"raw_text\""
			StationID     string    "xml:\"station_id\""
			IssueTime     time.Time "xml:\"issue_time\""
			BulletinTime  time.Time "xml:\"bulletin_time\""
			ValidTimeFrom time.Time "xml:\"valid_time_from\""
			ValidTimeTo   time.Time "xml:\"valid_time_to\""
			Remarks       string    "xml:\"remarks\""
			Latitude      float32   "xml:\"latitude\""
			Longitude     float32   "xml:\"longitude\""
			ElevationM    float32   "xml:\"elevation_m\""
			Forecast      []struct {
				FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
				FcstTimeTo          time.Time "xml:\"fcst_time_to\""
				ChangeIndicator     string    "xml:\"change_indicator\""
				TimeBecoming        time.Time "xml:\"time_becoming\""
				Probability         string    "xml:\"probability\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
				WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
				WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				VertVisFt           int       "xml:\"vert_vis_ft\""
				WxString            string    "xml:\"wx_string\""
				NotDecoded          string    "xml:\"not_decoded\""
				SkyCondition        []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				} "xml:\"sky_condition\""
				TurbulenceCondition []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				} "xml:\"turbulence_condition\""
				IcingCondition []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				} "xml:\"icing_condition\""
				Temperature []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				} "xml:\"temperature,omitempty\""
			} "xml:\"forecast\""
		}{RawText: "TAF URSS 070456Z 0706/0806 23005MPS 9999 FEW040 BECMG 0708/0709 28006G11MPS SCT030CB TEMPO 0709/0717 -TSRA BECMG 0717/0718 05005MPS BKN011 TEMPO 0718/0806 VRB06G11MPS -TSRA BKN007 SCT030CB", StationID: "URSS",
			IssueTime:     time.Date(2019, 06, 7, 4, 56, 0, 0, time.UTC),
			BulletinTime:  time.Date(2019, 06, 7, 5, 0, 0, 0, time.UTC),
			ValidTimeFrom: time.Date(2019, 06, 7, 6, 0, 0, 0, time.UTC),
			ValidTimeTo:   time.Date(2019, 06, 8, 6, 0, 0, 0, time.UTC), Remarks: "", Latitude: 43.45, Longitude: 39.95, ElevationM: 16, Forecast: []struct {
				FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
				FcstTimeTo          time.Time "xml:\"fcst_time_to\""
				ChangeIndicator     string    "xml:\"change_indicator\""
				TimeBecoming        time.Time "xml:\"time_becoming\""
				Probability         string    "xml:\"probability\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
				WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
				WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				VertVisFt           int       "xml:\"vert_vis_ft\""
				WxString            string    "xml:\"wx_string\""
				NotDecoded          string    "xml:\"not_decoded\""
				SkyCondition        []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				} "xml:\"sky_condition\""
				TurbulenceCondition []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				} "xml:\"turbulence_condition\""
				IcingCondition []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				} "xml:\"icing_condition\""
				Temperature []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				} "xml:\"temperature,omitempty\""
			}{struct {
				FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
				FcstTimeTo          time.Time "xml:\"fcst_time_to\""
				ChangeIndicator     string    "xml:\"change_indicator\""
				TimeBecoming        time.Time "xml:\"time_becoming\""
				Probability         string    "xml:\"probability\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
				WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
				WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				VertVisFt           int       "xml:\"vert_vis_ft\""
				WxString            string    "xml:\"wx_string\""
				NotDecoded          string    "xml:\"not_decoded\""
				SkyCondition        []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				} "xml:\"sky_condition\""
				TurbulenceCondition []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				} "xml:\"turbulence_condition\""
				IcingCondition []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				} "xml:\"icing_condition\""
				Temperature []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				} "xml:\"temperature,omitempty\""
			}{FcstTimeFrom: time.Date(2019, 6, 7, 6, 0, 0, 0, time.UTC),
				FcstTimeTo:      time.Date(2019, 06, 7, 8, 0, 0, 0, time.UTC),
				ChangeIndicator: "", TimeBecoming: time.Time{}, Probability: "", WindDirDegrees: 230, WindSpeedKt: 10, WindGustKt: 0, WindShearHgtFtAgl: 0, WindShearDirDegrees: 0, WindShearSpeedKt: 0, VisibilityStatuteMi: 6.21, AltimInHg: 0, VertVisFt: 0, WxString: "", NotDecoded: "", SkyCondition: []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{SkyCover: "FEW", CloudBaseFtAgl: 4000, CloudType: ""}}, TurbulenceCondition: []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				}(nil), IcingCondition: []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				}(nil), Temperature: []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				}(nil)}, struct {
				FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
				FcstTimeTo          time.Time "xml:\"fcst_time_to\""
				ChangeIndicator     string    "xml:\"change_indicator\""
				TimeBecoming        time.Time "xml:\"time_becoming\""
				Probability         string    "xml:\"probability\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
				WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
				WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				VertVisFt           int       "xml:\"vert_vis_ft\""
				WxString            string    "xml:\"wx_string\""
				NotDecoded          string    "xml:\"not_decoded\""
				SkyCondition        []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				} "xml:\"sky_condition\""
				TurbulenceCondition []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				} "xml:\"turbulence_condition\""
				IcingCondition []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				} "xml:\"icing_condition\""
				Temperature []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				} "xml:\"temperature,omitempty\""
			}{FcstTimeFrom: time.Date(2019, 06, 7, 8, 0, 0, 0, time.UTC),
				FcstTimeTo:      time.Date(2019, 06, 7, 17, 0, 0, 0, time.UTC),
				ChangeIndicator: "BECMG",
				TimeBecoming:    time.Date(2019, 06, 7, 9, 0, 0, 0, time.UTC),
				Probability:     "", WindDirDegrees: 280, WindSpeedKt: 12, WindGustKt: 21, WindShearHgtFtAgl: 0, WindShearDirDegrees: 0, WindShearSpeedKt: 0, VisibilityStatuteMi: 6.21, AltimInHg: 0, VertVisFt: 0, WxString: "", NotDecoded: "", SkyCondition: []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{SkyCover: "SCT", CloudBaseFtAgl: 3000, CloudType: "CB"}}, TurbulenceCondition: []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				}(nil), IcingCondition: []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				}(nil), Temperature: []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				}(nil)}, struct {
				FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
				FcstTimeTo          time.Time "xml:\"fcst_time_to\""
				ChangeIndicator     string    "xml:\"change_indicator\""
				TimeBecoming        time.Time "xml:\"time_becoming\""
				Probability         string    "xml:\"probability\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
				WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
				WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				VertVisFt           int       "xml:\"vert_vis_ft\""
				WxString            string    "xml:\"wx_string\""
				NotDecoded          string    "xml:\"not_decoded\""
				SkyCondition        []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				} "xml:\"sky_condition\""
				TurbulenceCondition []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				} "xml:\"turbulence_condition\""
				IcingCondition []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				} "xml:\"icing_condition\""
				Temperature []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				} "xml:\"temperature,omitempty\""
			}{FcstTimeFrom: time.Date(2019, 06, 7, 9, 0, 0, 0, time.UTC),
				FcstTimeTo:      time.Date(2019, 06, 7, 17, 0, 0, 0, time.UTC),
				ChangeIndicator: "TEMPO", TimeBecoming: time.Time{}, Probability: "", WindDirDegrees: 0, WindSpeedKt: 0, WindGustKt: 0, WindShearHgtFtAgl: 0, WindShearDirDegrees: 0, WindShearSpeedKt: 0, VisibilityStatuteMi: 0, AltimInHg: 0, VertVisFt: 0, WxString: "-TSRA", NotDecoded: "", SkyCondition: []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}(nil), TurbulenceCondition: []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				}(nil), IcingCondition: []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				}(nil), Temperature: []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				}(nil)}, struct {
				FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
				FcstTimeTo          time.Time "xml:\"fcst_time_to\""
				ChangeIndicator     string    "xml:\"change_indicator\""
				TimeBecoming        time.Time "xml:\"time_becoming\""
				Probability         string    "xml:\"probability\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
				WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
				WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				VertVisFt           int       "xml:\"vert_vis_ft\""
				WxString            string    "xml:\"wx_string\""
				NotDecoded          string    "xml:\"not_decoded\""
				SkyCondition        []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				} "xml:\"sky_condition\""
				TurbulenceCondition []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				} "xml:\"turbulence_condition\""
				IcingCondition []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				} "xml:\"icing_condition\""
				Temperature []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				} "xml:\"temperature,omitempty\""
			}{FcstTimeFrom: time.Date(2019, 06, 7, 17, 0, 0, 0, time.UTC),
				FcstTimeTo:      time.Date(2019, 06, 8, 6, 0, 0, 0, time.UTC),
				ChangeIndicator: "BECMG",
				TimeBecoming:    time.Date(2019, 06, 7, 18, 0, 0, 0, time.UTC),
				Probability:     "", WindDirDegrees: 50, WindSpeedKt: 10, WindGustKt: 0, WindShearHgtFtAgl: 0, WindShearDirDegrees: 0, WindShearSpeedKt: 0, VisibilityStatuteMi: 6.21, AltimInHg: 0, VertVisFt: 0, WxString: "", NotDecoded: "", SkyCondition: []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{SkyCover: "BKN", CloudBaseFtAgl: 1100, CloudType: ""}}, TurbulenceCondition: []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				}(nil), IcingCondition: []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				}(nil), Temperature: []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				}(nil)}, struct {
				FcstTimeFrom        time.Time "xml:\"fcst_time_from\""
				FcstTimeTo          time.Time "xml:\"fcst_time_to\""
				ChangeIndicator     string    "xml:\"change_indicator\""
				TimeBecoming        time.Time "xml:\"time_becoming\""
				Probability         string    "xml:\"probability\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				WindShearHgtFtAgl   int       "xml:\"wind_shear_hgt_ft_agl\""
				WindShearDirDegrees int       "xml:\"wind_shear_dir_degrees\""
				WindShearSpeedKt    int       "xml:\"wind_shear_speed_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				VertVisFt           int       "xml:\"vert_vis_ft\""
				WxString            string    "xml:\"wx_string\""
				NotDecoded          string    "xml:\"not_decoded\""
				SkyCondition        []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				} "xml:\"sky_condition\""
				TurbulenceCondition []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				} "xml:\"turbulence_condition\""
				IcingCondition []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				} "xml:\"icing_condition\""
				Temperature []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				} "xml:\"temperature,omitempty\""
			}{
				FcstTimeFrom:    time.Date(2019, 06, 7, 18, 0, 0, 0, time.UTC),
				FcstTimeTo:      time.Date(2019, 06, 8, 6, 0, 0, 0, time.UTC),
				ChangeIndicator: "TEMPO", TimeBecoming: time.Time{}, Probability: "", WindDirDegrees: 0, WindSpeedKt: 12, WindGustKt: 21, WindShearHgtFtAgl: 0, WindShearDirDegrees: 0, WindShearSpeedKt: 0, VisibilityStatuteMi: 0, AltimInHg: 0, VertVisFt: 0, WxString: "-TSRA", NotDecoded: "", SkyCondition: []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{SkyCover: "BKN", CloudBaseFtAgl: 700, CloudType: ""}, struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
					CloudType      string "xml:\"cloud_type,attr\""
				}{SkyCover: "SCT", CloudBaseFtAgl: 3000, CloudType: "CB"}}, TurbulenceCondition: []struct {
					TurbulenceIntensity   string "xml:\"turbulence_intensity,attr\""
					TurbulenceMinAltFtAgl int    "xml:\"turbulence_min_alt_ft_agl,attr\""
					TurbulenceMaxAltFtAgl int    "xml:\"turbulence_max_alt_ft_agl,attr\""
				}(nil), IcingCondition: []struct {
					IcingIntensity   string "xml:\"icing_intensity,attr\""
					IcingMinAltFtAgl int    "xml:\"icing_min_alt_ft_agl,attr\""
					IcingMaxAltFtAgl int    "xml:\"icing_max_alt_ft_agl,attr\""
				}(nil), Temperature: []struct {
					ValidTime time.Time "xml:\"valid_time,omitempty\""
					SfcTempC  float32   "xml:\"sfc_temp_c,omitempty\""
					MaxTempC  string    "xml:\"max_temp_c,omitempty\""
					MinTempC  string    "xml:\"min_temp_c,omitempty\""
				}(nil)}}}}, NumResults: 1}}

		si, err := UnmarshalTafs(input)
		Convey("struct should be builded correctly", func() {
			So(si, ShouldResemble, expected)
		})
		Convey("err must bi nil", func() {
			So(err, ShouldBeNil)
		})
	})
}

func TestUnmarshalMetars(t *testing.T) {
	Convey("Unmarshal METAR should work correctly", t, func() {
		input := []byte(`<response xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XML-Schema-instance" version="1.2" xsi:noNamespaceSchemaLocation="http://aviationweather.gov/adds/schema/metar1_2.xsd"><request_index>45754830</request_index><data_source name="metars"/><request type="retrieve"/><errors/><warnings/><time_taken_ms>4</time_taken_ms><data num_results="1"><METAR><raw_text>ULLI 100800Z 23007MPS 210V270 9999 FEW040 20/11 Q1022 R88/090060 NOSIG</raw_text><station_id>ULLI</station_id><observation_time>2019-06-10T08:00:00Z</observation_time><latitude>59.8</latitude><longitude>30.27</longitude><temp_c>20.0</temp_c><dewpoint_c>11.0</dewpoint_c><wind_dir_degrees>230</wind_dir_degrees><wind_speed_kt>14</wind_speed_kt><visibility_statute_mi>6.21</visibility_statute_mi><altim_in_hg>30.177166</altim_in_hg><quality_control_flags><no_signal>TRUE</no_signal></quality_control_flags><sky_condition sky_cover="FEW" cloud_base_ft_agl="4000"/><flight_category>VFR</flight_category><metar_type>METAR</metar_type><elevation_m>4.0</elevation_m></METAR></data></response>`)
		expected := &METARresponse{RequestIndex: 45754830, DataSource: struct {
			Name string "xml:\"name,attr\""
		}{Name: "metars"}, Request: struct {
			Type string "xml:\"type,attr\""
		}{Type: "retrieve"}, Errors: nil, Warnings: nil, TimeTakenMs: 4, Data: struct {
			METAR []struct {
				RawText             string    "xml:\"raw_text\""
				StationID           string    "xml:\"station_id\""
				ObservationTime     time.Time "xml:\"observation_time\""
				Latitude            float32   "xml:\"latitude\""
				Longitude           float32   "xml:\"longitude\""
				TempC               float32   "xml:\"temp_c\""
				DewpointC           float32   "xml:\"dewpoint_c\""
				WindDirDegrees      int       "xml:\"wind_dir_degrees\""
				WindSpeedKt         int       "xml:\"wind_speed_kt\""
				WindGustKt          int       "xml:\"wind_gust_kt\""
				VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
				AltimInHg           float32   "xml:\"altim_in_hg\""
				SeaLevelPressureMb  float32   "xml:\"sea_level_pressure_mb\""
				QualityControlFlags struct {
					Corrected               bool "xml:\"corrected\""
					Auto                    bool "xml:\"auto\""
					AutoStation             bool "xml:\"auto_station\""
					MaintenanceIndicatorOn  bool "xml:\"maintenance_indicator_on\""
					NoSignal                bool "xml:\"no_signal\""
					LightningSensorOff      bool "xml:\"lightning_sensor_off\""
					FreezingRainSensorOff   bool "xml:\"freezing_rain_sensor_off\""
					PresentWeatherSensorOff bool "xml:\"present_weather_sensor_off\""
				} "xml:\"quality_control_flags\""
				WxString     string "xml:\"wx_string\""
				SkyCondition []struct {
					SkyCover       string "xml:\"sky_cover,attr\""
					CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
				} "xml:\"sky_condition\""
				FlightCategory            string  "xml:\"flight_category\""
				ThreeHrPressureTendencyMb float32 "xml:\"three_hr_pressure_tendency_mb\""
				MaxTC                     float32 "xml:\"maxT_c\""
				MinTC                     float32 "xml:\"minT_c\""
				MaxT24HrC                 float32 "xml:\"maxT24hr_c\""
				MinT24HrC                 float32 "xml:\"minT24hr_c\""
				PrecipIn                  float32 "xml:\"precip_in\""
				Pcp3HrIn                  float32 "xml:\"pcp3hr_in\""
				Pcp6HrIn                  float32 "xml:\"pcp6hr_in\""
				Pcp24HrIn                 float32 "xml:\"pcp24hr_in\""
				SnowIn                    float32 "xml:\"snow_in\""
				VertVisFt                 int     "xml:\"vert_vis_ft\""
				MetarType                 string  "xml:\"metar_type\""
				ElevationM                float32 "xml:\"elevation_m\""
			} "xml:\"METAR\""
			NumResults int "xml:\"num_results,attr\""
		}{METAR: []struct {
			RawText             string    "xml:\"raw_text\""
			StationID           string    "xml:\"station_id\""
			ObservationTime     time.Time "xml:\"observation_time\""
			Latitude            float32   "xml:\"latitude\""
			Longitude           float32   "xml:\"longitude\""
			TempC               float32   "xml:\"temp_c\""
			DewpointC           float32   "xml:\"dewpoint_c\""
			WindDirDegrees      int       "xml:\"wind_dir_degrees\""
			WindSpeedKt         int       "xml:\"wind_speed_kt\""
			WindGustKt          int       "xml:\"wind_gust_kt\""
			VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
			AltimInHg           float32   "xml:\"altim_in_hg\""
			SeaLevelPressureMb  float32   "xml:\"sea_level_pressure_mb\""
			QualityControlFlags struct {
				Corrected               bool "xml:\"corrected\""
				Auto                    bool "xml:\"auto\""
				AutoStation             bool "xml:\"auto_station\""
				MaintenanceIndicatorOn  bool "xml:\"maintenance_indicator_on\""
				NoSignal                bool "xml:\"no_signal\""
				LightningSensorOff      bool "xml:\"lightning_sensor_off\""
				FreezingRainSensorOff   bool "xml:\"freezing_rain_sensor_off\""
				PresentWeatherSensorOff bool "xml:\"present_weather_sensor_off\""
			} "xml:\"quality_control_flags\""
			WxString     string "xml:\"wx_string\""
			SkyCondition []struct {
				SkyCover       string "xml:\"sky_cover,attr\""
				CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
			} "xml:\"sky_condition\""
			FlightCategory            string  "xml:\"flight_category\""
			ThreeHrPressureTendencyMb float32 "xml:\"three_hr_pressure_tendency_mb\""
			MaxTC                     float32 "xml:\"maxT_c\""
			MinTC                     float32 "xml:\"minT_c\""
			MaxT24HrC                 float32 "xml:\"maxT24hr_c\""
			MinT24HrC                 float32 "xml:\"minT24hr_c\""
			PrecipIn                  float32 "xml:\"precip_in\""
			Pcp3HrIn                  float32 "xml:\"pcp3hr_in\""
			Pcp6HrIn                  float32 "xml:\"pcp6hr_in\""
			Pcp24HrIn                 float32 "xml:\"pcp24hr_in\""
			SnowIn                    float32 "xml:\"snow_in\""
			VertVisFt                 int     "xml:\"vert_vis_ft\""
			MetarType                 string  "xml:\"metar_type\""
			ElevationM                float32 "xml:\"elevation_m\""
		}{struct {
			RawText             string    "xml:\"raw_text\""
			StationID           string    "xml:\"station_id\""
			ObservationTime     time.Time "xml:\"observation_time\""
			Latitude            float32   "xml:\"latitude\""
			Longitude           float32   "xml:\"longitude\""
			TempC               float32   "xml:\"temp_c\""
			DewpointC           float32   "xml:\"dewpoint_c\""
			WindDirDegrees      int       "xml:\"wind_dir_degrees\""
			WindSpeedKt         int       "xml:\"wind_speed_kt\""
			WindGustKt          int       "xml:\"wind_gust_kt\""
			VisibilityStatuteMi float32   "xml:\"visibility_statute_mi\""
			AltimInHg           float32   "xml:\"altim_in_hg\""
			SeaLevelPressureMb  float32   "xml:\"sea_level_pressure_mb\""
			QualityControlFlags struct {
				Corrected               bool "xml:\"corrected\""
				Auto                    bool "xml:\"auto\""
				AutoStation             bool "xml:\"auto_station\""
				MaintenanceIndicatorOn  bool "xml:\"maintenance_indicator_on\""
				NoSignal                bool "xml:\"no_signal\""
				LightningSensorOff      bool "xml:\"lightning_sensor_off\""
				FreezingRainSensorOff   bool "xml:\"freezing_rain_sensor_off\""
				PresentWeatherSensorOff bool "xml:\"present_weather_sensor_off\""
			} "xml:\"quality_control_flags\""
			WxString     string "xml:\"wx_string\""
			SkyCondition []struct {
				SkyCover       string "xml:\"sky_cover,attr\""
				CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
			} "xml:\"sky_condition\""
			FlightCategory            string  "xml:\"flight_category\""
			ThreeHrPressureTendencyMb float32 "xml:\"three_hr_pressure_tendency_mb\""
			MaxTC                     float32 "xml:\"maxT_c\""
			MinTC                     float32 "xml:\"minT_c\""
			MaxT24HrC                 float32 "xml:\"maxT24hr_c\""
			MinT24HrC                 float32 "xml:\"minT24hr_c\""
			PrecipIn                  float32 "xml:\"precip_in\""
			Pcp3HrIn                  float32 "xml:\"pcp3hr_in\""
			Pcp6HrIn                  float32 "xml:\"pcp6hr_in\""
			Pcp24HrIn                 float32 "xml:\"pcp24hr_in\""
			SnowIn                    float32 "xml:\"snow_in\""
			VertVisFt                 int     "xml:\"vert_vis_ft\""
			MetarType                 string  "xml:\"metar_type\""
			ElevationM                float32 "xml:\"elevation_m\""
		}{RawText: "ULLI 100800Z 23007MPS 210V270 9999 FEW040 20/11 Q1022 R88/090060 NOSIG",
			StationID:       "ULLI",
			ObservationTime: time.Date(2019, 06, 10, 8, 0, 0, 0, time.UTC),
			Latitude:        59.8, Longitude: 30.27, TempC: 20, DewpointC: 11, WindDirDegrees: 230, WindSpeedKt: 14, WindGustKt: 0, VisibilityStatuteMi: 6.21, AltimInHg: 30.177166, SeaLevelPressureMb: 0, QualityControlFlags: struct {
				Corrected               bool "xml:\"corrected\""
				Auto                    bool "xml:\"auto\""
				AutoStation             bool "xml:\"auto_station\""
				MaintenanceIndicatorOn  bool "xml:\"maintenance_indicator_on\""
				NoSignal                bool "xml:\"no_signal\""
				LightningSensorOff      bool "xml:\"lightning_sensor_off\""
				FreezingRainSensorOff   bool "xml:\"freezing_rain_sensor_off\""
				PresentWeatherSensorOff bool "xml:\"present_weather_sensor_off\""
			}{Corrected: false, Auto: false, AutoStation: false, MaintenanceIndicatorOn: false, NoSignal: true, LightningSensorOff: false, FreezingRainSensorOff: false, PresentWeatherSensorOff: false}, WxString: "", SkyCondition: []struct {
				SkyCover       string "xml:\"sky_cover,attr\""
				CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
			}{struct {
				SkyCover       string "xml:\"sky_cover,attr\""
				CloudBaseFtAgl int    "xml:\"cloud_base_ft_agl,attr\""
			}{SkyCover: "FEW", CloudBaseFtAgl: 4000}}, FlightCategory: "VFR", ThreeHrPressureTendencyMb: 0, MaxTC: 0, MinTC: 0, MaxT24HrC: 0, MinT24HrC: 0, PrecipIn: 0, Pcp3HrIn: 0, Pcp6HrIn: 0, Pcp24HrIn: 0, SnowIn: 0, VertVisFt: 0, MetarType: "METAR", ElevationM: 4}}, NumResults: 1}}

		si, err := UnmarshalMetars(input)
		Convey("struct should be builded correctly", func() {
			So(si, ShouldResemble, expected)
		})
		Convey("err must bi nil", func() {
			So(err, ShouldBeNil)
		})
	})
}
