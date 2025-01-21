package handlers

import (
	"fmt"
	"net/http"

	"github.com/CloudyKit/jet"

	"github.com/mo0Oonnn/bookings/internal/config"
	"github.com/mo0Oonnn/bookings/internal/forms"
	"github.com/mo0Oonnn/bookings/internal/helpers"
	"github.com/mo0Oonnn/bookings/internal/models"
	"github.com/mo0Oonnn/bookings/internal/render"
)

type Repository struct {
	AppConfig *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: a,
	}
}

func SetRepo(r *Repository) {
	Repo = r
}

func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.jet", jet.VarMap{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.jet", jet.VarMap{})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(jet.VarMap)
	data.Set("reservation", emptyReservation)
	data.Set("form", forms.New(nil))

	render.RenderTemplate(w, r, "make-reservation.jet", data)
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		DateFrom:  r.Form.Get("date_from"),
		DateTo:    r.Form.Get("date_to"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
		Room:      r.Form.Get("room"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "date_from", "date_to", "email")
	form.CheckLength("first_name", 3)
	form.CheckLength("last_name", 3)
	form.IsEmail("email")

	if !form.IsValid() {
		data := make(jet.VarMap)
		data.Set("reservation", reservation)
		data.Set("form", form)

		render.RenderTemplate(w, r, "make-reservation.jet", data)
		return
	}

	m.AppConfig.SessionManager.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// SingleRoom renders the room page
func (m *Repository) SingleRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "single-room.jet", jet.VarMap{})
}

// TwoBedRoom renders the room page
func (m *Repository) TwoBedRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "two-bed-room.jet", jet.VarMap{})
}

// DoubleBedRoom renders the room page
func (m *Repository) DoubleBedRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "double-bed-room.jet", jet.VarMap{})
}

// FamilyRoom renders the room page
func (m *Repository) FamilyRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "family-room.jet", jet.VarMap{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.jet", jet.VarMap{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Дата въезда: %s, дата выезда: %s", start, end)))
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.jet", jet.VarMap{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.AppConfig.SessionManager.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.AppConfig.ErrorLog.Println("Can't get item from session")
		m.AppConfig.SessionManager.Put(
			r.Context(),
			"error",
			"Невозможно получить сводку бронирования",
		)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.AppConfig.SessionManager.Remove(r.Context(), "reservation")

	data := make(jet.VarMap)
	data.Set("reservation", reservation)

	render.RenderTemplate(w, r, "reservation-summary.jet", data)
}
