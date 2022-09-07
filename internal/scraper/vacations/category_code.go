package vacations

type CategoryCodes struct {
	CustomerService          string
	ProductionEngineering    string
	Sales                    string
	Retail                   string
	HotelRestaurantTourism   string
	Administration           string
	IT                       string
	LogisticSupplyChain      string
	AutoTransport            string
	Healthcare               string
	EducationScientific      string
	Accounting               string
	MarketingAdvertisingPR   string
	OfficeSecretarial        string
	Telecommunications       string
	ConstructionArchitecture string
	BeautySports             string
	BankingFinance           string
	DesignArt                string
	PublishingMedia          string
	HRRecruitment            string
	Agriculture              string
	ManagementExecutive      string
	Security                 string
	Legal                    string
	RealEstate               string
	CultureMusicShowbiz      string
	Insurance                string
}

func NewCategoryCodes() *CategoryCodes {
	return &CategoryCodes{
		CustomerService:          "20",
		ProductionEngineering:    "14",
		Sales:                    "22",
		Retail:                   "23",
		HotelRestaurantTourism:   "4",
		Administration:           "2",
		IT:                       "1",
		LogisticSupplyChain:      "8",
		AutoTransport:            "24",
		Healthcare:               "10",
		EducationScientific:      "12",
		Accounting:               "3",
		MarketingAdvertisingPR:   "9",
		OfficeSecretarial:        "15",
		Telecommunications:       "6792",
		ConstructionArchitecture: "19",
		BeautySports:             "6",
		BankingFinance:           "26",
		DesignArt:                "5",
		PublishingMedia:          "17",
		HRRecruitment:            "25",
		Agriculture:              "30",
		ManagementExecutive:      "21",
		Security:                 "13",
		Legal:                    "27",
		RealEstate:               "11",
		CultureMusicShowbiz:      "7",
		Insurance:                "18",
	}
}
