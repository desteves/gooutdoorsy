package rental_test

import (
	"github.com/desteves/gooutdoorsy/database"
	. "github.com/desteves/gooutdoorsy/rental"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OutdoorsyRV\n", func() {
	var outdoorsy OutdoorsyRV
	Context("When Postgres is running\n", func() {
		var connection string

		BeforeEach(func() {
			connection = "user=root password=root dbname=testingwithrentals sslmode=disable connect_timeout=30 port=5434"
		})
		It("should create a new outdoorsy", func() {
			o, err := NewOutdoorsyProvider(connection)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(o.DB).ShouldNot(BeNil())
		})

		Context("with an invalid connection string\n", func() {
			BeforeEach(func() {
				connection = "mongo"
			})
			It("should fail to create a new outdoorsy", func() {
				_, err := NewOutdoorsyProvider(connection)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
	Context("When getting a single rental", func() {
		var id int
		BeforeEach(func() {
			id = 2
			outdoorsy = OutdoorsyRV{
				DB: &database.Mock{
					RentalData:  []database.RentalData{{ID: id}, {ID: id}},
					ReturnError: false,
				},
			}
		})
		It("can find the rental", func() {
			rental, err := outdoorsy.GetRental(id)

			Expect(rental.ID).Should(Equal(id))
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("can find rentals", func() {
			rentals, err := outdoorsy.GetRentals(database.Parameters{})

			Expect(id).Should(Equal(len(rentals)))
			Expect(err).ShouldNot(HaveOccurred())
		})

	})
})
