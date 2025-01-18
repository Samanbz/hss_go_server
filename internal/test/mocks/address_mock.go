package mocks

import "hss/internal/models"

var (
	MockAddress = models.Address{
		Username:  "user1",
		Password:  "501613c686d94b3f1faf0157ea5242f6e313cbc4520efdf9cd6b9f4917057280",
		CompanyID: 101,
		Street:    "123 Elm St",
		City:      "Metropolis",
		State:     "NY",
		Zip:       "10001",
		Country:   "USA",
		Latitude:  40.7128,
		Longitude: -74.0060,
	}

	MockAddress2 = models.Address{
		Username:  "user2",
		Password:  "d11a7fbe8f34311aeb11dfce4ac59679b90973fc44fbab2e11d4f902b8220e99",
		CompanyID: 102,
		Street:    "456 Oak St",
		City:      "Gotham",
		State:     "NJ",
		Zip:       "07030",
		Country:   "USA",
		Latitude:  40.7437,
		Longitude: -74.0324,
	}

	MockAddress3 = models.Address{
		Username:  "user3",
		Password:  "a4123be42aa90d448d81b6e327b64d8c6b89523e94f154743a85838600bf33ee",
		CompanyID: 103,
		Street:    "789 Maple St",
		City:      "Star City",
		State:     "CA",
		Zip:       "90001",
		Country:   "USA",
		Latitude:  34.0522,
		Longitude: -118.2437,
	}

	MockAddress4 = models.Address{
		Username:  "user4",
		Password:  "f4dd11f951f746ee10898fd5482cecadd016f90419ec94c48bb064128ed394a3",
		CompanyID: 104,
		Street:    "321 Pine St",
		City:      "Central City",
		State:     "CO",
		Zip:       "80422",
		Country:   "USA",
		Latitude:  39.7392,
		Longitude: -104.9903,
	}
)
