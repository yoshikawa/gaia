package database

import (
	"time"

	"github.com/Pluslab/fieldsensing/server/domain"
)

// VendorRepository model
type VendorRepository struct {
	SQLHandler
}

// Store insert values into vendor table
func (repo *VendorRepository) Store(vendor domain.Vendor) (id int64, err error) {
	result, err := repo.Execute(
		`INSERT INTO vendors (
			name,
			organization_name,
			voice,
			facsimile,
			delivary_point,
			city,
			postal_code,
			country,
			electronic_mail_address
		) VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)`,
		vendor.Name,
		vendor.OrganizationName,
		vendor.Voice,
		vendor.Facsimile,
		vendor.DelivaryPoint,
		vendor.City,
		vendor.PostalCode,
		vendor.Country,
		vendor.ElectronicMailAddress,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int64(id64)
	return
}

// FindByID find the vendor by id
func (repo *VendorRepository) FindByID(identifier int64) (vendor domain.Vendor, err error) {
	row, err := repo.Query("SELECT * FROM vendors WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int64
	var name string
	var organizationName string
	var voice string
	var facsimile string
	var delivaryPoint string
	var city string
	var postalCode string
	var country string
	var electronicMailAddress string
	var createdAt time.Time
	var updatedAt time.Time
	row.Next()
	if err = row.Scan(&id, &name, &organizationName, &voice, &facsimile, &delivaryPoint, &city, &postalCode, &country, &electronicMailAddress, &createdAt, &updatedAt); err != nil {
		return
	}
	vendor.ID = id
	vendor.Name = name
	vendor.OrganizationName = organizationName
	vendor.Voice = voice
	vendor.Facsimile = facsimile
	vendor.DelivaryPoint = delivaryPoint
	vendor.City = city
	vendor.PostalCode = postalCode
	vendor.Country = country
	vendor.ElectronicMailAddress = electronicMailAddress
	vendor.CreatedAt = createdAt
	vendor.UpdatedAt = updatedAt
	return
}

// FindAll find all vendors
func (repo *VendorRepository) FindAll() (vendors domain.Vendors, err error) {
	rows, err := repo.Query("SELECT * FROM vendors")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var name string
		var organizationName string
		var voice string
		var facsimile string
		var delivaryPoint string
		var city string
		var postalCode string
		var country string
		var electronicMailAddress string
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&id, &name, &organizationName, &voice, &facsimile, &delivaryPoint, &city, &postalCode, &country, &electronicMailAddress, &createdAt, &updatedAt); err != nil {
			continue
		}
		vendor := domain.Vendor{
			ID:                    id,
			Name:                  name,
			OrganizationName:      organizationName,
			Voice:                 voice,
			Facsimile:             facsimile,
			DelivaryPoint:         delivaryPoint,
			City:                  city,
			PostalCode:            postalCode,
			Country:               country,
			ElectronicMailAddress: electronicMailAddress,
			CreatedAt:             createdAt,
			UpdatedAt:             updatedAt,
		}
		vendors = append(vendors, vendor)
	}
	return
}
