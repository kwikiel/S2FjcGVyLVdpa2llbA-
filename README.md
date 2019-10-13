# S2FjcGVyLVdpa2llbA-

[![Build Status](https://travis-ci.com/kwikiel/S2FjcGVyLVdpa2llbA-.svg?branch=master)](https://travis-ci.com/kwikiel/S2FjcGVyLVdpa2llbA-)

Zadanie rekrutacyjne GWP

Uwagi: 
- Obsługa błędów typu 400, 413 nie jest zaimplementowana: prawdopodobnie jest to bardzo proste przy użyciu chi, 
tutaj jednak chciałem skupić się na nauczeniu się Go

- Worker jest uruchamiany poprzez /worker 

Zasadniczo to lepiej było by użyć bazy danych + stworzyć narzędzie command line do obsługi / włączania / wyłączanie workera

- Historia pobrań pod /api/fetcher/ zwraca wszystkie obiekty. 

- Worker jest stworzony w oparciu o goroutine 
