package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	r := gin.Default()

	// Setup CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/chat", func(c *gin.Context) {
		var input struct {
			Message string `json:"message"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
			return
		}

		// Pastikan Anda sudah menjalankan: export GEMINI_API_KEY="kunci_anda"
		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey == "" {
			fmt.Println("[ERROR] GEMINI_API_KEY tidak ditemukan di environment variable")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server configuration error: Missing API Key"})
			return
		}

		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			fmt.Printf("[ERROR] Gagal membuat client: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal terhubung ke layanan AI"})
			return
		}
		defer client.Close()

		model := client.GenerativeModel("models/gemini-2.5-flash-lite")
		
		// Proses Generate Content
		resp, err := model.GenerateContent(ctx, genai.Text(input.Message))
		if err != nil {
			fmt.Printf("[ERROR] Gemini API Error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "AI gagal merespon: " + err.Error()})
			return
		}

		var reply string
		if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
			for _, part := range resp.Candidates[0].Content.Parts {
				reply += fmt.Sprintf("%v", part)
			}
		} else {
			reply = "Maaf, AI tidak memberikan jawaban."
		}

		c.JSON(http.StatusOK, gin.H{"reply": reply})
	})

	fmt.Println("Server backend berjalan di port :8080...")
	r.Run(":8080")
}