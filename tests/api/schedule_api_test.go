package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL = "http://test-api:8080/api"
	Green   = "\033[32m"
	Red     = "\033[31m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Reset   = "\033[0m"
)

type TestCase struct {
	Name           string
	Method         string
	URL            string
	Body           interface{}
	ExpectedStatus int
	Description    string
}

type ScheduleCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectID   uint   `json:"project_id"`
}

type ScheduleCreateResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProjectID   uint      `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	fmt.Printf("%s🚀 Starting API Tests for Schedule Create Endpoint%s\n", Blue, Reset)
	fmt.Printf("%s=================================================%s\n", Blue, Reset)

	testCases := []TestCase{
		{
			Name:   "Create Schedule - Success",
			Method: "POST",
			URL:    "/schedule",
			Body: ScheduleCreateRequest{
				Name:        "Test Schedule",
				Description: "Test Description",
				ProjectID:   1,
			},
			ExpectedStatus: 201,
			Description:    "Should successfully create a schedule with valid data",
		},
		{
			Name:   "Create Schedule - Missing Name",
			Method: "POST",
			URL:    "/schedule",
			Body: ScheduleCreateRequest{
				Description: "Test Description",
				ProjectID:   1,
			},
			ExpectedStatus: 400,
			Description:    "Should return validation error when name is missing",
		},
		{
			Name:   "Create Schedule - Missing Description",
			Method: "POST",
			URL:    "/schedule",
			Body: ScheduleCreateRequest{
				Name:      "Test Schedule",
				ProjectID: 1,
			},
			ExpectedStatus: 400,
			Description:    "Should return validation error when description is missing",
		},
		{
			Name:   "Create Schedule - Missing Project ID",
			Method: "POST",
			URL:    "/schedule",
			Body: ScheduleCreateRequest{
				Name:        "Test Schedule",
				Description: "Test Description",
			},
			ExpectedStatus: 400,
			Description:    "Should return validation error when project_id is missing",
		},
		{
			Name:   "Create Schedule - Empty Name",
			Method: "POST",
			URL:    "/schedule",
			Body: ScheduleCreateRequest{
				Name:        "",
				Description: "Test Description",
				ProjectID:   1,
			},
			ExpectedStatus: 400,
			Description:    "Should return validation error when name is empty",
		},
		{
			Name:   "Create Schedule - Empty Description",
			Method: "POST",
			URL:    "/schedule",
			Body: ScheduleCreateRequest{
				Name:        "Test Schedule",
				Description: "",
				ProjectID:   1,
			},
			ExpectedStatus: 400,
			Description:    "Should return validation error when description is empty",
		},
		{
			Name:   "Create Schedule - Zero Project ID",
			Method: "POST",
			URL:    "/schedule",
			Body: ScheduleCreateRequest{
				Name:        "Test Schedule",
				Description: "Test Description",
				ProjectID:   0,
			},
			ExpectedStatus: 400,
			Description:    "Should return validation error when project_id is zero",
		},
		{
			Name:           "Create Schedule - Invalid JSON",
			Method:         "POST",
			URL:            "/schedule",
			Body:           `{"name": "Test", "description": "Test", "project_id": "invalid"}`,
			ExpectedStatus: 400,
			Description:    "Should return error when JSON is malformed",
		},
		{
			Name:           "Create Schedule - Empty Body",
			Method:         "POST",
			URL:            "/schedule",
			Body:           nil,
			ExpectedStatus: 400,
			Description:    "Should return error when request body is empty",
		},
	}

	passed := 0
	failed := 0

	for i, testCase := range testCases {
		fmt.Printf("\n%s[Test %d/%d]%s %s\n", Yellow, i+1, len(testCases), Reset, testCase.Name)
		fmt.Printf("%s📝 %s%s\n", Blue, testCase.Description, Reset)

		success := runTest(testCase)
		if success {
			passed++
			fmt.Printf("%s✅ PASSED%s\n", Green, Reset)
		} else {
			failed++
			fmt.Printf("%s❌ FAILED%s\n", Red, Reset)
		}
	}

	fmt.Printf("\n%s=================================================%s\n", Blue, Reset)
	fmt.Printf("%s📊 Test Results:%s\n", Blue, Reset)
	fmt.Printf("%s✅ Passed: %d%s\n", Green, passed, Reset)
	fmt.Printf("%s❌ Failed: %d%s\n", Red, failed, Reset)
	fmt.Printf("%s📈 Success Rate: %.1f%%%s\n", Blue, float64(passed)/float64(len(testCases))*100, Reset)

	if failed > 0 {
		fmt.Printf("\n%s⚠️  Some tests failed. Please check the API implementation.%s\n", Yellow, Reset)
	} else {
		fmt.Printf("\n%s🎉 All tests passed! The Schedule Create API is working correctly.%s\n", Green, Reset)
	}
}

func runTest(testCase TestCase) bool {
	client := &http.Client{Timeout: 10 * time.Second}

	var body io.Reader
	if testCase.Body != nil {
		if str, ok := testCase.Body.(string); ok {
			body = bytes.NewBufferString(str)
		} else {
			jsonBody, err := json.Marshal(testCase.Body)
			if err != nil {
				fmt.Printf("%s❌ Error marshaling request body: %v%s\n", Red, err, Reset)
				return false
			}
			body = bytes.NewBuffer(jsonBody)
		}
	}

	req, err := http.NewRequest(testCase.Method, BaseURL+testCase.URL, body)
	if err != nil {
		fmt.Printf("%s❌ Error creating request: %v%s\n", Red, err, Reset)
		return false
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s❌ Error making request: %v%s\n", Red, err, Reset)
		return false
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s❌ Error reading response: %v%s\n", Red, err, Reset)
		return false
	}

	fmt.Printf("🔍 Request: %s %s\n", testCase.Method, BaseURL+testCase.URL)
	if testCase.Body != nil {
		if str, ok := testCase.Body.(string); ok {
			fmt.Printf("📤 Body: %s\n", str)
		} else {
			bodyJSON, _ := json.MarshalIndent(testCase.Body, "", "  ")
			fmt.Printf("📤 Body: %s\n", string(bodyJSON))
		}
	}
	fmt.Printf("📥 Response Status: %d (Expected: %d)\n", resp.StatusCode, testCase.ExpectedStatus)
	fmt.Printf("📥 Response Body: %s\n", string(responseBody))

	if resp.StatusCode != testCase.ExpectedStatus {
		fmt.Printf("%s❌ Status code mismatch. Expected: %d, Got: %d%s\n", Red, testCase.ExpectedStatus, resp.StatusCode, Reset)
		return false
	}

	// Additional validation for successful creation
	if testCase.ExpectedStatus == 201 && resp.StatusCode == 201 {
		var response ScheduleCreateResponse
		if err := json.Unmarshal(responseBody, &response); err != nil {
			fmt.Printf("%s❌ Error parsing success response: %v%s\n", Red, err, Reset)
			return false
		}

		if response.ID == 0 {
			fmt.Printf("%s❌ Created schedule should have a valid ID%s\n", Red, Reset)
			return false
		}

		if response.CreatedAt.IsZero() {
			fmt.Printf("%s❌ Created schedule should have a valid CreatedAt timestamp%s\n", Red, Reset)
			return false
		}

		fmt.Printf("🆔 Created Schedule ID: %d\n", response.ID)
		fmt.Printf("📅 Created At: %s\n", response.CreatedAt.Format(time.RFC3339))
	}

	return true
}
