package browser

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"

	"github.com/yuansheng0111/Tbot/internal/config"
)

func New(cfg config.Config) {
	// Launch browser with custom path
	u := launcher.New().Bin(cfg.Browser_Path).Headless(false).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose() // Close browser on exit

	page := browser.MustPage(("https://tixcraft.com/"))
	page.MustWaitLoad()
	log.Println("Tixcraft page loaded successfully")

	// step 1: click on "立即購票"
	buyButton := page.MustElement(".buy a") // Match text
	buyButton.MustClick()
	log.Println("Buy button clicked")

	// step 2: select date)
	xpathQuery := fmt.Sprintf(`//tr[td[contains(text(), '%s')]]`, cfg.Date)
	row := page.MustElementX(xpathQuery) // should be one
	row.MustElement("button").MustClick()
	log.Println("Date selected")

	// step 3: select ticket area
	xpathQuery = fmt.Sprintf(`//li[@class='select_form_a' or @class='select_form_b'][a[contains(text(), '%s')]]`, cfg.Price)
	rows := page.MustElementsX(xpathQuery) // might be multiple

	found := false
	for _, element := range rows {
		span := element.MustElement("span")
		spanText := strings.TrimSpace(span.MustText())

		// Ensure no exclusion words exist
		exclude := false
		for _, word := range cfg.Exclude {
			if strings.Contains(spanText, word) {
				exclude = true
				break
			}
		}

		if !exclude {
			element.MustElement("a").MustClick()
			found = true
			break
		}
	}
	if found {
		log.Println("Ticket area selected")
	} else {
		log.Fatal("Ticket area not found")
	}

	// step 4: select ticket number
	ticketSelect := page.MustElement("select.form-select.mobile-select")
	ticketSelect.MustSelect(cfg.Ticket_number)
	log.Println("Ticket number selected")

	// step 5: fill the check box (agree to terms)
	page.MustElement("#TicketForm_agree").MustClick()

	// step 6: solve captcha
	captchaText, err := solveCaptcha(page, 10)
	if err != nil {
		log.Fatal(err)
	}

	inputField := page.MustElement("#TicketForm_verifyCode")
	inputField.MustClick()
	inputField.MustInput(captchaText)
	inputField.MustType(input.Enter)
	log.Println("Captcha submitted")

	// step 7: click on submit
	page.MustElement("button.btn.btn-primary.btn-green").MustClick()

	// page.MustWaitLoad()
	select {}
}

func solveCaptcha(page *rod.Page, maxRetry int) (string, error) {
	for attempt := 0; attempt < maxRetry; attempt++ {
		img := page.MustElement("#TicketForm_verifyCode-image")
		img.MustWaitVisible()
		img.MustWaitStable()
		img.MustWaitLoad()
		imgBytes, err := img.Screenshot(proto.PageCaptureScreenshotFormatPng, 1000)
		if err != nil {
			return "", err
		}
		os.WriteFile("captcha.png", imgBytes, 0644)

		out, err := exec.Command("tesseract", "captcha.png", "stdout", "--psm", "7").Output()
		if err != nil {
			log.Println("Error in solving captcha")
			log.Fatal(err)
		}
		captchaText := strings.TrimSpace(string(out))
		log.Printf("Captcha solved: %s\n", captchaText)

		if len(captchaText) == 4 {
			return captchaText, nil
		}
		img.MustClick()
	}
	log.Fatal("Failed to solve captcha after maximum retries")
	return "", fmt.Errorf("Failed to solve captcha after maximum retries")
}
