package dbcontroller

import (
	"fmt"
	"main/db/connectionhandler"
)

type WebhookEntry struct {
	ID    string
	Token string
	Host  string
}

func RegisterWebhook(id string, authToken string, host string) {
	db := connectionhandler.GetConnection()
	db.Query("insert into webhooks values (?, ?, ?)", id, authToken, host)
	db.Close()
}

func GetAllWebhooks() ([]WebhookEntry, error) {
	var entries []WebhookEntry
	db := connectionhandler.GetConnection()
	rows, err := db.Query("select * from webhooks")
	if err != nil {
		return nil, fmt.Errorf("webhooks %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var entry WebhookEntry
		if err := rows.Scan(&entry.ID, &entry.Token, &entry.Host); err != nil {
			return nil, fmt.Errorf("webhooks %v", err)
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("webhooks %v", err)
	}

	return entries, nil
}

type WebhookEntrySanitary struct {
	ID   string
	Host string
}

func GetAllWebhooksSanitary() ([]WebhookEntrySanitary, error) {
	var entries []WebhookEntrySanitary
	db := connectionhandler.GetConnection()
	rows, err := db.Query("select id, host from webhooks")
	if err != nil {
		return nil, fmt.Errorf("webhooks %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var entry WebhookEntrySanitary
		if err := rows.Scan(&entry.ID, &entry.Host); err != nil {
			return nil, fmt.Errorf("webhooks %v", err)
		}
		entries = append(entries, entry)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("webhooks %v", err)
	}

	return entries, nil
}
