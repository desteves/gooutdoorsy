package database_test

import (
	. "github.com/desteves/gooutdoorsy/database"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Postgres\n", func() {
	var db Postgres
	var conn string
	var err error
	Context("When the connection string is not set\n", func() {
		It("can generate an error\n", func() {
			err = db.Open("")
			Expect(err).Should(HaveOccurred())

		})
	})
	Context("When the connection string is invalid\n", func() {
		It("can generate an error\n", func() {
			err = db.Open("mongo")
			Expect(err).Should(HaveOccurred())
		})
	})
	Context("When the connection string is valid\n", func() {
		BeforeEach(func() {
			conn = "user=root password=root dbname=testingwithrentals sslmode=disable connect_timeout=30 port=5434"

		})

		Context("When the postgres daemon is running\n", func() {

			BeforeEach(func() {
				err = db.Open(conn)
			})
			It("can connect\n", func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(db.Client).ShouldNot(BeNil())
			})

			Context("When postgres' schema is defined\n", func() {
				Context("When rows are available\n", func() {
					It("can read a row\n", func() {
						_, err := db.ReadOne(2)
						Expect(err).ShouldNot(HaveOccurred())
					})

					Context("When parameters are present\n", func() {
						var p Parameters
						BeforeEach(func() {
							p = Parameters{}
						})
						Context("When parameters contain a sort", func() {
							BeforeEach(func() {
								p.Sort = "name"
							})
							It("can return sorted results\n", func() {
								_, err := db.ReadMany(p)
								Expect(err).ShouldNot(HaveOccurred())
							})
						})
						Context("When parameters contain a max price", func() {
							BeforeEach(func() {
								p.PriceMax = 3000
							})
							It("can return filtered results based on max price\n", func() {
								_, err := db.ReadMany(p)
								Expect(err).ShouldNot(HaveOccurred())
							})
						})
						Context("When parameters contain a min price", func() {
							BeforeEach(func() {
								p.PriceMin = 3000
							})
							It("can return filtered results based on min price\n", func() {
								_, err := db.ReadMany(p)
								Expect(err).ShouldNot(HaveOccurred())
							})
						})
						Context("When parameters contain a limit", func() {
							BeforeEach(func() {
								p.Limit = 2
							})
							It("can return limited results\n", func() {
								_, err := db.ReadMany(p)
								Expect(err).ShouldNot(HaveOccurred())
							})
						})
						Context("When parameters contain an offset", func() {
							BeforeEach(func() {
								p.Offset = 10
							})
							It("can return limited results\n", func() {
								_, err := db.ReadMany(p)
								Expect(err).ShouldNot(HaveOccurred())
							})
						})
						Context("When parameters contain Ids", func() {
							BeforeEach(func() {
								p.IDs = []int{2, 3}
								p.RawIDs = "2,3"

							})
							It("can return correct results\n", func() {
								_, err := db.ReadMany(p)
								Expect(err).ShouldNot(HaveOccurred())
							})
						})
						Context("When parameters contain valid Geo coordinates", func() {
							BeforeEach(func() {
								p.Near = []float64{33.64, -117.93}
								p.RawNear = "33.64,-117.93"

							})
							It("can return correct results\n", func() {
								_, err := db.ReadMany(p)
								Expect(err).ShouldNot(HaveOccurred())
							})
						})
						// price_min=9000&price_max=75000&limit=3&offset=6&sort=price
						Context("When multiple parameters are set", func() {
							BeforeEach(func() {
								p.Near = []float64{33.64, -117.93}
								p.RawNear = "33.64,-117.93"
								p.PriceMin = 9000
								p.PriceMax = 75000
								p.Limit = 3
								p.Offset = 6
								p.Sort = "price"

							})
							It("can return correct results\n", func() {
								_, err := db.ReadMany(p)
								Expect(err).ShouldNot(HaveOccurred())
							})
						})
					})
				})
			})
		})
	})

})
