package e2e_test

import (
	"backend/core/config"
	"backend/core/pkg/storage"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

const (
	invalidBotToken = "INVALID_TOKEN"
	invalidChatID   = "999999999"
)

func openDB(t *testing.T) *pgxpool.Pool {
	ctx := context.Background()

	cfg := config.Load()
	pool, err := pgxpool.New(ctx, storage.PostgresDsn(cfg))
	require.NoError(t, err)

	err = pool.Ping(ctx)
	require.NoError(t, err)

	return pool
}

func postJSONWithStatus(t *testing.T, url string, body any, expectedStatus int) {
	b, _ := json.Marshal(body)

	cfg := config.Load()
	req, err := http.NewRequest(http.MethodPost, "http://"+cfg.ApiHost+url, bytes.NewBuffer(b))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	require.NoError(t, err)

	require.Equal(t, expectedStatus, resp.StatusCode)
}

func randomOrderNumber() string {
	return fmt.Sprintf("TEST-%d", time.Now().UnixNano())
}

func randomOrderTotal() int {
	return rand.Intn(1000)
}

func TestOrder_SendsTelegram(t *testing.T) {
	db := openDB(t)
	defer db.Close()

	cfg := config.Load()
	postJSONWithStatus(t, "/api/shops/2/telegram/connect", map[string]any{
		"bot_token":  cfg.TestBotToken,
		"chat_id":    cfg.TestChatID,
		"is_enabled": true,
	}, http.StatusOK)

	orderNumber := randomOrderNumber()
	postJSONWithStatus(t, "/api/shops/2/orders", map[string]any{
		"number":        orderNumber,
		"total":         randomOrderTotal(),
		"customer_name": "Tanya",
	}, http.StatusOK)

	time.Sleep(1 * time.Second)

	ctx := context.Background()

	var orderID int64
	err := db.QueryRow(ctx, `SELECT id FROM orders WHERE number = $1`, orderNumber).Scan(&orderID)
	require.NoError(t, err)

	var status string
	err = db.QueryRow(ctx, `SELECT status FROM telegram_send_log WHERE order_id = $1`, orderID).Scan(&status)
	require.NoError(t, err)

	require.Equal(t, "SENT", status)
}

func TestOrder_Idempotent(t *testing.T) {
	db := openDB(t)
	defer db.Close()

	cfg := config.Load()
	postJSONWithStatus(t, "/api/shops/2/telegram/connect", map[string]any{
		"bot_token":  cfg.TestBotToken,
		"chat_id":    cfg.TestChatID,
		"is_enabled": true,
	}, http.StatusOK)

	orderNumber := randomOrderNumber()
	order := map[string]any{
		"number":        orderNumber,
		"total":         randomOrderTotal(),
		"customer_name": "Tanya",
	}

	postJSONWithStatus(t, "/api/shops/2/orders", order, http.StatusOK)
	postJSONWithStatus(t, "/api/shops/2/orders", order, http.StatusUnprocessableEntity)

	time.Sleep(1 * time.Second)

	ctx := context.Background()

	var orderID int64
	db.QueryRow(ctx, `SELECT id FROM orders WHERE number = $1`, orderNumber).Scan(&orderID)

	var count int
	db.QueryRow(ctx, `SELECT COUNT(*) FROM telegram_send_log WHERE order_id = $1`, orderID).Scan(&count)

	require.Equal(t, 1, count)
}

func TestOrder_TelegramFailed(t *testing.T) {
	db := openDB(t)
	defer db.Close()

	postJSONWithStatus(t, "/api/shops/2/telegram/connect", map[string]any{
		"bot_token":  invalidBotToken,
		"chat_id":    invalidChatID,
		"is_enabled": true,
	}, http.StatusOK)

	orderNumber := randomOrderNumber()
	postJSONWithStatus(t, "/api/shops/2/orders", map[string]any{
		"number":        orderNumber,
		"total":         randomOrderTotal(),
		"customer_name": "FailCase",
	}, http.StatusOK)

	time.Sleep(1 * time.Second)

	ctx := context.Background()

	var orderID int64
	db.QueryRow(ctx, `SELECT id FROM orders WHERE number = $1`, orderNumber).Scan(&orderID)

	var status string
	db.QueryRow(ctx, `SELECT status FROM telegram_send_log WHERE order_id = $1`, orderID).Scan(&status)

	require.Equal(t, "FAILED", status)
}
