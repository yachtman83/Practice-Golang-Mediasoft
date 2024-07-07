package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Car struct {
	ID      int    `json:"id"`
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Mileage int64  `json:"mileage"`
	Owners  int64  `json:"owners"`
}

var cars []Car

func saveCars() {
	data, _ := json.MarshalIndent(cars, "", "  ")
	os.WriteFile("cars.json", data, 0644)
}

func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}
	for _, item := range cars {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCar Car

	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newCar.ID = rand.Intn(10000)
	for _, car := range cars {
		if car.ID == newCar.ID {
			http.Error(w, "ID conflict, try again", http.StatusConflict)
			return
		}
	}

	cars = append(cars, newCar)
	saveCars()
	json.NewEncoder(w).Encode(newCar)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var updatedFields map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updatedFields); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for i, item := range cars {
		if item.ID == id {
			if brand, ok := updatedFields["brand"].(string); ok {
				cars[i].Brand = brand
			}
			if model, ok := updatedFields["model"].(string); ok {
				cars[i].Model = model
			}
			//приведение ко float64, так как в таком формате приходят
			//числа в json. Это касается числовых полей.
			if mileage, ok := updatedFields["mileage"].(float64); ok {
				cars[i].Mileage = int64(mileage)
			}
			if owners, ok := updatedFields["owners"].(float64); ok {
				cars[i].Owners = int64(owners)
			}
			saveCars()
			json.NewEncoder(w).Encode(cars[i])
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

func updateCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedCars []Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCars); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for _, updatedCar := range updatedCars {
		for i, item := range cars {
			if item.ID == updatedCar.ID {
				cars[i] = updatedCar
				break
			}
		}
	}
	saveCars()
	json.NewEncoder(w).Encode(updatedCars)
}

/*func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    for i, car := range cars {
        if car.ID == id {
            cars = append(cars[:i], cars[i+1:]...)
            saveCars()
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "Car not found", http.StatusNotFound)
}*/

func deleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range cars {
		if item.ID == id {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
	}
	saveCars()
	json.NewEncoder(w).Encode(cars)
}

func main() {
	r := mux.NewRouter()
	cars = append(cars, Car{ID: 1, Brand: "Toyota", Model: "Camry", Mileage: 55000, Owners: 1})
	//books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	r.HandleFunc("/cars", getCars).Methods("GET")
	r.HandleFunc("/cars/{id}", getCar).Methods("GET")
	r.HandleFunc("/cars", createCar).Methods("POST")
	r.HandleFunc("/cars", updateCars).Methods("PUT")
	r.HandleFunc("/cars/{id}", updateCar).Methods("PATCH")
	r.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
