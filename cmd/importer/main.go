package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"church_user/config"
	"church_user/domain"

	"github.com/joho/godotenv"
	"github.com/xuri/excelize/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDb()
	log.Println("db start")
	log.Println(db)
	log.Println("db end")

	f, err := excelize.OpenFile("Hoja de vida MIVD.xlsx")
	if err != nil {
		log.Fatal("Error abriendo el archivo:", err)
	}

	// Suponemos que los datos están en la primera hoja
	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Fatal("Error leyendo las filas:", err)
	}

	for i, row := range rows {
		if i == 0 {
			continue // saltamos el encabezado
		}

		// Cambia los índices según el orden de las columnas en tu archivo
		user := domain.User{
			FirstName:            strings.TrimSpace(row[1]),
			LastName:             strings.TrimSpace(row[2]),
			DocumentType:         strings.TrimSpace(row[3]),
			DocumentNumber:       strings.TrimSpace(row[4]),
			Birthdate:            parseDate(strings.TrimSpace(row[5])),
			Birthplace:           strings.TrimSpace(row[6]),
			Address:              strings.TrimSpace(row[7]),
			Landline:             strings.TrimSpace(row[8]),
			Mobile:               strings.TrimSpace(row[9]),
			Email:                strings.TrimSpace(row[10]),
			Neighborhood:         strings.TrimSpace(row[11]),
			Locality:             strings.TrimSpace(row[12]),
			SocioeconomicStratum: parseInt(strings.TrimSpace(row[13])),
			BloodType:            strings.TrimSpace(row[14]),
			Profession:           strings.TrimSpace(row[15]),
			Company:              strings.TrimSpace(row[16]),
			Position:             strings.TrimSpace(row[17]),
			MaritalStatus:        strings.TrimSpace(row[18]),
			LeaderName:           strings.TrimSpace(row[19]),
			ConversionDate:       parseDate(strings.TrimSpace(row[20])),
			InvitedBy:            strings.TrimSpace(row[21]),
			AttendedOtherChurch:  parseBool(strings.TrimSpace(row[22])),
			OtherChurchName:      strings.TrimSpace(row[23]),
			EncounterDate:        parseDate(strings.TrimSpace(row[24])),
			BaptismDate:          parseDate(strings.TrimSpace(row[25])),
			MinisterialFunction:  strings.TrimSpace(row[26]),
			LastSchoolLevel:      strings.TrimSpace(row[27]),
			PescaTeam:            strings.TrimSpace(row[28]),
			SpouseName:           strings.TrimSpace(row[29]),
			SpouseIsMember:       parseBool(strings.TrimSpace(row[30])),
			MarriageDate:         parseDate(strings.TrimSpace(row[31])),
			ParentsNames:         strings.TrimSpace(row[32]),
			ParentsAreMembers:    parseBool(strings.TrimSpace(row[33])),
			ChildrenNames:        strings.TrimSpace(row[34]),
			HasAllergy:           parseBool(strings.TrimSpace(row[35])),
			DataAuthorization:    parseBool(strings.TrimSpace(row[36])),

			// y así sucesivamente...

			// Ejemplo de fecha

		}

		log.Println(user.FirstName)

		// Assuming config.DB is undefined because it's not exported or a getter function is preferred.
		// The previous code attempted to use config.GetDB(), but the lint error indicates it's undefined.
		// Assuming config.DB is the exported GORM DB instance initialized by config.InitDb().
		err := db.Create(&user).Error
		if err != nil {
			fmt.Printf("Error guardando fila %d: %v\n", i+1, err)
		}
	}
	fmt.Println("Importación completada.")
}

func parseDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}
	}
	return t
}
func parseInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}
	return int(num)
}

func parseBool(s string) bool {
	if s == "Si" {
		return true
	}
	return false
}
