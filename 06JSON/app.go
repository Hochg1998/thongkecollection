package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

/* https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
"name":"Dee Leng",
"email":"dleng0@cocolog-nifty.com",
"job":"developer",
"gender":"Female",
"city":"London",
"salary":9662,
"birthdate":"2007-09-30" */
type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("person.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var people []Person

	json.Unmarshal(byteValue, &people)

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(&people[i])
		}
	*/
	peopleByCity := GroupPeopleByCity(people)
	for key, value := range peopleByCity {
		fmt.Println(key)
		for _, person := range value {
			fmt.Println("  ", (&person).Name)
		}
	}

	// fmt.Println(peopleByCity)
	peopleByJob := GroupPeopleByJob(people)
	fmt.Println(peopleByJob)
	// for index, value := range peopleByJob {
	// 	fmt.Println(key, "-", value)
	// }
	fmt.Println("Top 5 Jobs is:")
	for i := len(Top5JobsByNumer(peopleByJob)) - 1; i > len(Top5JobsByNumer(peopleByJob))-6; i-- {

		fmt.Printf("%v: %v \n", len(Top5JobsByNumer(peopleByJob))-i, Top5JobsByNumer(peopleByJob)[i].Name)
	}

	top5CitiesByNumber := Top5CitiesByNumber(GroupCityByNumberOfPeople(people))
	fmt.Println("Top 5 Cities is:")
	for i := len(top5CitiesByNumber) - 1; i > len(top5CitiesByNumber)-6; i-- {

		fmt.Printf("%v: %v \n", len(top5CitiesByNumber)-i, top5CitiesByNumber[i].Name)
	}

	jobInEachCity := JobInEachCity(people)
	result := make(map[string]map[string]int)

	for k, v := range jobInEachCity {
		r := CountNumberEachJob(v)
		result[k] = r
	}
	fmt.Println("Trong mỗi thành phố nghề được làm nhiều nhất là")
	topJobByNumerInEachCity := TopJobByNumerInEachCity(result)
	for k, v := range topJobByNumerInEachCity {
		fmt.Printf("%v : %v \n", k, v)

	}

	salaryInEachJob := SalaryInEachJob(people)
	result1 := make(map[string]float64)
	for k, v := range salaryInEachJob {
		r := CountAverageSalary(v)
		result1[k] = r
	}
	fmt.Println("Với mỗi nghề mức lương trung bình là:")
	for k, v := range result1 {
		fmt.Printf("%v có mức lương trung bình: %.3f \n", k, v)
	}
	fmt.Println("thành phố có mức lương trung bình cao nhất")
	salaryInEachCity := SalaryInEachCity(people)
	result2 := make(map[string]float64)
	for k, v := range salaryInEachCity {
		r := CountAverageSalary(v)
		result2[k] = r
	}
	convertSalaryInEachCity := ConvertSalaryInEachCity(result2)
	top5CitiesBySalary := Top5CitiesBySalary(convertSalaryInEachCity)
	fmt.Println("Top 5 Highest Paid Cities is:")
	for i := len(top5CitiesBySalary) - 1; i > len(top5CitiesBySalary)-6; i-- {

		fmt.Printf("%v: %v %.3f \n", len(top5CitiesBySalary)-i, top5CitiesBySalary[i].Name, top5CitiesBySalary[i].Salary)
	}
	result3 := make(map[string][]int)

	for _, city := range layTenTungThanhPho(people) {
		var salaries []int
		for _, person := range people {
			if person.Job == "developer" && person.City == city {
				salaries = append(salaries, person.Salary)

			}
			result3[city] = salaries
		}
	}
	// fmt.Println(result3)
	//roi tinh muc luong trung binh cua deverloper tren tung thanh pho
	result4 := make(map[string]float64)
	for k, v := range result3 {
		result4[k] = CountAverageSalary(v)
	}
	// fmt.Println(result4)
	//loc nhung thanh pho ko co salari for developer
	for k, v := range result4 {
		if math.IsNaN(v) {
			delete(result4, k)
		}
	}
	//sort tung thanh pho dua tren muc luong
	result5 := ConvertSalaryInEachCityByDeverloper(result4)
	// fmt.Println(result5)
	result6 := FiveCitiesHasTopSalaryForDeveloper(result5)
	// fmt.Println(result6)
	fmt.Println("Năm thành phố có mức lương trung bình của developer cao nhất")
	for i := len(result6) - 1; i > len(result6)-6; i-- {

		fmt.Printf("%v: %v %.3f \n", len(result6)-i, result6[i].Name, result6[i].AverageSalary)
	}
	//lay ten tung job
	result7 := layTenTungNghe(people)
	//cho slice if int cua tuoi
	result8 := make(map[string]float64)
	for _, job := range result7 {
		var AgesPerJob []int
		for _, person := range people {
			if person.Job == job {
				AgesPerJob = append(AgesPerJob, layTuoi(person.Birthday))

			}
			result8[job] = CountAverageSalary(AgesPerJob)
		}
	}
	fmt.Println("Tuoi trung binh cua tung nghe nghiep la:")
	for k, v := range result8 {
		fmt.Printf("%v : %.2f \n", k, v)
	}
	result9 := make(map[string]float64)
	for _, city := range layTenTungThanhPho(people) {
		var AgesPerCity []int
		for _, person := range people {
			if person.City == city {
				AgesPerCity = append(AgesPerCity, layTuoi(person.Birthday))

			}
			result9[city] = CountAverageSalary(AgesPerCity)
		}
	}
	fmt.Println("Tuoi trung binh cua tung thanh pho la:")
	for k, v := range result9 {
		fmt.Printf("%v : %.2f \n", k, v)
	}
}
