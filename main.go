package main

import (
	"encoding/csv"
	"go-fiber-swag/model/domain"
	"log"
	"strconv"

	"github.com/go-the-way/exl"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func main() {

	records := []domain.Merchant{
		{Id: uuid.New(), Email: "test1@mail.com", MdrPercentage: 12.5, InterestSharingBankPercentage: 5, InterestSharingMerchantPercentage: 5},
		{Id: uuid.New(), Email: "test2@mail.com", MdrPercentage: 12.5, InterestSharingBankPercentage: 5, InterestSharingMerchantPercentage: 7},
		{Id: uuid.New(), Email: "test3@mail.com", MdrPercentage: 12.5, InterestSharingBankPercentage: 5, InterestSharingMerchantPercentage: 6},
		{Id: uuid.New(), Email: "test4@mail.com", MdrPercentage: 12.5, InterestSharingBankPercentage: 5, InterestSharingMerchantPercentage: 4},
	}
	app := fiber.New()

	app.Get("/csv", func(c *fiber.Ctx) {
		c.Attachment("test.csv")
		writer := csv.NewWriter(c.Fasthttp.Response.BodyWriter())
		var data [][]string

		data = append(data, []string{"ID", "Email", "MDR Percentage", "InterestSharingBankPercentage", "InterestSharingMerchantPercentage"})
		for _, m := range records {
			row := []string{m.Id.String(), m.Email, strconv.FormatFloat(m.MdrPercentage, 'f', 6, 64), strconv.FormatFloat(m.InterestSharingBankPercentage, 'f', 6, 64), strconv.FormatFloat(m.InterestSharingMerchantPercentage, 'f', 6, 64)}
			data = append(data, row)
		}

		err := writer.WriteAll(data)
		if err != nil {
			log.Panic(err)
		}
	})

	app.Get("/xlsx", func(c *fiber.Ctx) {
		w := exl.NewWriter()
		w.Write("sheet1", records)
		_, err := w.WriteTo(c.Fasthttp.Response.BodyWriter())
		c.Attachment("test.xlsx")
		if err != nil {
			log.Panic(err)
		}
	})

	app.Listen(":3000")

	// csvFile, err := os.Create("merchant.csv")
	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }
	// defer csvFile.Close()

	// writer := csv.NewWriter(csvFile)
	// defer writer.Flush()

	// var data [][]string

	// for _, m := range records {
	// 	row := []string{m.Id.String(), m.Email, strconv.FormatFloat(m.MdrPercentage, 'f', 6, 64), strconv.FormatFloat(m.InterestSharingBankPercentage, 'f', 6, 64), strconv.FormatFloat(m.InterestSharingMerchantPercentage, 'f', 6, 64)}
	// 	data = append(data, row)
	// }

	// writer.WriteAll(data)
}

