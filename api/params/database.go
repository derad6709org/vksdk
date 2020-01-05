package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// DatabaseGetChairsBuilder builder
//
// Returns list of chairs on a specified faculty.
//
// https://vk.com/dev/database.getChairs
type DatabaseGetChairsBuilder struct {
	api.Params
}

// NewDatabaseGetChairsBuilder func
func NewDatabaseGetChairsBuilder() *DatabaseGetChairsBuilder {
	return &DatabaseGetChairsBuilder{api.Params{}}
}

// FacultyID id of the faculty to get chairs from
func (b *DatabaseGetChairsBuilder) FacultyID(v int) {
	b.Params["faculty_id"] = v
}

// Offset offset required to get a certain subset of chairs
func (b *DatabaseGetChairsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count amount of chairs to get
func (b *DatabaseGetChairsBuilder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetCitiesBuilder builder
//
// Returns a list of cities.
//
// https://vk.com/dev/database.getCities
type DatabaseGetCitiesBuilder struct {
	api.Params
}

// NewDatabaseGetCitiesBuilder func
func NewDatabaseGetCitiesBuilder() *DatabaseGetCitiesBuilder {
	return &DatabaseGetCitiesBuilder{api.Params{}}
}

// CountryID Country ID.
func (b *DatabaseGetCitiesBuilder) CountryID(v int) {
	b.Params["country_id"] = v
}

// RegionID Region ID.
func (b *DatabaseGetCitiesBuilder) RegionID(v int) {
	b.Params["region_id"] = v
}

// Q Search query.
func (b *DatabaseGetCitiesBuilder) Q(v string) {
	b.Params["q"] = v
}

// NeedAll parameter
//
// * 1 — to return all cities in the country,
//
// * 0 — to return major cities in the country (default),
func (b *DatabaseGetCitiesBuilder) NeedAll(v bool) {
	b.Params["need_all"] = v
}

// Offset Offset needed to return a specific subset of cities.
func (b *DatabaseGetCitiesBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of cities to return.
func (b *DatabaseGetCitiesBuilder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetCitiesByIDBuilder builder
//
// Returns information about cities by their IDs.
//
// https://vk.com/dev/database.getCitiesById
type DatabaseGetCitiesByIDBuilder struct {
	api.Params
}

// NewDatabaseGetCitiesByIDBuilder func
func NewDatabaseGetCitiesByIDBuilder() *DatabaseGetCitiesByIDBuilder {
	return &DatabaseGetCitiesByIDBuilder{api.Params{}}
}

// CityIDs City IDs.
func (b *DatabaseGetCitiesByIDBuilder) CityIDs(v []int) {
	b.Params["city_ids"] = v
}

// DatabaseGetCountriesBuilder builder
//
// Returns a list of countries.
//
// https://vk.com/dev/database.getCountries
type DatabaseGetCountriesBuilder struct {
	api.Params
}

// NewDatabaseGetCountriesBuilder func
func NewDatabaseGetCountriesBuilder() *DatabaseGetCountriesBuilder {
	return &DatabaseGetCountriesBuilder{api.Params{}}
}

// NeedAll parameter
//
// * 1 — to return a full list of all countries,
//
// * 0 — to return a list of countries near the current user's country (default).
func (b *DatabaseGetCountriesBuilder) NeedAll(v bool) {
	b.Params["need_all"] = v
}

// Code Country codes in [vk.com/dev/country_codes|ISO 3166-1 alpha-2] standard.
func (b *DatabaseGetCountriesBuilder) Code(v string) {
	b.Params["code"] = v
}

// Offset Offset needed to return a specific subset of countries.
func (b *DatabaseGetCountriesBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of countries to return.
func (b *DatabaseGetCountriesBuilder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetCountriesByIDBuilder builder
//
// Returns information about countries by their IDs.
//
// https://vk.com/dev/database.getCountriesById
type DatabaseGetCountriesByIDBuilder struct {
	api.Params
}

// NewDatabaseGetCountriesByIDBuilder func
func NewDatabaseGetCountriesByIDBuilder() *DatabaseGetCountriesByIDBuilder {
	return &DatabaseGetCountriesByIDBuilder{api.Params{}}
}

// CountryIDs Country IDs.
func (b *DatabaseGetCountriesByIDBuilder) CountryIDs(v []int) {
	b.Params["country_ids"] = v
}

// DatabaseGetFacultiesBuilder builder
//
// Returns a list of faculties (i.e., university departments).
//
// https://vk.com/dev/database.getFaculties
type DatabaseGetFacultiesBuilder struct {
	api.Params
}

// NewDatabaseGetFacultiesBuilder func
func NewDatabaseGetFacultiesBuilder() *DatabaseGetFacultiesBuilder {
	return &DatabaseGetFacultiesBuilder{api.Params{}}
}

// UniversityID University ID.
func (b *DatabaseGetFacultiesBuilder) UniversityID(v int) {
	b.Params["university_id"] = v
}

// Offset Offset needed to return a specific subset of faculties.
func (b *DatabaseGetFacultiesBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of faculties to return.
func (b *DatabaseGetFacultiesBuilder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetMetroStationsBuilder builder
//
// Get metro stations by city
//
// https://vk.com/dev/database.getMetroStations
type DatabaseGetMetroStationsBuilder struct {
	api.Params
}

// NewDatabaseGetMetroStationsBuilder func
func NewDatabaseGetMetroStationsBuilder() *DatabaseGetMetroStationsBuilder {
	return &DatabaseGetMetroStationsBuilder{api.Params{}}
}

// CityID parameter
func (b *DatabaseGetMetroStationsBuilder) CityID(v int) {
	b.Params["city_id"] = v
}

// Offset parameter
func (b *DatabaseGetMetroStationsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count parameter
func (b *DatabaseGetMetroStationsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Extended parameter
func (b *DatabaseGetMetroStationsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// DatabaseGetMetroStationsByIDBuilder builder
//
// Get metro station by his id
//
// https://vk.com/dev/database.getMetroStationsById
type DatabaseGetMetroStationsByIDBuilder struct {
	api.Params
}

// NewDatabaseGetMetroStationsByIDBuilder func
func NewDatabaseGetMetroStationsByIDBuilder() *DatabaseGetMetroStationsByIDBuilder {
	return &DatabaseGetMetroStationsByIDBuilder{api.Params{}}
}

// StationIDs parameter
func (b *DatabaseGetMetroStationsByIDBuilder) StationIDs(v []int) {
	b.Params["station_ids"] = v
}

// DatabaseGetRegionsBuilder builder
//
// Returns a list of regions.
//
// https://vk.com/dev/database.getRegions
type DatabaseGetRegionsBuilder struct {
	api.Params
}

// NewDatabaseGetRegionsBuilder func
func NewDatabaseGetRegionsBuilder() *DatabaseGetRegionsBuilder {
	return &DatabaseGetRegionsBuilder{api.Params{}}
}

// CountryID Country ID, received in [vk.com/dev/database.getCountries|database.getCountries] method.
func (b *DatabaseGetRegionsBuilder) CountryID(v int) {
	b.Params["country_id"] = v
}

// Q Search query.
func (b *DatabaseGetRegionsBuilder) Q(v string) {
	b.Params["q"] = v
}

// Offset Offset needed to return specific subset of regions.
func (b *DatabaseGetRegionsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of regions to return.
func (b *DatabaseGetRegionsBuilder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetSchoolClassesBuilder builder
//
// Returns a list of school classes specified for the country.
//
// https://vk.com/dev/database.getSchoolClasses
type DatabaseGetSchoolClassesBuilder struct {
	api.Params
}

// NewDatabaseGetSchoolClassesBuilder func
func NewDatabaseGetSchoolClassesBuilder() *DatabaseGetSchoolClassesBuilder {
	return &DatabaseGetSchoolClassesBuilder{api.Params{}}
}

// CountryID Country ID.
func (b *DatabaseGetSchoolClassesBuilder) CountryID(v int) {
	b.Params["country_id"] = v
}

// DatabaseGetSchoolsBuilder builder
//
// Returns a list of schools.
//
// https://vk.com/dev/database.getSchools
type DatabaseGetSchoolsBuilder struct {
	api.Params
}

// NewDatabaseGetSchoolsBuilder func
func NewDatabaseGetSchoolsBuilder() *DatabaseGetSchoolsBuilder {
	return &DatabaseGetSchoolsBuilder{api.Params{}}
}

// Q Search query.
func (b *DatabaseGetSchoolsBuilder) Q(v string) {
	b.Params["q"] = v
}

// CityID City ID.
func (b *DatabaseGetSchoolsBuilder) CityID(v int) {
	b.Params["city_id"] = v
}

// Offset Offset needed to return a specific subset of schools.
func (b *DatabaseGetSchoolsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of schools to return.
func (b *DatabaseGetSchoolsBuilder) Count(v int) {
	b.Params["count"] = v
}

// DatabaseGetUniversitiesBuilder builder
//
// Returns a list of higher education institutions.
//
// https://vk.com/dev/database.getUniversities
type DatabaseGetUniversitiesBuilder struct {
	api.Params
}

// NewDatabaseGetUniversitiesBuilder func
func NewDatabaseGetUniversitiesBuilder() *DatabaseGetUniversitiesBuilder {
	return &DatabaseGetUniversitiesBuilder{api.Params{}}
}

// Q Search query.
func (b *DatabaseGetUniversitiesBuilder) Q(v string) {
	b.Params["q"] = v
}

// CountryID Country ID.
func (b *DatabaseGetUniversitiesBuilder) CountryID(v int) {
	b.Params["country_id"] = v
}

// CityID City ID.
func (b *DatabaseGetUniversitiesBuilder) CityID(v int) {
	b.Params["city_id"] = v
}

// Offset needed to return a specific subset of universities.
func (b *DatabaseGetUniversitiesBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of universities to return.
func (b *DatabaseGetUniversitiesBuilder) Count(v int) {
	b.Params["count"] = v
}
