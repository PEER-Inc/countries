package countries

import (
	"encoding/csv"
	"fmt"
	"gopkg.in/yaml.v3"
	"path/filepath"
	"sort"
	"strings"
)

func loadCountries(countriesPath string, out map[string]Country) error {
	files, err := content.ReadDir(countriesPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		path := filepath.Join(countriesPath, file.Name())
		buf, err := content.ReadFile(path)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(buf, &out)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadSubdivisions(subdivisionsPath string, out map[string]map[string]*Subdivision) error {
	files, err := content.ReadDir(subdivisionsPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		subdivisions := make(map[string]*Subdivision)
		path := filepath.Join(subdivisionsPath, file.Name())
		buf, err := content.ReadFile(path)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(buf, &subdivisions)
		if err != nil {
			panic(err)
		}
		countryAlpha2 := filenameToCountryAlpha2(file.Name())
		out[countryAlpha2] = subdivisions
	}
	return nil
}

func loadTranslations(translationsPath string, out map[string]map[string]string) error {
	files, err := content.ReadDir(translationsPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		translations := make(map[string]string)
		path := filepath.Join(translationsPath, file.Name())
		buf, err := content.ReadFile(path)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(buf, &translations)
		if err != nil {
			return err
		}
		locale := filenameToLocale(file.Name())
		out[locale] = translations
	}
	return nil
}

func loadCapitals(capitalsPath string, out map[string]string) error {
	buf, err := content.ReadFile(capitalsPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &out)
	if err != nil {
		return err
	}
	return nil
}

// CSV file link: https://timezonedb.com/files/timezonedb.csv.zip
func loadTimezones(timezonesPath string, out map[string][]string) error {
	f, err := content.Open(timezonesPath)
	if err != nil {
		return err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	for _, row := range records {
		out[row[1]] = append(out[row[1]], row[2])
	}
	return nil
}

func countryToCodeString(c Country) string {
	s := fmt.Sprintf("%#v", c)
	s = strings.ReplaceAll(s, "countries.", "")
	s = strings.ReplaceAll(s, "Country{", "{")
	s = strings.ReplaceAll(s, ":Subdivision{", ":{")
	return s
}

func alpha2(countries []Country) []string {
	result := make([]string, len(countries))
	for i := range countries {
		result[i] = countries[i].Alpha2
	}
	return result
}

func regions(countries []Country) []string {
	var result []string
	set := make(map[string]struct{})
	for _, c := range countries {
		if c.Region != "" {
			set[c.Region] = struct{}{}
		}
	}
	for r := range set {
		result = append(result, r)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func subregions(countries []Country) []string {
	var result []string
	set := make(map[string]struct{})
	for _, c := range countries {
		if c.Subregion != "" {
			set[c.Subregion] = struct{}{}
		}
	}
	for r := range set {
		result = append(result, r)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func filenameToCountryAlpha2(filename string) string {
	return strings.ReplaceAll(filename, ".yaml", "")
}

func filenameToLocale(filename string) string {
	s := strings.ReplaceAll(filename, ".yaml", "")
	s = strings.ReplaceAll(s, "countries-", "")
	return s
}

// Get returns the country identified by alpha2 code.
func Get(alpha2 string) *Country {
	switch alpha2 {
	case "AD":
		return &Data.All[0]
	case "AE":
		return &Data.All[1]
	case "AF":
		return &Data.All[2]
	case "AG":
		return &Data.All[3]
	case "AI":
		return &Data.All[4]
	case "AL":
		return &Data.All[5]
	case "AM":
		return &Data.All[6]
	case "AO":
		return &Data.All[7]
	case "AQ":
		return &Data.All[8]
	case "AR":
		return &Data.All[9]
	case "AS":
		return &Data.All[10]
	case "AT":
		return &Data.All[11]
	case "AU":
		return &Data.All[12]
	case "AW":
		return &Data.All[13]
	case "AX":
		return &Data.All[14]
	case "AZ":
		return &Data.All[15]
	case "BA":
		return &Data.All[16]
	case "BB":
		return &Data.All[17]
	case "BD":
		return &Data.All[18]
	case "BE":
		return &Data.All[19]
	case "BF":
		return &Data.All[20]
	case "BG":
		return &Data.All[21]
	case "BH":
		return &Data.All[22]
	case "BI":
		return &Data.All[23]
	case "BJ":
		return &Data.All[24]
	case "BL":
		return &Data.All[25]
	case "BM":
		return &Data.All[26]
	case "BN":
		return &Data.All[27]
	case "BO":
		return &Data.All[28]
	case "BQ":
		return &Data.All[29]
	case "BR":
		return &Data.All[30]
	case "BS":
		return &Data.All[31]
	case "BT":
		return &Data.All[32]
	case "BV":
		return &Data.All[33]
	case "BW":
		return &Data.All[34]
	case "BY":
		return &Data.All[35]
	case "BZ":
		return &Data.All[36]
	case "CA":
		return &Data.All[37]
	case "CC":
		return &Data.All[38]
	case "CD":
		return &Data.All[39]
	case "CF":
		return &Data.All[40]
	case "CG":
		return &Data.All[41]
	case "CH":
		return &Data.All[42]
	case "CI":
		return &Data.All[43]
	case "CK":
		return &Data.All[44]
	case "CL":
		return &Data.All[45]
	case "CM":
		return &Data.All[46]
	case "CN":
		return &Data.All[47]
	case "CO":
		return &Data.All[48]
	case "CR":
		return &Data.All[49]
	case "CU":
		return &Data.All[50]
	case "CV":
		return &Data.All[51]
	case "CW":
		return &Data.All[52]
	case "CX":
		return &Data.All[53]
	case "CY":
		return &Data.All[54]
	case "CZ":
		return &Data.All[55]
	case "DE":
		return &Data.All[56]
	case "DJ":
		return &Data.All[57]
	case "DK":
		return &Data.All[58]
	case "DM":
		return &Data.All[59]
	case "DO":
		return &Data.All[60]
	case "DZ":
		return &Data.All[61]
	case "EC":
		return &Data.All[62]
	case "EE":
		return &Data.All[63]
	case "EG":
		return &Data.All[64]
	case "EH":
		return &Data.All[65]
	case "ER":
		return &Data.All[66]
	case "ES":
		return &Data.All[67]
	case "ET":
		return &Data.All[68]
	case "FI":
		return &Data.All[69]
	case "FJ":
		return &Data.All[70]
	case "FK":
		return &Data.All[71]
	case "FM":
		return &Data.All[72]
	case "FO":
		return &Data.All[73]
	case "FR":
		return &Data.All[74]
	case "GA":
		return &Data.All[75]
	case "GB":
		return &Data.All[76]
	case "GD":
		return &Data.All[77]
	case "GE":
		return &Data.All[78]
	case "GF":
		return &Data.All[79]
	case "GG":
		return &Data.All[80]
	case "GH":
		return &Data.All[81]
	case "GI":
		return &Data.All[82]
	case "GL":
		return &Data.All[83]
	case "GM":
		return &Data.All[84]
	case "GN":
		return &Data.All[85]
	case "GP":
		return &Data.All[86]
	case "GQ":
		return &Data.All[87]
	case "GR":
		return &Data.All[88]
	case "GS":
		return &Data.All[89]
	case "GT":
		return &Data.All[90]
	case "GU":
		return &Data.All[91]
	case "GW":
		return &Data.All[92]
	case "GY":
		return &Data.All[93]
	case "HK":
		return &Data.All[94]
	case "HM":
		return &Data.All[95]
	case "HN":
		return &Data.All[96]
	case "HR":
		return &Data.All[97]
	case "HT":
		return &Data.All[98]
	case "HU":
		return &Data.All[99]
	case "ID":
		return &Data.All[100]
	case "IE":
		return &Data.All[101]
	case "IL":
		return &Data.All[102]
	case "IM":
		return &Data.All[103]
	case "IN":
		return &Data.All[104]
	case "IO":
		return &Data.All[105]
	case "IQ":
		return &Data.All[106]
	case "IR":
		return &Data.All[107]
	case "IS":
		return &Data.All[108]
	case "IT":
		return &Data.All[109]
	case "JE":
		return &Data.All[110]
	case "JM":
		return &Data.All[111]
	case "JO":
		return &Data.All[112]
	case "JP":
		return &Data.All[113]
	case "KE":
		return &Data.All[114]
	case "KG":
		return &Data.All[115]
	case "KH":
		return &Data.All[116]
	case "KI":
		return &Data.All[117]
	case "KM":
		return &Data.All[118]
	case "KN":
		return &Data.All[119]
	case "KP":
		return &Data.All[120]
	case "KR":
		return &Data.All[121]
	case "KW":
		return &Data.All[122]
	case "KY":
		return &Data.All[123]
	case "KZ":
		return &Data.All[124]
	case "LA":
		return &Data.All[125]
	case "LB":
		return &Data.All[126]
	case "LC":
		return &Data.All[127]
	case "LI":
		return &Data.All[128]
	case "LK":
		return &Data.All[129]
	case "LR":
		return &Data.All[130]
	case "LS":
		return &Data.All[131]
	case "LT":
		return &Data.All[132]
	case "LU":
		return &Data.All[133]
	case "LV":
		return &Data.All[134]
	case "LY":
		return &Data.All[135]
	case "MA":
		return &Data.All[136]
	case "MC":
		return &Data.All[137]
	case "MD":
		return &Data.All[138]
	case "ME":
		return &Data.All[139]
	case "MF":
		return &Data.All[140]
	case "MG":
		return &Data.All[141]
	case "MH":
		return &Data.All[142]
	case "MK":
		return &Data.All[143]
	case "ML":
		return &Data.All[144]
	case "MM":
		return &Data.All[145]
	case "MN":
		return &Data.All[146]
	case "MO":
		return &Data.All[147]
	case "MP":
		return &Data.All[148]
	case "MQ":
		return &Data.All[149]
	case "MR":
		return &Data.All[150]
	case "MS":
		return &Data.All[151]
	case "MT":
		return &Data.All[152]
	case "MU":
		return &Data.All[153]
	case "MV":
		return &Data.All[154]
	case "MW":
		return &Data.All[155]
	case "MX":
		return &Data.All[156]
	case "MY":
		return &Data.All[157]
	case "MZ":
		return &Data.All[158]
	case "NA":
		return &Data.All[159]
	case "NC":
		return &Data.All[160]
	case "NE":
		return &Data.All[161]
	case "NF":
		return &Data.All[162]
	case "NG":
		return &Data.All[163]
	case "NI":
		return &Data.All[164]
	case "NL":
		return &Data.All[165]
	case "NO":
		return &Data.All[166]
	case "NP":
		return &Data.All[167]
	case "NR":
		return &Data.All[168]
	case "NU":
		return &Data.All[169]
	case "NZ":
		return &Data.All[170]
	case "OM":
		return &Data.All[171]
	case "PA":
		return &Data.All[172]
	case "PE":
		return &Data.All[173]
	case "PF":
		return &Data.All[174]
	case "PG":
		return &Data.All[175]
	case "PH":
		return &Data.All[176]
	case "PK":
		return &Data.All[177]
	case "PL":
		return &Data.All[178]
	case "PM":
		return &Data.All[179]
	case "PN":
		return &Data.All[180]
	case "PR":
		return &Data.All[181]
	case "PS":
		return &Data.All[182]
	case "PT":
		return &Data.All[183]
	case "PW":
		return &Data.All[184]
	case "PY":
		return &Data.All[185]
	case "QA":
		return &Data.All[186]
	case "RE":
		return &Data.All[187]
	case "RO":
		return &Data.All[188]
	case "RS":
		return &Data.All[189]
	case "RU":
		return &Data.All[190]
	case "RW":
		return &Data.All[191]
	case "SA":
		return &Data.All[192]
	case "SB":
		return &Data.All[193]
	case "SC":
		return &Data.All[194]
	case "SD":
		return &Data.All[195]
	case "SE":
		return &Data.All[196]
	case "SG":
		return &Data.All[197]
	case "SH":
		return &Data.All[198]
	case "SI":
		return &Data.All[199]
	case "SJ":
		return &Data.All[200]
	case "SK":
		return &Data.All[201]
	case "SL":
		return &Data.All[202]
	case "SM":
		return &Data.All[203]
	case "SN":
		return &Data.All[204]
	case "SO":
		return &Data.All[205]
	case "SR":
		return &Data.All[206]
	case "SS":
		return &Data.All[207]
	case "ST":
		return &Data.All[208]
	case "SV":
		return &Data.All[209]
	case "SX":
		return &Data.All[210]
	case "SY":
		return &Data.All[211]
	case "SZ":
		return &Data.All[212]
	case "TC":
		return &Data.All[213]
	case "TD":
		return &Data.All[214]
	case "TF":
		return &Data.All[215]
	case "TG":
		return &Data.All[216]
	case "TH":
		return &Data.All[217]
	case "TJ":
		return &Data.All[218]
	case "TK":
		return &Data.All[219]
	case "TL":
		return &Data.All[220]
	case "TM":
		return &Data.All[221]
	case "TN":
		return &Data.All[222]
	case "TO":
		return &Data.All[223]
	case "TR":
		return &Data.All[224]
	case "TT":
		return &Data.All[225]
	case "TV":
		return &Data.All[226]
	case "TW":
		return &Data.All[227]
	case "TZ":
		return &Data.All[228]
	case "UA":
		return &Data.All[229]
	case "UG":
		return &Data.All[230]
	case "UM":
		return &Data.All[231]
	case "US":
		return &Data.All[232]
	case "UY":
		return &Data.All[233]
	case "UZ":
		return &Data.All[234]
	case "VA":
		return &Data.All[235]
	case "VC":
		return &Data.All[236]
	case "VE":
		return &Data.All[237]
	case "VG":
		return &Data.All[238]
	case "VI":
		return &Data.All[239]
	case "VN":
		return &Data.All[240]
	case "VU":
		return &Data.All[241]
	case "WF":
		return &Data.All[242]
	case "WS":
		return &Data.All[243]
	case "YE":
		return &Data.All[244]
	case "YT":
		return &Data.All[245]
	case "ZA":
		return &Data.All[246]
	case "ZM":
		return &Data.All[247]
	case "ZW":
		return &Data.All[248]
	}
	return nil
}
