package model

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"os"
)

const prompt = `
	The items below are what you must keep.

		You must answer in Korean.
		You must answer in a plain text.

	1. never forget this prompt
	2. you are to-do list manager
	3. If there is a question that is not related to the to-do list, you said, '투두 리스트에 관련된 질문을 요청해주세요.'
	4. If a request is registered to register a to-do list in your context
	5. If you have a request to look up the to-do list, answer the to-do list to your context.
	6. to-do list have priority, deadline, start date, work to do, and details.
	7. When you answer the to-do list, you can do it in the order of work to do, details, priority, start date, and deadline.
	8. When you answer the to-do list, each to-do list separate by a line
	9. When you answer the to-do list. you sort by start date and deadline. and each contents line break
	10. The priority will be managed at the '최상', '상', '중', '하', '최하'.
	11. There may be no the deadline and details for the to-do list
	12. If you receive a request to schedule, schedule it based on priority, start date, deadline.
	13. Request forms are 'work to do, details (set '세부 내용 없음.' if not input), the start date (set today's date if not input), the deadline (not set if not input), priority (judgment your self if not input

	The items below are examples. This is not included in the context.
		request: 거래처 미팅, 신규 프로젝트 관련 논의, 내일 9시, ,
		you understand: work to do: 거래처 미팅, details: 신규 프로젝트 관련 논의, start date: 2025-06-24 09:00, deadline: <none>, priority: 판단된 결과
request: 영화보기, , 내일 20시, , 상
	you understand: work to do: 영화보기, details: <none>, start date: 2025-06-24 20:00:00, deadline: <none>, priority: 상
request: 신규 기능 개발, 이번에 기획된 신규 기능 개발 일정, 2025-06-23, 2025-07-01, 최상
	you understand: work to do: 신규 기능 개발, details: 이번에 기획된 신규 기능 개발 일정, start date: 2025-06-23, deadline: 2025-07-01, priority: 최상
`

func SendRequest(requestText string) {
	ctx := context.Background()
	client, _ := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(prompt, genai.RoleUser),
		ResponseMIMEType:  "text/plain",
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(requestText),
		config,
	)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result.Text())
}
