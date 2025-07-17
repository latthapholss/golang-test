package controllers

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"

	"golang-training/database"
	m "golang-training/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func BodyParser(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func Hello(c *fiber.Ctx) error {
	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func Search(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}
func Valid(c *fiber.Ctx) error {
	//Connect to database
	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("number must be non-negative")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func Fact(c *fiber.Ctx) error {
	numberStr := c.Params("number")
	if numberStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Number is required",
		})
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid number",
		})
	}

	fact, err := factorial(number)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"number":    number,
		"factorial": fact,
	})
}

func Ascii(c *fiber.Ctx) error {
	input := c.Query("tax_id")
	if input == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Query parameter 'tax_id' is required",
		})
	}

	asciiCodes := []int{}
	for _, ch := range input {
		asciiCodes = append(asciiCodes, int(ch))
	}

	return c.JSON(fiber.Map{
		"input": input,
		"ascii": asciiCodes,
	})
}

func Register(c *fiber.Ctx) error {
	var req m.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "อ่านข้อมูลไม่สำเร็จ",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "กรอกข้อมูลผิดพลาด",
			"errors":  err.Error(),
		})
	}

	if req.Password != req.InlinePassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "รหัสผ่านและยืนยันรหัสผ่านไม่ตรงกัน",
		})
	}
	nameWebsiteValid, _ := regexp.MatchString("^[a-z0-9-]{2,30}$", req.NameWebsite)
	if !nameWebsiteValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ชื่อเว็บไซต์ต้องมีความยาว 2-30 ตัว ใช้ได้เฉพาะ a-z, 0-9, และ - เท่านั้น",
		})
	}
	usernameValid, _ := regexp.MatchString("^[a-zA-Z0-9]+$", req.Username)
	if !usernameValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username ต้องไม่มีช่องว่างและเป็นตัวอักษรหรือตัวเลขเท่านั้น",
		})
	}

	return c.JSON(fiber.Map{
		"message": "ลงทะเบียนสำเร็จ",
		"data":    req,
	})
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	countRed := 0
	countGreen := 0
	countPink := 0
	countNoColor := 0
	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""

		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			countRed++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			countGreen++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			countPink++
		} else {
			typeStr = "no color"
			countNoColor++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:       dataResults,
		Name:       "golang-test",
		Count:      len(dogs), //หาผลรวม,
		SumRed:     countRed,
		SumGreen:   countGreen,
		SumPink:    countPink,
		SumNoColor: countNoColor,
	}

	return c.Status(200).JSON(r)
}

func GetDeletedDocs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	result := db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(dogs)
}

func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.Company

	db.Find(&companies)

	return c.Status(200).JSON(companies)
}

func GetDogsFilter50(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	result := db.Where("dog_id > ? AND dog_id < ?", 50, 100).Find(&dogs)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(dogs)
}

func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company
	result := db.First(&company, id)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(company)
}

func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if err := validate.Struct(company); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "กรอกข้อมูลผิดพลาด",
			"errors":  err.Error(),
		})
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")

	var existingCompany m.Company
	if err := db.First(&existingCompany, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "ไม่พบข้อมูลบริษัท",
		})
	}

	var updateData m.Company
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "รูปแบบข้อมูลไม่ถูกต้อง",
			"error":   err.Error(),
		})
	}

	if err := db.Model(&existingCompany).Updates(updateData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "อัปเดตข้อมูลไม่สำเร็จ",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(existingCompany)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
