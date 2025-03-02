## Known Issues
- Lack of Captcha Retry Mechanism: If captcha solving fails, the bot does not attempt to retry, leading to purchase failures.
- Low Captcha Solving Accuracy: The current implementation using Tesseract OCR is not highly reliable and may struggle with ambiguous captchas.
- Hard-Coded Configurations: Settings such as browser path, ticket details, and exclusion keywords are predefined in the code rather than being configurable through a user-friendly interface.
- Limited Platform Compatibility: The bot is specifically designed for tixcraft.com and may not work with other ticketing platforms without modifications.
- Slower Performance Compared to Selenium: The bot operates slower than Selenium-based implementations, potentially reducing its effectiveness in high-demand ticketing scenarios.
- Ambiguous Naming Conventions: Some package names, functions, and variables lack clarity, making the codebase harder to read and maintain.
- Incorrect Exclusion Logic: The filtering mechanism for excluding unwanted ticket options does not function as expected, which could lead to unintended ticket selections.
- Inefficient Blocking Mechanism: The use of select{} in browser.go causes the program to block indefinitely without a structured exit strategy.
- Limited Error Handling: The bot lacks robust error handling, leading to abrupt failures instead of implementing retries or fallback mechanisms.

## Prerequisites
- Go environment
- Tesseract