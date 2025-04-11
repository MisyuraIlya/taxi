package geo

import "net/http"

type Handler struct {
	service Service
}

func NewHandler(router *http.ServeMux, service Service) {
	handler := &Handler{
		service: service,
	}

	router.HandleFunc("PUT /updateLocation", handler.UpdateLocation)
	router.HandleFunc("GET /GetLocation", handler.GetLocation)
}

func (h *Handler) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	driverId := r.URL.Query().Get("driverId")
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	err := h.service.UpdateLocation(driverId, latitude, longitude)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Location updated successfully"))
}

func (h *Handler) GetLocation(w http.ResponseWriter, r *http.Request) {
	driverId := r.URL.Query().Get("driverId")

	latitude, longitude, err := h.service.GetLocation(driverId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Latitude: " + latitude + ", Longitude: " + longitude))
}
