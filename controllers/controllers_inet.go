package controllers

import (
	"errors"
	"log"
	"regexp"
	"strconv"

	"golang-training/models"
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
	// รับค่า query param ชื่อ tax_id
	input := c.Query("tax_id")
	if input == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Query parameter 'tax_id' is required",
		})
	}

	// แปลงแต่ละตัวอักษรเป็น ASCII
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
	var req models.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "อ่านข้อมูลไม่สำเร็จ",
		})
	}

	// validate struct ด้วย validator
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "กรอกข้อมูลผิดพลาด",
			"errors":  err.Error(),
		})
	}

	// รหัสผ่านไม่ตรงกัน
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
	// ตรวจ username ด้วย regexp (ตัวอย่าง: ห้ามมีช่องว่าง และต้องเป็นตัวอักษร a-z, A-Z, 0-9 เท่านั้น)
	usernameValid, _ := regexp.MatchString("^[a-zA-Z0-9]+$", req.Username)
	if !usernameValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username ต้องไม่มีช่องว่างและเป็นตัวอักษรหรือตัวเลขเท่านั้น",
		})
	}

	// สำเร็จ
	return c.JSON(fiber.Map{
		"message": "ลงทะเบียนสำเร็จ",
		"data":    req,
	})
}
