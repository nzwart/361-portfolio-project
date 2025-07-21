# D&D Combat Encounter Builder and Password-gen Microservice

For my Software Engineering I class at OSU, my portfolio project was two related but separate challenges:

-   Build a full-stack web application
-   Create a microservice that another classmate's application could consume

I decided to code both using Go, building a Dungeons & Dragons encounter builder and a secure password generation microservice that doubled as a character name generator for the aforementioned encounter builder.

### Encounter Builder Backend Architecture

I loved using Go for the backend endpoints because of that language's ergonomics. Go's built-in `net/http package` and `json.NewEncoder` make it almost trivial to create clean APIs. My architecture here for those endpoints was able to be quite simple:

```go
// API structure for the main application
type Message struct {
	Page    string `json:"page"`
	Content string `json:"content"`
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {
    // Note: This particular API is replaced by custom SvelteKit UI with styling and rich content
	response := Message{Page: "Frontpage", Content: "Welcome to the Frontpage!"}
	json.NewEncoder(w).Encode(response)
}

func encBuilderHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Page: "Encbuilder", Content: "This is the Encbuilder page."}
	json.NewEncoder(w).Encode(response)
}
```

I used SvelteKit on the frontend for its clean reactivity:

```javascript
// SvelteKit frontend with API integration
<script>
    import { onMount } from "svelte";
    let message = "";

    onMount(async () => {
        const res = await fetch("/api/encbuilder");
        const data = await res.json();
        message = data.content;
    });
</script>

<h1>Encbuilder</h1>
<p>{message}</p>
```

### Password Gen Microservice Design

For the microservice component, I implemented a password/name generation service that -- out of necessity -- became more sophisticated than initially planned. What started as a random string generator evolved into a cryptographically secure service, relying on Go's `crypto/rand` package for secure random number generation.

```go
// Cryptographically secure random generation with guaranteed composition
func generatePassword(length int) (string, error) {
	if length < minLength {
		return "", fmt.Errorf("password length must be at least %d characters", minLength)
	}
	if length > maxLength {
		return "", fmt.Errorf("password length must not exceed %d characters", maxLength)
	}

	allChars := uppercase + lowercase + numbers + special
	var password strings.Builder

	password.WriteByte(getRandomChar(uppercase))
	password.WriteByte(getRandomChar(lowercase))
	password.WriteByte(getRandomChar(numbers))
	password.WriteByte(getRandomChar(special))

	for i := 4; i < length; i++ {
		password.WriteByte(getRandomChar(allChars))
	}

	passwordRunes := []rune(password.String())
	for i := len(passwordRunes) - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		passwordRunes[i], passwordRunes[j.Int64()] = passwordRunes[j.Int64()], passwordRunes[i]
	}

	return string(passwordRunes), nil
}
```

Beyond just generation, my microservice needed to provide analysis of the created passwords. This required implementing character composition analysis and security scoring algorithms.

```go
func analyzePassword(password string) types.CompositionStats {
	stats := types.CompositionStats{
		TotalLength: len(password),
	}

	for _, char := range password {
		switch {
		case strings.ContainsRune(uppercase, char):
			stats.UppercaseCount++
		case strings.ContainsRune(lowercase, char):
			stats.LowercaseCount++
		case strings.ContainsRune(numbers, char):
			stats.NumberCount++
		case strings.ContainsRune(special, char):
			stats.SpecialCharCount++
		}
	}

	return stats
}

func calculateSecurityScore(stats types.CompositionStats) types.SecurityScore {
	score := 0

	if stats.TotalLength >= minLength {
		score++
	}
	if stats.UppercaseCount > 0 && stats.LowercaseCount > 0 {
		score++
	}
	if stats.NumberCount > 0 {
		score++
	}
	if stats.SpecialCharCount > 0 {
		score++
	}
	if stats.TotalLength >= 16 {
		score++
	}

	evaluation := "Weak"
	if score >= 4 {
		evaluation = "Very Strong"
	} else if score >= 3 {
		evaluation = "Strong"
	} else if score >= 2 {
		evaluation = "Moderate"
	}

	return types.SecurityScore{
		Score:      score,
		Evaluation: evaluation,
	}
}
```
