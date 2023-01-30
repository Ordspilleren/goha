package nordpool

type NordpoolResponse struct {
	Data     Data   `json:"data,omitempty"`
	CacheKey string `json:"cacheKey,omitempty"`
	Conf     Conf   `json:"conf,omitempty"`
	Header   Header `json:"header,omitempty"`
	EndDate  string `json:"endDate,omitempty"`
	Currency string `json:"currency,omitempty"`
	PageID   int    `json:"pageId,omitempty"`
}
type Columns struct {
	Index                            int         `json:"Index,omitempty"`
	Scale                            int         `json:"Scale,omitempty"`
	SecondaryValue                   interface{} `json:"SecondaryValue,omitempty"`
	IsDominatingDirection            bool        `json:"IsDominatingDirection,omitempty"`
	IsValid                          bool        `json:"IsValid,omitempty"`
	IsAdditionalData                 bool        `json:"IsAdditionalData,omitempty"`
	Behavior                         int         `json:"Behavior,omitempty"`
	Name                             string      `json:"Name,omitempty"`
	Value                            string      `json:"Value,omitempty"`
	GroupHeader                      string      `json:"GroupHeader,omitempty"`
	DisplayNegativeValueInBlue       bool        `json:"DisplayNegativeValueInBlue,omitempty"`
	CombinedName                     string      `json:"CombinedName,omitempty"`
	DateTimeForData                  string      `json:"DateTimeForData,omitempty"`
	DisplayName                      string      `json:"DisplayName,omitempty"`
	DisplayNameOrDominatingDirection string      `json:"DisplayNameOrDominatingDirection,omitempty"`
	IsOfficial                       bool        `json:"IsOfficial,omitempty"`
	UseDashDisplayStyle              bool        `json:"UseDashDisplayStyle,omitempty"`
}
type Rows struct {
	Columns         []Columns   `json:"Columns,omitempty"`
	Name            string      `json:"Name,omitempty"`
	StartTime       string      `json:"StartTime,omitempty"`
	EndTime         string      `json:"EndTime,omitempty"`
	DateTimeForData string      `json:"DateTimeForData,omitempty"`
	DayNumber       int         `json:"DayNumber,omitempty"`
	StartTimeDate   string      `json:"StartTimeDate,omitempty"`
	IsExtraRow      bool        `json:"IsExtraRow,omitempty"`
	IsNtcRow        bool        `json:"IsNtcRow,omitempty"`
	EmptyValue      string      `json:"EmptyValue,omitempty"`
	Parent          interface{} `json:"Parent,omitempty"`
}
type Data struct {
	Rows                      []Rows        `json:"Rows,omitempty"`
	IsDivided                 bool          `json:"IsDivided,omitempty"`
	SectionNames              []string      `json:"SectionNames,omitempty"`
	EntityIDs                 []string      `json:"EntityIDs,omitempty"`
	DataStartdate             string        `json:"DataStartdate,omitempty"`
	DataEnddate               string        `json:"DataEnddate,omitempty"`
	MinDateForTimeScale       string        `json:"MinDateForTimeScale,omitempty"`
	AreaChanges               []interface{} `json:"AreaChanges,omitempty"`
	Units                     []string      `json:"Units,omitempty"`
	LatestResultDate          string        `json:"LatestResultDate,omitempty"`
	ContainsPreliminaryValues bool          `json:"ContainsPreliminaryValues,omitempty"`
	ContainsExchangeRates     bool          `json:"ContainsExchangeRates,omitempty"`
	ExchangeRateOfficial      interface{}   `json:"ExchangeRateOfficial,omitempty"`
	ExchangeRatePreliminary   string        `json:"ExchangeRatePreliminary,omitempty"`
	ExchangeUnit              string        `json:"ExchangeUnit,omitempty"`
	DateUpdated               string        `json:"DateUpdated,omitempty"`
	CombinedHeadersEnabled    bool          `json:"CombinedHeadersEnabled,omitempty"`
	DataType                  int           `json:"DataType,omitempty"`
	TimeZoneInformation       int           `json:"TimeZoneInformation,omitempty"`
}
type ResolutionPeriod struct {
	ID           string `json:"Id,omitempty"`
	Resolution   int    `json:"Resolution,omitempty"`
	Unit         int    `json:"Unit,omitempty"`
	PeriodNumber int    `json:"PeriodNumber,omitempty"`
}
type ResolutionPeriodY struct {
	ID           string `json:"Id,omitempty"`
	Resolution   int    `json:"Resolution,omitempty"`
	Unit         int    `json:"Unit,omitempty"`
	PeriodNumber int    `json:"PeriodNumber,omitempty"`
}
type Attributes struct {
	ID       string   `json:"Id,omitempty"`
	Name     string   `json:"Name,omitempty"`
	Role     string   `json:"Role,omitempty"`
	HasRoles bool     `json:"HasRoles,omitempty"`
	Value    string   `json:"Value,omitempty"`
	Values   []string `json:"Values,omitempty"`
}
type ProductType struct {
	ID          string       `json:"Id,omitempty"`
	Attributes  []Attributes `json:"Attributes,omitempty"`
	Name        string       `json:"Name,omitempty"`
	DisplayName string       `json:"DisplayName,omitempty"`
}
type SecondaryProductType struct {
	ID          string      `json:"Id,omitempty"`
	Attributes  interface{} `json:"Attributes,omitempty"`
	Name        string      `json:"Name,omitempty"`
	DisplayName string      `json:"DisplayName,omitempty"`
}
type Entities struct {
	ProductType                 ProductType          `json:"ProductType,omitempty"`
	SecondaryProductType        SecondaryProductType `json:"SecondaryProductType,omitempty"`
	SecondaryProductBehavior    int                  `json:"SecondaryProductBehavior,omitempty"`
	ID                          string               `json:"Id,omitempty"`
	Name                        string               `json:"Name,omitempty"`
	GroupHeader                 string               `json:"GroupHeader,omitempty"`
	DataUpdated                 string               `json:"DataUpdated,omitempty"`
	Attributes                  []Attributes         `json:"Attributes,omitempty"`
	Drillable                   bool                 `json:"Drillable,omitempty"`
	DateRanges                  []interface{}        `json:"DateRanges,omitempty"`
	Index                       int                  `json:"Index,omitempty"`
	IndexForColumn              int                  `json:"IndexForColumn,omitempty"`
	MinMaxDisabled              bool                 `json:"MinMaxDisabled,omitempty"`
	DisableNumberGroupSeparator int                  `json:"DisableNumberGroupSeparator,omitempty"`
	TimeserieID                 interface{}          `json:"TimeserieID,omitempty"`
	SecondaryTimeserieID        string               `json:"SecondaryTimeserieID,omitempty"`
	HasPreliminary              bool                 `json:"HasPreliminary,omitempty"`
	TimeseriePreliminaryID      interface{}          `json:"TimeseriePreliminaryID,omitempty"`
	Scale                       int                  `json:"Scale,omitempty"`
	SecondaryScale              int                  `json:"SecondaryScale,omitempty"`
	DataType                    int                  `json:"DataType,omitempty"`
	SecondaryDataType           int                  `json:"SecondaryDataType,omitempty"`
	LastUpdate                  string               `json:"LastUpdate,omitempty"`
	Unit                        string               `json:"Unit,omitempty"`
	IsDominatingDirection       bool                 `json:"IsDominatingDirection,omitempty"`
	DisplayAsSeparatedColumn    bool                 `json:"DisplayAsSeparatedColumn,omitempty"`
	EnableInChart               bool                 `json:"EnableInChart,omitempty"`
	BlueNegativeValues          bool                 `json:"BlueNegativeValues,omitempty"`
}
type ExtraRows struct {
	ID             string   `json:"Id,omitempty"`
	Header         string   `json:"Header,omitempty"`
	ColumnProducts []string `json:"ColumnProducts,omitempty"`
}
type Filters struct {
	ID            string   `json:"Id,omitempty"`
	AttributeName string   `json:"AttributeName,omitempty"`
	Values        []string `json:"Values,omitempty"`
	DefaultValue  string   `json:"DefaultValue,omitempty"`
}
type NtcProductType struct {
	ID          string      `json:"Id,omitempty"`
	Attributes  interface{} `json:"Attributes,omitempty"`
	Name        string      `json:"Name,omitempty"`
	DisplayName string      `json:"DisplayName,omitempty"`
}
type Conf struct {
	ID                       string            `json:"Id,omitempty"`
	Name                     interface{}       `json:"Name,omitempty"`
	Published                string            `json:"Published,omitempty"`
	ShowGraph                bool              `json:"ShowGraph,omitempty"`
	ResolutionPeriod         ResolutionPeriod  `json:"ResolutionPeriod,omitempty"`
	ResolutionPeriodY        ResolutionPeriodY `json:"ResolutionPeriodY,omitempty"`
	Entities                 []Entities        `json:"Entities,omitempty"`
	TableType                int               `json:"TableType,omitempty"`
	ExtraRows                []ExtraRows       `json:"ExtraRows,omitempty"`
	Filters                  []Filters         `json:"Filters,omitempty"`
	IsDrillDownEnabled       bool              `json:"IsDrillDownEnabled,omitempty"`
	DrillDownMode            int               `json:"DrillDownMode,omitempty"`
	IsMinValueEnabled        bool              `json:"IsMinValueEnabled,omitempty"`
	IsMaxValueEnabled        bool              `json:"IsMaxValueEnabled,omitempty"`
	ValidYearsBack           int               `json:"ValidYearsBack,omitempty"`
	TimeScaleUnit            string            `json:"TimeScaleUnit,omitempty"`
	IsNtcEnabled             bool              `json:"IsNtcEnabled,omitempty"`
	NtcProductType           NtcProductType    `json:"NtcProductType,omitempty"`
	NtcHeader                string            `json:"NtcHeader,omitempty"`
	ShowTimelineGraph        int               `json:"ShowTimelineGraph,omitempty"`
	ExchangeMode             int               `json:"ExchangeMode,omitempty"`
	IsPivotTable             int               `json:"IsPivotTable,omitempty"`
	IsCombinedHeadersEnabled int               `json:"IsCombinedHeadersEnabled,omitempty"`
	NtcFormat                int               `json:"NtcFormat,omitempty"`
	DisplayHourAlsoInUKTime  bool              `json:"DisplayHourAlsoInUKTime,omitempty"`
}
type Header struct {
	Title              string `json:"title,omitempty"`
	Description        string `json:"description,omitempty"`
	QuestionMarkInfo   string `json:"questionMarkInfo,omitempty"`
	HideDownloadButton string `json:"hideDownloadButton,omitempty"`
}