package api_test

import (
	. "github.com/desteves/gooutdoorsy/api"
	"github.com/desteves/gooutdoorsy/database"
	"github.com/desteves/gooutdoorsy/rental"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api\n", func() {

	Context("When the GET rental by ID route is set\n", func() {
		var mockAPI API
		var r *gin.Engine
		var w *httptest.ResponseRecorder
		BeforeEach(func() {
			mockAPI = API{
				RVRentalProvider: &rental.MockRV{},
			}
			w = httptest.NewRecorder()
			_, r = gin.CreateTestContext(w)
			r.GET(RentalsEndpointByID, mockAPI.GetRVRentalByIDHandler)
		})
		It("cannot find a string-based rental id\n", func() {
			req, _ := http.NewRequest("GET", "/rentals/two", nil)
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusBadRequest))
			// var actual map[string]interface{}
			// json.Unmarshal([]byte(w.Body.Bytes()), &actual)
			// Expect(actual[ErrorResponseField]).Should(Equal(ErrorMsgStatusBadRequest))
		})
		It("cannot find non-existing rental\n", func() {
			req, _ := http.NewRequest("GET", "/rentals/2", nil)
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusNotFound))
		})
		Context("when a rental is stored\n", func() {
			BeforeEach(func() {
				editableRental, _ := mockAPI.RVRentalProvider.(*rental.MockRV)
				editableRental.RentalData = []database.RentalData{{ID: 2}}
				mockAPI.RVRentalProvider = editableRental
			})
			It("can gracefully handle the backend being unavailable\n", func() {
				// TODO --
			})
			It("can find an existing rental\n", func() {
				req, _ := http.NewRequest("GET", "/rentals/2", nil)
				r.ServeHTTP(w, req)
				Expect(w.Code).Should(Equal(http.StatusOK))
				var actual database.RentalData
				err := json.Unmarshal([]byte(w.Body.Bytes()), &actual)
				Expect(err).NotTo(HaveOccurred())
				Expect(actual.ID).Should(Equal(2))
			})
		})

	})

	Context("When the GET rentals route is set\n", func() {
		var mockAPI API
		var r *gin.Engine
		var w *httptest.ResponseRecorder
		BeforeEach(func() {
			mockAPI = API{
				RVRentalProvider: &rental.MockRV{},
			}
			w = httptest.NewRecorder()
			_, r = gin.CreateTestContext(w)
			r.GET(RentalsEndpoint, mockAPI.GetRVRentalsHandler)
		})
		It("cannot find a string-based rental id\n", func() {
			req, _ := http.NewRequest("GET", "/rentals?ids=two", nil)
			req.ParseForm()
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusBadRequest))

		})

		It("can send bad req responses with bad ids\n", func() {
			req, _ := http.NewRequest("GET", "/rentals?ids=two", nil)
			req.ParseForm()
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusBadRequest))

		})
		It("can send bad req responses with an illegal near filter\n", func() {

			req, _ := http.NewRequest("GET", "/rentals?near=33.64,-117.93,20", nil)
			req.ParseForm()
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusBadRequest))

			req, _ = http.NewRequest("GET", "/rentals?near=A,B", nil)
			req.ParseForm()
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusBadRequest))

		})
		It("cannot find non-existing rental\n", func() {
			req, _ := http.NewRequest("GET", "/rentals?ids=2", nil)
			req.ParseForm()
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusNotFound))
		})

		It("can find an existing rental via rentals\n", func() {

			editableRental, _ := mockAPI.RVRentalProvider.(*rental.MockRV)
			editableRental.RentalData = []database.RentalData{{ID: 2}}
			mockAPI.RVRentalProvider = editableRental

			req, _ := http.NewRequest("GET", "/rentals?ids=2", nil)
			req.ParseForm()
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusOK))

		})

		It("can find existing rentals\n", func() {
			editableRental, _ := mockAPI.RVRentalProvider.(*rental.MockRV)
			editableRental.RentalData = []database.RentalData{{ID: 2}, {ID: 1}, {ID: 3}}
			mockAPI.RVRentalProvider = editableRental

			req, _ := http.NewRequest("GET", "/rentals?ids=2,3", nil)
			req.ParseForm()
			r.ServeHTTP(w, req)

			Expect(w.Code).Should(Equal(http.StatusOK))
			// var actual database.RentalData
			// err := json.Unmarshal([]byte(w.Body.Bytes()), &actual)
			// Expect(err).NotTo(HaveOccurred())
			// Expect(actual.ID).Should(Equal(2))
		})

		// It("can find rentals based on the near filter\n", func() {
		// 	editableRental, _ := mockAPI.RVRentalProvider.(*rental.MockRV)
		// 	editableRental.RentalData = []database.RentalData{}
		// 	mockAPI.RVRentalProvider = editableRental

		// 	req, _ := http.NewRequest("GET", "/rentals?near=33.64,-117.93", nil)
		// 	req.ParseForm()
		// 	r.ServeHTTP(w, req)
		// 	// Expect(req).Should(Equal(http.StatusOK))
		// 	Expect(w.Code).Should(Equal(http.StatusOK))
		// })

	})

})
