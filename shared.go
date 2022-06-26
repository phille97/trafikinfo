package trafikinfo

import (
	"fmt"
	"time"
)

type Geometry struct {
	SWEREF99TM *string `json:"SWEREF99TM,omitempty"`
	WGS84      *string `json:"WGS84,omitempty"`
}

type GeometryWithMod struct {
	Geometry
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
}

type Info struct {
	LastModified *struct {
		DateTime *time.Time `json:"_attr_datetime,omitempty"`
	} `json:"LASTMODIFIED,omitempty"`
	LastChangeID *string `json:"LASTCHANGEID,omitempty"`
	EvalResult   []any   `json:"EVALRESULT,omitempty"`
	SSEURL       *string `json:"SSEURL,omitempty"`
}

// County is a numerical ID assigned to a county in Sweden
type County uint

const (
	CountyStockholm County = iota + 1
	CountyStockholmDeprecated
	CountyUppsala
	CountySodermanland
	CountyOstergotland
	CountyJonkoping
	CountyKronobergs
	CountyKalmar
	CountyGotland
	CountyBlekinge
	_
	CountySkane
	CountyHallands
	CountyVastraGotaland
	_
	_
	CountyVarmland
	CountyOrebro
	CountyVastmanland
	CountyDalarna
	CountyGavleborg
	CountyVasternorrland
	CountyJamtland
	CountyVasterbotten
	CountyNorrbotten
)

func (c County) String() string {
	switch c {
	case CountyStockholm, CountyStockholmDeprecated:
		return "Stockholms län"
	case CountyUppsala:
		return "Uppsala län"
	case CountySodermanland:
		return "Södermanlands län"
	case CountyOstergotland:
		return "Östergötlands län"
	case CountyJonkoping:
		return "Jönköpings län"
	case CountyKronobergs:
		return "Kronobergs län"
	case CountyKalmar:
		return "Kalmar län"
	case CountyGotland:
		return "Gotlands län"
	case CountyBlekinge:
		return "Blekinge län"
	case CountySkane:
		return "Skåne län"
	case CountyHallands:
		return "Hallands län"
	case CountyVastraGotaland:
		return "Västra Götalands län"
	case CountyVarmland:
		return "Värmlands län"
	case CountyOrebro:
		return "Örebro län"
	case CountyVastmanland:
		return "Västmanlands län"
	case CountyDalarna:
		return "Dalarnas län"
	case CountyGavleborg:
		return "Gävleborgs län"
	case CountyVasternorrland:
		return "Västernorrlands län"
	case CountyJamtland:
		return "Jämtlands län"
	case CountyVasterbotten:
		return "Västerbottens län"
	case CountyNorrbotten:
		return "Norrbottens län"
	default:
		return fmt.Sprintf("Okänt län (%d)", c)
	}
}

type Country string

const (
	CountryGermany Country = "DE"
	CountryDenmark Country = "DK"
	CountryNorway  Country = "NO"
	CountrySweden  Country = "SE"
)

type Icon1Dot0 struct {
	Base64       *string    `json:"Base64,omitempty"`
	Deleted      *bool      `json:"Deleted,omitempty"`
	Description  *string    `json:"Description,omitempty"`
	ID           *string    `json:"Id,omitempty"`
	ModifiedTime *time.Time `json:"ModifiedTime,omitempty"`
	URL          *string    `json:"Url,omitempty"`
}

type Region uint

const (
	RegionNorth Region = iota + 1
	RegionMiddle
	RegionEast
	RegionStockholm
	RegionWest
	RegionSouth
)

func (r Region) String() string {
	switch r {
	case RegionNorth:
		return "Region Norr"
	case RegionMiddle:
		return "Region Mitt"
	case RegionEast:
		return "Region Öst"
	case RegionStockholm:
		return "Region Stockholm"
	case RegionWest:
		return "Region Väst"
	case RegionSouth:
		return "Region Syd"
	default:
		return fmt.Sprintf("Okänd region (%d)", r)
	}
}

type Vehicle string

const (
	VehicleAgricultural                       Vehicle = "agriculturalVehicle"
	VehicleAny                                Vehicle = "anyVehicle"
	VehicleArticulatedVehicle                 Vehicle = "articulatedVehicle"
	VehicleBicycle                            Vehicle = "bicycle"
	VehicleBus                                Vehicle = "bus"
	VehicleCar                                Vehicle = "car"
	VehicleACaravan                           Vehicle = "caravan"
	VehicleCarOrLightVehicle                  Vehicle = "carOrLightVehicle"
	VehicleCarWithCaravan                     Vehicle = "carWithCaravan"
	VehicleCarWithTrailer                     Vehicle = "carWithTrailer"
	VehicleConstructionOrMaintenanceVehicle   Vehicle = "constructionOrMaintenanceVehicle"
	VehicleFourWheelDrive                     Vehicle = "fourWheelDrive"
	VehicleHighSidedVehicle                   Vehicle = "highSidedVehicle"
	VehicleLorry                              Vehicle = "lorry"
	VehicleMoped                              Vehicle = "moped"
	VehicleMotorcycle                         Vehicle = "motorcycle"
	VehicleMotorcycleWithSideCar              Vehicle = "motorcycleWithSideCar"
	VehicleScooter                            Vehicle = "motorscooter"
	VehicleTanker                             Vehicle = "tanker"
	VehicleThreeWheeledVehicle                Vehicle = "threeWheeledVehicle"
	VehicleTrailer                            Vehicle = "trailer"
	VehicleTram                               Vehicle = "tram"
	VehicleTwoWheeledVehicle                  Vehicle = "twoWheeledVehicle"
	VehicleVan                                Vehicle = "van"
	VehicleWithCatalyticConverter             Vehicle = "vehicleWithCatalyticConverter"
	VehicleWithoutCatalyticConverter          Vehicle = "vehicleWithoutCatalyticConverter"
	VehicleWithCaravan                        Vehicle = "vehicleWithCaravan"
	VehicleWithTrailer                        Vehicle = "vehicleWithTrailer"
	VehicleWithEvenNumberedRegistrationPlates Vehicle = "withEvenNumberedRegistrationPlates"
	VehicleWithOddNumberedRegistrationPlates  Vehicle = "withOddNumberedRegistrationPlates"
	VehicleOther                              Vehicle = "other"
)

type Accessibility string

const (
	AccessibilityBarrierFreeAccessible           Accessibility = "barrierFreeAccessible"
	AccessibilityHandicappedAccessible           Accessibility = "handicappedAccessible"
	AccessibilityHandicappedEasements            Accessibility = "handicappedEasements"
	AccessibilityHandicappedMarked               Accessibility = "handicappedMarked"
	AccessibilityOrientationSystemForBlindPeople Accessibility = "orientationSystemForBlindPeople"
	AccessibilityWheelChairAccessible            Accessibility = "wheelChairAccessible"

	AccessibilityNone    Accessibility = "none"
	AccessibilityUnknown Accessibility = "unknown"
	AccessibilityOther   Accessibility = "other"
)

type Equipment string

const (
	EquipmentBikeParking             Equipment = "bikeParking"
	EquipmentCashMachine             Equipment = "cashMachine"
	EquipmentCopyMachineOrService    Equipment = "copyMachineOrService"
	EquipmentDefibrillator           Equipment = "defibrillator"
	EquipmentDumpingStation          Equipment = "dumpingStation"
	EquipmentElectricChargingStation Equipment = "electricChargingStation"
	EquipmentElevator                Equipment = "elevator"
	EquipmentFaxMachineOrService     Equipment = "faxMachineOrService"
	EquipmentFireExtingiusher        Equipment = "fireExtingiusher"
	EquipmentFireHose                Equipment = "fireHose"
	EquipmentFireHydrant             Equipment = "fireHydrant"
	EquipmentFirstAidEquipment       Equipment = "firstAidEquipment"
	EquipmentFreshWater              Equipment = "freshWater"
	EquipmentIceFreeScaffold         Equipment = "iceFreeScaffold"
	EquipmentInformationPoint        Equipment = "informationPoint"
	EquipmentInformatonStele         Equipment = "informatonStele"
	EquipmentInternetTerminal        Equipment = "internetTerminal"
	EquipmentInternetWireless        Equipment = "internetWireless"
	EquipmentLuggageLocker           Equipment = "luggageLocker"
	EquipmentPayDesk                 Equipment = "payDesk"
	EquipmentPaymentMachine          Equipment = "paymentMachine"
	EquipmentPicnicFacilities        Equipment = "picnicFacilities"
	EquipmentPlayground              Equipment = "playground"
	EquipmentPublicCardPhone         Equipment = "publicCardPhone"
	EquipmentPublicCoinPhone         Equipment = "publicCoinPhone"
	EquipmentPublicPhone             Equipment = "publicPhone"
	EquipmentRefuseBin               Equipment = "refuseBin"
	EquipmentSafeDeposit             Equipment = "safeDeposit"
	EquipmentShower                  Equipment = "shower"
	EquipmentToilet                  Equipment = "toilet"
	EquipmentTollTerminal            Equipment = "tollTerminal"
	EquipmentVendingMachine          Equipment = "vendingMachine"
	EquipmentWasteDisposal           Equipment = "wasteDisposal"

	EquipmentNone    Equipment = "none"
	EquipmentOther   Equipment = "other"
	EquipmentUnknown Equipment = "unknown"
)

type Facility string

const (
	FacilityHotel                   Facility = "hotel"
	FacilityMotel                   Facility = "motel"
	FacilityOvernightAccommodation  Facility = "overnightAccommodation"
	FacilityShop                    Facility = "shop"
	FacilityKiosk                   Facility = "kiosk"
	FacilityFoodShopping            Facility = "foodShopping"
	FacilityCafe                    Facility = "cafe"
	FacilityRestaurant              Facility = "restaurant"
	FacilityRestaurantSelfService   Facility = "restaurantSelfService"
	FacilityMotorwayRestaurant      Facility = "motorwayRestaurant"
	FacilityMotorwayRestaurantSmall Facility = "motorwayRestaurantSmall"
	FacilitySparePartsShopping      Facility = "sparePartsShopping"
	FacilityPetrolStation           Facility = "petrolStation"
	FacilityVehicleMaintenance      Facility = "vehicleMaintenance"
	FacilityTyreRepair              Facility = "tyreRepair"
	FacilityTruckRepair             Facility = "truckRepair"
	FacilityTruckWash               Facility = "truckWash"
	FacilityCarWash                 Facility = "carWash"
	FacilityPharmacy                Facility = "pharmacy"
	FacilityMedicalFacility         Facility = "medicalFacility"
	FacilityPolice                  Facility = "police"
	FacilityTouristInformation      Facility = "touristInformation"
	FacilityBikeSharing             Facility = "bikeSharing"
	FacilityDocstop                 Facility = "docstop"
	FacilityLaundry                 Facility = "laundry"
	FacilityLeisureActivities       Facility = "leisureActivities"

	FacilityUnknown Facility = "unknown"
	FacilityOther   Facility = "other"
)

type Cause string

const (
	CauseFog                 Cause = "Dimma"
	CauseFallingTemperatures Cause = "Fallande temperatur"
	CauseFrost               Cause = "Frost"
	CauseRain                Cause = "Regn"
	CauseMeltedWater         Cause = "Smältvatten"
	CauseSnowdrift           Cause = "Snödrev"
	CauseSnow                Cause = "Snöfall"
	CauseFreezingRain        Cause = "Underkylt regn"
)

type Condition uint

const (
	ConditionNormal Condition = iota + 1
	ConditionDifficult
	ConditionVeryDifficult
	ConditionIceAndSnow
)

func (c Condition) String() string {
	switch c {
	case ConditionNormal:
		return "normalt"
	case ConditionDifficult:
		return "besvärligt (risk för)"
	case ConditionVeryDifficult:
		return "mycket besvärligt"
	case ConditionIceAndSnow:
		return "is- och snövägbana"
	default:
		return fmt.Sprintf("Okänd väglagskod (%d)", c)
	}
}

type Measurement string

const (
	MeasurementPreventiveAntiSlip Measurement = "Förebyggande halkbekämpning"
	MeasurementAntiSlip           Measurement = "Halkbekämpning"
	MeasurementIcePlowing         Measurement = "Ishyvling"
	MeasurementSnowPlowing        Measurement = "Plogning"
	MeasurementSanding            Measurement = "Sandning"
	MeasurementOther              Measurement = "Annat"
)

type Warning string

const (
	WarningSlippingRisk Warning = "Risk för halka"
	WarningSlipping     Warning = "Halka"
	WarningSnowSmoke    Warning = "Snörök"
	WarningSnowdrift    Warning = "Snödrev"
	WarningHarshWind    Warning = "Hård vind"
	WarningSnow         Warning = "Snöfall"
	WarningOther        Warning = "Annat"
)

type Status uint

const (
	StatusFreeAccess Status = iota + 1
	StatusPassabilityDifficult
	StatusPassabilityImpossile
	StatusPassabilityUnknown
)

func (s Status) String() string {
	switch s {
	case StatusFreeAccess:
		return "fri framkomlighet"
	case StatusPassabilityDifficult:
		return "svår framkomlighet"
	case StatusPassabilityImpossile:
		return "framkomlighet omöjlig"
	case StatusPassabilityUnknown:
		return "framkomlighet okänd"
	default:
		return fmt.Sprintf("Okänd status (%d)", s)
	}
}

type TrafficStatus string

const (
	TrafficStatusFreeflow   TrafficStatus = "freeflow"
	TrafficStatusHeavy      TrafficStatus = "heavy"
	TrafficStatusCongested  TrafficStatus = "congested"
	TrafficStatusImpossible TrafficStatus = "impossible"
	TrafficStatusUnknown    TrafficStatus = "unknown"
)
