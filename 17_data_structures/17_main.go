package main

import (
	"fmt"
	"errors"
)

type person struct {
	name string
	age  int
}

type Level uint8

const (
	Junior Level = iota + 1
	SemiSenior
	Senior
)

type worker struct {
	id uint32
	active bool
	person
	yearsInCompany uint8
	level Level
}

func main() {
	p := createPerson("John Doe", 30)
	w := createWorker(1, true, p)

	printWorkerInfo(w)

	separator()

	err := registerAniversary(&w)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	fmt.Println("After 1st anniversary:")
	printWorkerInfo(w)

	separator()
	
	deactivateWorker(&w)
	fmt.Println("After deactivation:")
	printWorkerInfo(w)

	separator()
	
	err = registerAniversary(&w)
	if err != nil {
		fmt.Println("Error:", err)
	}

	separator()
	
	activateWorker(&w)
	fmt.Println("After activation:")
	printWorkerInfo(w)
	
	separator()
	
	err = registerAniversary(&w)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("After 2nd anniversary:")
	printWorkerInfo(w)

	separator()

	registerAniversary(&w)
}

func createPerson(name string, age int) person {
	if age < 18 || age > 70 || age == 0 {
		_ = errors.New("Invalid age")
		return person{}
	}

	if name == "" {
		_ = errors.New("Name cannot be empty")
		return person{}
	}

	return person{name: name, age: age}
}

func createWorker(id uint32, active bool, p person) worker {
	return worker{id: id, active: active, person: p}
}

func printWorkerInfo(w worker) {
	fmt.Printf("Worker ID: %d\n", w.id)
	fmt.Printf("Active: %t\n", w.active)
	fmt.Printf("Name: %s\n", w.name)
	fmt.Printf("Age: %d\n", w.age)
	fmt.Printf("Years in Company: %d\n", w.yearsInCompany)
	fmt.Printf("Level: %d\n", w.level)
}

func registerAniversary(w *worker) error {

	if !w.active {
		err := errors.New("Cannot register anniversary for inactive worker")
		return err
	}

	w.yearsInCompany++
	w.person.age++

	if w.yearsInCompany >= 2 && w.yearsInCompany < 5 {
		w.level = SemiSenior
	} else if w.yearsInCompany >= 5 {
		w.level = Senior
	}
	return nil
}

func deactivateWorker(w *worker) error {

	if !w.active {
		err := errors.New("Worker is already inactive")
		return err
	}

	w.active = false
	return nil
}

func activateWorker(w *worker) error {
	
	if w.active {
		err := errors.New("Worker is already active")
		return err
	}

	w.active = true
	return nil
}

func separator() {
	fmt.Println("----------------------------")
}