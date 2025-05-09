package models

import (
	"reflect"
	"testing"
)

func TestLessonPlayerToPlayer(t *testing.T) {
	lessonPlayer := LessonPlayer{
		UserID:                "user-123",
		PaymentID:             "payment-456",
		RegistrationPrice:     "30.00",
		PaymentMethodType:     "CREDIT_CARD",
		FullName:              "John Doe",
		LevelValue:            3.5,
		Picture:               "profile.jpg",
		PaidAtMerchant:        true,
		PaymentB2bBillingType: "INVOICE",
	}

	player := LessonPlayerToPlayer(&lessonPlayer)

	if player.UserID != lessonPlayer.UserID {
		t.Errorf("UserID not correctly mapped. Expected %s, got %s", lessonPlayer.UserID, player.UserID)
	}

	if player.Name != lessonPlayer.FullName {
		t.Errorf("Name not correctly mapped. Expected %s, got %s", lessonPlayer.FullName, player.Name)
	}

	if player.LevelValue != lessonPlayer.LevelValue {
		t.Errorf("LevelValue not correctly mapped. Expected %f, got %f", lessonPlayer.LevelValue, player.LevelValue)
	}

	if player.Picture != lessonPlayer.Picture {
		t.Errorf("Picture not correctly mapped. Expected %s, got %s", lessonPlayer.Picture, player.Picture)
	}
}

func TestLessonTenantToTenant(t *testing.T) {
	address := Address{
		Street:                "123 Main St",
		PostalCode:            "12345",
		City:                  "Cityville",
		SubAdministrativeArea: "County",
		AdministrativeArea:    "State",
		Country:               "Country",
		CountryCode:           "CC",
		Coordinate: Coordinate{
			Lat: 12.345,
			Lon: 67.890,
		},
		Timezone: "UTC+1",
	}

	lessonTenant := LessonTenant{
		TenantID:      "tenant-123",
		TenantName:    "Test Club",
		TenantAddress: address,
		TenantImages:  []string{"image1.jpg", "image2.jpg"},
		Properties:    map[string]interface{}{"amenity": "courts"},
	}

	tenant := LessonTenantToTenant(&lessonTenant)

	if tenant.TenantID != lessonTenant.TenantID {
		t.Errorf("TenantID not correctly mapped. Expected %s, got %s", lessonTenant.TenantID, tenant.TenantID)
	}

	if tenant.TenantName != lessonTenant.TenantName {
		t.Errorf("TenantName not correctly mapped. Expected %s, got %s", lessonTenant.TenantName, tenant.TenantName)
	}

	if !reflect.DeepEqual(tenant.Address, lessonTenant.TenantAddress) {
		t.Errorf("Address not correctly mapped. Expected %+v, got %+v", lessonTenant.TenantAddress, tenant.Address)
	}

	if !reflect.DeepEqual(tenant.Images, lessonTenant.TenantImages) {
		t.Errorf("Images not correctly mapped. Expected %v, got %v", lessonTenant.TenantImages, tenant.Images)
	}

	if !reflect.DeepEqual(tenant.Properties, lessonTenant.Properties) {
		t.Errorf("Properties not correctly mapped. Expected %v, got %v", lessonTenant.Properties, tenant.Properties)
	}
}
