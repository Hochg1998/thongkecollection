package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bearbin/go-age"
)

type job struct {
	Name   string
	Number int
}
type city struct {
	Name           string
	NumberOfPeople int
}
type citySalary struct {
	Name   string
	Salary float64
}

// type JobsInCity struct {
// 	City string
// 	Job  []string
// }
type AllJobsinCity struct {
	Name    string
	AllJobs []map[string]int
}
type AverageSalaryForDeverloperInEachCity struct {
	Name          string
	AverageSalary float64
}

func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}

func GroupCityByNumberOfPeople(p []Person) []city {
	cityMap := make(map[string]int)
	var result []city
	for _, person := range p {
		cityMap[person.City]++
	}
	for k, v := range cityMap {
		var city1 city
		city1.Name = k
		city1.NumberOfPeople = v
		result = append(result, city1)
	}
	return result
}
func GroupPeopleByJob(p []Person) []job {
	jobMap := make(map[string]int)
	var result []job
	for _, person := range p {
		jobMap[person.Job]++
	}
	for k, v := range jobMap {
		var job1 job
		job1.Name = k
		job1.Number = v
		result = append(result, job1)
	}
	return result
}

func Top5JobsByNumer(jobs []job) []job {
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].Number < jobs[j].Number })
	return jobs
}

func Top5CitiesByNumber(cities []city) []city {
	sort.Slice(cities, func(i, j int) bool { return cities[i].NumberOfPeople < cities[j].NumberOfPeople })
	return cities
}

func JobInEachCity(p []Person) map[string][]string {
	result := make(map[string][]string)
	fmt.Println(result)

	for _, person := range p {
		result[person.City] = append(result[person.City], person.Job)

	}
	return result
}
func CountNumberEachJob(input []string) (result map[string]int) {
	result = make(map[string]int)
	for _, job := range input {
		result[job]++
	}
	return result
}

func TopJobByNumerInEachCity(j map[string]map[string]int) map[string][]string {
	result := make(map[string][]string)
	for key1, value1 := range j {
		maxCount := 0
		for _, value2 := range value1 {
			if maxCount < value2 {
				maxCount = value2
			}
		}
		for key3, value3 := range value1 {
			if value3 == maxCount {
				result[key1] = append(result[key1], key3)
			}
		}

	}
	return result

	// }
}
func SalaryInEachJob(p []Person) map[string][]int {
	result := make(map[string][]int)

	for _, person := range p {
		result[person.Job] = append(result[person.Job], person.Salary)

	}
	return result
}
func CountAverageSalary(input []int) (result float64) {
	count := 0.0
	for _, salary := range input {
		count++
		result += float64(salary)
	}
	result = result / count
	return result
}
func SalaryInEachCity(p []Person) map[string][]int {
	result := make(map[string][]int)

	for _, person := range p {
		result[person.City] = append(result[person.City], person.Salary)

	}
	return result
}
func ConvertSalaryInEachCity(p map[string]float64) []citySalary {
	// salaryMap := make(map[string]float64)
	var result1 []citySalary

	for k, v := range p {
		var Salary1 citySalary
		Salary1.Name = k
		Salary1.Salary = v
		result1 = append(result1, Salary1)
	}
	return result1
}
func Top5CitiesBySalary(citySalaryS []citySalary) []citySalary {
	sort.Slice(citySalaryS, func(i, j int) bool { return citySalaryS[i].Salary < citySalaryS[j].Salary })
	return citySalaryS
}

// FiveCitiesHasTopSalaryForDeveloper
// func AverageSalaryOfDeverloprer()
//lay ten tung thanh pho
//Neu trung ten trung developer thi append vao []int cua luong
//roi tinh muc luong trung binh cua deverloper tren tung thanh pho
//sort tung thanh pho dua tren muc luong

func layTenTungThanhPho(p []Person) []string {
	var results []string
	results = append(results, p[0].City)
	for _, person := range p {
		count := 0
		for _, result := range results {
			if person.City == result {
				count++
			}
		}
		if count == 0 {
			results = append(results, person.City)
		}
	}
	return results

}
func ConvertSalaryInEachCityByDeverloper(p map[string]float64) []AverageSalaryForDeverloperInEachCity {
	// salaryMap := make(map[string]float64)
	var result1 []AverageSalaryForDeverloperInEachCity

	for k, v := range p {
		var Salary1 AverageSalaryForDeverloperInEachCity
		Salary1.Name = k
		Salary1.AverageSalary = v
		result1 = append(result1, Salary1)
	}
	return result1
}
func FiveCitiesHasTopSalaryForDeveloper(a []AverageSalaryForDeverloperInEachCity) []AverageSalaryForDeverloperInEachCity {
	sort.Slice(a, func(i, j int) bool { return a[i].AverageSalary < a[j].AverageSalary })
	return a
}
func layTenTungNghe(p []Person) []string {
	var results []string
	results = append(results, p[0].Job)
	for _, person := range p {
		count := 0
		for _, result := range results {
			if person.Job == result {
				count++
			}
		}
		if count == 0 {
			results = append(results, person.Job)
		}
	}
	return results

}
func getAge(year, month, day int) time.Time {
	Age := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return Age
}
func layTuoi(p string) int {
	var tuoi int

	txt := strings.Split(p, "-")
	txtY, _ := strconv.Atoi(txt[0])
	txtM, _ := strconv.Atoi(txt[1])
	txtD, _ := strconv.Atoi(txt[2])
	agePerson := getAge(txtY, txtM, txtD)
	tuoi = age.Age(agePerson)

	return tuoi
}
